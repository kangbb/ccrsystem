/*
* A simple logs management.
* It will process all of the errors about the system and write them to file.
 */
package logs

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var Log = &log.Logger{}
var SWITCH_BRANCH_ERROR = errors.New("The branch shouldn't appear.")
var PERMISSION_DENY = errors.New("Permission deny.")

type ErrorMsg struct {
	Msg string
}

/*
* Excute when the system initial
* It will create a log file to store the error logs.
 */
func init() {
	// Define a file to store the logs.
	// Open log function of xorm, and write the logs to file
	fname := "./data/logs/ccrsystemlog/" + time.Now().Format("2006-01-02-15:04:05") + ".log"
	f, err := os.Create(fname)
	if err != nil {
		println(err.Error())
		return
	}

	// New a logger object.
	Log = log.New(f, "[Error]", log.Ldate|log.Ltime|log.Lmicroseconds)
}

/*
* Process the sql operation error, judge whether the error is exist.
* If there is a error, return true; else, return false.
 */
func SqlError(err error, w http.ResponseWriter, hasResult bool, arg ...interface{}) bool {
	var msg []byte
	var pat *regexp.Regexp

	// for some conditions, err = nil, hasResult = true, needn't too much judge
	if err == nil && hasResult {
		return false
	}

	// Get the file who calls the function.
	_, file, line, _ := runtime.Caller(1)
	pathArray := strings.Split(file, "/")
	fname := pathArray[len(pathArray)-1]

	// Judge the err pripority.
	if err != nil {
		// The insert information includes a incorrect string value.
		pat = regexp.MustCompile("Error 1366")
		if res := pat.FindString(err.Error()); len(res) != 0 {
			w.WriteHeader(500)
			msg, _ = json.Marshal(ErrorMsg{Msg: "您插入的信息字段包含非法字符"})
			w.Write(msg)
			Log.Println(fname + fmt.Sprintf(" %d: ", line) + err.Error())
			return true
		}
		// The insert information has a duplicate primary key.
		pat = regexp.MustCompile("Error 1062")
		if res := pat.FindString(err.Error()); len(res) != 0 {
			w.WriteHeader(500)
			msg, _ = json.Marshal(ErrorMsg{Msg: "该条信息已经存在"})
			w.Write(msg)
			Log.Println(fname + fmt.Sprintf(" %d: ", line) + err.Error())
			return true
		}
		// Other errors, regard all of them as exception.
		w.WriteHeader(500)
		msg, _ = json.Marshal(ErrorMsg{Msg: "服务器内部错误"})
		w.Write(msg)
		Log.SetPrefix("[Exception]")
		Log.Println(fname + fmt.Sprintf(" %d: ", line) + err.Error())
		Log.SetPrefix("[Error]")
		return true
	}

	// If the query result from mysql is empty, hasResult = false;
	// else, hasResult = false
	if !hasResult {
		w.WriteHeader(404)
		msg, _ = json.Marshal(ErrorMsg{Msg: "您查询的信息不存在"})
		w.Write(msg)
		Log.Println(fname + fmt.Sprintf(" %d: ", line) + "The query result is empty.")
		return true
	}

	return false
}

/*
* Process the normal error.judge whether the error is exist.
* If there is a error, return true; else, return false.
* As a web application, the system will not panic anytime.
 */
func NormalError(err error, arg ...interface{}) bool {
	if err != nil {
		//panic(err)
		// Get the file who calls the function.
		_, file, line, _ := runtime.Caller(1)
		pathArray := strings.Split(file, "/")
		fname := pathArray[len(pathArray)-1]

		Log.Println(fname + fmt.Sprintf(" %d: ", line) + err.Error())

		if len(arg) != 0 {
			w := arg[0].(http.ResponseWriter)
			w.WriteHeader(500)
			msg, _ := json.Marshal(ErrorMsg{Msg: "服务器内部错误"})
			w.Write(msg)
		}
		return true
	}
	return false
}

/*
* Return the custom error which is result in the wrong parameters or a wrong request.
* Accept a custom error and a http.ResponseWriter, and return nothing.
 */
func RequestError(statusCode int, msg ErrorMsg, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	data, _ := json.Marshal(msg)
	w.Write(data)
}
