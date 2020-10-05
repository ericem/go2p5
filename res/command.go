// cerner_2^5_2020
package res

// A function to change state, run by an Applier
type Command func() error

// Set Resource.cmd to the supplied Command
func (r *Resource) SetCommand(cmd Command) {
	r.cmd = cmd
}

// Create a function that does nothing for use when state is unchaged
func NoOp() Command {
	return func() error { return nil }
}

// Create an new empty Resource with default NoOp command
func NewResource() Resource {
	return Resource{cmd: NoOp()}
}
