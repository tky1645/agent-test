package main

import (
	"fmt"
	"log"
	"net/http"
	"sqlboiler-test/slices/utility"
	"time"
)
func main (){
	utility.DB.SetConnMaxLifetime(0)

	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/create", createHandler)

	if err := http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal("fail to start listen ")
	}
	fmt.Println("start listen")
}


func readHandler(w http.ResponseWriter, r *http.Request){

}

func createHandler(w http.ResponseWriter, r *http.Request){
	post_fix :=  time.Now().Format("20060102 15:04:05")
	title := fmt.Sprintf("test_title_%s", post_fix)
	body := fmt.Sprintf("test_body_%s", post_fix)
	
	insert, err :=  utility.DB.Prepare("INSERT INTO Article(title, body) VALUES(?,?)")
	if err !=nil{
		log.Fatal(err.Error())
	}
	insert.Exec(title, body)
}