package controllers

import (
	"github.com/andre-haniak/go-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão de quantidade:", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para inteiro:", err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float64:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade para inteiro:", err)
		}

		models.UpdateProduct(convertedId, convertedQuantity, name, description, convertedPrice)
	}
	http.Redirect(w, r, "/", 301)
}
