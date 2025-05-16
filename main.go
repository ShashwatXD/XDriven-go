package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/home", handleHome)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"widgets": []map[string]interface{}{
			{
				"uiType": "UikText",
				"props": map[string]interface{}{
					"text":       "Welcome to Xdriven",
					"fontSize":   18,
					"color":      "#333333",
					"fontWeight": "bold",
				},
			},
			{
				"uiType": "UikText",
				"props": map[string]interface{}{
					"text":     "Another text below",
					"fontSize": 14,
					"color":    "#777777",
				},
			},
		},
	})

}
