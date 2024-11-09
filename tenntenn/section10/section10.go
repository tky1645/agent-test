package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

type testHandler struct{

}

func (t testHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "this is handler")

}
func main() {
	// http.Handlerを継承したtestHandlerを作ってHandle関数にセット
	//  http.Handlerを継承するにはServeHTTPの実装が必要
	var testHandler testHandler 
	http.Handle("/handle", testHandler)
	
	//  http.Handler継承したオブジェクトの定義をすっ飛ばしてServeHTTPに相当する関数を引数にとる
	// 内部的にhttp.Handlerをとして登録している
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/head", headerHandler)
	http.HandleFunc("/tmpl", templateEngineHandler)
	http.ListenAndServe(":3000", nil)
}

func init(){
	// rand.Seed(time.Now().UnixNano())

}

func handler(w http.ResponseWriter, r *http.Request) {

	num:= rand.Intn(3)

	switch num {
	case 1:
		fmt.Fprint(w,"吉")
	case 2:
		fmt.Fprint(w,"大吉")
	default:
		fmt.Fprint(w,"凶")
	}
	fmt.Println("handle!")
	fmt.Fprintln(w, "hello server")
}

type Person struct{
	Name string `json:"name"`
	Age int`json:"age"`

}
func jsonHandler(w http.ResponseWriter, r *http.Request){
	person := &Person{
		Name: "name",
		Age: 1,
	}

	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	if err:= 	encoder.Encode(person); err != nil {
		fmt.Fprintln(w, "fail")

	}

	fmt.Fprintln(w, "done")
	fmt.Fprint(w, buff.String())

}
type response struct{
	Param string `json:"message"`
	Body Person `json:"body"`
	Header string	`json:"header"`
}
func headerHandler(w http.ResponseWriter, r *http.Request){
	var p Person
	// パラメタ
	param := r.FormValue("msg")
	// ヘッダ
	header := r.Header.Get("Content-Type")
	// ボディ
	dec := 	json.NewDecoder(r.Body)
	dec.Decode(&p)

	w.Header().Set("Content-Type", "application/json; charset=ytf-8")
	v := response{
		Param: param,
		Body: p,
		Header: header,
	}
	if err := json.NewEncoder(w).Encode(v); err != nil{
		fmt.Println("fail")
	}
}

func templateEngineHandler(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.New("sign").
		Parse("<html><body>{{.}}</html></body>"))	

	tmpl.Execute(w,r.FormValue("content"))
}