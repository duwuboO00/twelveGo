package api

// import (
// 	"net/http"
// 	"pkg/models"
// )

// func CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	product := models.Product{}
// 	err := product.Create()
// 	if err != nil {
// 		http.Error(w, "Failed to create product", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Product created successfully"))
// }
