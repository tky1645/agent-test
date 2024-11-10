package main

import (
	"errors"
	"fmt"
	"net/http"
)

func refreshAccount()error {
	// DBにテーブル作成
	const sqlCreate = `
	CREATE TABLE IF NOT EXISTS account (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	amount integer NOT NULL)
	;
	`
	_,err := db.Exec(sqlCreate)
	if err != nil{
		fmt.Printf("failed to create table")
		return errors.New("failed to create table")
	}

	const insertAccountData = `
	INSERT OR IGNORE INTO account (name, amount) VALUES ('hoge', 1000), ('fuga', 2000);`
	_, err = db.Exec(insertAccountData)
	if err != nil {
		return errors.New("failed to insert account data")
	}
	return nil
}

func transportHandler(w http.ResponseWriter, r *http.Request) {
	refreshAccount()
	fmt.Fprintln(w, "this is handler")
	// トランザクションを張る
	tx, err :=db.Begin()
	if err != nil{
		fmt.Printf("failed to begin transaction")
		return
	}

	sqlIncrease :=`
	UPDATE account SET amount = amount +  ? WHERE name = ?;	 `
	if _, err := tx.Exec(sqlIncrease, "+100", "hoge");err != nil{
		tx.Rollback()
		fmt.Printf("failed to update")
		return
	}

	
	if _,err:=tx.Exec(sqlIncrease, "-100", "fuga"); err != nil{
		tx.Rollback()
		fmt.Printf("failed to update")
		return
	}

	// トランザクションのコミット
	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}
	
	fmt.Fprintln(w, "success to update")

}