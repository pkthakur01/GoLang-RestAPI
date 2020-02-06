package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go-login/config/db"
	"go-login/model"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	logger "go-login/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

/*
* POST API for user registration
* Input : body {username, firstname, lastname & password}
* Proccess : 1. Check if user alreday exist or not. 2. Encrypt hte passsword. 3. Save user details to database
* Output : User registered successfully if there is no Error Occurred.
 */

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Reached Here")

	var user model.User                //geting User model from model structure
	body, _ := ioutil.ReadAll(r.Body)  //
	err := json.Unmarshal(body, &user) //parsing body input to user variable
	var res model.ResponseResult       //getting response model from model structure
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		logger.ErrorLogger.Println("Error occurred " + res.Error) //sending error as response message
		return
	}

	collection, err := db.GetDBCollection() //getdbCollection function to connect database and create table

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.User
	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			logger.ErrorLogger.Println("No documents in result")
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			user.Password = string(hash)

			_, err = collection.InsertOne(context.TODO(), user)
			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				logger.ErrorLogger.Println("Error While Creating User, Try Again")
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Registration Successful"
			logger.GeneralLogger.Println("Registration Successful")
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		logger.ErrorLogger.Println(res)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Username already Exists!!"
	logger.ErrorLogger.Println("User already Exists")
	json.NewEncoder(w).Encode(res)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	collection, err := db.GetDBCollection()

	if err != nil {
		log.Fatal(err)
	}
	var result model.User
	var res model.ResponseResult

	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)

	if err != nil {
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result.Username,
		"firstname": result.FirstName,
		"lastname":  result.LastName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Token = tokenString
	result.Password = "Can't be disclosed as it's confidential !!"

	json.NewEncoder(w).Encode(result)

}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	var result model.User
	var res model.ResponseResult
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Username = claims["username"].(string)
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)

		json.NewEncoder(w).Encode(result)
		return
	} else {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

}
