package source

import (
	"fmt"
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
type Category struct {
    Title string
    ImageURL string
}

var (
	Products   []PProfile
	Categories []Category 
)

type CategoryProduct struct {
	MatchedProducts []PProfile
	MatchedCategory Category
}


func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	if err := LoadProduct(); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	Categories = []Category{
		{Title: "Ice Cream", ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTjZ5QpVr0Yq0taYNdA7vg1AfD7KUfEbk_NXQ&s"},
		{Title: "Frozen Yogurt", ImageURL: "https://images.immediate.co.uk/production/volatile/sites/30/2020/08/recipe-image-legacy-id-1029452_10-563fda8.jpg"},
		{Title: "Sorbet", ImageURL: "https://www.lecremedelacrumb.com/wp-content/uploads/2014/06/raspberry-sorbet-1.jpg"},
		{Title: "Gelato", ImageURL: "https://emmaduckworthbakes.com/wp-content/uploads/2023/06/Chocolate-Gelato-Recipe.jpg"},
		{Title: "Sundaes", ImageURL: "https://www.keep-calm-and-eat-ice-cream.com/wp-content/uploads/2022/08/Ice-cream-sundae-hero-11.jpg"},
	}

	
	homeTempl := template.Must(template.ParseFiles("./static/home.html"))

	execErr := homeTempl.Execute(w, &Categories)
	if execErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}

func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}

func AdminDashbordPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/dashbord" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}

func MenuPageHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/menu" {
        http.Error(w, "404 Page Not Found", http.StatusNotFound)
        return
    }

    Categories = []Category{
        {Title: "Ice Cream", ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTjZ5QpVr0Yq0taYNdA7vg1AfD7KUfEbk_NXQ&s"},
        {Title: "Frozen Yogurt", ImageURL: "https://images.immediate.co.uk/production/volatile/sites/30/2020/08/recipe-image-legacy-id-1029452_10-563fda8.jpg"},
        {Title: "Sorbet", ImageURL: "https://www.lecremedelacrumb.com/wp-content/uploads/2014/06/raspberry-sorbet-1.jpg"},
        {Title: "Gelato", ImageURL: "https://emmaduckworthbakes.com/wp-content/uploads/2023/06/Chocolate-Gelato-Recipe.jpg"},
        {Title: "Sundaes", ImageURL: "https://www.keep-calm-and-eat-ice-cream.com/wp-content/uploads/2022/08/Ice-cream-sundae-hero-11.jpg"},
    }

    parsedCategories := make([]CategoryProduct, len(Categories))

    for i := 0; i < len(Categories); i++ {
		parsedCategories[i].MatchedCategory = Categories[i] // Store the matched category
		for j := 0; j < len(Products); j++ {
			if Products[j].Category == Categories[i].Title {
				parsedCategories[i].MatchedProducts = append(parsedCategories[i].MatchedProducts, Products[j])
			}
		}
	}
	
	
	fmt.Printf("Parsed Categories: %+v\n", parsedCategories)


    menuTempl := template.Must(template.ParseFiles("./static/menu.html"))

    // Execute the template with parsedCategories
    if execErr := menuTempl.Execute(w, &parsedCategories); execErr != nil {
        // Log the error and return a 500 status code
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return // Exit to avoid further writes
    }
}




func ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/product" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}

func AboutUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about-us" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}

	aboutUsTempl := template.Must(template.ParseFiles("./static/about-us.html"))

	// LoadProduct(w)

	execErr := aboutUsTempl.Execute(w, nil)
	if execErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func ContactUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact-us" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}
}


func GeoLocalisationPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/geo-localisation" {
		http.Error(w, "404 Page Not Found", http.StatusInternalServerError)
		return
	}

	glTempl := template.Must(template.ParseFiles("./static/geo-localisation.html"))

	// LoadProduct(w)

	execErr := glTempl.Execute(w, nil)
	if execErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}