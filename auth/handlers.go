package auth

import (
	"encoding/json"
	"fmt"
	"github.com/claudioontheweb/go-blog/db"
	"github.com/claudioontheweb/go-blog/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid Request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)

}

func FindOne(email, password string) map[string]interface{} {
	user := &User{}

	if err := db.DB.Where("email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := models.Token{
		UserID: user.ID,
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Email: user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)

	createdUser := db.DB.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}


//FetchUser function
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.DB.Preload("auths").Find(&users)

	json.NewEncoder(w).Encode(users)
}

// Update User
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	params := mux.Vars(r)
	var id = params["id"]
	db.DB.First(&user, id)
	json.NewDecoder(r.Body).Decode(user)
	db.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user User
	db.DB.First(&user, id)
	db.DB.Delete(&user)
	json.NewEncoder(w).Encode("User deleted")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user User
	db.DB.First(&user, id)
	json.NewEncoder(w).Encode(&user)
}