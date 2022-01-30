format : main.go 
	@$(MAKE) format -C database
	@go vet main.go
	@go fmt main.go
	@golint main.go

build : format
	$(MAKE) build -C database
	@go build main.go

run : build 
	@./main

all : run 

clean : 
	$(MAKE) clean -C database
	@rm main 