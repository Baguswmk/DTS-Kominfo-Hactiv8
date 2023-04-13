package test

import (
	"DTS-Kominfo/Chapter3/Challange3/enums"
	"encoding/json"
	"go-jwt/controllers"
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	product := models.Product{
		Title:       "New Product",
		Description: "HEHEHEH",
	}
	db := database.GetDB()
	db.Create(&product)

	token := helpers.GenerateToken(product.ID, "user@example.com", enums.User)

	req, err := http.NewRequest("GET", "/products/"+strconv.Itoa(int(product.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/products/:productId", controllers.GetProductById)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, float64(product.ID), response["id"])
	assert.Equal(t, product.Title, response["name"])
	assert.Equal(t, product.Description, response["description"])
}


func TestGetProductByIdNotFound(t *testing.T) {
	product := models.Product{
		Title:       "New Product",
		Description: "HEHEHEH",
	}
	db := database.GetDB()
	db.Create(&product)

	token := helpers.GenerateToken(product.ID, "user@example.com", enums.User)


	req, err := http.NewRequest("GET", "/products/"+strconv.Itoa(int(product.ID+1)), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)


	rr := httptest.NewRecorder()

	
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		controllers.GetProductById(c)
	})


	handler.ServeHTTP(rr, req)


	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, "Not found", response["error"])
	assert.Equal(t, "Product not found", response["message"])
}


func TestGetAllProductFound(t *testing.T) {
	product1 := models.Product{
		Title:       "New Product 1",
		Description: "HEHEHEH 1",
	}
	product2 := models.Product{
		Title:       "New Product 2",
		Description: "HEHEHEH 2",
	}
	db := database.GetDB()
	db.Create(&product1)
	db.Create(&product2)

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()


	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		controllers.GetProductById(c)
	})

	handler.ServeHTTP(rr, req)


	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	
	var response []models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, product1.Title, response[0].Title)
	assert.Equal(t, product1.Description, response[0].Description)
	assert.Equal(t, product2.Title, response[1].Title)
	assert.Equal(t, product2.Description, response[1].Description)
}



func TestGetAllProductNotFound(t *testing.T) {
	db := database.GetDB()
	db.Exec("DELETE FROM products")

	token := helpers.GenerateToken(1, "user@example.com", enums.User)

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		c.Request = r
		controllers.GetAllProduct(c)
	}

	handler(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, "Not found", response["error"])
	assert.Equal(t, "No products found", response["message"])
}
