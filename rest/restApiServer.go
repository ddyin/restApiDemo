package rest

import (
	"time"
	"encoding/json"
	"net/http"
	"fmt"
)

//获取json数据
func getJsonData()  ([]byte,error){
	rst1 := make(map[string]int)
	rst1["exp1"]=99
	rst1["exp2"]=88
	resp1 := Response{200,rst1}

	rst2 := make(map[string]int)
	rst2["exp1"]=77
	rst2["exp2"]=66
	resp2 := Response{200,rst2}

	detail := []Response{resp1,resp2}

	msg := Data{"ddyin","xxx",time.Now().Unix(),detail}

	return json.MarshalIndent(msg,"","")
}

type Response struct {
	//返回码
	ReturnCode int
	//返回信息
	Result map[string]int
}

type Data struct {
	UserName string
	PassWord string
	Time int64
	DetailMsg []Response
}

func RestfulHandler(w http.ResponseWriter,r *http.Request)  {
	resp,err :=getJsonData()
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w,string(resp))
}











