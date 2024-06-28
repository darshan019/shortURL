package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func godot_env_var(key string) string {
	ok := godotenv.Load(".env")
	if ok != nil {
		fmt.Println("error loading env")
	}
	return os.Getenv(key)
}

func Create_token(username, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	token_string, err := token.SignedString([]byte(godot_env_var("SECRETKEY")))
	if err != nil {
		return "", err
	}
	return token_string, nil
}

func Verify_token(token_string string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return []byte(godot_env_var("SECRETKEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid expiration claim")
	}

	//USE EXP TO FORM REFRESH TOKEN HERE

	fmt.Println(time.Unix(int64(exp), 0))
	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return nil, fmt.Errorf("token is expired")
	}

	return claims, nil
}

func Protect_routes(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token_string := r.Header.Get("Authorization")
		if token_string == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing auth header")
			return
		}
		token_string = token_string[len("Bearer "):]

		claims, err := Verify_token(token_string)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid Token")
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
