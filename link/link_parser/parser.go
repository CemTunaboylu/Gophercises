package link_parser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func Parse_Links(r io.Reader) map[string][]string {

	html_node, err := html.Parse(r)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	// iterative dfs
	var href_node_map map[string]*html.Node = map[string]*html.Node{}
	stack := []*html.Node{}
	for {
		if html_node != nil {
			stack = append(stack, html_node)
			html_node = html_node.FirstChild

		} else if len(stack) > 0 {
			html_node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// use html_node
			if html_node.Type == html.ElementNode && html_node.Data == "a" {
				for _, a := range html_node.Attr {
					if a.Key == "href" {
						href_node_map[a.Val] = html_node
						break
					}
				}
			}
			html_node = html_node.NextSibling
		} else {
			break
		}

	}
	var hrefs_and_texts map[string][]string = map[string][]string{}
	for k, v := range href_node_map {
		hrefs_and_texts[k] = gather_everything_underneath(v)
	}
	return hrefs_and_texts

}

func gather_everything_underneath(node *html.Node) (str_arr []string) {
	stack := []*html.Node{}
	node = node.FirstChild
	for {
		if node != nil {
			stack = append(stack, node)
			node = node.FirstChild

		} else if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if node.Type != html.CommentNode && node.Type != html.ElementNode {
				cleaned := strings.TrimSpace(node.Data)
				if cleaned != "" {
					// fmt.Printf("'%v' from %v is getting added\n", cleaned, node.Type)
					str_arr = append(str_arr, cleaned)
				}
			}
			node = node.NextSibling

		} else {
			break
		}
	}

	return

}
