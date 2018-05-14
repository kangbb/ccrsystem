/*
* Validdate some information of business.
* The data from user input maybe often need a validation.
 */
package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/kangbb/ccrsystem/logs"
)

type UserErrorMsg struct {
	Id  string
	Pwd string
}

/*
* Validate the user information format.
* If pass, return true; else return false.
 */
func UserIfoFormatValidate(id string, pwd string, w http.ResponseWriter) bool {
	var id_msg string
	var pwd_msg string
	match_id, err := regexp.MatchString(`^\d{8}$`, id)
	if logs.NormalError(err) {
		w.WriteHeader(500)
		return false
	}
	match_pwd, err := regexp.MatchString(`^[a-zA-Z0-9][^\s]{5,15}`, pwd)
	if logs.NormalError(err) {
		w.WriteHeader(500)
		return false
	}
	if match_id && match_pwd {
		return true
	}
	if !match_id {
		id_msg = "您输入的学工号码不符合格式"
	}
	if !match_pwd {
		pwd_msg = "您输入的密码不符合格式"
	}
	w.WriteHeader(500)
	data, _ := json.Marshal(UserErrorMsg{id_msg, pwd_msg})
	w.Write(data)
	return false
}

/*
* Validate the user password whether as same as the password in database.
* if same, return true; else return false.
 */
func UserPasswordValidate(signinPwd string, databasePwd string, w http.ResponseWriter) bool {
	if signinPwd != databasePwd {
		w.WriteHeader(500)
		data, _ := json.Marshal(UserErrorMsg{"", "您输入的密码不正确"})
		w.Write(data)
		return false
	}
	return true
}
