package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Header|Payload|Signature
// Claims ~ Payload
type Claims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
	Email string `json:"email"`
	
}

func CreateToken(jwtSecret string ,email string,role string)(string,error){
	now:=time.Now().UTC()
	exp:=now.Add(7*24*time.Hour)
	

	claims:=Claims{

		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt:jwt.NewNumericDate(exp),
			Issuer:"cmrohityadavappltd",
		},
		Role: role,
		Email: email,
	}

	jwtToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	sToken,err:=jwtToken.SignedString([]byte(jwtSecret))

	if err!=nil{
		return "",fmt.Errorf("signed token faild: %w",err)
	}

	return sToken,nil

	

}

func ParseToken(jwtSecret string,tokenString string)(Claims,error){
	var claims Claims
	parsed,err:=jwt.ParseWithClaims(
			tokenString,
	
			&claims,
		
			func(t *jwt.Token) (any, error) {
				if t.Method.Alg()!=string(jwt.SigningMethodHS256.Alg()){
					return nil,fmt.Errorf("unexptected signing metheod : %v",t.Header["alg"])
				}

				return []byte(jwtSecret),nil
			},

		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)


	if err!=nil{
		return  Claims{}, fmt.Errorf("parse token failed: %w",err)
	}

	if !parsed.Valid {
		return  Claims{},errors.New("Invalid token")
	}

	return claims,nil


}
