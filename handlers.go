package main

import (
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

func newHandler(w http.ResponseWriter, r *http.Request) {
	entry := NewEntry("", "", "")
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmplStr))
	template.Must(tmpl.Parse(AddTmplStr))
	tmpl.Execute(w, entry)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	entry := Store.GetEntry(id)
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmplStr))
	template.Must(tmpl.Parse(EditTmplStr))
	tmpl.Execute(w, entry)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	entry := NewEntryFromReq(r)
	Store.SaveEntry(entry)
	Skeddy.ReStart(Store.AllEntries())
	http.Redirect(w, r, "/", 301)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	entry := NewEntryFromReq(r)
	Store.SaveEntry(entry)
	Skeddy.AddEntry(entry)
	http.Redirect(w, r, "/", 301)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	Store.DeleteEntry(id)
	Skeddy.ReStart(Store.AllEntries())
	http.Redirect(w, r, "/", 301)
}
