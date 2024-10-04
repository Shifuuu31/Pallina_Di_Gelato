package main

import (
	"net/http"

	"Pallina_Di_Gelato/source"
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
	http.ListenAndServe(source.Host+source.Port, nil)
}
