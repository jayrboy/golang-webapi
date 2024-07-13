# Middleware

```go
func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler middleware start")
		handler.ServeHTTP(w, r)
		fmt.Println("middleware finished")
	})
}

func main() {
	productItemHandler := http.HandlerFunc(productHandler)
	productListHandler := http.HandlerFunc(productsHandler)

	http.Handle("/product/", middlewareHandler(productItemHandler))
	http.Handle("/product", middlewareHandler(productListHandler))

	port := "5000"
	log.Printf("Server running at http://localhost:%s", port)
	http.ListenAndServe(":" + port, nil)
}
```

## Logging

```c
before handler middleware start
middleware finished
```
