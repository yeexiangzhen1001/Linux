# Linux 查看端口使用情况

使用命令：netstat -aptn

查看所有的进程和端口使用情况，其中最后一栏是PID/Program name
![](/imags/1.PNG)

关闭某个端口使用命令：kill -9 pid

启动虚拟机etcd指令
```
nohup /vagrant/go/bin/etcd  --data-dir=/data/etcd/node1/logs --name=node1 --listen-peer-urls=http://192.168.110.31:2379 --initial-advertise-peer-urls=http://192.168.110.31:2379 --listen-client-urls=http://192.168.110.31:2479 --advertise-client-urls=http://192.168.110.31:2479 --initial-cluster-token=etcd-cluster --initial-cluster=node1=http://192.168.110.31:2379 --initial-cluster-state=new --auto-compaction-retention=1 --max-request-bytes=33554432 --quota-backend-bytes=335544320 &
```

git 指令
```
git branch -a //查看所有分支
git branch origin :分支名  //删除远程分支
git push origin --delete [branchname] //删除远程分支
git branch -D 分支名 //删除本地分支
```

git 提交代码到远程分支
```
git add . //提交全部文件
git add  工程名的下一级开始写路径直到文件名  //提交单个文件
git commit -m "文字描述"（单引号和双引号都可以） //执行commit提交
git push origin feature/debugservice20200930:feature/debugservice20200930 //本地分支上传到远程分支上
```
git status检测不到文件变化
```
1. 首先, 提交全部更新

2. 执行 git rm -r --cached . //从 index 内删除所有变更过的文件

3. 执行 git add .            // 添加所有文件

4. 执行 git commit -m " "

5. 执行 git push origin feature/debugservice20200930:feature/debugservice20200930 //本地分支上传到远程分支上
```
