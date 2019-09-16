package formatter

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type JSONFormat struct {
	Status  int         `json:"status" binding:"required"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func FailResponse(w http.ResponseWriter, status int, msg string) JSONFormat {
	w.WriteHeader(status)
	return JSONFormat{Status: status, Message: msg}
}

func SuccessResponse(w http.ResponseWriter, status int, data interface{}) JSONFormat {
	w.WriteHeader(status)
	return JSONFormat{Status: status, Message: "Success", Data: data}
}

func (j JSONFormat) Stringify() string {
	toJSON, err := json.Marshal(j)
	if err != nil {
		logrus.WithError(err).Errorln("Unable to stringify JSON")
		return ""
	}
	return string(toJSON)
}
