package main

import (
	"net/http"

	"github.com/Shifuuu31/Palline_Di_Gelato/source"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", source.HomePageHandler)
	http.HandleFunc("/admin", source.AdminPageHandler)
	http.HandleFunc("/admin/dashboard", source.AdminDashboardPageHandler)
	http.HandleFunc("/admin/dashboard/add/product", source.AddNewProductPageHandler)
	http.HandleFunc("/admin/dashboard/add/", source.AddingHandler)
	
	http.HandleFunc("/menu", source.MenuPageHandler)
	http.HandleFunc("/menu/product", source.ProductPageHandler)
	http.HandleFunc("/about-us", source.AboutUsPageHandler)
	http.HandleFunc("/contact-us", source.ContactUsPageHandler)
	http.HandleFunc("/geo-localisation", source.GeoLocalisationPageHandler)
	if err := source.LoadProduct(); err!= nil {
		println(err)
	}
	source.Open()
	http.ListenAndServe(source.Port, nil)
}
