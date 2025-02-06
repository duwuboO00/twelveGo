package api

// import (
// 	"net/http"
// 	"pkg/models"
// )

// func CreateCart(w http.ResponseWriter, r *http.Request) {
// 	cart := models.Cart{}
// 	err := cart.Create()
// 	if err != nil {
// 		http.Error(w, "Failed to create cart", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Cart created successfully"))
// }

// func AddItemToCart(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to add product to cart
// }

// func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
// 	// Implement logic to remove product from cart
// }
