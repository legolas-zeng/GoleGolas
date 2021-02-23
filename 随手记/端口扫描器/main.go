package main

import (
	"GoleGolas/随手记/端口扫描器/lib"
	"GoleGolas/随手记/端口扫描器/scan"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	startTime = time.Now()
	timeout   = int(200)
	process   = int(1000)
	path      = string("log")
	ip        = string("192.168.3.5")
	port      = string("1-65535")
	//ip        = flag.String("ip", "192.168.3.5", "ip地址 例如:-ip=192.168.0.1-255 或直接输入域名 xs25.cn")
	//port      = flag.String("p", "1-65535", "端口号范围 例如:-p=80,81,88-1000")
	//path      = flag.String("path", "log", "日志地址 例如:-path=log")
	//timeout   = flag.Int("t", 200, "超时时长(毫秒) 例如:-t=200")
	//process   = flag.Int("n", 1000, "进程数 例如:-n=10")
	//h         = flag.Bool("h", false, "帮助信息")
)

//go run client.go -h
func main() {
	//flag.Parse()
	//帮助信息
	//if *h == true {
	//	lib.Usage("scanPort version: scanPort/1.10.0\n Usage: scanPort [-h] [-ip ip地址] [-n 进程数] [-p 端口号范围] [-t 超时时长] [-path 日志保存路径]\n\nOptions:\n")
	//	return
	//}
	fi, _ := os.Open("log/host.txt")
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		ip = string(a)
		fmt.Printf("========== Start %v ip:%v,port:%v ==================== \n", time.Now().Format("2006-01-02 15:04:05"), ip, port)

		//创建目录
		lib.Mkdir(path)

		//初始化
		scanIP := scan.ScanIp{
			Debug:   true,
			Timeout: timeout,
			Process: process,
		}
		//扫所有的ip
		fileName := path + "/" + ip + "_port.txt"
		//for i := 0; i < len(ips); i++ {
		ports := scanIP.GetIpOpenPort(ip, port)
		lib.Contrast(ports, fileName, ip)
		if len(ports) > 0 {
			f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
			if err != nil {
				fmt.Println(err)
			}
			var str = fmt.Sprintf("%v \n", ports)
			if _, err := f.WriteString(str); err != nil {
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		//}
		fmt.Printf("========== End %v 总执行时长：%.2fs ================ \n", time.Now().Format("2006-01-02 15:04:05"), time.Since(startTime).Seconds())

	}

}
