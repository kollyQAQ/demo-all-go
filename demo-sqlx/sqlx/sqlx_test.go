package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/11
*/

var db *sqlx.DB

type Place struct {
	id      int
	country string
	city    string
	telcode int
}

func TestMain(m *testing.M) {
	//定义mysql数据源，配置数据库地址，帐号以及密码， dsn格式下面会解释
	dsn := "root:root1234@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True&loc=Local"

	//根据数据源dsn和mysql驱动, 创建数据库对象
	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	//设置连接池最大连接数
	db.SetMaxOpenConns(100)

	//设置连接池最大空闲连接数
	db.SetMaxIdleConns(20)

	m.Run()
}

func TestCreateTable(t *testing.T) {
	//定义表结构
	schema := `CREATE TABLE place (
	id int primary key auto_increment,
	country varchar(50),
	city varchar(50) NULL default '',
	telcode int);`

	// 调用Exec函数执行sql语句，创建表
	_, err := db.Exec(schema)
	//错误处理
	if err != nil {
		panic(err)
	}
}

/*
插入数据
*/
func TestInsertTable(t *testing.T) {
	//定义sql语句, 通过占位符 问号（ ? ) 定义了三个参数
	countryCitySql := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`

	//通过Exec插入数据, 这里传入了三个参数，对应sql语句定义的三个问号所在的位置
	result, err := db.Exec(countryCitySql, "中国", "香港", 852)
	//错误处理
	if err != nil {
		fmt.Println("插入失败!")
	}
	//插入成功后，获取insert id
	id, _ := result.LastInsertId()
	println(id)

	//通过MustExec插入数据, 如果sql语句出错，则直接抛出panic错误
	result2 := db.MustExec(countryCitySql, "South Africa", "Johannesburg", 27)

	//插入成功后，获取插入id
	id2, _ := result2.LastInsertId()
	println(id2)
}

/*
更新数据
*/
func TestUpdateTable(t *testing.T) {
	//定义sql语句，通过问号定义了三个参数
	sql := "update place set telcode=?, city=? where id=?"

	//通过Exec更新数据, 这里传入了三个参数，对应sql语句定义的三个问号所在的位置
	result1, err := db.Exec(sql, 100, "香港", 1)
	//错误处理
	if err != nil {
		fmt.Println("更新失败!")
	}

	//查询更新影响行数
	rowsAffected, _ := result1.RowsAffected()
	println(rowsAffected)
}

/*
通过 Get 和 Select 函数查询数据
Get 函数主要用于查询一条记录，Select 用于查询多条记录。
*/
func TestGetSelect(t *testing.T) {
	//定义保存查询结果的struct变量
	p := Place{}
	// 查询一条记录, 并且往sql语句传入参数 1，替换sql语句中的问号，最后将查询结果保存到struct对象中
	err := db.Get(&p, "SELECT * FROM place LIMIT ?", 1)
	if err != nil {
		panic(err)
	}
	println(&p)

	var total int
	//统计表的总记录数，并将查询结果保存到一个变量中
	err = db.Get(&total, "SELECT count(*) FROM place")
	if err != nil {
		panic(err)
	}
	println(total)

	//定义一个保存多条记录的struct数组变量
	var pp []Place
	// 通过Select查询多条记录，并且将结果保存至pp变量中
	// 这里相当于将一条记录的字段值都映射到struct字段中
	err = db.Select(&pp, "SELECT * FROM place WHERE telcode > ?", 50)
	if err != nil {
		panic(err)
	}
	println(pp)

	var names []string
	// 通过Select查询多条记录，并且将结果保存至names变量中
	// 这里仅查询一个字段
	err = db.Select(&names, "SELECT name FROM place LIMIT 10")
	if err != nil {
		panic(err)
	}
	println(names)
}

// 通过 Queryx 和 QueryRowx 查询数据
// 相对于 Get 和 Select 函数，Queryx 和 QueryRowx 函数要繁琐一些。
// Queryx 可以用于查询多条记录，QueryRowx 函数用于查询一条记录。
func TestQueryx1(t *testing.T) {
	// 查询所有的数据，这里返回的是sqlx.Rows对象
	rows, err := db.Queryx("SELECT country, city, telcode FROM place")
	//错误检测
	if err != nil {
		panic(err)
	}

	// 循环遍历每一行记录，rows.Next()函数用于判断是否还有下一行数据
	for rows.Next() {
		//这里定义三个变量用于接收每一行数据
		var country string
		var city string
		var telcode int

		//调用Scan函数，将当记录的数据保存到变量中，这里参数的顺序跟上面sql语句中select后面的字段顺序一致。
		err = rows.Scan(&country, &city, &telcode)
		println(country, city, telcode)
	}
}

func TestQueryx2(t *testing.T) {
	//定义保存数据的结构体， 默认struct字段名（小写）跟表的字段名一致。
	type Place struct {
		Country string
		//因为city字段允许null，所以这里可以使用sql.NullString类型
		City sql.NullString
		//如果struct字段名跟表的字段名不一样，可以通过db标签设置数据库字段名
		TelephoneCode int `db:"telcode"`
	}

	//查询数据
	rows, err := db.Queryx("SELECT * FROM place")
	if err != nil {
		panic(err)
	}

	//遍历数据
	for rows.Next() {
		//下面演示如何将数据保存到struct、map和数组中
		//定义struct对象
		var p Place

		//定义map类型
		m := make(map[string]interface{})

		//使用StructScan函数将当前记录的数据保存到struct对象中
		err = rows.StructScan(&p)
		//保存到map
		err = rows.MapScan(m)

		//保存到数组
		var s []interface{}
		s, err = rows.SliceScan()
		println(s)
	}
}

func TestQueryRowx(t *testing.T) {
	//查询数据
	row := db.QueryRowx("SELECT country, city, telcode FROM place where id = ?", 1)

	//定义保存数据的结构体， 默认struct字段名（小写）跟表的字段名一致。
	type Place struct {
		Country string
		City    sql.NullString
		Telcode int
	}

	var p Place

	//使用StructScan函数将当前记录的数据保存到struct对象中
	err := row.StructScan(&p)
	if err != nil {
		panic(err)
	}
}

func TestDelete(t *testing.T) {
	//定义sql语句，通过问号定义了一个参数
	sql := "delete from place where id=?"

	//通过Exec删除数据, 这里传入了一个参数，对应sql语句定义的问号所在的位置
	result1, err := db.Exec(sql, 4)

	//获取删除影响行数
	rowsAffected, _ := result1.RowsAffected()

	//错误处理
	if err != nil {
		fmt.Println("更新失败!")
	}

	println(rowsAffected)
}

func TestTx(t *testing.T) {
	//开始一个事务，返回一个事务对象tx
	tx, err := db.Beginx()

	//使用事务对象tx, 执行事务
	//err = tx.Queryx(...)
	//err = tx.Exec(...)
	//err = tx.Exec(...)

	if err != nil {
		//回滚事务
		tx.Rollback()
	}

	_, err1 := tx.Exec("delete from place where id=?", 1)
	_, err2 := tx.Exec("delete from place where id=?", 2)

	if err1 != nil || err2 != nil {
		//回滚事务
		tx.Rollback()
	}

	//提交事务
	err = tx.Commit()
}
