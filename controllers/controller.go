package controllers

import (
	// "context"
	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(cont *gin.Context){
	db := connect()
	defer db.Close()

	query := "SELECT * FROM movies"
	
	title := cont.Request.URL.Query()["title"]

	if title != nil{
		query += " WHERE title ='" + title[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendResponse(cont, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	var movie Movie
	var movies []Movie

	for rows.Next() {
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Rating, &movie.Director); err != nil {
			log.Println(err)
			sendResponse(cont,400,"Error Scan Movie")
			return
		} else {
			movies = append(movies, movie)
		}
	}

	if len(movies) == 0 {
		sendResponse(cont,400,"Movie Not Found!!")
	}else{
		sendDataResponse(cont,200,"Success",movies)
	}
}

func InsertMovies(cont *gin.Context){
	db := connect()
	defer db.Close()

	err := cont.Request.ParseForm()
	if err != nil {
		sendResponse(cont,400,"Something Went Wrong")
		return
	}
	title := cont.Request.Form.Get("title")
	rating, _ := strconv.Atoi(cont.Request.Form.Get("rating"))
	director := cont.Request.Form.Get("director")
	_, errQuery := db.Exec("INSERT INTO movies(title, rating, director) values (?,?,?)", title, rating, director)
	var movie Movie
	movie.Title = title
	movie.Rating = rating
	movie.Director = director

	if errQuery == nil{
		sendDataResponse(cont,200,"Insert Success",movie)
	}else{
		sendResponse(cont, 400, "Insert Failed!")
	}
}

func UpdateMovies(cont *gin.Context){
	db := connect()
	defer db.Close()

	id := cont.Param("id")

	err := cont.Request.ParseForm()
	if err != nil {
		sendResponse(cont, 400, "Something Went Wrong!!")
		return
	}
	title := cont.Request.Form.Get("title")
	rating, _ := strconv.Atoi(cont.Request.Form.Get("rating"))
	director := cont.Request.Form.Get("director")

	_, errQuery := db.Exec("UPDATE movies SET title = ?, rating = ?, director = ? WHERE id = ?",title,rating,director,id)
	var movie Movie
	movie.Id, _ = strconv.Atoi(id)
	movie.Title = title
	movie.Rating = rating
	movie.Director = director
	if errQuery == nil {
		sendDataResponse(cont,200,"Update Success",movie)
	}else{
		log.Println(errQuery)
		sendResponse(cont,400,"Update Failed!")
	}
}

func DeleteMovies(cont *gin.Context){
	db := connect()
	defer db.Close()

	id := cont.Param("id")

	_, errQuery := db.Exec("DELETE FROM movies WHERE id=?",id)

	if errQuery == nil{
		sendResponse(cont,200,"Delete Success")
	}else{
		sendResponse(cont,400,"Delete Failed!")
	}
}