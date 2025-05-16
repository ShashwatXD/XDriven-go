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
		"widgets": []interface{}{
			map[string]interface{}{
				"uiType": "UikIconButton",
				"props": map[string]interface{}{
					"icon":    "question_mark",
					"tooltip": "About",
					"action": map[string]interface{}{
						"type":  "navigate",
						"route": "/about",
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 16,
					"color":   "#A9BCD0",
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":       "Welcome to Xdriven",
								"fontSize":   24,
								"fontWeight": "bold",
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "The screen you’re seeing right now is built entirely from backend JSON.",
								"fontSize": 16,
								"color":    "#333333",
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "Not a single widget is hardcoded.",
								"fontSize": 14,
								"color":    "#333333",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 16,
					"color":   "#D8DBE2",
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "✅ No app store approvals",
								"fontSize": 14,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "✅ Layouts controlled from server",
								"fontSize": 14,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "✅ New pages with zero redeploy",
								"fontSize": 14,
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 16,
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":       "Just update JSON. Boom. UI changes instantly.",
								"fontSize":   16,
								"fontWeight": "bold",
							},
						},
					},
				},
			},
		},
	})

}
