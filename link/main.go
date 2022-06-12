package main

import (
	"flag"
	"fmt"
	"link/link_parser"
	"os"
	"path/filepath"
)

func main() {

	var f_path_for_htmls string
	flag.StringVar(&f_path_for_htmls, "path", "htmls/", "The file path for htmls to be parsed, we find htmls by ourselves so only the path will suffice.")
	flag.Parse()

	// I find every HTML file in the given path
	html_files := pick_file_type(f_path_for_htmls, ".html")

	fmt.Printf("%v\n", html_files)

	var doc_parsed_maps map[string]map[string][]string = map[string]map[string][]string{}

	// I parse them all and put them in a file ? or a DB
	for _, file := range html_files {
		r, err := os.OpenFile(file, os.O_RDONLY, 444)
		defer r.Close()
		check(err)
		doc_parsed_maps[file] = link_parser.Parse_Links(r)
	}

	for k, v := range doc_parsed_maps {
		fmt.Printf("'%v' : %v\n", k, v)
	}

}

func pick_file_type(f_path string, f_type string) (files []string) {
	err := filepath.Walk(f_path, func(path string, info os.FileInfo, err error) error {
		check(err)
		// if not a directory and ends with f_type -> can be more than .<type> if you want
		if !info.IsDir() && filepath.Ext(path) == f_type {
			files = append(files, path)
		}
		return nil
	})

	check(err)
	return
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error occured : %v", e)
		panic(e)
	}
}
