build : main.go 
	@go vet main.go
	@go fmt main.go
	@golint main.go
	@go build main.go

run : build 
	@./main

all : run 

clean : 
	@rm main 