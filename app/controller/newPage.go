package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

//Функция «StartPage» реализует интерфейс хендлера, который обрабатывает запрос,
//приходящий по маршруту «/» и др. для запросов типа GET.

func NewPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//указываем путь к нужному файлу
	path := filepath.Join("public", "html", "index.html")
	common := filepath.Join("public", "html", "newHtml.html")
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//выводим шаблон клиенту в браузер
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
