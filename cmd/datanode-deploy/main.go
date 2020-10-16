// cerner_2^5_2020
package main

import (
	"github.com/ericem/go2p5/deploy"
	"github.com/ericem/go2p5/hadoop"
	"github.com/ericem/go2p5/hdfs"
	"github.com/ericem/go2p5/node"
)

// An application to deploy an HDFS datanode
func main() {
	// Create a Node instance and read the current state of the node
	node := node.New()

	// Create a parallel deploy group for deploying a datanode
	datanode := deploy.NewGroup("Datanode")

	// Add a hadoop deployment to the datanode
	datanode.Add(hadoop.New(&node))

	// Add an hdfs deployment to the datanode
	datanode.Add(hdfs.New(&node))

	// Start the datanode deployment and run all deployments in parallel
	datanode.Deploy()
}
