.PHONY: all

all: bin/uptime bin/disks bin/osinfo bin/fsinfo bin/nics bin/hostinfo \
	bin/rpms bin/ohey bin/wipedisk bin/partdisk bin/formatxfs bin/wipepart \
	bin/pkginstall bin/pkgremove bin/getfile bin/rmfile bin/users \
	bin/addusertest bin/groups bin/addgrouptest bin/sudoertest \
	bin/nginxstate bin/userstate

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

bin/formatxfs: ./examples/formatxfs/main.go ./cmds/fstool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/wipepart: ./examples/wipepart/main.go ./cmds/fstool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/pkginstall: ./examples/pkginstall/main.go ./cmds/pkgtool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/pkgremove: ./examples/pkgremove/main.go ./cmds/pkgtool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/getfile: ./examples/getfile/main.go ./cmds/file.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/rmfile: ./examples/rmfile/main.go ./cmds/filetool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/users: ./examples/users/main.go ./cmds/user.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/addusertest: ./examples/addusertest/main.go ./cmds/usertool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/groups: ./examples/groups/main.go ./cmds/group.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/addgrouptest: ./examples/addgrouptest/main.go ./cmds/grouptool.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/sudoertest: ./examples/sudoertest/main.go ./cmds/sudoers.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/nginxstate: ./examples/nginxstate/main.go ./res/resource.go ./res/package.go ./node/packages.go ./res/command.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

bin/userstate: ./examples/userstate/main.go ./res/user.go ./node/users.go ./cmds/userkey.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<
