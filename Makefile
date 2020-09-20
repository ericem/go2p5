.PHONY: all

all: bin/uptime bin/disks bin/osinfo bin/fsinfo bin/nics bin/hostinfo

bin/uptime: ./examples/uptime/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/uptime $<

bin/disks: ./examples/disks/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/disks $<

bin/nics: ./examples/nics/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/nics $<

bin/fsinfo: ./examples/fsinfo/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/fsinfo $<

bin/hostinfo: ./examples/hostinfo/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/hostinfo $<

bin/osinfo: ./examples/osinfo/main.go
	 GOOS=linux GOARCH=amd64 go build -o bin/osinfo $<
