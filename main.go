package main

import (
	"github.com/gin-gonic/gin"
)

// Movie结构体
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Tittle   string    `json:"tittle"`
	Director *Director `json:"director"`
}

// 导演结构体
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// 定义一个切片存储数据
var movies []Movie

func getMovies(c *gin.Context) {
}
func getMovie(c *gin.Context) {
}
func createMovie(c *gin.Context) {

}
func updateMovie(c *gin.Context) {

}
func deleteMovie(c *gin.Context) {

}
func main() {
	// 创建路由
	r := gin.Default()

	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovie)
	r.POST("/movies", createMovie)
	r.PUT("/movies/:id", updateMovie)
	r.DELETE("/movies/:id", deleteMovie)

	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "123",
		Tittle:   "美人鱼",
		Director: &Director{FirstName: "星驰", LastName: "周"},
	})
	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "456",
		Tittle:   "分手大师",
		Director: &Director{FirstName: "超", LastName: "邓"},
	})

}
