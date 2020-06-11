### kafka

```
start kafka:
    bin/zookeeper-server-start.sh config/zookeeper.properties
    bin/kafka-server-start.sh config/server.properties

list all topic:
    bin/kafka-topics.sh --list --zookeeper localhost:2181

list all groups:
    bin/kafka-consumer-groups.sh --all-groups --bootstrap-server localhost:9092 --describe

delete group:
    bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --delete --group group1

delete topic:
    bin/kafka-topics.sh  --delete --zookeeper localhost:2181  --topic topic1 // config: delete.topic.enable=true
```

### openssl

```
openssl genrsa -out server.key 2048 // use 'openssl rsa -in server.key - pubout -out server.key.public' export public key
openssl req -new -x509 -key server.key -out server.crt -days 365
```

### mysql

```
mysql> select * from a;
+------+------+
| c1   | c2   |
+------+------+
|    1 |    2 |
|    3 |    4 |
|    5 |    6 |
+------+------+
3 rows in set (0.00 sec)

mysql> select * from b;
+------+------+
| c3   | c4   |
+------+------+
|    3 |   33 |
|    5 |   55 |
|    7 |    8 |
|    9 |   10 |
+------+------+
4 rows in set (0.00 sec)

mysql> select * from a inner join b on a.c1=b.c3;
+------+------+------+------+
| c1   | c2   | c3   | c4   |
+------+------+------+------+
|    3 |    4 |    3 |   33 |
|    5 |    6 |    5 |   55 |
+------+------+------+------+
2 rows in set (0.00 sec)

mysql> select * from a left join b on a.c1=b.c3;
+------+------+------+------+
| c1   | c2   | c3   | c4   |
+------+------+------+------+
|    3 |    4 |    3 |   33 |
|    5 |    6 |    5 |   55 |
|    1 |    2 | NULL | NULL |
+------+------+------+------+
3 rows in set (0.00 sec)

mysql> select * from a right join b on a.c1=b.c3;
+------+------+------+------+
| c1   | c2   | c3   | c4   |
+------+------+------+------+
|    3 |    4 |    3 |   33 |
|    5 |    6 |    5 |   55 |
| NULL | NULL |    7 |    8 |
| NULL | NULL |    9 |   10 |
+------+------+------+------+
4 rows in set (0.00 sec)

mysql> select * from a left join b on a.c1=b.c3 union select * from a right join b on a.c1=b.c3;
+------+------+------+------+
| c1   | c2   | c3   | c4   |
+------+------+------+------+
|    3 |    4 |    3 |   33 |
|    5 |    6 |    5 |   55 |
|    1 |    2 | NULL | NULL |
| NULL | NULL |    7 |    8 |
| NULL | NULL |    9 |   10 |
+------+------+------+------+
5 rows in set (0.00 sec)

mysql> select * from a, b;
+------+------+------+------+
| c1   | c2   | c3   | c4   |
+------+------+------+------+
|    1 |    2 |    3 |   33 |
|    3 |    4 |    3 |   33 |
|    5 |    6 |    3 |   33 |
|    1 |    2 |    5 |   55 |
|    3 |    4 |    5 |   55 |
|    5 |    6 |    5 |   55 |
|    1 |    2 |    7 |    8 |
|    3 |    4 |    7 |    8 |
|    5 |    6 |    7 |    8 |
|    1 |    2 |    9 |   10 |
|    3 |    4 |    9 |   10 |
|    5 |    6 |    9 |   10 |
+------+------+------+------+
12 rows in set (0.00 sec)


after update or insert, 'RowsAffected' and 'LastInsertId' can only be used as a reference 

```

### docker cmds

```
docker run --name ubuntu1 -p 10022:22 -d ubuntu /bin/bash -c "while true; service ssh start; do sleep 3; done"
docker exec -t -i ubuntu1 /bin/bash
docker commit ubuntu1 foobar/myubuntu:v1.0

docker run --name mysql1 -p 3308:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:latest --character- set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci --default-authentication-plugin=mysql_native_password
docker run --net=host --name redis1 -p 6379:6379 -d redis:latest redis-server --appendonly yes
docker run -d --name mongodb1 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -p 27017:27017 mongo:3.6.6
```

### etcd
```
sudo docker exec -e "ETCDCTL_API=3" etcd1 etcdctl get "" --prefix
sudo docker exec -e "ETCDCTL_API=3" etcd1 etcdctl lease list
export "ETCDCTL_API=3" && ./etcdctl -w table endpoint --cluster health
./etcdctl -w table endpoint --cluster status
./etcdctl defrag --cluster
./etcdctl alarm list
./etcdctl alarm disarm
./etcd --auto-compaction-retention '10m' --quota-backend-bytes '8589934592' --max-request-bytes ‘1572864‘
```

### linux cmds

```
// 0 stdin, 1 stdout, 2 stderr, /dev/null
nc -c -k -l 1234
nc -c -k localhost 1234
df -h
du -h --max-depth=0 . (mac: du -h -d=0 .)
```

### golang

```
go tool pprof http://localhost:3000/debug/pprof/profile?seconds=30

/foo1/ -> [/foo1 (301) - > /foo1/, /foo1/]

base type: bool, string, int, uint, float, complex, uintptr, unsafe.Pointer, array, slice, chan, func, interface, map, pointer, struct

default named type: bool, string, int, uint, float, complex, uintptr, unsafe.Pointer

use 'type' let unnamed type can be named type

if method reciver is not pointer, compiler can auto build a pointer version

go install -i -v -ldflags "-X 'main.Version=123' -X 'main.Version1=`date`'" test2

Print(r.RequestURI)   // /hello?a=b&a=c1
Print(r.URL.Path)     // /hello
Print(r.URL.RawQuery) // a=b&a=c1

export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GOPROXY=https://goproxy.io
go build -mod=vendor

sudo rm -rf /home/foobar1/.cache/go-build /home/foobar1/work/gopath1/pkg/linux_amd64
```

### linux signal

```
'9 SIGKILL' and '15 SIGTERM' can not be catch, and ignore
```

### time zone

```
<- (minus time)
-> (add time)
```

### docker ubuntu

```
apt-get update
unminimize
apt-get install g++
apt-get install ssh
apt-get install net-tools
```

edit '/etc/ssh/sshd_config' and edit'PermitRootLogin yes'

'ClientAliveInterval 60' server send client info in 60 seconds

'ClientAliveCountMax 3' send client, not replay 3 times, close it

### vim config

```
set nu
set nobackup
set shiftwidth=2
set tabstop=2
set autoindent
set cindent
filetype plugin indent on
if has("autocmd")
  au BufReadPost * if line("'\"") > 1 && line("'\"") <= line("$") | exe "normal! g'\"" | endif
endif
```

'shift + k' open man with word

'12 shift + g' jump line

'gg' jump first line

'shift + g' jump last line

'2 dd' delete 2 row

'x' delete a char

'r' replace a char

'yy' copy a row

'2 yy' copy 2 row

'v' inter visual mode


### git config

```
git config --global user.name qshe
git config --global user.email qs-he@outlook.com
git config --global core.autocrlf false
git config --global mergetool.keepBackup false
git config --global difftool.prompt false
git config --global merge.tool vimdiff
git config --global merge.conflictstyle diff3
git config --global http.sslVerify false
```

'git fetch origin && git rebase origin/master' fetch origin all branch, and rebase origin/master branch to local branch

'git log dev ^master' show dev have commit, but master not have commit

'git mergetool'

'git config --global rerere.enabled true' what means?

### bash

```
#!/usr/bin/env bash // bash script
set -u // do not ignore undeclared variables
set -x // display commands before executing commands
set -e // if an error occurs to terminate the execution
set -o pipefail // as long as a sub command fails, the entire pipeline command fails, and the script terminates execution.
```

### cros

```
quote from: http://www.ruanyifeng.com/blog/2016/04/cors.html

// once the browser finds the AJAX request across the source,
// it will automatically add some additional header information and sometimes additional requests,
// but the user will not feel.

simple request:
    method is one of HEAD GET POST and not exceeding Accept Accept-Language Content-Language Last-Event-ID Content-Type(application/x-www-form-urlencoded multipart/form-data text/plain)

    ~~~~~~~~
    -> Origin: xxx
    <- Access-Control-Allow-Origin: xxx
    <- Access-Control-Allow-Credentials: true
    <- Access-Control-Expose-Headers: FooBar // default allow: Cache-Control Content-Language Content-Type Expires Last-Modified Pragma
    ~~~~~~~~
    
not simple request: 
    method is OPTIONS
    
    ~~~~~~~~
    -> Origin: xxx
    -> Access-Control-Request-Method: PUT
    -> Access-Control-Request-Headers: FooBar
    <- Access-Control-Allow-Origin: xxx
    <- Access-Control-Allow-Methods: GET, POST, PUT <- Access-Control-Allow-Headers: X-Custom-Header <- Access-Control-Allow-Credentials: true
    <- Access-Control-Max-Age: 1728000
    ~~~~~~~~
```

### curl

```
curl -v -X POST --data "foo=foo bar" http://localhost
curl -v -X POST --data @hello.txt http://localhost
curl -v -X POST --data-urlencode "k=hello world" --data-urlencode "k1=hello world1" http://localhost
curl -v -X POST --form upload1=@/Users/foobar/hello.js --form a=b --form a=b1 --form a1=b1 http://localhost
curl -v -X POST -c c1.txt --header "Content-Type:application/json" --data '{"username": "foo","password": "bar"}' http://localhost
curl -v -X GET -b c1.txt http://localhost
```

### proxy

```
case1:

client --> proxy --> server
  ^                   |
  |                   |
  |-------------------|
  
case2:

client --> proxy --> server1
                 --> server2
                 --> server3
client <-- proxy <-- server1 or server2 or server3
```

### goland

```
ctrl + k // show git commit and diff dialog
ctrl + shift + k // show push origin branch dialog
ctrl + alt + ; // fetch origin branch dialog
ctrl + q // show function document
ctrl + alt + l // format code
ctrl + alt + <- or -> // back or forward code view
```