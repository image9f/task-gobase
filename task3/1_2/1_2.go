package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint `gorm:"primary_key;auto_increment"`
	Balance float64
}
type Transaction struct {
	ID            uint `gorm:"primary_key;auto_increment"`
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func transfer(db *gorm.DB, srcID, dstID uint, val float64) error {
	return db.Transaction(func(tx *gorm.DB) error {

		//查询转入账户
		var account Account
		if err := tx.First(&account, srcID).Error; err != nil {
			return fmt.Errorf("转出账户错误", err)
		}

		if account.Balance < val {
			return errors.New("账户余额不足")
		}

		//转账
		if err := tx.Model(&account).Where("id = ?", srcID).Update("balance", account.Balance-val).Error; err != nil {
			return fmt.Errorf("转账失败", err)
		}

		//查询接收转账账户
		var recvAccount Account
		if err := tx.First(&recvAccount, srcID).Error; err != nil {
			return fmt.Errorf("转入账户错误", err)
		}

		//接收转账
		if err := tx.Model(&recvAccount).Where("id=?", dstID).Update("balance", gorm.Expr("balance + ?", val)).Error; err != nil {
			return fmt.Errorf("接收转账失败", err)
		}

		//交易记录
		trans := Transaction{
			FromAccountId: srcID,
			ToAccountId:   dstID,
			Amount:        val,
		}

		if err := tx.Create(&trans).Error; err != nil {
			return fmt.Errorf("记录交易失败", err)
		}

		return nil

	})
}

func main() {
	fmt.Println("事务语句")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("open test.db err", err)
	}

	db.AutoMigrate(&Account{}, &Transaction{})

	//a := Account{ID: 1, Balance: 100}
	//b := Account{ID: 2, Balance: 200}
	//
	//db.Create(&a)
	//db.Create(&b)

	//余额充足转账
	fmt.Println("余额重组转账,a->b 50")
	err = transfer(db, 1, 2, 50)
	if err != nil {
		fmt.Println("转账失败", err)
	} else {
		fmt.Println("转账成功")
	}

	//余额不足转账
	fmt.Println("余额不足转账,a->b 500")
	err = transfer(db, 1, 2, 500)
	if err != nil {
		fmt.Println("转账失败", err)
	} else {
		fmt.Println("转账成功")
	}

	//查询所有交易信息
	fmt.Println("查询所有交易信息")
	var transaction []Transaction
	err = db.Find(&transaction).Error
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range transaction {
			fmt.Printf("%+v\n", v)
		}
	}

}
