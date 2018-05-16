/*
* This file include some middlewares.
* Some of them are used by gorilla/mux.
 */
package middlewares

import (
	"encoding/gob"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/sessions"
	"github.com/kangbb/ccrsystem/logs"
)

var store *sessions.FilesystemStore

/*
* Initial the session store.
 */
func init() {
	gob.Register(time.Time{})
	store = sessions.NewFilesystemStore("./data/sessions", []byte("ccrsystem"))
}

/*
* The middleware is used by gorilla/mux to resolve the Cross-Domain problems.
* And the next middleware will be used to validate the session.
 */
func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		w.Header().Set("Content-Type", "application/json")

		//not-so-simple request will initiate(发起) a preflight request("预检"请求) before the formal request
		//such as PUT, DELETE, and the vaule of Content-Type is application/json.
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
			w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
			w.WriteHeader(200)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

/*
* The middleware is used by gorilla/mux to validate the users' legality
* The url that is '/' needn't validate
* If the user is illegal, then will redirect to the 'index.html', which for user to signin
* else the next will process the users' business request.
 */

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Some url needn't to validate the session
		match_signin, _ := regexp.MatchString(".*signin", r.URL.Path)
		match_static, _ := regexp.MatchString("^/static.*", r.URL.Path)
		if r.URL.Path == "/" || match_signin || match_static {
			next.ServeHTTP(w, r)
			return
		}

		// The others need to validate the session
		session, _ := store.Get(r, "user")

		// If session doesn't exist, redirect to the 'index.htm'
		if session.Values["id"] == nil || session.Values["type"] == nil || session.Values["accessTime"] == nil {
			req, _ := http.NewRequest("GET", "/", r.Body)
			http.Redirect(w, req, "/", 302)
		}
		// If session expired, redirect to the 'index.htm'
		now := time.Now()
		sub := now.Sub(session.Values["accessTime"].(time.Time)).Seconds()
		if int(sub)-session.Options.MaxAge >= 0 || session.Values["type"] == nil {
			session.Options.MaxAge = -1
			session.Save(r, w)
			req, _ := http.NewRequest("GET", "/", r.Body)
			http.Redirect(w, req, "/", 302)
		}

		// Give the  user type to the business handle functions, for further verification of the permission
		r.Header.Set("userType", session.Values["type"].(string))
		r.Header.Set("userId", session.Values["id"].(string))
		next.ServeHTTP(w, r)
	})
}

/*
* A smiple middleware just used to process the session in sigin and signout
* When used in sigin, save some information in session and save the session.
* When used in signout, delete the session
* If process sucessfully, return true; else return false.
 */
func SessionProcess(w http.ResponseWriter, r *http.Request, arg ...interface{}) bool {
	var err error
	session, _ := store.Get(r, "user")
	if len(arg) == 0 {
		session.Options.MaxAge = -1
		err = session.Save(r, w)
		if logs.NormalError(err, w) {
			return false
		}
	} else {
		session.Values["id"] = arg[0].(string)
		session.Values["type"] = arg[1].(string)
		session.Values["name"] = arg[2].(string)
		session.Values["accessTime"] = time.Now()
		session.Options.MaxAge = 86400
		err = session.Save(r, w)
		if logs.NormalError(err, w) {
			return false
		}
	}
	return true
}
