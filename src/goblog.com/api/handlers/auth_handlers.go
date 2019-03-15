package handlers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/husobee/vestigo"
	"goblog.com/api/utils"
	"net/http"
	"os"
	"strings"
)

func RequireLogin() vestigo.Middleware {
	return requiresLoginMiddleware
}

var requiresLoginMiddleware = func(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ValidateJWT(writer, request, next)
	}
}

func ValidateJWT(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	tokenParts := strings.Split(request.Header.Get("Authorization"), "Bearer")
	if len(tokenParts) != 2 {
		utils.RenderErrorWithStatus(writer, errors.New("validateJWT error: got invalid authorization header"), http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(strings.TrimSpace(tokenParts[1]), getAndValidJWTSecret)

	if err != nil {
		utils.RenderErrorWithStatus(writer, err, http.StatusUnauthorized)
		return
	}

	if isTokenValidAfterClaims(token) {
		next(writer, request)
	} else {
		utils.RenderErrorWithStatus(writer, errors.New("invalid JWT token"), http.StatusUnauthorized)
	}
}

func getAndValidJWTSecret(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if !isSigningMethodOkay(token) {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	secret, ok := getJWTSecret()
	if !ok {
		return nil, fmt.Errorf("JWT secret not found")
	}

	return secret, nil
}

func getJWTSecret() ([]byte, bool) {
	secret, ok := os.LookupEnv("JWT_SECRET")
	return []byte(secret), ok
}

func isSigningMethodOkay(token *jwt.Token) bool {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	return ok
}

func isTokenValidAfterClaims(token *jwt.Token) bool {
	_, ok := token.Claims.(jwt.MapClaims)
	return ok && token.Valid
}
