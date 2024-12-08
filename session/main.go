package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key = []byte("my-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/secret", secret)
}

func login(w http.ResponseWriter, r *http.Request) {
	// クッキーにセッションIDを設定
	session,err := store.Get(r, "session-name")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// セッションに値を設定
	// TODO　セッション確立済みであることをハックされないようにするために、sessionIDにUUIDを使う
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func secret(w http.ResponseWriter, r *http.Request) {
	// セッションから値を取得する
	session,err := store.Get(r, "session-name")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sessionValue := session.Values["authenticated"]
	if sessionValue == nil{
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}else{
		w.Write([]byte("secret"))
	}
}
