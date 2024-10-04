package main

import (
	"github.com/Shifuuu31/Palline_Di_Gelato/source"
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", source.HomePageHandler)
	http.HandleFunc("/admin", source.AdminPageHandler)
	http.HandleFunc("/admin/dashbord", source.AdminDashbordPageHandler)
	http.HandleFunc("/menu", source.MenuPageHandler)
	http.HandleFunc("/menu/product", source.ProductPageHandler)
	http.HandleFunc("/about-us", source.AboutUsPageHandler)
	http.HandleFunc("/contact-us", source.ContactUsPageHandler)
	http.HandleFunc("/geo-localisation", source.GeoLocalisationPageHandler)

	source.Open()
	http.ListenAndServe(source.Port, nil)
}
