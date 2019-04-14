package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"goblog.com/api/models"
	"goblog.com/api/utils"
	"net/http"
)

func CreateUser(auth models.Authorizer, users models.UserCRUD) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var newUser models.NewUserRequest
		err := json.NewDecoder(request.Body).Decode(&newUser)
		if err != nil {
			utils.RenderErrorAndDeriveStatus(writer, utils.NewValidationError("Invalid user JSON."))
			return // bail
		}

		// The contract of this method will always only result in a 400
		validator := utils.ValidateNewUserRequest(newUser)
		if !validator.Ok {
			utils.RenderErrorAndDeriveStatus(writer, err)
			return
		}

		// If we get an error here we have a serious problem
		user, err := users.InsertUser(newUser)
		if err != nil {
			utils.RenderErrorAndDeriveStatus(writer, err)
			return
		}

		// good to go.
		writer.WriteHeader(http.StatusCreated)

		token, err := getJWTToken(user)
		if err != nil {
			utils.RenderErrorAndDeriveStatus(writer, err)
			return
		}

		err = users.DoLoginBookkeeping(user.Email)
		if err != nil {
			// Not a critical failure.
			log.WithError(err).WithFields(log.Fields{
				"email": user.Email,
			}).Warn("Unable to call bookkeep for user")
		}

		preFetchUser := user
		user, err = users.SelectUserByEmail(newUser.Email)
		if err != nil {
			user = preFetchUser

			log.WithError(err).WithFields(log.Fields{
				"email": user.Email,
			}).Warn("Unable to retrieve the newly created user from the database.")
		}

		json.NewEncoder(writer).Encode(LoginResponse{user, token})
	}
}

func GetUserById(users models.UserCRUD) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		userId, err := utils.ParseUserId(request)
		if err != nil {
			utils.RenderErrorAndDeriveStatus(writer, err)
			return
		}

		user, err := users.SelectUserById(*userId)
		if err != nil {
			utils.RenderErrorAndDeriveStatus(writer, err)
			return
		}

		writer.WriteHeader(http.StatusOK)

		json.NewEncoder(writer).Encode(user)
	}
}
