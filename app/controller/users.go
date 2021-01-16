package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"goSiteProjectDB/app/model"
	"html/template"
	"net/http"
	"path/filepath"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем пути к файлам с шаблонами
	main := filepath.Join("public", "html", "usersDynamicPage.html")
	common := filepath.Join("public", "html", "common.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем путь к файлу с шаблоном
	/*main := filepath.Join("public", "html", "usersDynamicPage.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	*/
	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра name, переданного в форме запроса
	name := r.FormValue("name")
	//получаем значение из параметра surname, переданного в форме запроса
	surname := r.FormValue("surname")

	//проверяем на пустые значения
	if name == "" || surname == "" {
		http.Error(rw, "Имя и фамилия не могут быть пустыми", 400)
		return
	}

	//создаем новый объект
	user := model.NewUser(name, surname)
	//записываем нового пользователя в таблицу БД
	err := user.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра userId, переданного в строке запроса
	userId := p.ByName("userId")

	//получаем пользователя из БД по его id
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//удаляем строку из таблицы
	err = user.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь был успешно удален")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра userId, переданного в строке запроса
	userId := p.ByName("userId")
	//получаем значения из параметров name и surname, переданных в форме запроса
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	//получаем пользователя из БД по его id
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//заменяем старые значения на новые
	user.Name = name
	user.Surname = surname

	//обновляем данные в таблице
	err = user.Update()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь был успешно изменен")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
