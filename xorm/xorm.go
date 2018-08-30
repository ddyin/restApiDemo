package xorm

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"github.com/go-xorm/core"
)

func ConnectionAndOperation()  {
	//创建orm引擎
	engine,err := xorm.NewEngine("mysql","用户名:密码@tcp(127.0.0.1:3306)/ddyin")

	if err!=nil {
		fmt.Println(err)
		return
	}

	//连接测试
	if err := engine.Ping();err!=nil{
		fmt.Println(err)
		return
	}

	//日志打印sql
	engine.ShowSQL(true)

	//设置连接池空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)

	//名称映射规则主要是负责结构体到表名和表中字段名称的映射
	engine.SetTableMapper(core.SnakeMapper{})

	//addData(*engine)
	//delData(*engine)
	//updateData(*engine)
	selectData(*engine)

}
//新增
func addData(engine xorm.Engine)  {
	hoteldict := new(HotelDict)
	hoteldict.HotelName="链家777"
	affected,err := engine.Insert(hoteldict)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}

//删除
func delData(engine xorm.Engine)  {
	hoteldict := new(HotelDict)
	hoteldict.Id=4
	affected,err := engine.Delete(hoteldict)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}

//修改
func updateData(engine xorm.Engine)  {
	hoteldict := new(HotelDict)
	hoteldict.HotelName="ctrip"
	affected,err := engine.Id(2).Update(hoteldict)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}

//查询
func selectData(engine xorm.Engine)  {
	hoteldict := new(HotelDict)
	//判断数据是否存在，存在返回true，不存在返回false
	//res,err := engine.Id(6).Get(hoteldict)
	res,err := engine.Where("id=?",3).Get(hoteldict)

	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println(hoteldict)
	fmt.Println(hoteldict.HotelName)
}

type HotelDict struct {
	Id int `xorm:"not null pk autoincr INT(11)"`
	HotelName string `xorm:"not null VARCHAR(200)"`
}



