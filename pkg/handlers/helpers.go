package handlers

import "net/http"

func isUserAuthenticated(cookies []*http.Cookie) bool {
	for _, cookie := range cookies {
		if cookie.Name == "user_id" {
			return true
		}
	}
	return false
}
