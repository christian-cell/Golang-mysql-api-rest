package commons

import (
	"encoding/json"
	"log"
	"net/http"
)

/* func SendResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(status)
	w.Write(data)
}
*/

func JsonRespond(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func SendError(w http.ResponseWriter, status int) {
	data := []byte(`{message : "error en la peticion}`)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func Catch(err error) {
	if err != nil {
		log.Println(err)
	}
}
