package main

import (
	"cyoa/arc"
	"cyoa/http_handlers"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	var stories_json_file string
	flag.StringVar(&stories_json_file, "json", "gopher.json", "the json file that the story arcs are in")
	flag.Parse()
	// open the file
	s_f, err := ioutil.ReadFile(stories_json_file)
	check(err)

	story_arcs, err := arc.Story_Arcs_From_JSON(s_f)
	check(err)

	mux := default_mux()
	static_file_regexp := regexp.MustCompile(`^/(js|css)/.*`)
	f_handler := http_handlers.Static_File_Handler(static_file_regexp, mux)

	handler := http_handlers.Form_Handler(story_arcs, http_handlers.With_Fallback(f_handler))
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func default_mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http_handlers.NoArcHandler)
	return mux
}
