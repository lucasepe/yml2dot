package renderer

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestRender(t *testing.T) {
	tmp := []yaml.MapItem{
		{Key: "apiVersion", Value: "v3"},
		{Key: "kind", Value: "Pod"},
		{Key: "metadata", Value: map[string]interface{}{
			"name": "rss-site",
			"labels": map[string]interface{}{
				"app": "web",
			},
		},
		},
	}

	got := Render(tmp)

	id := "n1"
	if n := got.FindNodeByID(id); n == nil {
		t.Errorf("node with label [%v] not found", id)
	}

	id = "n2"
	if n := got.FindNodeByID(id); n == nil {
		t.Errorf("node with label [%v] not found", id)
	}

	id = "n43"
	if n := got.FindNodeByID(id); n != nil {
		t.Errorf("node with label [%v] should not exist", id)
	}
}
