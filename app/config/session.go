package config

import "github.com/gorilla/sessions"

const SESSION_ID = "letsadopt_session"

var Store = sessions.NewCookieStore([]byte("anfabwr08q3fnseiof"))