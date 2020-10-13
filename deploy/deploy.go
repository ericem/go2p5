// cerner_2^5_2020
package deploy

import (
	"fmt"

	"github.com/ericem/go2p5/res"
)

type Deployable interface {
	Deploy() error
}

type Deploy struct {
	Name      string
	resources []*res.Resource
}

// Create a new deployment
func New(name string) *Deploy {
	return &Deploy{Name: name}
}

// Add a Resource to the resources for deployment
func (d *Deploy) Add(r *res.Resource) {
	d.resources = append(d.resources, r)
}

// Do the deployment by applying each Resource in turn
func (d *Deploy) Deploy() error {
	fmt.Printf("Starting %s deployment\n", d.Name)
	for _, r := range d.resources {
		err := r.Apply()
		if err != nil {
			return err
		}
	}
	fmt.Println("Deployment complete.")
	return nil
}
