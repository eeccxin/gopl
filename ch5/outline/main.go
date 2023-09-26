// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// !+
func main() {
	//doc, err := html.Parse(os.Stdin)
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
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data+"|"+strconv.Itoa(int(n.Type))) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-
