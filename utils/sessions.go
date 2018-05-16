package utils

import "github.com/gorilla/sessions"

var ppioSessions *sessions.CookieStore = sessions.NewCookieStore([]byte("something-very-secret"))

func GetSessionStore() *sessions.CookieStore {
	 return ppioSessions
}
