package renderer

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tmp := map[string]interface{}{
		"apiVersion": "v3",
		"kind":       "Pod",
		"metadata": map[string]interface{}{
			"name": "rss-site",
			"labels": map[string]interface{}{
				"app": "web",
			},
		},
	}

	want := `digraph  {nodesep="0.8";pad="1";rankdir="LR";ranksep="1.2 equally";node [fontname="monospace",fontsize="10",margin="0.4,0.2",penwidth="1.5",shape="rectangle",style="rounded,filled"];n1[label="apiVersion"];n2[label="v3"];n3[label="kind"];n4[label="Pod"];n5[label="metadata"];n6[label="name"];n7[label="rss-site"];n8[label="labels"];n9[label="app"];n10[label="web"];n1->n2;n3->n4;n5->n6;n5->n8;n6->n7;n8->n9;n9->n10;}`

	if got := flatten(Render(tmp)); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}
