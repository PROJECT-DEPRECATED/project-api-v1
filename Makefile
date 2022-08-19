src = main.go
target = project-api

$(target): $(src)
	go mod tidy	
	go build -o $(target) $(src)

docker:
	go mod tidy	
	go build -o $(target) $(src)
	docker build --no-cache --tag project-api:latest .

clean:
	rm $(target)