// cerner_2^5_2020
package deploy

import "fmt"

type Group struct {
	Name    string
	deploys []Deployable
}

// Create a new parallel deployment group
func NewGroup(name string) *Group {
	return &Group{Name: name}
}

// Add a Deploy to the Group
func (g *Group) Add(d Deployable) {
	g.deploys = append(g.deploys, d)
}

// Do the Group deployment by
func (g *Group) Deploy() error {
	fmt.Printf("Starting Group: %s\n", g.Name)
	err := make(chan error, len(g.deploys))
	for _, d := range g.deploys {
		go func(d Deployable) {
			deployErr := d.Deploy()
			err <- deployErr
		}(d)
	}
	for range g.deploys {
		groupErr := <-err
		if groupErr != nil {
			fmt.Printf("%s deployment error.\n", g.Name)
			return groupErr
		}
	}
	fmt.Printf("Group %s complete.\n", g.Name)
	return nil
}
