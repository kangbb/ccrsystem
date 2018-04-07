package controller

import (
	"encoding/json"
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kangbb/todolist/core/models/service"
)

func dataProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postData(w, r)
	} else if r.Method == "GET" {
		getData(w)
	} else {
		http.Error(w, "An Error taken place!", 501)
	}
}

//处理Post数据
//这里的数据要和表格提交验证区分开
//数据格式 {method:add/update/delete, label: string, isFinished: bool, time: string }
func postData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	method := js.Get("method").MustString()
	label := js.Get("Label").MustString()
	isFinished := js.Get("IsFinished").MustBool()
	time := js.Get("CreateAt").MustString()
	if method == "add" {
		addData(w, label, isFinished, time)
	} else if method == "update" {
		updateData(w, label, isFinished, time)
	} else {
		deleteData(w, label, isFinished, time)
	}
}

func getData(w http.ResponseWriter) {
	data := service.GetAllItems()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	b, _ := json.Marshal(data)
	w.Write(b)
}

/*----------------------------------------------*/
func addData(w http.ResponseWriter, label string, isFinished bool, time string) {
	_, err := service.CreateItem(label, isFinished, time)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("unknown wrong!"))
	}
	w.WriteHeader(200)
}
func updateData(w http.ResponseWriter, label string, isFinished bool, time string) {
	_, err := service.UpdateItem(label, isFinished, time)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("unknown wrong!"))
	}
	w.WriteHeader(200)
}
func deleteData(w http.ResponseWriter, label string, isFinished bool, time string) {
	_, err := service.DeleteItem(label, isFinished, time)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("unknown wrong!"))
	}
	w.WriteHeader(200)
}
