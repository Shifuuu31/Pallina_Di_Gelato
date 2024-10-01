package source

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)


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
