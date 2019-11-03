package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/dynastymasra/mujib/delivery/http/formatter"
	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/mujib/config"
	"github.com/urfave/negroni"
)

func Authorization(secret string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		secretKey := r.Header.Get(config.Authorization)
		encoded := base64.StdEncoding.EncodeToString([]byte(secret))
		serverKey := fmt.Sprintf("Bearer %s", encoded)

		if secretKey != serverKey {
			logrus.WithField("secret_key", secretKey).WithField("server_key", serverKey).Warningln("User doesn't have right access")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, formatter.FailResponse(config.ErrNotAuthorized.Error()).Stringify())
			return
		}
		next(w, r)
	}
}
