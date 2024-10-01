package source

import (
	"net/http"
	"text/template"
)


type PProfile struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
	Price       float64  `json:"price"`
	Category    string   `json:"category"`
	PublishDate string   `json:"publishDate"`
	IsNew       bool     `json:"isNew"`
}

var (
	Products   []PProfile
	Categories []string
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	homeTempl := template.Must(template.ParseFiles("./static/index.html"))

	LoadProduct(w)

	execErr := homeTempl.Execute(w, nil)
	if execErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func AdminPageHandler(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/admin" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}

func AdminDashbordPageHandler(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/admin/dashbord" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}

func MenuPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/menu" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}

func ProductPageHandler(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/product" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}


func AboutUsPageHandler(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/about-us" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}

	aboutUsTempl := template.Must(template.ParseFiles("./static/about-us.html"))

	LoadProduct(w)

	execErr := aboutUsTempl.Execute(w, nil)
	if execErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func ContactUsPageHandler(w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/contact-us" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}
