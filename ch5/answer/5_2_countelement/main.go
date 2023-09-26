/*
递归实现+ map记录
*/

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
					<span>Span 1</span>
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

	counts := make(map[string]int)
	countElements(counts, doc)

	for tag, count := range counts {
		fmt.Printf("%s: %d\n", tag, count)
	}
}

func countElements(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElements(counts, c)
	}
}
