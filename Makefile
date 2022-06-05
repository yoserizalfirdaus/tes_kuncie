start: test run

test: 
	go test ./...

run: 
	docker-compose up -d && echo "application is running"

stop:
	docker-compose down