package renderer

import (
	"fmt"

	"github.com/lucasepe/dot"
	"gopkg.in/yaml.v2"
)

var childCache map[string]*dot.Node

// Render returns a GraphViz representation of a YAML tree.
func Render(v yaml.MapSlice, isCacheEnabled bool) *dot.Graph {
	g := dot.NewGraph(dot.Directed)
	g.Attr("nodesep", "0.4")
	g.Attr("rankdir", "LR")
	g.Attr("pad", "0.5")
	g.Attr("ranksep", "0.25 equally")
	g.Attr("fontname", "Fira Mono")
	g.Attr("fontsize", "14")

	g.NodeBaseAttrs().
		Attr("fontname", "Fira Mono").
		Attr("fontsize", "10").
		Attr("margin", "0.3,0.1").
		Attr("fillcolor", "#fafafa").
		Attr("shape", "box").
		Attr("penwidth", "2.0").
		Attr("style", "rounded,filled")

	for _, el := range v {
		renderMapItem(el, g, nil, isCacheEnabled)
	}

	return g
}

func renderMapItem(v yaml.MapItem, g *dot.Graph, parent *dot.Node, isCacheEnabled bool) {
	child := g.Node()
	if parent != nil {
		child.Attr("label", fmt.Sprintf("%v", v.Key))

		link := g.Edge(parent, child)
		link.Attr("arrowhead", "none")
		link.Attr("penwidth", "2.0")
	} else {
		child.Attr("label", dot.HTML(fmt.Sprintf("<b>%v</b>", v.Key)))
		child.Attr("shape", "plaintext")
		child.Attr("style", "")
	}

	renderVal(v.Value, g, child, isCacheEnabled)
}

// getVal - return new Nodes, or use an optional cache
func getVal(g *dot.Graph, label string, isCacheEnabled bool) *dot.Node {
	child, present := childCache[label]
	if !present {
		child = g.Node()
		child.Attr("label", label)
		if isCacheEnabled {
			// A nil map is always empty, create a cache only when needed
			if childCache == nil {
				childCache = make(map[string]*dot.Node)
			}
			childCache[label] = child
		}
	}
	return child
}

func renderVal(v interface{}, g *dot.Graph, parent *dot.Node, isCacheEnabled bool) {
	switch v.(type) {
	case []interface{}:
		renderSlice(v.([]interface{}), g, parent, isCacheEnabled)
	case yaml.MapSlice:
		for _, el := range v.(yaml.MapSlice) {
			renderMapItem(el, g, parent, isCacheEnabled)
		}
	case map[string]interface{}:
		renderMap(v.(map[string]interface{}), g, parent, isCacheEnabled)
	default:
		label := fmt.Sprintf("%v", v)
		child := getVal(g, label, isCacheEnabled)
		if parent != nil {
			link := g.Edge(parent, child)
			link.Attr("arrowhead", "none")
			link.Attr("penwidth", "2.0")
		}
	}
}

func renderMap(m map[string]interface{}, g *dot.Graph, parent *dot.Node, isCacheEnabled bool) {
	for k, v := range m {
		child := g.Node(dot.WithLabel(k))
		if parent != nil {
			link := g.Edge(parent, child)
			link.Attr("arrowhead", "none")
			link.Attr("penwidth", "2.0")
		}
		renderVal(v, g, child, isCacheEnabled)
	}
}

func renderSlice(slc []interface{}, g *dot.Graph, parent *dot.Node, isCacheEnabled bool) {
	for _, v := range slc {
		renderVal(v, g, parent, isCacheEnabled)
	}
}
