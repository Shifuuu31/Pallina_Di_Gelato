package source

import (
	"encoding/json"
	"os"
	"strings"
)

func LoadProduct() error {
	const filePath string = "/root/Desktop/Palline_Di_Gelato/product_details.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(strings.NewReader(string(jsonData)))
	if err := decoder.Decode(&Products); err != nil {
		// http.Error(w, "Error product details cannot be displayed", http.StatusInternalServerError)
		return err
	}
	return nil
}
