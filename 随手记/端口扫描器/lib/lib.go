package lib

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Usage(str string) {
	fmt.Fprintf(os.Stderr, str)
	flag.PrintDefaults()
}

func Mkdir(path string) {
	f, err := os.Stat(path)
	if err != nil || f.IsDir() == false {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			fmt.Println("创建目录失败！", err)
			return
		}
	}
}

func Contrast(newports []int, fileName string, ip string) {
	if exists(fileName) {
		str := Openlog(fileName)
		reg := regexp.MustCompile(`\d+`)
		portslist := reg.FindAllStringSubmatch(str, -1)
		oldports := make([]int, len(portslist))
		for i := range oldports {
			outdata, _ := strconv.Atoi(portslist[i][0])
			oldports[i] = outdata
		}
		new2old(newports, oldports, ip)
	}

}

func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Openlog(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(content))
	return string(content)
}

func new2old(newports []int, oldports []int, ip string) {
	fmt.Println(newports)
	fmt.Println(oldports)
	for i := range newports {
		if InIntSlice(oldports, newports[i]) {
			fmt.Println("相同的端口")
		} else {
			fmt.Println("开启端口：", newports[i])
			msg := fmt.Sprintf("%s开启端口：%s", ip, strconv.Itoa(newports[i]))
			sendmsg(msg)
		}
	}
	for i := range oldports {
		if InIntSlice(newports, oldports[i]) {
			fmt.Println("相同的端口")
		} else {
			fmt.Println("关闭端口：", oldports[i])
			msg := fmt.Sprintf("%s关闭端口：%s", ip, strconv.Itoa(oldports[i]))
			sendmsg(msg)
		}
	}

}

func InIntSlice(haystack []int, needle int) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}

func sendmsg(msg string) {
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=beddb4248f9c93ec537676544265d49f23a9d773f6d55b6689ff387c334592c1"
	content := `{"msgtype": "text",
      	"text": {"content": "` + msg + `"},
                "at": {
                     "atMobiles": [
                         "18126473169"
                     ],
                     "isAtAll": false
                }
		}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println("handle error")
	}
}
