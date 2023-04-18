package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB = DB()

func AddBook(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("desc")
	cost := r.FormValue("cost")

	c, err := strconv.Atoi(cost)
    if err != nil {
        fmt.Println(err)
    }

	var response = JsonResponse{}
	if id == "" || title == "" || description == "" || c == 0{
        response = JsonResponse{Type: "error", Message: "No parameters"}
    } else {
		db.Create(&Book{Id: id, Title: title, Description: description, Cost: c})
		response = JsonResponse{Type: "success", Message: "You have successfully added!"}
	}
	json.NewEncoder(w).Encode(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var response = JsonResponse{}
	book := Book{}

	db.Where("id = ? and id != '' and deleted_at is null", id).Find(&book)

	if id == "" {
        response = JsonResponse{Type: "error", Message: "No id"}
    } else if book.Id != ""{
		response = JsonResponse{Type: "success", Message: "Books: " , Data: book}
	} else {
		response = JsonResponse{Type: "error", Message: "No book with such id"}
	}
	json.NewEncoder(w).Encode(response)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var response = JsonResponse{}
	book := []Book{}
	db.Find(&book)
	fmt.Println(book)
	response = JsonResponse{Type: "success", Message: "Books: " , Datas: book}
	json.NewEncoder(w).Encode(response)
}

func SortOrderByAsc(w http.ResponseWriter, r *http.Request){
	var response = JsonResponse{}
	book := []Book{}
	db.Order("cost asc").Find(&book)
	fmt.Println(book)
	response = JsonResponse{Type: "success", Message: "Books: ", Datas: book}
	json.NewEncoder(w).Encode(response)
}

func SortOrderByDesc(w http.ResponseWriter, r *http.Request){
	var response = JsonResponse{}
	book := []Book{}
	db.Order("cost desc").Find(&book)
	fmt.Println(book)
	response = JsonResponse{Type: "success", Message: "Books: ", Datas: book}
	json.NewEncoder(w).Encode(response)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	book := Book{}
	db.Where("id = ? and id != '' ", id).Find(&book)
	book.Title = r.FormValue("title")
	book.Description = r.FormValue("desc")
	db.Save(&book)
	fmt.Println(book)
	var response = JsonResponse{}

	if id == "" || book.Title == "" || book.Description == "" {
        response = JsonResponse{Type: "error", Message: "No parameters"}
	} else if book.Id != ""{
		response = JsonResponse{Type: "success", Message: "You have successfully updated!" , Data: book}
	} else {
		response = JsonResponse{Type: "error", Message: "No book with such id"}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var response = JsonResponse{}
	book := Book{}
	db.Where("id=? and id != '' ", id).Find(&book)
	deleted := book
	db.Delete(&book, id)
	fmt.Println(book)

	i, err := strconv.Atoi(id)
    if err != nil {
        panic(err)
    }

	if i == -1 {
        response = JsonResponse{Type: "error", Message: "No id"}
    } else if deleted.Id != "" {
		response = JsonResponse{Type: "success", Message: "You have successfully deleted!"}
	} else {
		response = JsonResponse{Type: "error", Message: "No book with such id"}
	}
	json.NewEncoder(w).Encode(response)
}

func SearchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	var response = JsonResponse{}
	book := []Book{}
	db.Where("title = ?", title).Find(&book)
	if title == "" {
        response = JsonResponse{Type: "error", Message: "No title"}
    } else if len(book) !=0 {
		response = JsonResponse{Type: "success", Message: "Found book with such title" , Datas: book}
	} else {
		response = JsonResponse{Type: "error", Message: "No book with such title"}
	}
	json.NewEncoder(w).Encode(response)
}