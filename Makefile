
export CGO_ENABLED=0
linux: export GOOS=linux
linux: export GOARCH=amd64
linux: 
	go build main/esImport.go
	go build main/esService.go
	go build main/soldierImportService.go

initgo: export GO111MODULE=on
initgo:
	go mod download