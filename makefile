build : main.go 
	@go build main.go

run : build 
	@./main

all : run 

clean : 
	@rm main 