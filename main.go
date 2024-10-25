package main

import "fmt"

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

func main() {
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
	fmt.Println(movies)
}
