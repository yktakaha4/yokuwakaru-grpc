# よくわかるgRPC

https://booth.pm/ja/items/1557285

```
$ make server

$ make client message=hoge
```

## tips

### 証明書作成について

https://ichi.pro/https-o-shiyoshite-golang-de-anzenna-kuraianto-to-sa-ba-o-sakuseisuru-187565712042515

go1.15以降こうなるっぽいので避け方

```
$ make client message=ahoaho
cd ./hello-grpc/client/ && go run client.go ahoaho
2021/05/16 20:02:48 gRPC Error (message: connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0")
exit status 1
make: *** [Makefile:9: client] エラー 1
```

証明書作成

- 必要要件
    - certstrap

```
$ certstrap init --common-name ca
$ certstrap request-cert --domain localhost
$ certstrap sign localhost --CA ca

$ cp -vp "$PWD/out/"localhost.{crt,key} "$PWD/hello-grpc/server"
$ ln -s "$PWD/hello-grpc/server/localhost.crt" "$PWD/hello-grpc/client/localhost.crt"
```
