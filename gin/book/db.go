package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(192.168.2.32:3306)/go?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	//最大连接
	db.SetMaxOpenConns(20)
	//最大空闲
	db.SetMaxIdleConns(10)
	return
}

func querAllBook()(bookList []*Book,err error)  {
	sqlStr := "select * from book"
	err = db.Select(&bookList,sqlStr)
	if err != nil {
		fmt.Println("失败")
		return
	}
	return
}

func insertBook(title string,price int64)(err error)  {
	sqlStr := "insert into book(title,price) value (?,?)"
	_, err = db.Exec(sqlStr,title,price)
	if err != nil {
		fmt.Println("失败")
		return
	}
	return
}

func deleteBook(id int64)(err error)  {
	sqlStr := "delete from book when id = ?"
	_, err = db.Exec(sqlStr,id)
	if err != nil {
		fmt.Println("失败")
		return
	}
	return
}