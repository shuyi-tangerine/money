

``` shell
# server
go run *.go -server=true -P=compact -buffered=true -framed=false -addr=localhost:31080 -secure=false -web_addr=localhost:31081
# client rpc
go run *.go -server=false -P=compact -buffered=true -framed=false -addr=localhost:31080 -secure=false
# client web
curl -XGET localhost:31081/ping
```





