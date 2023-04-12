compose_files=$(shell find cmd -maxdepth 2 -exec readlink -f {} \; | grep docker-compose.yml | xargs -I {} echo -f {})
# compose_root=/home/darrenb/projects/tryroll/docker-compose.yml
compose_up=docker compose ${compose_files} up --remove-orphans --force-recreate --build

test:
	go test ./...

gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	protoc \
	--go_out=. \
	--go_opt=paths=import \
	--go-grpc_out=. \
	--go-grpc_opt=paths=import api/grpc/*.proto

gen-gql:
	go run github.com/99designs/gqlgen generate

dep:
	go mod tidy
	go mod vendor

down:
	docker compose ${compose_files} down --remove-orphans

run: down
	$(compose_up)