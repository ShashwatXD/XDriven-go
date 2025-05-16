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
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 16,
					"action": map[string]interface{}{
						"type":  "navigate",
						"route": "/profile",
					},
					"children": []map[string]interface{}{
						{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":       "Welcome to Xdriven",
								"fontSize":   24,
								"fontWeight": "bold",
							},
						},
						{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "Build UI from server-side config",
								"fontSize": 16,
								"color":    "#666666",
							},
						},
					},
				},
			},
			{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 12,
					"color":   "#F5F5F5",
					"action": map[string]interface{}{
						"type":  "navigate",
						"route": "/profile",
					},
					"children": []map[string]interface{}{
						{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":       "Why Xdriven?",
								"fontSize":   18,
								"fontWeight": "bold",
							},
						},
						{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "No more app releases for layout updates.",
								"fontSize": 14,
							},
						},
					},
				},
			},
		},
	})

}
