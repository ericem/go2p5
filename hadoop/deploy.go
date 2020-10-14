// cerner_2^5_2020
package hadoop

import (
	"fmt"

	"github.com/ericem/go2p5/deploy"
	"github.com/ericem/go2p5/node"
	"github.com/ericem/go2p5/res"
)

func bigtopRepoURL(version string) string {
	return fmt.Sprintf("https://mirrors.gigenet.com/apache/bigtop/bigtop-%s/repos/centos7/bigtop.repo", version)

}

// Create a Deployable for Hadoop
func New(node *node.Node) deploy.Deployable {
	// Create a File resource for the Apache Bigtop RPM repository
	repo := res.File(node, bigtopRepoURL("1.4.0"), "/etc/yum.repos.d/bigtop.repo", res.Present)

	// Create a Package resource for the hadoop package
	hdp := res.Package(node, "hadoop", res.Present)

	// Create a Package resource for the java openjdk package
	java := res.Package(node, "java-1.8.0-openjdk-headless", res.Present)

	// Create the deployment
	hadoop := deploy.New("Hadoop")

	// Add the resources to the deployment
	hadoop.Add(repo)
	hadoop.Add(hdp)
	hadoop.Add(java)

	return hadoop
}
