build: 
	cd task-backend/cmd/server/ && go build -o ../../../target/task-backend

run: build 
	./target/task-backend 

vet: 
	go vet ./task-backend/...

lint:
	golint ./task-backend/...

test:
	go test -cover ./task-backend/...

cover: clean vet build coveralls

clean:
	rm ./target/task-backend 
	go clean ./...
	
coveralls:
	go test -v -covermode=count -coverprofile=coverage.out ./task-backend/...
