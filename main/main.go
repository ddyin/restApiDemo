package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"GoDemo/rest"
)

func main()  {
	fmt.Println("999")

	//db,err := sql.Open("mysql","用户名:密码@tcp(127.0.0.1:3306)/ddyin")

	//if err !=nil{
	//	fmt.Println(err)
	//}
	fmt.Println("operation start...")
	//fmt.Println(db)
	//xorm.Test()

	//普通的数据库操作(类似java的jdbc)
	//insert(db)
	//selectData(db)
	//updateData(db)
	//deleteData(db)

	//xorm框架
	//xorm.ConnectionAndOperation()

	//restful API设计
	//服务端启动
	//http.HandleFunc("/",rest.RestfulHandler)
	//http.ListenAndServe("localhost:8080",nil)
	//客户端调用
	rest.HttpTest()

	fmt.Println("operation end...")
}

//新增
func insert(db *sql.DB)  {
	stmt,err := db.Prepare("insert into hoteldict(id,HotelName) values (?,?)")

	defer stmt.Close()

	if err!=nil {
		fmt.Println(err)
		return
	}
	stmt.Exec(1,"儒家")
	stmt.Exec(2,"道家")
}

//查询
func selectData(db *sql.DB)  {

	rows,err := db.Query("select * from hoteldict")
	defer rows.Close()
	if err != nil{
		log.Print(err)
		return
	}

	for rows.Next() {
		var id int
		var HotelName string

		rows.Columns()
		err = rows.Scan(&id,&HotelName)

		fmt.Println("序号",id)
		fmt.Println("名称",HotelName)
	}

	fmt.Println(rows)
}

//修改
func updateData(db *sql.DB)  {

	stmt,err := db.Prepare("update hoteldict set HotelName = ? where id =?")
	defer stmt.Close()

	if err!=nil {
		log.Println(err)
		return
	}

	res,err := stmt.Exec("哲学家1",1)
	if err!=nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

}

//删除
func deleteData(db *sql.DB)  {

	stmt,err := db.Prepare("delete from hoteldict where id =?")

	defer stmt.Close()
	if err!=nil {
		log.Println(err)
		return
	}

	res,err := stmt.Exec(1)

	num,err := res.RowsAffected()
	fmt.Println(num)
}

