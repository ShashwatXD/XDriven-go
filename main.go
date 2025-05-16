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
					"icon":    "info",
					"color":   "#5C6B73",
					"tooltip": "About Xdriven",
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
					"color":   "#5C6B73",
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":       "🌌 Xdriven is Live",
								"fontSize":   28,
								"fontWeight": "bold",
								"color":      "#FFFFFF",
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "Your UI, dynamically configured",
								"fontSize": 18,
								"color":    "rgba(255,255,255,0.85)",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 20,
					"margin": map[string]interface{}{
						"top":    16,
						"bottom": 16,
						"left":   12,
						"right":  12,
					},
					"borderRadius": 8,
					"color":        "#FFFFFF",
					"border": map[string]interface{}{
						"width": 1,
						"color": "#C2DFE3",
					},
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":      "\"This interface materialized from JSON - no app update needed\"",
								"fontSize":  15,
								"color":     "#5C6B73",
								"fontStyle": "italic",
								"textAlign": "center",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 20,
					"margin": map[string]interface{}{
						"bottom": 16,
						"left":   12,
						"right":  12,
					},
					"borderRadius": 8,
					"color":        "#F8FBFB",
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "🛠️ How It Works",
								"fontSize":     20,
								"fontWeight":   "600",
								"color":        "#5C6B73",
								"marginBottom": 12,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "▸ Instant UI updates without deployments",
								"fontSize":     16,
								"color":        "#5C6B73",
								"marginBottom": 8,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "▸ Complete layout control from backend",
								"fontSize":     16,
								"color":        "#5C6B73",
								"marginBottom": 8,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "▸ Real-time interface experimentation",
								"fontSize": 16,
								"color":    "#5C6B73",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 20,
					"margin": map[string]interface{}{
						"left":  12,
						"right": 12,
					},
					"borderRadius": 8,
					"color":        "#9DB4C0",
					"action": map[string]interface{}{
						"type":  "navigate",
						"route": "/about",
					},
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":         "Discover More →",
								"fontSize":     18,
								"fontWeight":   "600",
								"color":        "#FFFFFF",
								"marginBottom": 8,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "See the technology behind dynamic UI generation",
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
