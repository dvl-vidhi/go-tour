package auth

import (
	"fmt"
	"net/http"
	"online-election-system/config"

	"github.com/golang-jwt/jwt"
)

type User_Claims struct {
	Username   string
	Authorized bool
	Role       string
}

var JWT_SECRET_KEY = []byte(config.APP_CONFIG.JWT_SECRET_KEY)

func GenerateJWT(user_claims User_Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      user_claims.Username,
		"role":       user_claims.Role,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(JWT_SECRET_KEY)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte("You're Unauthorized"))
					if err != nil {
						return nil, err
					}
				}
				return "", nil

			})
			// parsing errors result
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err2 := writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
				if err2 != nil {
					return
				}

			}
			// if there's a token
			if token.Valid {
				endpointHandler(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}
		// response for if there's no token header
	})
}

func ExtractClaims(_ http.ResponseWriter, request *http.Request) (User_Claims, error) {
	var user_claims User_Claims
	if request.Header["Token"] != nil {
		tokenString := request.Header["Token"][0]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return JWT_SECRET_KEY, nil
		})
		if err != nil {
			return user_claims, err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			user_claims.Username = claims["email"].(string)
			user_claims.Role = claims["role"].(string)
			user_claims.Authorized = claims["authorized"].(bool)
			return user_claims, nil
		}
	}

	return user_claims, nil
}
