package primary


import (
	"net/http"
	"encoding/json"
	"testing"
	"net/http/httptest"
	"log"
	"io/ioutil"
	"fmt"
	"flysnow.org/hello/common"
)

func Routes(){
	http.HandleFunc("/sendjson",SendJSON)
}
func SendJSON(rw http.ResponseWriter,r *http.Request){
	u := struct {
		Name string
	}{
		Name:"张三",
	}

	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(u)
}
func TestSendJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("创建Request失败")
	}
	rw:=httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw,req)

	log.Println("code:",rw.Code)

	log.Println("body:",rw.Body.String())
}



func mockServer() *httptest.Server {
	//API调用处理函数
	sendJson := func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "张三",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(u)
	}
	//适配器转换
	return httptest.NewServer(http.HandlerFunc(sendJson))
}
func TestSendJSON1(t *testing.T) {
	//创建一个模拟的服务器
	server := mockServer()
	defer server.Close()
	//Get请求发往模拟服务器的地址
	resq, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("创建Get失败")
	}
	defer resq.Body.Close()

	log.Println("code:", resq.StatusCode)
	json, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("body:%s\n", json)
}


func Tag(tag int){
	switch tag {
	case 1:
		fmt.Println("Android")
	case 2:
		fmt.Println("Go")
	case 3:
		fmt.Println("Java")
	default:
		fmt.Println("C")


	}
}








