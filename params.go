package httprouter

type Param struct {
	Name  string
	Value string
}

type Params []Param

func (ps Params) ByName(name string) string {
	for _, p := range ps {
		if p.Name == name {
			return p.Value
		}
	}
	return ""
}

func (ps Params) ByPosition(pos int) string {
	if pos < 0 || pos >= len(ps) {
		return ""
	}
	return ps[pos].Value
}

func (ps Params) Names() []string {
	names := make([]string, len(ps))
	for i, p := range ps {
		names[i] = p.Name
	}
	return names
}

func (ps Params) Values() []string {
	values := make([]string, len(ps))
	for i, p := range ps {
		values[i] = p.Value
	}
	return values
}
