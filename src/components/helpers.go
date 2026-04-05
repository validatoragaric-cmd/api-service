package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

// JSONResponse sends a JSON response with the given status code and payload.
func JSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// ReadJSON decodes the request body into the given destination.
func ReadJSON(r *http.Request, dest interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dest)
}

// ErrorResponse sends a JSON error response with the given status code and message.
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	JSONResponse(w, statusCode, map[string]string{"error": message})
}

// IsValidEmail checks if the given string is a valid email address.
func IsValidEmail(email string) bool {
	// Simple regex for basic email validation
	// For production use, consider a more robust solution
	return len(email) >= 3 && len(email) <= 254 && 
		regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

// GenerateRandomString generates a random string of the given length.
func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil
}