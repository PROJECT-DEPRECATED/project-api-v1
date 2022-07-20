src = main.go
target = project-api

$(target): $(src)
	go mod tidy	
	go build -o $(target) $(src)

clean:
	rm $(target)