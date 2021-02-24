package renderer

import (
	"fmt"
	"reflect"

	"github.com/lucasepe/dot"
)

// Render returns a GraphViz representation of a YAML tree.
func Render(v interface{}) string {
	g := dot.NewGraph(dot.Directed)
	g.Attrs("nodesep", "0.8", "pad", "1", "rankdir", "LR", "ranksep", "1.2 equally")
	g.NodeGlobalAttrs("fontname", "monospace", "fontsize", "10", "margin", "0.4,0.2", "penwidth", "1.5", "shape", "rectangle", "style", "rounded,filled")

	renderVal(v, g, nil)

	return g.String()
}

func renderVal(v interface{}, g *dot.Graph, parent *dot.Node) {
	typ := reflect.TypeOf(v).Kind()

	if typ == reflect.Slice {
		renderSlice(v.([]interface{}), g, parent)
		return
	}

	if typ == reflect.Map {
		renderMap(v.(map[string]interface{}), g, parent)
		return
	}

	child := g.Node(fmt.Sprintf("%v", v))
	g.Edge(parent, child)
}

func renderMap(m map[string]interface{}, g *dot.Graph, parent *dot.Node) {
	for k, v := range m {
		child := g.Node(k)
		if parent != nil {
			g.Edge(parent, child)
		}
		renderVal(v, g, child)
	}
}

func renderSlice(slc []interface{}, g *dot.Graph, parent *dot.Node) {
	for _, v := range slc {
		renderVal(v, g, parent)
	}
}
