package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB
func main() {
	db =NewDb()
	if db == nil{
		fmt.Printf("no db")
		return
	}
	defer db.Close()
	// サーバ起動
	http.HandleFunc("/", handler)
	http.HandleFunc("/insert", transportHandler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"welcome to my server/n")


	// パラメータから各種情報を取り出し
	name := r.FormValue("name")
	tel := r.FormValue("tel")
	if name == "" || tel == ""{
		fmt.Fprintf(w,"no param")	
	}
	// 構造体にセット
	user := struct{
		Name string
		Tel string
	}{
		Name: name,
		Tel: tel,
	}
	fmt.Fprintf(w,"name:%s,tel:%s",user.Name,user.Tel)
	// DBにインサート
	sql := `INSERT INTO usr(name,tel) VALUES(?,?)`
	_, err := db.Exec(sql,user.Name,user.Tel)
	if err != nil{
		fmt.Fprintf(w,"fail to insert, %s", err.Error())

		return
	}
	getAll(db)
}

func NewDb()*sql.DB{
	// dbコネクションを作成
	db, err :=sql.Open("sqlite","sqlite.db")
	if err != nil{
		fmt.Printf("failed to open db")
		return nil
	}

	// テーブル作成
	const sql = `
	CREATE TABLE IF NOT EXISTS usr (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		tel INTEGER NOT NULL
	);
	`
	if _,err :=db.Exec(sql);err !=nil{
		fmt.Printf("failed to create table")
		return nil
	}

	return db
}

func getAll(db *sql.DB){
	// db全ての情報を取得
	sql := `SELECT * FROM usr`
	rows, err :=db.Query(sql)
	if err != nil{
		fmt.Printf("failed to get all")
		return 
	}
	defer rows.Close()

	for rows.Next(){
		var id int
		var name string
		var tel string
		if rows.Scan(&id,&name,&tel); err != nil{
			fmt.Printf("failed to scan")
			return 
		}
		fmt.Printf("id:%d,name:%s,tel:%s",id,name,tel)
		fmt.Println()

	}
}