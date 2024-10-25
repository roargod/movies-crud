package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
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
	c.JSON(http.StatusOK, movies)
}
func getMovie(c *gin.Context) {
	// 切出其中的id字段
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "无效的id"})
		return
	}
	// 遍历寻找该电影id
	for i := 0; i < len(movies); i++ {
		// 找到就返回该电影
		if movies[i].ID == id {
			c.JSON(http.StatusOK, movies[i])
			return
		}
	}
	// 走到这，没找到
	c.JSON(http.StatusOK, "没找到对应电影")
}
func createMovie(c *gin.Context) {
	var movie Movie
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	c.JSON(http.StatusOK, movie)
}
func updateMovie(c *gin.Context) {

}
func deleteMovie(c *gin.Context) {
	// 切出其中的id字段
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "无效的id"})
		return
	}
	// 遍历寻找该电影id
	for i := 0; i < len(movies); i++ {
		// 找到就删除该电影
		if movies[i].ID == id {
			//c.JSON(http.StatusOK, movies[i])
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	c.JSON(http.StatusOK, movies)
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
