package utils

import (
	"github.com/husobee/vestigo"
	"net/http"
	"strconv"
)

func ParseUserId(request *http.Request) (*int64, error) {
	userId, err := strconv.ParseInt(vestigo.Param(request, "userId"), 10, 64)
	if err != nil {
		return nil, NewValidationErrorFromError(err)
	}
	return &userId, nil
}
