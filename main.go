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
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
}
