format : database.go 
	@go vet database.go
	@go fmt database.go
	@golint database.go

test: database_test.go 
	@go test -timeout 30s -coverprofile=/tmp/vscode-goybgeto/go-code-cover

build : format
	@go build database.go

clean : database 
	@rm database