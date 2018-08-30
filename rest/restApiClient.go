package rest

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
)

//客户端调用
func HttpTest()  {
	url := "http://localhost:8080"
	rst,err := http.Get(url)
	if err!=nil {
		fmt.Println(err)
	}

	defer rst.Body.Close()

	body,err := ioutil.ReadAll(rst.Body)

	var data Data
	err = json.Unmarshal(body,&data)
	timeResult := time.Unix(data.Time,0).Format("2018-08-30 15:02:02")

	fmt.Println("username:",data.UserName)
	fmt.Println("password:",data.PassWord)
	fmt.Println("time:",timeResult,"\n",data.DetailMsg)
	fmt.Println("detailMsg:",data.DetailMsg)
}















