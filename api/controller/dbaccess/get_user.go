package dbaccess

import (
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

// GetUserInput : Input參數
type GetUserInput struct {
	UserID int32 `form:"UserID"`
}

// GetUser API
func GetUser(c *gin.Context) {
	res := &protocol.Response{}
	input := &GetUserInput{}

	// 綁定Input參數至結構中
	if err := c.Bind(input); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	u, err := getUserByID(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, protocol.SomethingWrongRes(err))
		return
	}

	res.Result = u

	c.JSON(http.StatusOK, res)
	return
}

// User : 使用者物件
type User struct {
	ID       int
	Account  string
	Password string
}

// getUserByID : 從資料庫中撈取使用者資料
func getUserByID(ID int32) (user *User, err error) {
	fn := "getUserByID"

	dbS := database.GetConn(env.AccountDB)

	sql := " SELECT "
	sql += "   `id`, "
	sql += "   `account`, "
	sql += "   `password` "
	sql += " FROM `account_db`.`user` "
	sql += " WHERE `id` = ? ;"

	var params []interface{}
	params = append(params, ID)

	rows, err := dbS.Query(sql, params...)
	if err != nil {
		log.Fatalf("Exec Query Failed. fn:%s , err:%s", fn, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		user = &User{}
		if err := rows.Scan(
			&user.ID,
			&user.Account,
			&user.Password,
		); err != nil {
			log.Fatalf("Fatch Data Error. fn:%s , err:%s", fn, err.Error())
			break
		}
	}

	return
}
