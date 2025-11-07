package httprouter

import (
	"slices"
	"strings"

	"github.com/cespare/mph"
)

type routeDef struct {
	method  string
	path    string
	handler Handler
}

type routeCompiled struct {
	method       string
	key          string
	paramIndices []int
	paramNames   []string
	handler      Handler
}

type mphGroup struct {
	table  *mph.Table
	routes []routeCompiled
}

type mphGroups map[int]mphGroup

func build(routes []routeDef) mphGroups {
	grouped := make(map[int][]routeCompiled)

	for _, def := range routes {
		var paramIndices []int
		var paramNames []string

		segments := strings.Split(def.path, "/")
		for i, segment := range segments {
			if strings.HasPrefix(segment, ":") {
				paramIndices = append(paramIndices, i)
				paramNames = append(paramNames, strings.TrimPrefix(segment, ":"))
				segments[i] = ":"
			}
		}

		normalized := def.method + ":" + strings.Join(segments, "/")
		depth := len(segments)

		grouped[depth] = append(grouped[depth], routeCompiled{
			method:       def.method,
			key:          normalized,
			paramIndices: paramIndices,
			paramNames:   paramNames,
			handler:      def.handler,
		})
	}

	groups := make(mphGroups)

	for depth, routes := range grouped {
		keys := make([]string, len(routes))
		for i, rc := range routes {
			keys[i] = rc.key
		}

		table := mph.Build(keys)
		groups[depth] = mphGroup{
			table:  table,
			routes: routes,
		}
	}

	return groups
}

func find(groups mphGroups, method, path string) (*routeCompiled, Params) {
	segments := strings.Split(path, "/")
	depth := len(segments)

	group, ok := groups[depth]
	if !ok {
		return nil, nil
	}

	paramMask := group.routes[0].paramIndices
	normalizedSegments := make([]string, len(segments))

	for i, segment := range segments {
		if segment == "" {
			continue
		}
		if slices.Contains(paramMask, i) {
			normalizedSegments[i] = ":"
		} else {
			normalizedSegments[i] = segment
		}
	}

	key := method + ":" + strings.Join(normalizedSegments, "/")

	idx, ok := group.table.Lookup(key)
	if !ok {
		return nil, nil
	}

	route := group.routes[int(idx)]
	params := make(Params, 0, len(route.paramIndices))

	for i, pos := range route.paramIndices {
		if pos < len(segments) {
			params = append(params, Param{
				Name:  route.paramNames[i],
				Value: segments[pos],
			})
		}
	}

	return &route, params
}
