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
					"icon":    "help_outline",
					"color":   "#5F7AFF",
					"tooltip": "Learn more about Xdriven",
					"action": map[string]interface{}{
						"type":  "navigate",
						"route": "/about",
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 24,
					"color":   "#5F7AFF",
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":       "✨ Xdriven is Live",
								"fontSize":   28,
								"fontWeight": "bold",
								"color":      "#FFFFFF",
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "Zero-code UI delivered instantly",
								"fontSize": 18,
								"color":    "rgba(255,255,255,0.9)",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding":      20,
					"margin":       12,
					"borderRadius": 12,
					"color":        "#FFFFFF",
					"shadow":       2,
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":      "\"This entire interface was generated from backend JSON - no app update required\"",
								"fontSize":  15,
								"color":     "#666666",
								"fontStyle": "italic",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding":      16,
					"margin":       12,
					"borderRadius": 12,
					"color":        "#F8FAFF",
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "🚀 Key Benefits",
								"fontSize":     20,
								"fontWeight":   "600",
								"color":        "#5F7AFF",
								"marginBottom": 12,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "• Instant UI updates without app releases",
								"fontSize":     16,
								"color":        "#444444",
								"marginBottom": 8,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "• Dynamic layouts controlled from backend",
								"fontSize":     16,
								"color":        "#444444",
								"marginBottom": 8,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "• A/B test interfaces in real-time",
								"fontSize": 16,
								"color":    "#444444",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding":      20,
					"margin":       12,
					"borderRadius": 12,
					"color":        "#5F7AFF",
					"action": map[string]interface{}{
						"type":  "navigate",
						"route": "/about",
					},
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "👉 See How It Works",
								"fontSize":     18,
								"fontWeight":   "600",
								"color":        "#FFFFFF",
								"marginBottom": 8,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "Tap to explore the technology behind instant UI delivery",
								"fontSize": 14,
								"color":    "rgba(255,255,255,0.9)",
							},
						},
					},
				},
			},
		},
	})
}
