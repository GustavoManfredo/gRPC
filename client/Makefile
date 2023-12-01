build:
	go build main.go

run:
	go run main.go

clean:
	rm -rf pb

gen-files:
	protoc --go_out=. \
	--go-grpc_out=. \
	proto/*.proto