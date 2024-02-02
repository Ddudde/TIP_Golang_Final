

go get github.com/spf13/viper
go get google.golang.org/grpc
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt

go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger

swag init -g mainRoute/main.go

go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
