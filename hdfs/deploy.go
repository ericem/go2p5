package hdfs

import (
	"fmt"

	"github.com/ericem/go2p5/deploy"
	"github.com/ericem/go2p5/node"
	"github.com/ericem/go2p5/res"
)

func bigtopRepoURL(version string) string {
	return fmt.Sprintf("https://mirrors.gigenet.com/apache/bigtop/bigtop-%s/repos/centos7/bigtop.repo", version)

}

// Create a Deployment for HDFS
func New(node *node.Node) deploy.Deployable {
	// Create a File resource for the Apache Bigtop RPM repository
	repo := res.File(node, bigtopRepoURL("1.4.0"), "/etc/yum.repos.d/bigtop.repo", res.Present)

	// Create a Package resource for the hadoop-hdfs package
	hdfs := res.Package(node, "hadoop-hdfs", res.Present)

	// Create a Package resource for the hadoop-hdfs-datanode package
	datanode := res.Package(node, "hadoop-hdfs-datanode", res.Present)

	// Create a Package resource for the java openjdk package
	java := res.Package(node, "java-1.8.0-openjdk-headless", res.Present)

	// Create the deployment
	d := deploy.New("HDFS")

	// Add the resources to the deployment
	d.Add(repo)
	d.Add(hdfs)
	d.Add(datanode)
	d.Add(java)

	return d
}
