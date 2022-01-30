build : main.go 
	$(MAKE) -C database
	@go vet main.go
	@go fmt main.go
	@golint main.go
	@go build main.go

run : build 
	@./main

all : run 

clean : 
	$(MAKE) clean -C database
	@rm main 