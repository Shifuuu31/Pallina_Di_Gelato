package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
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
	Host, Port = "localhost", ":8080"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	homeTempl := template.Must(template.ParseFiles("./static/index.html"))
	
	LoadProduct(w)

	// fmt.Println(Products)
	execErr := homeTempl.Execute(w, nil)
	if execErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/menu", MenuPageHandler)
	http.HandleFunc("/menu/product", ProductPageHandler)
	http.HandleFunc("/about-us", AboutUsPageHandler)
	http.HandleFunc("/contact-us", ContactUsPageHandler)

	open()
	http.ListenAndServe(Host+Port, nil)
}

func LoadProduct(w http.ResponseWriter) {
	const filePath string = "/root/Desktop/Palline_Di_Gelato/product_details.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v\n", err), http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(strings.NewReader(string(jsonData)))
	if err := decoder.Decode(&Products); err != nil {
		http.Error(w, "Error product details cannot be displayed", http.StatusInternalServerError)
		return
	}
}

func open() {
	fmt.Println("server details:")
	fmt.Println("\tstatus: \033[1m\033[92m• Live\033[0m")
	fmt.Println("\t" + Host + Port)

	// // Command to run
	// cmd := exec.Command("bash", "openBrowser.sh") // Example: running 'grep main'

	// // Set up pipes for standard input and output
	// // cmd.Stdin = os.Stdin
	// // cmd.Stdout = os.Stdout
	// // cmd.Stderr = os.Stderr

	// // Run the command
	// cmderr := cmd.Run()
	// if cmderr != nil {
	// 	log.Fatalln(cmderr)
	// }
}
