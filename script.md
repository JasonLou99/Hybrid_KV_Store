install Go:
```bash
wget https://studygolang.com/dl/golang/go1.15.4.linux-amd64.tar.gz
sudo tar -zxvf go1.15.4.linux-amd64.tar.gz -C ~/
mkdir ~/Go
echo "export GOPATH=~/Go" >>  ~/.bashrc 
echo "export GOROOT=~/go" >> ~/.bashrc 
echo "export GOTOOLS=\$GOROOT/pkg/tool" >> ~/.bashrc
echo "export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc
go env -w GOPROXY=https://goproxy.cn
```


start kvserver cluster: 
`go run kvstore/kvserver/kvserver.go -address 192.168.10.120:3088 -internalAddress 192.168.10.120:30881 -peers 192.168.10.120:30881,192.168.10.121:30881,192.168.10.122:30881`

start kvclient: 
`go run ./kvstore/kvclient/kvclient.go -cnums 1 -mode RequestRatio -onums 100 -getratio 4 -servers 192.168.10.120:3088,192.168.10.121:3088,192.168.10.122:3088`
* cums: number of clients, goroutines simulate
* onums: number of operations, each client will do onums operations
* mode: only support RequestRatio (put/get ratio is changeable)
* getRatio: get times per put time
* servers: kvserver address
operation times = cnums * onums * (1+getRatio)