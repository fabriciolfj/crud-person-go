# crud-person-go

# GRPC
- para gerar os arquivos go com base no proto, execute:
```
protoc --go_out=. --go_opt=paths=source_relative    --go-grpc_out=. --go-grpc_opt=paths=source_relative    person.proto
```
