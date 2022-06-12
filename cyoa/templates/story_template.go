package templates

import (
	"text/template"
)

func Form_Template() *template.Template {
	allFiles := []string{"content.tmpl", "footer.tmpl", "header.tmpl", "page.tmpl"}
	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, "./templates/"+tmpl)
	}

	// templates, _ := template.New("page").ParseFiles(allPaths...)
	templates := template.Must(template.New("page").ParseFiles(allPaths...))
	return templates
}
