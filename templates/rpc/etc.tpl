Name: {{.serviceName}}.rpc
ListenOn: 127.0.0.1:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc

Mysql:
  DataSource: root:123456@tcp(localhost:3306)/queojv4?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai

Redis:
  Host: localhost:6379