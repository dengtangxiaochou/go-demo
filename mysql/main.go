package	main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
// 定义一个全局对象db
var db *sql.DB
var db1 *sqlx.DB

// 定义一个初始化数据库的函数
//func initDB() (err error) {
//	// DSN:Data Source Name
//	dsn := "root:root@tcp(192.168.2.32:3306)/go?charset=utf8mb4&parseTime=True"
//	// 不会校验账号密码是否正确
//	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	// 尝试与数据库建立连接（校验dsn是否正确）
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	//设置数据库连接池的最大值
//	db.SetMaxOpenConns(30)
//	//设置最大空闲连接数
//	db.SetMaxIdleConns(20)
//	return nil
//}
//初始化连接数据库用salx进行
func initDB() (err error) {
	dsn := "root:root@tcp(192.168.2.32:3306)/go?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db1, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db1.SetMaxOpenConns(20)
	db1.SetMaxIdleConns(10)
	return
}
//
type user struct {
	Id int
	Name string
	Age int
}
//1.查询单条数据
func querOne(id int)  {
	var u1 user
	//1.查询单条数据
	sqlStr := `select id ,name,age from user where id=?`
	rowObj := db.QueryRow(sqlStr,id)
	//测试最大连接数
	//for i := 0; i <10 ; i++ {
	//	db.QueryRow(sqlStr,1)
	//}
	rowObj.Scan(&u1.Id,&u1.Name,&u1.Age)
	fmt.Println(u1)
}

func queryMultiRowDemo(n int)  {
	sqlStr := `select id ,name,age from user where id > ?`
	rows, err := db.Query(sqlStr,n)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	//4.循环取值

	for rows.Next() {
		var u  user
		err := rows.Scan(&u.Id,&u.Name,&u.Age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u)
		return
	}
}

// 插入数据
func insertRowDemo(b string,a int) {
	sqlStr := `insert into user(name,age) values (?,?)`
	ret, err := db.Exec(sqlStr,b,a)
	if  err != nil{
		fmt.Println(err)
		return
	}
	//插入数据的操作会有id值
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)

}
// 更新数据
func updateRowDemo(x int,y int) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, x, y)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo(p int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, p)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"哈哈":30,
		"黑河":23,
		"拉拉":34,
	}
	for k ,v := range m {
		_, err = stmt.Exec(k, v)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
		fmt.Println("insert success.")
	}
}

// 事务操作示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 1)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}

// 查询多条数据示例
func queryMultiRowDemo1() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db1.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for _,v := range users {
		fmt.Println(v)
	}
	fmt.Printf("users:%#v\n", users)
}

func main()  {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("成功")

	//querOne(2)
	//queryMultiRowDemo(2)
	//insertRowDemo("小屋",33)
	//updateRowDemo(38,1)
	//deleteRowDemo(2)
	//prepareInsertDemo()
	//transactionDemo()
	queryMultiRowDemo1()
}