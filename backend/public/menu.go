package public

import (
	"net/http"
	"time"

	"github.com/Shifuuu31/Pallina_Di_Gelato/backend"
)

// MenuPageHandler handles requests for the menu page
func MenuPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/menu" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	parsedCategories, err := createCategoryProducts()
	if err != nil {
		backend.RenderTemplate(w, backend.Template, "menu.html", "error getting products/Categories")
	}
	backend.RenderTemplate(w, backend.Template, "menu.html", &parsedCategories)
}

// createCategoryProducts organizes products into categories
func createCategoryProducts() ([]backend.CategoryProduct, error) {
	parsedCategories := make([]backend.CategoryProduct, len(backend.Categories))
	for i := range backend.Categories {
		// fmt.Println(backend.Categories[i].Title)
		parsedCategories[i].MatchedCategory = backend.Categories[i]
		for _, product := range backend.Products {
			parsedTime, err := time.Parse("02-01-2006 15:04:05", product.PublishDate)
			if err != nil {
				// print("here")
				// fmt.Println(product.PublishDate)
				// log.Fatalln(err)
				return nil, err
			}
			if product.Category == backend.Categories[i].Title && product.IsVisible && !time.Now().Before(parsedTime) {
				parsedCategories[i].MatchedProducts = append(parsedCategories[i].MatchedProducts, product)
			}
		}
	}
	return parsedCategories, nil
}
