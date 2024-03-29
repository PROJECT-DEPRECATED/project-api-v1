MAIN = main.go
SRC = *.go **/*.go
TARGET = project-api
TAG = default

$(TARGET): $(SRC)
	go build -o $(TARGET) $(MAIN)

docker:
	docker-compose build --no-cache

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

publish:
	./scripts/publish.sh $(TAG)

run:
	go run $(SRC)

clean:
	rm $(TARGET)
