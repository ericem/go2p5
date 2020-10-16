# Hacking on go2p5



## Install Prerequisites

* macOS
* Home Brew
* VirtualBox 
* Vagrant
* Go
* Task



### macOS



I did all of the development for this competition on macOS. However, it should cross-compile just fine on other platforms but have not been tested.

### Home Brew

Home Brew makes it easy to install Go, VirtualBox and Vagrant on macOS. It is a package manager similar to Yum or Apt on Linux.

Install Home Brew:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```



### VirtualBox

I used VirtualBox to execute and test the code for this competition. It is a Virtual Machine hypervisor that runs on many different operating systems, including macOS.

Install VirtualBox

```bash
brew cask install virtualbox
```

 

### Vagrant

Vagrant is an automation tool for virtual machines that makes creating and destroying the test VM for the code super easy.

Install Vagrant:

```bash
brew cash install vagrant
```



### Go

All of the competition entries in this repository are written in the Go programming language. Home Brew usually has the latest version of Go.

Install Go:

```bash
brew install go
```



### Task

I used Task to build the Go code in this repository. Task is a task runner and build tool similar to Make, but is written in Go and therefore fits nicely with a Go project.

Install Task:

```bash
brew install go-task/tap/go-task
```



## Build Code

To build the the examples and other binaries in this repository use the included Taskfile.yml and the Task utility.

```bash
task build
```

All binaries created during the build will be placed in the bin/ directory. The Taskfile uses environment variables to turn on cross-compilation for Linux, the target platform for the binaries in this repository. 

## Create Virtual Machine

Once the binaries have been built they can tested by running them on any Enterprise Linux distribution including Redhat, Oracle, and CentOS. This repository includes a Vagrantfile for creating an Oracle Linux 7 virtual machine. To boot up the example virtual machine use vagrant.

```bash
vagrant up
```



## Run Examples

Once the virtual machine has finished booting up, use the vagrant utility to ssh into the virtual machine. This will login to the VM using the default *vagrant* account.

```bash
vagrant ssh
```

To run any of the examples included in this repository requires root access; use sudo to gain root access.

```bash
sudo su
```

The VirtualBox virtual machine contains the necessary software to automatically mount the local repository directory into a directory inside the virtual machine. You can run the examples by using cd to switch to the /vagrant/bin directory.

```bash
cd /vagrant/bin
```

The first example created for entry 1 is *uptime*:

```bash
./uptime
```

