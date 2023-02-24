package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Caknoooo/to-do-list-api/entities"
	"github.com/Caknoooo/to-do-list-api/models"
)

var userModel = models.NewUserModel()

func Get(response http.ResponseWriter, request *http.Request){
	// Untuk mengambil id dari endpoint atau parameter http request
	id, err := strconv.Atoi(request.URL.Path[len("/api/getData/"):]) // id -> int
	// id, err := strconv.Atoi(request.URL.Query().Get("id"))
	// fmt.Print(strconv.Atoi(request.URL.Path[len("/api/getData/"):]))
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// Ambil user dengan id tersebut dari database
	data, err := userModel.GetUserById(id)
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// Kirim data kepada user dalam bentuk JSON sebagai response nya
	user := map[string]interface{}{
		"data": data,
		"message": "Berhasil mendapatkan ID tersebut",
	}

	b, err := json.Marshal(user)
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-type", "application/json") // 
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, string(b))
}


func GetAll(response http.ResponseWriter, request *http.Request){
	user, err := userModel.GetAllUser()

	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"user": user,
		"message": "Berhasil",
	}

	b, err := json.Marshal(data);
	if err != nil{
		fmt.Println(err)
	}

	response.Header().Set("Content-type", "application/json");
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, string(b))
}	

func Create(response http.ResponseWriter, request *http.Request){
	// Ambil data dari request body
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user) // Untuk menampung data dari request postman
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic untuk aplikasi utama, gaada cek waktu mulai ama selesai ada batasan atau gimana, bisa masukkin todonya, Buat logicnya jika gabut
	// Menambahkan data user ke database
	err = userModel.Insert(user) // Udah ada data dimasukkan ke dalam Insert
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"message": "Berhasil Menambahkan data",
	}

	b, err := json.Marshal(data)
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, string(b))
}

func Update(response http.ResponseWriter, request *http.Request){
	// Mengambil id yang akan di update
	id, err := strconv.Atoi(request.URL.Path[len("/api/updateData/"):]) // int 
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// Ambil data user dari request body
	var user entities.User
	err = json.NewDecoder(request.Body).Decode(&user) // 
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user di database
	err = userModel.Update(id, user)
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Kirim pesan sebagai response
	data := map[string]interface{}{
		"message": "Data Berhasil update",
	}

	b, err := json.Marshal(data)
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, string(b))
}

func Delete(response http.ResponseWriter, request *http.Request){
	// Mengambil data dengan Id yang akan di delete
	id, err := strconv.Atoi(request.URL.Path[len("/api/deleteData/"):]) // 1
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// hapus user dari database
	err = userModel.Delete(id) // 1
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// Kirim pesan berhasil sebagai response
	data := map[string]interface{}{
		"message": "Berhasil menghapus data",
	}

	b, err := json.Marshal(data)
	if err != nil{
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, string(b))
}