echo '启动用户 api 8888'
go run user/userapi/userapi.go -f user/userapi/etc/users.yaml

#goctl api plugin -plugin goctl-swagger="swagger -filename app.json -host localhost:8888 -basepath /" -api v1.api -dir ./doc