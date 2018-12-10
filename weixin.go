package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	//	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"
	"time"
)

//这几个URL的Cookies都打印一下

func GetQR() (QR string) {

	url := "https://login.wx.qq.com/jslogin?appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_CN&_=1544413228164"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Cookies())  //打印Cookies
	aaa := string(body)
	bbb := strings.Split(aaa, ";")
	ccc := strings.Fields(bbb[1])
	ddd := ccc[2]
	eee := ddd[1 : len(ddd)-1]
	//fmt.Println(eee)
	return eee

}

func Getlogin() (login string) {
	aaa := GetQR()
	url01 := fmt.Sprintf("https://login.weixin.qq.com/qrcode/%s", aaa)
	fmt.Println(url01)
	resp, err := http.Get(url01)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Cookies())   //打印Cookies
	file, _ := os.Create(aaa + ".jpg")
	io.Copy(file, resp.Body)

	for {
		url02 := fmt.Sprintf("https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=%s&tip=0&r=1870135597&_=1544318065401", aaa)
		time.Sleep(1)
		resp01, err := http.Get(url02)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp01.Cookies())    //打印Cookies
		body, _ := ioutil.ReadAll(resp01.Body)
		bbb := string(body)
		//	fmt.Println(string(body))
		ccc := strings.Split(bbb, ";")[0]
		ddd := strings.Split(ccc, "=")[1]
		eee, _ := strconv.Atoi(ddd)
		if eee == 200 {
			fff := strings.Split(bbb, ";")[1]
			ggg := strings.Split(fff, "\"")[1]
			return ggg
		}
	}
}

//搞
func Getxml() {
	url := Getlogin()
	fmt.Println(url)
	resp01, _ := http.Get(url)

	file, _ := os.Create("login.xml")
	io.Copy(file, resp01.Body)

	// 	req, _ := http.NewRequest("GET", url, nil)

	// 	res, _ := http.DefaultClient.Do(req)

	// 	defer res.Body.Close()
	// 	body, _ := ioutil.ReadAll(res.Body)

	// 	fmt.Println(res)
	// //	fmt.Println(string(body))
	fmt.Println(resp01.Cookies())  //打印Cookies
}

func main() {
	Getxml()
}
