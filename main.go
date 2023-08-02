package main

import (
	"log"
	"net/http"

	"github.com/Glomen/goblog/app/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// запускаем роутер для обработки запросов
	r := httprouter.New()
	routes(r)
	// прикрепляем роутер к свободному порту
	// вторым параметром передаем роутер для работы с запросами
	err := http.ListenAndServe("localhost:4444", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	// путь к папке с внешними файлами
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	// выполняем запросы по указанным адресам
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
}
