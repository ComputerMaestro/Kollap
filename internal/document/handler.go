package document

import (
	"io"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Service Working")
}

func GetDocuments(w http.ResponseWriter, r *http.Request) {

}
