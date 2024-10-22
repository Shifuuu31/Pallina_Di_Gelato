package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Shifuuu31/Palline_Di_Gelato/backend"
	"github.com/Shifuuu31/Palline_Di_Gelato/backend/private"
	"github.com/Shifuuu31/Palline_Di_Gelato/backend/public"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))
	_, err := os.Stat("." + r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/") || err != nil {
		http.Error(w, "Forbidden: Access Denied", http.StatusForbidden)
		// CheckError(w, http.StatusForbidden, "Forbidden: Access Denied")
		return
	}
	fs.ServeHTTP(w, r)
}

func main() {
	

	http.HandleFunc("/assets/", StaticHandler)

	http.HandleFunc("/login", private.LoginPageHandler)
	http.HandleFunc("/dashboard", private.DashboardPageHandler)
	http.HandleFunc("/dashboard/add/product", private.AddProductPageHandler)
	http.HandleFunc("/dashboard/add-product", private.AddingHandler)
	http.HandleFunc("/dashboard/delete/product", private.DeleteProduct)
	http.HandleFunc("/dashboard/delete-product", private.DeletingHandler)
	http.HandleFunc("/dashboard/edit/product", private.EditProduct)
	http.HandleFunc("/dashboard/add/category", private.AddCategory)
	http.HandleFunc("/dashboard/delete/category", private.DeleteCategory)
	http.HandleFunc("/dashboard/edit/category", private.EditCategory)

	http.HandleFunc("/", public.HomePageHandler)
	http.HandleFunc("/menu", public.MenuPageHandler)
	http.HandleFunc("/menu/product", public.ProductPageHandler)
	http.HandleFunc("/about-us", public.AboutUsPageHandler)
	http.HandleFunc("/contact-us", public.ContactUsPageHandler)
	http.HandleFunc("/find-us", public.FindUsPageHandler)

	// tt()

	if err := backend.LoadProducts(); err != nil {
		log.Fatal(err)
	}
	if err := backend.LoadCategories(); err != nil {
		log.Fatal(err)
	}
	backend.Open()
	if err := http.ListenAndServe(backend.Port, nil); err != nil {
		log.Fatal(err)
	}
}

func tt() {
	// Define two time strings
	timeString1 := "2024-10-22 14:09:50"
	timeString2 := "2024-10-22 14:09:50"

	// Define the layout
	layout := "2006-01-02 15:04:05"

	// Parse the strings into time.Time variables
	time1, err1 := time.Parse(layout, timeString1)
	time2, err2 := time.Parse(layout, timeString2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing time:", err1, err2)
		return
	}

	// Compare the times
	if time1.Equal(time2) {
		fmt.Println("time1 is equal to time2")
	} else if time1.Before(time2) {
		fmt.Println("time1 is before time2")
	} else if time1.After(time2) {
		fmt.Println("time1 is after time2")
	}
}
