# Cerner 2^5 Programming Competition



<img src="https://raw.githubusercontent.com/egonelbre/gophers/master/vector/friends/crash-dummy.svg" align="right"
     alt="Crash dummy by @egonelbre" width="120" height="178">


This repository contains my entries for the 2020 edition of the Cerner 2^5 programming competition. For this years competition I have chosen to learn the Go programming language while working on my submissions. Why Go you ask? Well for starters, Go has been proclaimed [the language of cloud infrastructure](https://hackernoon.com/go-has-indeed-become-the-language-of-cloud-infrastructure-rob-pike-q51a3wle), and I work a team supporting cloud infrastructure. It's a language that I have been wanting to learn and now seems like a great time to jump in. For this years competition, each entry will have a cloud and devops focus. Go can be a little verbose, but I hope that I can still make something useful within the constraints of the rules.



## The Rules
* Any Cerner associate in any role or office location can participate.  You just have to know or learn how to write code.
* Submissions must be on a public repository on github.com.
* Only 1 submission per day will count.
* Submissions must be 32 lines of code or less.  (Lines will be counted based on the terminal symbols for the language, not based on a number of characters per line.)
* Comments will not count as lines.  (This way you are encouraged to comment your code.) However, the code must be able to stand on its own without comments.
* Please include a comment in your code of "cerner_2^5_2020". This will assist in a live feed of code examples which are pulled from GitHub.
* The code must be syntactically correct.
* Pair or more programming is okay, but keep in mind, we award prizes based on submission or repository, not on team size.  You may have to share a prize should your team win.
* The competition runs September 14 – October 16, 2020.  No submissions will be accepted after October 16th at Midnight.
* Winners will be announced 2 weeks after the close of the competition.



## The Entries

For this years competition, I used the Go programming language for all of my submissions. Since Go is fast becoming the defacto language for cloud native infrastructure, I wanted to learn Go by building something useful for a devops engineer working on cloud infrastructure. In last year's competition each entry that I submitted was a standalone application, this year I took a different approach. Each entry is a tiny piece of a larger API that once completed can be used for configuration management of a Linux server. Since this was my first attempt at writing Go code you can see my style evolve over time as I learn more. Also, due to the constraints of the competition, the code does not have as robust error checking as it should, so don't use this on real servers!



### Current State of Configuration Management

My overall goal was to build a configuration management API for configuring servers in an idempotent manner. There are already many config management solutions like [Chef](https://www.chef.io/products/chef-infra), [Ansible](https://www.ansible.com), [Saltstack](https://www.saltstack.com), [Puppet](https://puppet.com), and [others](https://en.wikipedia.org/wiki/Comparison_of_open-source_configuration_management_software). So why create a new one? 

All of the current solutions were created before cloud ops was common place and were designed for managing [pets instead of cattle](http://cloudscaling.com/blog/cloud-computing/the-history-of-pets-vs-cattle/). For instance Chef, Saltstack, and Puppet all require an agent to be installed on the server and for the agent to communicate with a central controller. An architecture that requires constantly grooming to maintain the association between active servers and the central controller. An architecture that is at odds with the cattle based philosophy of continuously deployment to upgrade functionality or upon failure.  

By contrast, Ansible does not require a central server for configuration management (often referred to as masterless architecture). However, Ansible (like the others) requires a runtime to be installed on every server (in this case Python). By requiring a runtime for configuration management, you run into a chicken and egg scenario. The config management requires a runtime, the runtime needs to be installed by config management.  The current config management solutions break this dependency by either using a bootstrap method to install the agent, or by installing the agent and its dependencies during the initial OS installation (through OS dependent techniques).

Another axis that config management solutions can be measured on is their method for describing server configuration. For instance, Chef and Puppet use a Ruby based DSL while Ansible and Saltstack use declarative YAML files to describe configuration. Each technique has it's benefits and problems. The YAML based methods are simple, declarative and easy to reason about. However, it is difficult (and clumsy) to express the looping structures that are necessary to describe dynamically generated configuration. The Ruby based DSLs are infinitely flexible and handily provide for dynamically generated content, but at the cost of a very steep learning curve and immense complexity. Also, since Ruby is a dynamically typed language and the configuration can be modified at many different layers there is a considerable amount of defensive coding required which adds to the technical debt of these DSL based solutions. 

For these reasons, my ideal config management system for cloud based infrastructure would 1) not require a runtime to be installed on the server, 2) not depend on a central controller for configuration management, 3) use code written in a statically typed language to express the desired state of the server.



### A Different Way

After thinking about what my ideal configuration management system would look like, I came up with a rough sketch of how such a system would work:

* Use Go as the configuration language to take advantage of its statically typed, compiled nature to generate a dependency free binary

* Represent server state as types that can be checked at compile type
* Encapsulate OS specific shell commands as functions with 1) arguments that can be type checked 2) have a defined unit of behavior, 3) a typed return value
* Collect current server state just once so that it can be re-used efficiently across many operations without a performance penalty
* Combine server state and shell functions into a higher-level construct that can diff current state against the desired state and then apply a new state in an idempotent operation
* Combine a sequence of states to be applied to a server into a deployment that can be re-used in different servers, environments and contexts
* Group together several atomic deployments into a single unit that can then be applied to a server using parallel execution to speed the configuration process
* Generate a single binary that can be copied to a server to apply the desired state to the server; requiring no runtime, or central controller

Using this rough outline, I set out to see if I could build something, 32 lines at a time, that could achieve this. I had major reservations if this would work out selecting Go as my language for implementation. From what I had read before I started, Go was known for requiring excess boilerplate that the dynamic languages did not. However, I told myself that this was mainly an exercise to learn Go and if it failed, at least I would still have learned something! Well, to my surprise, Go turned to out to be a great language for this type of project. Yes, it has some boilerplate, but it's main method of code re-use is composition of interfaces and this turns out to work great with small snippets of code. I could create an interface in one file and then split the methods across multiple files, keeping the code size < 32 lines. 

The most important thing I learned using Go was what an incredibly productive language it is. I use Neovim as my editor and it has great support for the Go language, providing type completion, and diagnostic feedback of syntax errors. With these two features combined, I was able to edit, save, compile and it would always "Just Work". My usual practice with Ruby or Python is to have a REPL open to check bits of code to see if it will work. Not the case with Go. I was able to write directly from idea to code without using much reference materials, and usually worked the first time, truly amazing. 

So, here are the entries. You will see that since I was writing this one bit at a time in the evenings, I sometimes forgot a piece of functionality and would have to implement that the next night before I could carry on with progress towards the ultimate goal. I probably spent too much time at the beginning on the OS specific shell commands instead of moving onto more important features. If I had more time I would have fleshed out a remote execution story including a built-in remote file copy over SSH. I did not use any 3rd party libraries, all functionality was implemented from scratch using the Go standard library. This keeps the code base light, dependency free, and efficient.

### 1 cmds/uptime.go

This is a function to return the current uptime of a server. The return value is the current uptime exposed as a time.Duration value from the Go standard library. By using a Duration, the consumers of this function cand do date math with the result. I chose this as the first feature to build because I would learn to use the Go file operations as well as time and string manipulation.

   

### 2 cmds/disks.go

This is a function to list all the block devices on a server that are disks. The return value is a slice of Disk objects. A slice in Go terminology is like an expandable array. I used this function in a later entry to gather the current state of the disks in a server. By having a type checked value for all of disks on a system, it makes it really easy to re-use the disk data for other purposes; for config management, system metrics reporting, etc.



### 3 cmds/os.go

This is a function to parse the Name, Version, and ID of the OS on the system. This comes from the `/etc/os-release` file which is a semi-standard file to store OS version information. This function makes it possible to build a config management system that can alter its behavior based on the OS that it is running on. For instance, when creating a Package resource, the system can use RPMs on an Enterprise Linux system or DEBs on a Debian based system.

### 4 parse/kvpair.go

This entry was a case of trying to be DRY. After writing the disks.go and os.go functions, I realized that many Linux commands return output in a key value pair format and that I needed a generic key value parser that could be re-used across all shell functions. However, while there are several parsers in the Go standard library, there is no key value parser, so this entry was created. This function takes a string (a line from a file for example), parses it and returns a map of key names to values. A map is Go terminology for a Hash or Dict. 

### 5 cmds/net.go

This entry parses the IP configuration of a server and returns a slice of Nic objects. The Nic objects contain the interface name and IP address information. This entry uses the Go JSON parser to parse the data since the `ip` commands have an option to display their output in JSON format. Very handy, I wish all Linux commands had this feature!

### 6 cmds/host.go

This entry parses and returns several machine specific bits of information including the Model, Family, and Vendor information taken from the BIOS of the machine. When running on virtual machines, this information can be used identify the type of hypervisor for example VirtualBox. This function could be used to alter the behavior of the config system based on the type of machine (virtual versus physical, etc.).

### 7 cmds/fsinfo.go

This entry returns a slice of Filesystem objects, enumerating all the formatted filesystems on a server. It uses the key value parser from entry 4 to do the parsing. This is a critical piece of information for any configuration management system.

### 8 cmds/rpm.go

This entry will enumerate all of the RPM packages installed on a server. You can see that the data structure for holding RPM data is called Package. My thinking was to support multiple different OS distributions with their different package formats, but use the same struct to make the data parsing consistent. Also, you will note that I learned how JSON encoding works while writing this entry. The Package struct is the first struct in this repo to include the JSON annotations, which describe how the field names should look after encoding. I used the lower-level *rpm* tool to gather the data instead of *yum* because it was much faster 300ms vs 4 seconds!

### 9 node/node.go

The previous entries were building a set of data structures and functions for gathering facts about a server. This entry combines those data structures into a single Node object. This allows the config management system to cache the pieces of data that it collects about the server improving efficiency in larger deployments. This entry also includes a method to convert the Node object to a JSON string. This is one of my favorite parts about Go. It's built-in JSON encoder can automatically convert data structures to JSON. I included an example executable called `ohey` that dumps the state of a server in JSON format to the terminal, a knock-off of Chef's `ohai` command that does the same thing. The ohey source is only 8 lines, that's efficiency!

### 10 cmds/disktool.go

Starting with this entry, I began to add behaviors to the data structures created in the previous entries. This entry also shows an important feature of Go programs, methods. These are functions attached to types. You can attach methods to any type, but in this case I attached a *Wipe* method and *Partition* method to the Disk type. Because the type definition's don't have to be in the same file as the methods, its trivial to split applications into small bits and spread them across multiple files, perfect for a competition like this. 

### 11 cmds/fstool.go

Like the disktool.go entry, this entry provides methods for working with filesystems. It has a *Wipe* method and a *Format* method. If I had more time, I would have liked to have made a *Wiper* interface that could call the *Wipe* method on an object. This would allow some more code re-use when dealing with block devices in general. Those are the kinds of things you notice only after you start to see the repetition of methods in the implementation. 

### 12 cmds/pkgtool.go

This entry adds methods to *Install* and *Remove* yum packages. I thought about adding a method to *Upgrade* packages as well. However, I did not have enough lines to complete the feature and also thought it might not be necessary. An upgrade operation could probably be implemented on top of Install by tracking the state of the package version.

### 13 cmds/file.go

This entry adds a data structure *File* and a method *Download* for downloading files to the server from a URL. Why Download? My thinking was in a cloud based environment, you would most likely store configuration assets and files in a blob storage like S3. Therefore, it simplifies the architecture if the server downloads files to itself instead of moving or copying files to the server. After I submitted the entry, I realized that I should have used the type *FileInfo* from the Go standard library instead of creating my own. This would allow tighter integration with other methods in the *os* package.

### 14 cmds/filetool.go

This was a small entry to add a *Remove* method to the *File* type. Even though these methods in the last few entries are OS specific, I was trying to keep the method names consistent and generic so that they could be abstracted across OS flavors. I would like to have added a method to change file permissions, but because I did not have enough lines in *file.go* for the struct fields to track the permissions, I could not. If I had just used *FileInfo*, darn it!

### 15 cmds/user.go

You can probably see the pattern starting to emerge; a type defining a piece of server state and methods to operate on the server to change that state. This entry follows the same pattern, a *User* type to store user account information and a *Users* function to read the state of all users on a server. The type information is small and you might think it would have been better to pack those into a *types* package, but I've learned from reading about Go development that it is not idiomatic to do that. It is better to define the types near where they are going to be used. 

### 16 cmds/usertool.go

This entry adds methods for working with User accounts. It adds an *Add* method to create a new user, and a *SetPassword* method to set the password for an account. The *SetPassword* method uses a feature of Command execution that allows creating pipes to stdin and stdout of a process. In this case stdin of the *passwd* is connected to the method so that the password can be set without it leaking to the process table on the server.

### 17 cmds/group.go

To complement the Users function, the *Groups* function enumerates all of the group accounts on the server. It returns a slice of Group types which can then be used for state processing with Group information.

### 18 cmds/grouptool.go

Like *usertool.go* this entry has an *Add* method and a *Remove* method for working with Group types. If I had more lines available, I would have liked to add functionality to differentiate system from user level groups.

### 19 cmds/sudoers.go

This entry adds a type *Sudoers* and methods for working with the sudoers configuration on a server. The *Add* method can be used to add an arbitrary sudoers file to the server in the `/etc/sudoers.d` directory. The *NoPasswd* method follows another idiom in Go where a wrapper method is used make it easier to work with an API in a commonly used case. So here, the *NoPasswd* method is a shortcut for generating an a sudoers snippet allowing a specific user to access root permissions without a password. I wash thinking about making it easier to create Vagrant boxes when I added this feature.

### 20 res/resource.go

This entry marks another milestone where I began to create a higher-level abstraction on top of the simple data types and methods from the preceding entries. In this entry I introduced the *Resource* type which is used to encapsulate the state of a node and the domain knowledge required to convert from one state to another. The basic idea is; when the *Apply* method on a Resource is called it can read a Node object, determine the current state of the server and then call OS specific methods to convert the server from the current state to the new desired state. This entry also adds types for the common states *Absent* and *Present*.  

### 21 res/command.go

This entry adds some additional helper functionality for working with Resources. The *SetCommand* method is used to set the internal *cmd* field on a *Resource* to the function that will be called during the *Apply*. This function is responsible for changing the state of the *Resource*.

### 22 node/packages.go

I added the methods in this entry to the Node object, to make it easier to query and get the current state of packages on a server. I wanted to have a way to abstract over the low-level OS specific functionality for getting package state and have methods for common types of state queries.

### 23 res/package.go

This entry adds the *Package* function which creates a new Resource that is responsible for maintaining the state of a single OS package on a server. In my opinion this Package function shows how powerful and elegant a type based config management system can be. Since the state of the Resource is represented using the types Absent and Present, it is very readable and apparent what the code is doing and how it is modifying the state. You will also note that if the state has not changed then the function will set the Resource's internal *cmd* to a *NoOp*. This makes the Resource very efficient when the state has not changed, incurring no performance penalty.

 ### 24 cmds/userkey.go

This entry was a case of me realizing that I was missing a couple key features when working with Users. I added a *Remove* method for deleting users from a server and added *SetPubKey* method for setting an SSH Public Key for a user to access the server. Similar to the File's *Download* method the *SetPubKey* method downloads a key from a URL to the user's home directory in the standard location (~/.ssh) for keys.

### 25 res/user.go

I got the git comment wrong for this entry, it *is* entry number 25. This entry adds a *User* function for creating a Resource to manage the state of a user on a server. Again, because of the layers of abstraction, this function is very readable and looks very similar to the Package function. The anonymous function passed into the SetCommand method of the Resource shows how multiple steps can be used to modify the state of the Resource. In this case a User is added to the server, then the password is set for the user and finally a public key is set for the user. 

### 26 node/users.go

This entry adds some missing functionality on the Node type needed for the the *res.User* function. It adds a *GetUSer* method for querying the current state of a single user. It also adds a *PathExists* method which returns true/false if a a file or directory exists. It doesn't really fit in this file, but hey when you only have 32 lines to work with, you have to add functionality in any place you can. Since Go has package level visibility for functions and names, it makes it really easy to move methods and types around to different files.

### 27 node/file.go

This entry adds a method to the Node type to get the status of a single file. This is really just a help method that makes writing the code for a Resource more succinct. By putting this helper method on the Node type instead of in the Resource, it can be re-used across multiple Resource definitions.

### 28 res/file.go

This entry adds the *File* function for creating a Resource for managing the state of a file on the server. If the desired state is *Present* and the current state is *Absent* the Resource will download the file to the appropriate destination path. If the desired state changes at some point to *Absent* after the file has been downloaded, then the next time that the *Apply* method is called, the file will be deleted by calling its *Remove* method.

### 29 deploy/deploy.go

This entry adds another layer of abstraction to the config system. It creates a new *Deploy* type which is a data structure for aggregating multiple Resources into a single deployable unit. For instance, if you are deploying an Nginx web server, you would need a Package Resource to manage the state of the OS package and a File Resource to manage the state of the configuration file. These could then be grouped into an Nginx deployment that can be deployed as a single unit. This Deploy can then be re-used on any server that needs Nginx, for instance a web based dashboard and a product launch web site. 

### 30 hadoop/deploy.go

Up until this point all of the entries were building tiny bits of critical functionality that can be used to create a config management system. This entry pulls all of those pieces together to show how the API could be used to deploy Hadoop on a server. The *New* function in the *hadoop* package is used to create a new Hadoop deployment. It creates new Resources for the necessary Packages and Files and then adds these Resources to the Deploy. Since this is just regular Go code, it uses Go packages for organization and re-use. If someone wanted to re-use what I have developed here, they could just import *github.com/ericem/go2p5/hadoop* into their project and programmatically install Hadoop onto their server. Now that is efficient! 

### 31 deploy/group.go

This entry shows off one of Go's other great features, easy concurrent execution. In this entry I created a type called *Group* that is used to aggregate multiple deployments into a single unit. There is a method named *Add* that is used to add a Deployable to the Group. In Go terminology the Deployable type is an interface. An interface just describes the behavior that a type has. So in this case any type that has a Deploy method can be added to the Group. This adds flexibility to extend the API in the future to support other types that at this time are unimaginable. The *Deploy* method uses go routines to execute all the deployments in parallel. This means the order of execution is not guaranteed. Use a regular Deploy type to guarantee the sequential application of Resources and then use a Group to speed up deployment of several independent Deploys.   

### 32 cmd/datanode-deploy/main.go

My final entry demonstrates how the configuration management API that I have created can be used to create a binary that can be copied to a server and executed. As you can see in this entry, due to the abstractions, the end user of the API has a very simple and composable way to build a self-contained executable for configuring a server. First a Node is created to retrieve the current state of the server, then a Group is created to which will contain each of the independent Deploys. Next, a Hadoop and an HDFS Deploy are added to the Group. Finally, the entire Group state is applied to the server by executing the *Deploy* method. When run on a sever this executable will apply the state in an idempotent and consistent manner with no need for an agent or runtime installed on the server. All the state and logic required for configuration are contained in the binary. That is engineering efficiency!  



## Examples

For each entry I also created an example application to go along with it. I probably should have written Go tests using the API instead, but each example has the same purpose to exercise the API. 