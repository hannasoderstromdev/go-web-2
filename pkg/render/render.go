package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// RenderTemplate parses a template file and renders it as HTML
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// create template cache
	tc, err := createTemplateCache()
	if err !=  nil {
		log.Fatal(err)
	}

	// get template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	
	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)

	myCache := map[string]*template.Template{}

	// get all files *.page.tmpl from ./templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		fileName := filepath.Base(page)
		ts, err := template.New(fileName).ParseFiles(page)
		if err !=  nil {
			return myCache, err
		}

		// Get all layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err !=  nil {
			return myCache, err
		}

		// If there's a match, parse the layout file
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err !=  nil {
				return myCache, err
			}
		}

		myCache[fileName] = ts
	}

	return myCache, nil
}