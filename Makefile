src=main.go
target=project-api
TAG=default

$(target): $(src)
	go mod tidy	
	go build -o $(target) $(src)

docker:
	docker-compose build --no-cache

publish:
    ./publish.sh $(TAG)

clean:
	rm $(target)