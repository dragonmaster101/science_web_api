format : main.go 
	@$(MAKE) format -C database
	@go vet main.go
	@go fmt main.go
	@golint main.go

test : 
	@$(MAKE) test -C database

build : format
	$(MAKE) build -C database
	@go build main.go

run : build 
	@./main

all : 
	@$(MAKE) test
	@$(MAKE) build

clean : 
	@rm main