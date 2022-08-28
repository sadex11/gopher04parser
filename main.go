package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type NodeLink struct {
	Href string
	Text string
}

func createNodeLinks(node *html.Node, links *[]NodeLink) {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			*links = append(*links, NodeLink{attr.Val, getNodeText(node)})
		}
	}
}

func getNodeText(node *html.Node) string {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			// get the first text data
			return strings.TrimSpace(child.Data)
		}
	}

	return ""
}

func iterateNode(node *html.Node, links *[]NodeLink) {
	if node.Type == html.ElementNode && node.Data == "a" {
		createNodeLinks(node, links)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		iterateNode(child, links)
	}
}

func getNodeLinks(reader io.Reader) *[]NodeLink {
	rootNode, err := html.Parse(reader)

	if err != nil {
		panic(err)
	}

	var links *[]NodeLink = new([]NodeLink)
	iterateNode(rootNode, links)
	return links
}

func main() {
	reader, err := os.Open("ex3.html")

	if err != nil {
		panic(err)
	}

	links := getNodeLinks(reader)

	for _, link := range *links {
		fmt.Println(link)
	}
}
