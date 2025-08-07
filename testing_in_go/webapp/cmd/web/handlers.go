package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"path"
	"time"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	td := make(map[string]any)
	if app.Session.Exists(r.Context(), "test") {
		msg := app.Session.GetString(r.Context(), "test")
		td["test"] = msg
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at "+time.Now().UTC().String())
	}
	app.render(w, r, "home.page.html", &TemplateData{Data: td})
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse the template from disk
	slog.Debug("calling application render", "t", t)
	p1 := path.Join(pathToTemplates, t)
	p2 := path.Join(pathToTemplates, "base.layout.html")

	// beware, the order matters for parse files
	// in Go, the order of the provided filenames is significant for determining
	// the default template of the returned *template.Template object.
	parsedTemplate, err := template.ParseFiles(p1, p2)
	if err != nil {
		slog.Error("error in application render", "err", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// execute the template, passing it data, if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		slog.Error("error in application render, execute template", "t", t, "err", err)
		return err
	}

	slog.Debug("application render success", "t", t)
	return nil
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// validate data
	form := NewForm(r.PostForm)
	form.Required("email", "password")
	if !form.Valid() {
		fmt.Fprint(w, "failed validation")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	log.Println(email, password)

	fmt.Fprint(w, email)
}
