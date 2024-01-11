package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "webcrawler: %v\n", err)
		os.Exit(1)
	}	
	// outline(nil, doc)
    // forEachNode(doc, startElement, endElement)
    fmt.Println(Extract(doc))
}

func Extract(root *html.Node) ([]string, error) {
    links := []string{}
    var visitNode func(*html.Node) = nil
    visitNode = func(node *html.Node) {
        if node.Type == html.ElementNode && node.Data == "a" {
		    for _, a := range node.Attr {
			    if a.Key == "href" {
				    links = append(links, a.Val)
			    }
		    }

	    }
    }
    forEachNode(root, visitNode, nil)
    return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }
    if post != nil {
        post(n)
    }
}

func startElement(n *html.Node) {
    if n.Type == html.ElementNode {
        fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
        depth++
    }
}

func endElement(n *html.Node) {
    if n.Type == html.ElementNode {
        depth--
        fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
    }
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
