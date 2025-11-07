package httprouter

import (
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

type mphResult map[int]mphGroup

func build(routes []routeDef) mphResult {

}

func find(groups mphResult, method, path string) (*routeCompiled, []string) {

}
