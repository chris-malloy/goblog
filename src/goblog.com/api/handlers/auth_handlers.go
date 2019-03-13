package handlers

import (
	"errors"
	"github.com/husobee/vestigo"
	"goblog.com/api/utils"
	"net/http"
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

func ValidateJWT(writer http.ResponseWriter, request *http.Request, handlerFunc http.HandlerFunc) {
	tokenParts := strings.Split(request.Header.Get("Authorization"), "Bearer")
	if len(tokenParts) != 2 {
		utils.RenderErrorWithStatus(writer, errors.New("validateJWT error: got invalid authorization header"), http.StatusUnauthorized)
		return
	}

	return
}
