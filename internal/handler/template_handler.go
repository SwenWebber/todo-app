package handler

import (
	"net/http"
	"text/template"

	"github.com/swenwebber/todo-app/internal/service"
)

type TemplateHandler struct {
	template *template.Template
	service  *service.TaskService
}

func NewTemplateHandler(service *service.TaskService) *TemplateHandler {
	tmpl := template.Must(template.ParseFiles("/home/umid/Documents/todo-app/templates/index.html"))

	return &TemplateHandler{
		template: tmpl,
		service:  service,
	}
}

func (h *TemplateHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.template.Execute(w, nil) //JS handles data fetching
}

func (h *TemplateHandler) ServeFiles(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path

	if p == "./" {
		p = "./templates/index.html"
	}
	http.ServeFile(w, r, p)
}
