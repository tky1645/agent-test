package utility

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	user:="root"
	pw := "root"
	db_name:= "test"

	path := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, pw, db_name)
	var err error
	if DB, err = sql.Open("mysql", path); err !=nil{
		log.Fatal(err.Error())
		return
	}

	dberr := checkconnect(10)
	if dberr !=nil{
		log.Println(dberr.Error())
		return
	}

	fmt.Println("db connected!!")
}
func checkconnect(count uint)error{
	if count == 0{
		return nil
	}
	if err := DB.Ping(); err!=nil{
		time.Sleep(time.Second *2)
		count --
		fmt.Printf("retry...count:%v\n",count)
		checkconnect(count)
	}else{
		return nil
	}
	log.Fatal("failed to connect")
	return errors.New("")
}