
export CGO_ENABLED=0
linux: export GOOS=linux
linux: export GOARCH=amd64
linux: 
	go build -gcflags "all=-N -l" main/esImport.go
	go build -gcflags "all=-N -l" main/esService.go
	go build -gcflags "all=-N -l" main/soldierImportService.go

mac: export GOOS=darwin
mac: export GOARCH=amd64
mac: 
	go build -gcflags "all=-N -l" main/esImport.go
	go build -gcflags "all=-N -l" main/esService.go
	go build -gcflags "all=-N -l" main/soldierImportService.go

initgo: export GO111MODULE=on
initgo:
	go mod download