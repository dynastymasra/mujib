package formatter

import (
	"fmt"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, r *http.Request, statusCode int, data string) {

	w.WriteHeader(statusCode)
	fmt.Fprint(w, data)
}
