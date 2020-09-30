# Linux 查看端口使用情况

使用命令：netstat -aptn

查看所有的进程和端口使用情况，其中最后一栏是PID/Program name
![](/imags/1.PNG)

关闭某个端口使用命令：kill -9 pid

启动虚拟机etcd指令
```
nohup /vagrant/go/bin/etcd  --data-dir=/data/etcd/node1/logs --name=node1 --listen-peer-urls=http://192.168.110.31:2379 --initial-advertise-peer-urls=http://192.168.110.31:2379 --listen-client-urls=http://192.168.110.31:2479 --advertise-client-urls=http://192.168.110.31:2479 --initial-cluster-token=etcd-cluster --initial-cluster=node1=http://192.168.110.31:2379 --initial-cluster-state=new --auto-compaction-retention=1 --max-request-bytes=33554432 --quota-backend-bytes=335544320 &
```