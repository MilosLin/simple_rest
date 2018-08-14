package fixme

import (
	"fmt"
	"log"
	"net/http"
	"simple_rest/api/protocol"
	"simple_rest/database"
	"simple_rest/env"

	"github.com/gin-gonic/gin"
)

// GetDepostitInput : Input參數
type GetDepostitInput struct {
	// 玩家ID
	UserID int32 `form:"UserID" binding:"required"`

	// 扣款餘額，若是要扣款 1 元，則代入 1
	Amount int32 `form:"Amount" binding:"required,min=0"`
}

// GetDepostitOutput : Output參數
type GetDepostitOutput struct {
	// 玩家ID，一個玩家的ID可對應至一筆錢包紀錄
	UserID int32 `json:"UserID"`

	// 錢包餘額
	Balance int32 `json:"Balance"`
}

// GetDepostit API
func GetDepostit(c *gin.Context) {
	res := &protocol.Response{}
	input := &GetDepostitInput{}
	output := &GetDepostitOutput{}

	// 綁定Input參數至結構中
	if err := c.Bind(input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// 從資料庫中取得目前錢包餘額
	w, err := getBalanceByID(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, protocol.SomethingWrongRes(err))
		return
	}

	if w == nil {
		c.JSON(http.StatusInternalServerError, protocol.SomethingWrongRes(
			fmt.Errorf("Wallet Not Found. UserID:%d", input.UserID),
		))
		return
	}

	// 若扣款後餘額非負數，才做扣款動作
	afterBalance := (w.Balance - input.Amount)
	if afterBalance >= 0 {
		if err := deposit(input.UserID, input.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, protocol.SomethingWrongRes(err))
			return
		}

		output.UserID = input.UserID
		output.Balance = afterBalance
	} else {
		output.UserID = input.UserID
		output.Balance = w.Balance
	}

	res.Result = output

	c.JSON(http.StatusOK, res)
	return
}

// Wallet : 錢包物件
type Wallet struct {
	ID      int32
	Balance int32
}

// getBalanceByID : 從資料庫中撈取錢包資料
func getBalanceByID(ID int32) (wallet *Wallet, err error) {
	fn := "getBalanceByID"

	dbS := database.GetConn(env.AccountDB)

	sql := " SELECT "
	sql += "   `id`, "
	sql += "   `balance` "
	sql += " FROM `account_db`.`wallet` "
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
		wallet = &Wallet{}
		if err := rows.Scan(
			&wallet.ID,
			&wallet.Balance,
		); err != nil {
			log.Fatalf("Fatch Data Error. fn:%s , err:%s", fn, err.Error())
			break
		}
	}

	return
}

// deposit : 從錢包扣款
func deposit(UserID, Amount int32) (err error) {
	fn := "deposit"

	dbM := database.GetConn(env.AccountDB)

	sql := " UPDATE `account_db`.`wallet`"
	sql += " SET `balance` = `balance` - ?"
	sql += " WHERE `id` = ? ;"

	var param []interface{}
	param = append(param, Amount)
	param = append(param, UserID)

	if _, err = dbM.Exec(sql, param...); err != nil {
		log.Fatalf("Exec Query Failed. fn:%s , err:%s", fn, err.Error())
		return err
	}

	return nil
}
