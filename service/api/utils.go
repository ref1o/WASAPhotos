package api

import (
	"net/http"
	"strings"
)

// Function that verifies if the identifier of a user has the right lenght
func validIdentifier(identifier string) bool {
	trimmedId := strings.TrimSpace(identifier)

	// Controlla se il trimmedId è vuoto
	if trimmedId == "" {
		return false
	}

	// Controlla la lunghezza del trimmedId
	if len(trimmedId) < 3 || len(trimmedId) > 16 {
		return false
	}

	// Controlla la presenza di caratteri non consentiti
	if strings.ContainsAny(trimmedId, "?_") {
		return false
	}

	return true
}

// Function that extracts the bearer token from the Authorization header
func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}

// Function that checks if the requesting user has a valid token for the specified endpoint. Returns 0 if it's valid, the error (as a int, representing the http status) otherwise
func validateRequestingUser(identifier string, bearerToken string) int {

	// If the requesting user has an invalid token then respond with a fobidden status
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}

	//  If the requesting user's id is different than the one in the path then respond with a unathorized status.

	if identifier != bearerToken {
		return http.StatusUnauthorized
	}
	return 0
}

// Function that checks if a user is logged
func isNotLogged(auth string) bool {

	return auth == ""
}
