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
								"text":       "🚀 Xdriven is Live",
								"fontSize":   26,
								"fontWeight": "bold",
								"color":      "#222222",
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "This app is not built. It's configured.",
								"fontSize": 18,
								"color":    "#333333",
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "What you're seeing is crafted entirely by backend JSON.",
								"fontSize": 16,
								"color":    "#333333",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikText",
				"props": map[string]interface{}{
					"text":     "\"This UI didn’t exist 5 seconds ago.\"",
					"fontSize": 14,
					"color":    "#555555",
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
								"text":     "✅ Ship UI without redeploys",
								"fontSize": 16,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "✅ Add new screens with 0 code changes",
								"fontSize": 16,
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "✅ Control layout from backend in real-time",
								"fontSize": 16,
							},
						},
					},
				},
			},
			map[string]interface{}{
				"uiType": "UikContainer",
				"props": map[string]interface{}{
					"padding": 16,
					"color":   "#F0F4F8",
					"action": map[string]interface{}{
						"type":  "navigate",
						"route": "/about",
					},
					"children": []interface{}{
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":       "🎯 Want to know how it works?",
								"fontSize":   16,
								"fontWeight": "bold",
							},
						},
						map[string]interface{}{
							"uiType": "UikText",
							"props": map[string]interface{}{
								"text":     "Tap anywhere on this block to dive into the magic of Xdriven.",
								"fontSize": 14,
								"color":    "#333333",
							},
						},
					},
				},
			},
		},
	})

}
