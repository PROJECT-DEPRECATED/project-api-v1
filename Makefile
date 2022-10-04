src = main.go
target = project-api

$(target): $(src)
	go mod tidy	
	go build -o $(target) $(src)

docker:
	docker-compose build --no-cache

clean:
	rm $(target)