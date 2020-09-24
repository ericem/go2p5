.PHONY: all

all: bin/uptime bin/disks bin/osinfo bin/fsinfo bin/nics bin/hostinfo \
	bin/rpms bin/ohey bin/wipedisk bin/partdisk

bin/uptime: ./examples/uptime/main.go ./cmds/uptime.go
	 GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/disks: ./examples/disks/main.go ./cmds/disks.go
	 GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/nics: ./examples/nics/main.go ./cmds/net.go
	 GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/fsinfo: ./examples/fsinfo/main.go ./cmds/fsinfo.go
	 GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/hostinfo: ./examples/hostinfo/main.go ./cmds/host.go
	 GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/osinfo: ./examples/osinfo/main.go ./cmds/os.go
	 GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/rpms: ./examples/rpms/main.go ./cmds/rpm.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/ohey: ./examples/ohey/main.go ./node/node.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/wipedisk: ./examples/wipedisk/main.go ./cmds/disktool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/partdisk: ./examples/partdisk/main.go ./cmds/disktool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<
