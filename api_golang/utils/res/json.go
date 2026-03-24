package res

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// helper func code base -> but current using gin.H for response, so not use this func for now

// interface{} empty == any data type
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
		err := json.NewEncoder(w).Encode(data)

		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}