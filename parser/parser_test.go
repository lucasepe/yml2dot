package parser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	yaml := `apiVersion: v3
kind: Pod
metadata:
  name: rss-site
  labels:
    app: web
`

	res, err := Parse(strings.NewReader(yaml), "", "")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(res), 3; got != want {
		t.Fatalf("expected length [3], got [%d]", got)
	}
}

func TestFullScan(t *testing.T) {
	yaml := `apiVersion: v3
kind: Pod
metadata:
  name: rss-site
  labels:
    app: web
`

	res, err := fetchYAML(strings.NewReader(yaml), "", "")
	if err != nil {
		t.Fatal(err)
	}

	if got := string(res); got != yaml {
		t.Errorf("got [%v] want [%v]", got, yaml)
	}
}

func TestDelimitedScan(t *testing.T) {
	input := `
/*** INFO ***
apiVersion: v3
kind: Pod
metadata:
  name: rss-site
  labels:
    app: web
***/

public class Box {
    private Object object;

    public void set(Object object) { 
		this.object = object; 
	}
    public Object get() { return object; }
}
`

	res, err := fetchYAML(strings.NewReader(input), "/*** INFO ***", "***/")
	if err != nil {
		t.Fatal(err)
	}

	want := `apiVersion: v3
kind: Pod
metadata:
  name: rss-site
  labels:
    app: web`

	if got := strings.TrimSpace(string(res)); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
