package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.New("index.html").ParseFiles("templates/index.html"))
	var data struct {
		A1 string
	}
	data.A1 = "'aaa'"
	if err := temp.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

type ResData struct {
	State int  `json:"state"`
	Data  User `json:"data"`
}

type User struct {
	ChannelId string `json:"channelId"`
	OpenId    string `json:"openId"`
	OpenName  string `json:"openName"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	user = User{ChannelId: "iapp", OpenId: "abcdefg", OpenName: "aaa"}
	var resdata ResData
	resdata = ResData{State: 0, Data: user}
	b, err := json.Marshal(resdata)
	if err != nil {
		fmt.Println("json error", err)
	}

	fmt.Fprint(w, string(b))
}
