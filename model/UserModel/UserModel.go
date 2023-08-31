package UserModel

//import (
//	"database/sql"
//	"go-web-employee/config"
//	"go-web-employee/entity"
//)
//
//type UserModel struct {
//	db *sql.DB
//}
//
//func NewUserModel() *UserModel {
//	conn, err := config.ConnectDB()
//
//	if err != nil {
//		panic(err)
//	}
//
//	return &UserModel{
//		db: conn,
//	}
//}
//
//func (u UserModel) Where(user *entity.User, fieldName, fieldValue string) error {
//
//	row, err := u.db.Query("select id, name, username, password, is_deleted from users where "+fieldName+" = ? limit 1", fieldValue)
//
//	if err != nil {
//		return err
//	}
//
//	defer row.Close()
//
//	for row.Next() {
//		row.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.IsDeleted)
//	}
//
//	return nil
//}
//
//func (u UserModel) Create(user entity.User) (int64, error) {
//
//	result, err := u.db.Exec("insert into t1_user (name, username, password,is_deleted) values(?,?,?,?)",
//		user.Name, user.Username, user.Password, 0)
//
//	if err != nil {
//		return 0, err
//	}
//
//	lastInsertId, _ := result.LastInsertId()
//
//	return lastInsertId, nil
//
//}
