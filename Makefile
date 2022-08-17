build:
	go build -o bin/ps-notification-service ./cmd/ps-notification-service

run:
	./bin/ps-notification-service start

clean:
	rm -rf bin

all:
	clean build run