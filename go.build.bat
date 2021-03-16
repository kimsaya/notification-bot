go env -w GOOS=linux GOARCH=arm
go build
go env -w GOOS=windows GOARCH=amd64