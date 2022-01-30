build : database.go
	@go vet database.go
	@go fmt database.go
	@golint database.go
	@go build database.go

clean : database 
	@rm database