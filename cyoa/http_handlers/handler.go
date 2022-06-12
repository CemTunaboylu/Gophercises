package http_handlers

import (
	"bytes"
	"cyoa/arc"
	"cyoa/templates"
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

type Handler struct {
	stories     arc.Stories
	path_parser func(*http.Request) string
	template    *template.Template
	fallback    http.Handler
}
type HandlerOption func(h *Handler)

func With_Path_Parser(p_parser func(*http.Request) string) HandlerOption {
	return func(h *Handler) {
		h.path_parser = p_parser
	}
}

func With_Custom_Template(c_temp *template.Template) HandlerOption {
	return func(h *Handler) {
		h.template = c_temp
	}
}

func With_Fallback(fb http.Handler) HandlerOption {
	return func(h *Handler) {
		h.fallback = fb
	}
}

func Form_Handler(story arc.Stories, handler_opts ...HandlerOption) http.Handler {
	h := Handler{stories: story, path_parser: parse_path, template: templates.Form_Template()}
	for _, h_opts := range handler_opts {
		h_opts(&h)
	}
	return h
}
func Static_File_Handler(paths_regexp *regexp.Regexp, fallback http.Handler) http.HandlerFunc {

	fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	return func(w http.ResponseWriter, r *http.Request) {
		k := r.URL.Path
		matched := paths_regexp.MatchString(k)
		if matched {
			fs.ServeHTTP(w, r)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}

}

func NoArcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Where are you!? There is nothing here ...")
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	k := h.path_parser(r)

	if s, ok := h.stories[k]; ok {
		var b bytes.Buffer
		h.template.ExecuteTemplate(&b, "page", s)
		w.Write(b.Bytes())
	} else {
		h.fallback.ServeHTTP(w, r)
	}

}

func parse_path(r *http.Request) string {
	p := r.URL.Path
	if p == "" || p == "/" {
		p = "/intro"
	}
	return p[1:]

}
