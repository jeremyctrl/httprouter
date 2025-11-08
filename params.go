package httprouter

// Param represents a single URL parameter extracted from the path.
type Param struct {
	Name  string
	Value string
}

// Params is a slice of Param values representing all parameters
// captured in a route. It provides helper methods for access.
type Params []Param

// ByName returns the value of the parameter with the given name,
// or an empty string if not present.
func (ps Params) ByName(name string) string {
	for _, p := range ps {
		if p.Name == name {
			return p.Value
		}
	}
	return ""
}

// ByPosition returns the parameter value at the given position,
// or an empty string if the index is out of range.
func (ps Params) ByPosition(pos int) string {
	if pos < 0 || pos >= len(ps) {
		return ""
	}
	return ps[pos].Value
}

// Names returns a slice containing all parameter names in order.
func (ps Params) Names() []string {
	names := make([]string, len(ps))
	for i, p := range ps {
		names[i] = p.Name
	}
	return names
}

// Values returns a slice containing all parameter values in order.
func (ps Params) Values() []string {
	values := make([]string, len(ps))
	for i, p := range ps {
		values[i] = p.Value
	}
	return values
}
