package models

import(
	"database/sql"

	"github.com/Caknoooo/to-do-list-api/config"
	"github.com/Caknoooo/to-do-list-api/entities"
)

type UserModel struct{
	connection *sql.DB
}

func NewUserModel() *UserModel{
	conn, err := config.ConnectDatabase()
	if err != nil{
		panic(err) 
	}

	return &UserModel{
		connection: conn,
	}
}

func (u *UserModel) GetUserById(id int)(entities.User, error){
	var user entities.User

	row := u.connection.QueryRow("SELECT * FROM to_do_list WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Nama, &user.Aktifitas, &user.Mulai, &user.Selesai, &user.Done)
	if err != nil{
		return entities.User{}, err
	} // 

	return user, nil
}

func (u *UserModel) GetAllUser()([]entities.User, error){
	var result []entities.User

	rows, err := u.connection.Query("SELECT * FROM to_do_list") // 
	if err != nil{
		return nil, err
	}

	for rows.Next(){
		var user entities.User
		rows.Scan(&user.Id, &user.Nama, &user.Aktifitas, &user.Mulai, &user.Selesai, &user.Done)
		result = append(result, user)
	}
	
	return result, nil
}

func (u *UserModel) Insert(user entities.User) error{
	_, err := u.connection.Exec("INSERT INTO to_do_list(nama, aktifitas, mulai, selesai, isdone) VALUES($1, $2, $3, $4, $5)", user.Nama, user.Aktifitas, user.Mulai, user.Selesai, user.Done)
	if err != nil{
		return err
	}
	// 
	return nil
} 

func (u *UserModel) Update(id int, user entities.User) error{
	_, err := u.connection.Exec("UPDATE to_do_list SET nama = $1, aktifitas = $2, mulai = $3, selesai = $4, isdone = $5 WHERE id = $6", user.Nama, user.Aktifitas, user.Mulai, user.Selesai, user.Done, id)
	if err != nil{
		return err
	}

	return nil
}

func (u *UserModel) Delete(id int) error{
	_, err := u.connection.Exec("DELETE FROM to_do_list WHERE id = $1", id)
	if err != nil{
		return err
	}

	return nil
} 