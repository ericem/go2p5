// cerner_2^5_2020
package res

// Interface for applying state to a Resource
type Applier interface {
	Apply() error
}

// A Resource containg state and commands to change the state
type Resource struct {
	cmd func() error
}

// Method to apply (change) the state of a Resource
func (r *Resource) Apply() error {
	return r.cmd()
}

// Type representing possible state of Resource
type State int

// Some common Resource states
const (
	Present State = iota
	Absent
)
