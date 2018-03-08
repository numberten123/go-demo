# go-demo
Simple go-demo include tcp,web,protobuf ,mysql,log...
go version go1.9.3
GOPATH /mnt/hgfs/mahjong/golang/go-test

## mysql 驱动
	go get -u github.com/go-sql-driver/mysql

## proto 驱动
### 1.安装依赖
	git clone https://github.com/google/protobuf.git
	cd protobuf
	./autogen.sh
	./configure
	make
	make install
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib
		
### 2.安装go-protobuf	
	go get -u github.com/golang/protobuf/protoc-gen-go

### 3.生成proto.go文件		
	cd proto && ./create_proto.sh

## 生成go 配置文件		
	cd tool && ./make_all_go.sh
	test
