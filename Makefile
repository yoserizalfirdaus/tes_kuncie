start: test run

test: 
	go mod vendor
	go test ./...

run: 
	docker-compose up -d && echo "application is running"

stop:
	docker-compose down