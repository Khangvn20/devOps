package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Product struct
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product
var nextID = 1

func main() {
	router := gin.Default()

	// Routes
	router.GET("/products", getProducts)
	router.POST("/products", addProduct)

	// Start server
	router.Run(":8080") // Server chạy ở localhost:8080
}

// GET: Lấy danh sách sản phẩm
func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// POST: Thêm sản phẩm mới
func addProduct(c *gin.Context) {
	var newProduct Product

	// Parse body từ JSON
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tạo ID mới cho sản phẩm
	newProduct.ID = nextID
	nextID++

	// Lưu vào danh sách sản phẩm
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, newProduct)
}
