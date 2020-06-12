#### 项目说明
#### k8s插件
kubernetes的命令很长，通过这个小命令快速输入常用的命令。

编译可执行文件：
```
#在windows下
set GOARCH=amd64
set GOOS=linux
go build kubectl-fcom.go

#在linux下
go build kubectl-fcom.go
```

将编译的可执行文件拷贝到可执行路径中
```
[root@k8s-master-2 ceph-demo]# ll /etc/kubernetes/bin/kubectl-fcom 
-rwxr-xr-x 1 root root 4483024 4月  28 10:07 /etc/kubernetes/bin/kubectl-fcom
```
使用：

![image](https://s1.ax1x.com/2020/06/12/tXSmCj.png)
