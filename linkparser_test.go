package linkparser

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

var (
	hrefNode  = html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{html.Attribute{Key: "href", Val: "testhref"}}}
	childNOde = html.Node{Type: html.TextNode, Data: "test content"}
)

func TestCreateNodeLinks(t *testing.T) {
	var links []NodeLink
	createNodeLinks(&hrefNode, &links)

	if len(links) != 1 {
		t.Error("Can't get node text")
	}
}

func TestGetNodeText(t *testing.T) {
	hrefNode.FirstChild = &childNOde
	nodeText := getNodeText(&hrefNode)

	if nodeText != "test content" {
		t.Error("Get invalid node text:", nodeText)
	}
}

func TestIterateNode(t *testing.T) {
	var links []NodeLink
	hrefNode.FirstChild = &childNOde
	iterateNode(&hrefNode, &links)

	if len(links) != 1 {
		t.Error("Can't get node text")
	}
}

func TestGetNodeLinks(t *testing.T) {
	reader, err := os.Open("testdata/test.html")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	links := GetNodeLinks(reader)

	if len(*links) != 2 {
		t.Error("Invalid links lenght, expected 2, got", len(*links))
	}
}
