package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/11
*/

func main() {
	// 连接 DB
	// sql.Open的第一个参数是 driver 名称，第二个参数是 driver 连接数据库的信息，各个 driver 可能不同
	db, err := sql.Open("mysql",
		"root:root1234@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// DB 不是连接，并且只有当需要使用时才会创建连接，如果想立即验证连接，需要用Ping()方法，如下：
	err = db.Ping()
	if err != nil {
		// do something here
	}
	// sql.DB 的设计就是用来作为长连接使用的。不要频繁 Open, Close。比较好的做法是，为每个不同的 datastore 建一个 DB 对象，
	// 保持这些对象 Open。如果需要短连接，那么把 DB 作为参数传入 function，而不要在 function 中 Open, Close。

	// 读取 DB
	// 如果方法包含Query，那么这个方法是用于查询并返回 rows 的。其他情况应该用Exec()。
	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// 单行 Query
	// err 在Scan后才产生，所以可以如下写：
	var name2 string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	// 修改数据
	// 一般用 Prepared Statements 和Exec()完成INSERT, UPDATE, DELETE操作。
	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Dolly")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	// 事务
	// db.Begin()开始事务，Commit() 或 Rollback()关闭事务。Tx从连接池中取出一个连接，在关闭之前都是使用这个连接。Tx 不能和 DB 层的BEGIN, COMMIT混合使用。

	// 在 Transaction 中使用 PS
	// PS 在 Tx 中唯一绑定一个连接，不会 re-prepare。
	// Tx 和 statement 不能分离，在 DB 中创建的 statement 也不能在 Tx 中使用，因为他们必定不是使用同一个连接
	// 使用 Tx 必须十分小心，例如下面的代码：
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt2, err := tx.Prepare("INSERT INTO users(name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt2.Close() // danger!
	for i := 0; i < 10; i++ {
		_, err = stmt2.Exec(i)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	// stmt.Close() runs here!
	// *sql.Tx一旦释放，连接就回到连接池中，这里 stmt 在关闭时就无法找到连接。所以必须在 Tx commit 或 rollback 之前关闭 statement。
}
