/*
 * @File: utils.jwtMgr.go
 * @Description: Reusable jwt manager stuffs for services
 * @Author: Yuda Fatah (yudafatah@gmail.com)
 */
package utils

import (
	"errors"
	"time"
	"github.com/yudafatah/golang-microservices-boilerplate/tree/main/Identity/common"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	LastName string `json:"last_name"`
	FirstName string `json:"first_name"`
	Username string `json:"username"`
	Role string `json:"role"`
	Nationality string `json:"nationality"`
	IsCompany bool `json:"is_company"`
	UserId string `json:"user_id"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	ClientId string `json:"client_id"`
	jwt_lib.StandardClaims
}

type Utils struct {
}

// GenerateJWT generates token from the given information
func (u *Utils) GenerateJWT(LastName string, FirstName string, Username string, Role string,
	Nationality string, IsCompany bool, UserId string, Email string, 
	PhoneNumber string, ClientId string) (string, error) {
	claims := SdtClaims{
		LastName,
		FirstName,
		Username,
		Role, 
		Nationality,
		IsCompany,
		UserId,
		Email,
		PhoneNumber,
		ClientId,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 999999).Unix(),
			Issuer:    common.Config.Issuer,
			Audience: common.Config.Audience,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common.Config.JwtSecretPassword))

	return tokenString, err
}

// ValidateObjectID checks the given ID if it's an object id or not
func (u *Utils) ValidateObjectID(id string) error {
	if bson.IsObjectIdHex(id) != true {
		return errors.New(common.ErrNotObjectIDHex)
	}

	return nil
}
