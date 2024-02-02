protoc mainRoute/pb/*.proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. --go_opt=paths=source_relative --openapiv2_out ./docs

protoc authM/pb/*.proto --go_out=. --go-grpc_out=. --go_opt=paths=source_relative
protoc profileM/pb/*.proto --go_out=. --go-grpc_out=. --go_opt=paths=source_relative
protoc productM/pb/*.proto --go_out=. --go-grpc_out=. --go_opt=paths=source_relative