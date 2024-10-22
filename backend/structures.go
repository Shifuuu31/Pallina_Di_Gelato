package backend

import (
	"sync"
	"text/template"
)

// PProfile represents a product's profile
type (
	PProfile struct {
		ID           string   `json:"id"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		Price        float64  `json:"price"`
		Category     string   `json:"category"`
		PublishDate  string   `json:"publishDate"`
		CreationDate string   `json:"creationDate"`
		IsNew        bool     `json:"isNew"`
		IsVisible    bool     `json:"isVisible"` // New field for visibility
		ImageUrls    []string `json:"imageUrls"`
	}

	// Category represents a product category
	Category struct {
		Title           string `json:"title"`
		Description     string `json:"description"`
		ImageURL        string `json:"imageUrl"`
		CreationDate    string `json:"creationDate"`
		DeletedCategory bool   `json:"deletedCategory"`
	}

	// CategoryProduct associates products with a category
	CategoryProduct struct {
		MatchedProducts []PProfile
		MatchedCategory Category
	}
)

var (
	Products   []PProfile
	Categories []Category
	NewProduct PProfile
	Template   = template.Must(template.ParseGlob("./pages/*.html"))
	Mutex      sync.Mutex // Protect shared resources
)
