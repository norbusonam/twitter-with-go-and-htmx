package handlers

import (
	"net/http"

	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/db"
)

func getAuthenticatedUser(cookies []*http.Cookie) (*db.User, error) {
	for _, cookie := range cookies {
		if (cookie.Name == "user_id") && (cookie.Value != "") {
			return db.GetUserById(cookie.Value)
		}
	}
	return nil, nil
}
