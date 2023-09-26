package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

func main() {
	htmlStr := `
		<html>
			<body>
				<div>
					<p>Paragraph 1</p>
					<p>Paragraph 2</p>
					<script>alert("Hello, World!")</script>
					<span>Span 1</span>
					<style>body { color: red; }</style>
					<span>Span 2</span>
				</div>
				<div>
					<p>Paragraph 3</p>
					<span>Span 3</span>
				</div>
			</body>
		</html>
	`
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		log.Fatal(err)
	}

	textNodes := getTextNodes(doc)
	for _, node := range textNodes {
		fmt.Println(node.Data, "|", node.Type)
	}
}

func getTextNodes(n *html.Node) []*html.Node {
	var textNodes []*html.Node

	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" { //去掉空text
		textNodes = append(textNodes, n)
	} else if n.Type != html.ElementNode || (n.Data != "script" && n.Data != "style") { //条件用于过滤掉script标签
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			textNodes = append(textNodes, getTextNodes(c)...) //...参数展平
		}
	}

	return textNodes
}
