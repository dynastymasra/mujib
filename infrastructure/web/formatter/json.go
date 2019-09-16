package formatter

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type JSONFormat struct {
	Status  int         `json:"status" binding:"required"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func FailResponse(status int, msg string) JSONFormat {
	return JSONFormat{Status: status, Message: msg}
}

func SuccessResponse(status int, data interface{}) JSONFormat {
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
