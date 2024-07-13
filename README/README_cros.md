# Ngrok

https://ngrok.com/download

# CROS

```go
func enableCorsMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, X-Content-Security-Policy")
        handler.ServeHTTP(w, r)
	})
}

func main() {
	productItemHandler := http.HandlerFunc(productHandler)
	productListHandler := http.HandlerFunc(productsHandler)

	http.Handle("/product/", enableCorsMiddleware(productItemHandler))
	http.Handle("/product", enableCorsMiddleware(productListHandler))

	port := "5000"
	log.Printf("Server running at http://localhost:%s", port)
	http.ListenAndServe(":" + port, nil)
}
```
