# sqlite3

## 删除SQLite3数据库字段内含有 `陌生人`的ID

- 数据库路径
PC: `E:\\SQLiteStudio\\data\\RecordDB2.db`
armv7 : `/home/data/sqlite/RecordDB2.db`

- 数据库文件名 `RecordDB2.db`
- 表名 `RecordDB2`
- 字段 `RecordID PersoIndex PersonID RecordPicture RecordTime RecordType RecordPass Temperature RecordData`


- 已实现功能

1. 删除含有陌生人字段的ID
2. 删除ID的同时删除该ID对应的图片文件.jpeg格式

## build for armv7

**win10系统开启WSL**

```bash
 sudo tar xzvf go1.17.1.linux-amd64.tar.gz -C /usr/local/ 
 echo "export PATH=\$PATH:/usr/local/go/bin">> .bashrc  
 go env -w GOPROXY="https://goproxy.cn,direct" 
 go env -w GO111MODULE="on"  

sudo apt-get update  
sudo apt-get install build-essential                      
sudo apt-get install gcc-arm-linux-gnueabihf    
sudo apt-get install g++-arm-linux-gnueabihf   
```

```bash
env CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++  CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7  go build -v  -o deleteUnkonwPerson-armv7
```

## 使用

FileZilla 传到目标板

- 直接删除
```bash
chmod +x deleteUnkonwPersonAndImg-armv7 
./deleteUnkonwPersonAndImg-armv7
```

- web页面删除
```bash
./deleteUnkonwPersonAndImg-armv7-web
```
打开浏览器 输入目标机器ip地址:端口

