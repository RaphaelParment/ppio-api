package utils

import "github.com/gorilla/sessions"

var ppioSessions *sessions.CookieStore = sessions.NewCookieStore([]byte("something-very-secret"))

func getSessionStore() *sessions.CookieStore {
	 return ppioSessions
}
