package main

import(
  "html/template"
  "io"
	"net/http"
	"path"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	entries := Store.AllEntries()
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmplStr))
	template.Must(tmpl.Parse(ViewTmplStr))
	tmpl.Execute(w, entries)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	asset, ok := AssetMap[name]

	if ok {
		w.Header().Set("Content-Type", "text/css")
		io.WriteString(w, asset)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "asset not defined")
	}
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	asset, ok := AssetMap[name]
	if ok {
		w.Header().Set("Content-Type", "application/javascript")
		io.WriteString(w, asset)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "asset not defined")
	}
}
