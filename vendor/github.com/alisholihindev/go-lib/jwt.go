package lib

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

func TokenValidation(TokenString string) (jwt.MapClaims, error) {

	secret := []byte(GetConfig("JWT_KEY"))

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, nil
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	if err != nil {
		CommonLogger().Error(err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		err = errors.New("invalid token")
		CommonLogger().Error(err)
		return nil, err
	}

}
