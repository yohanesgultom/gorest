package models

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"	
	u "main/utils"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"; sql:"-"`
}

func GenerateToken(payload *Token) (string) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), payload)
	tokenString, _ := token.SignedString([]byte(os.Getenv("jwt_secret")))
	return tokenString
}

func (user *User) Validate() (map[string] interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	// Email must be unique
	temp := &User{}

	// check for errors and duplicate emails
	e := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if e != nil && e != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error, please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address is already used by other"), false
	}

	return u.Message(true, "Requirement passed"), true
}

func (user *User) Create() (map[string] interface{}) {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	GetDB().Create(user)
	if user.ID <= 0 {
		return u.Message(false, "Connection error, failed to create user")
	}

	user.Token = GenerateToken(&Token{UserId: user.ID})
	user.Password = "" // Remove password from response
	res := u.Message(true, "User successfully created")
	res["user"] = user
	return res
}

func Login(email, password string) (map[string] interface{}) {

	user := &User{}
	e := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if e != nil {
		if e == gorm.ErrRecordNotFound {
			return u.Message(false, "Invalid credential")
		}
		return u.Message(false, "Connection error, please retry")
	}

	e = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if e != nil && e == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid credential")
	}

	user.Password = ""
	user.Token = GenerateToken(&Token{UserId: user.ID})
	res := u.Message(true, "Logged in")
	res["user"] = user
	return res
}

func GetUserById(id uint) (*User, error) {
	user := &User{}
	e := GetDB().Table("users").Where("id = ?", id).First(user).Error
	return user, e
}