package renderer

import (
	"fmt"
	"testing"

	"github.com/lucasepe/dot"
	"gopkg.in/yaml.v2"
)

func TestRender(t *testing.T) {
	tmp := []yaml.MapItem{
		{Key: "apiVersion", Value: "v3"},
		{Key: "web", Value: "not cached"},
		{Key: "kind", Value: "Pod"},
		{Key: "metadata", Value: map[string]interface{}{
			"name": "rss-site",
			"labels": map[string]interface{}{
				"app":     "web",
				"other":   "network",
				"another": "web",
			},
		},
		},
	}

	got := Render(tmp, true)
	fmt.Println(got)

	id := "n1"
	if n := got.FindNodeByID(id); n == nil {
		t.Errorf("node with id [%v] not found", id)
	}

	id = "n2"
	if n := got.FindNodeByID(id); n == nil {
		t.Errorf("node with id [%v] not found", id)
	}

	id = "n43"
	if n := got.FindNodeByID(id); n != nil {
		t.Errorf("node with id [%v] should not exist", id)
	}

	// Test caching of "web" leaf node
	label := "web"
	if nn := FindAllByLabel(got, label); len(nn) != 1 {
		t.Errorf("Incorrect number '%d' of nodes with label [%v]", len(nn), label)
	}
}

// FindAllByLabel returns all nodes with a label
// TODO: move to dot package
func FindAllByLabel(g *dot.Graph, label string) (founds []*dot.Node) {
	g.VisitNodes(func(node *dot.Node) (done bool) {
		if l := node.Value("label"); l != nil {
			// l.(string) type assertion causes panic: interface conversion: interface {} is dot.HTML, not string
			if fmt.Sprintf("%v", l) == label {
				founds = append(founds, node)
			}
		}
		return false
	})
	return
}
