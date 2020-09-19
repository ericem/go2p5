.PHONY: all

all: bin/uptime bin/disks bin/osinfo bin/fsinfo bin/nics

bin/uptime: ./examples/uptime/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/uptime $<

bin/disks: ./examples/disks/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/disks $<

bin/osinfo: ./examples/osinfo/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/osinfo $<

bin/fsinfo: ./examples/fsinfo/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/fsinfo $<

bin/nics: ./examples/nics/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/nics $<
