package utils


import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string] interface{}) {
	return map[string] interface{} {"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string] interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
