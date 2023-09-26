package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	htmlStr := `
		<html>
			<body>
				<div>
					<p>Paragraph 1</p>
					<p>Paragraph 2</p>
					<img src="image.jpg" alt="Image">
					<script src="script.js"></script>
					<style>body { color: red; }</style>
				</div>
			</body>
		</html>
	`

	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		log.Fatal(err)
	}

	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "img":
			handleImageNode(n)
		case "script":
			handleScriptNode(n)
		case "style":
			handleStyleNode(n)
		default:
			// 处理其他元素节点
			fmt.Println(n.Data)
		}
	} else if n.Type == html.TextNode {
		// 处理文本节点
		if strings.TrimSpace(n.Data) != "" {
			fmt.Println(n.Data)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}

func handleImageNode(n *html.Node) {
	// 处理图片节点
	fmt.Println("Image Node:", n.Data)
}

func handleScriptNode(n *html.Node) {
	// 处理脚本节点
	fmt.Println("Script Node:", n.Data)
}

func handleStyleNode(n *html.Node) {
	// 处理样式表节点
	fmt.Println("Style Node:", n.Data)
}
