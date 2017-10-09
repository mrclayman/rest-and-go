package handlers

import (
	"net/http"
)

// LoginHandler handles login POST requests by verifying
// the client's credentials. Upon successful verification,
// the client is issued an authorization token used for
// further communication. In the case of an unsuccessful
// authorization, the HTTP code 403/Forbidden is returned
type LoginHandler struct {
}

// ProcessRequest handles the login POST request
func (handler *LoginHandler) ProcessRequest(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}

	const byteLimit := 256
	bodyData := make([]byte, byteLimit)
	bytesRead, err := request.Body.Read(bodyData)
	if bytesRead == byteLimit && err != nil {
		http.Error(response, "Request body too long", http.StatusBadRequest)
	}
}
