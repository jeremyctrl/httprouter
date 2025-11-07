package httprouter

import (
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

		segments := strings.Split(def.path, "/")
		for i, segment := range segments {
			if strings.HasPrefix(segment, ":") {
				paramIndices = append(paramIndices, i)
				segments[i] = ":"
			}
		}

		normalized := def.method + ":" + strings.Join(segments, "/")
		depth := len(segments)

		grouped[depth] = append(grouped[depth], routeCompiled{
			method:       def.method,
			key:          normalized,
			paramIndices: paramIndices,
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

func find(groups mphGroups, method, path string) (*routeCompiled, []string) {

}
