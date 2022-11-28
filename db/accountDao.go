package db

import (
	"database/sql"
	"fmt"
	"time"
)

type AccountDao struct {
	ID                 int64          `gorm:"column:id;primaryKey"`
	AccountType        string         `gorm:"column:account_type"`
	Currency           string         `gorm:"column:currency"`
	OpeningBalance     string         `gorm:"column:opening_balance;type:decimal(20,4)"`
	BaseOpeningBalance string         `gorm:"column:base_opening_balance;type:decimal(20,4)"`
	AccountClass       uint8          `gorm:"column:account_class"`
	AccountName        string         `gorm:"column:account_name;unique_index"`
	Description        sql.NullString `gorm:"column:description;type:text"`
	Created            time.Time      `gorm:"column:created;autoCreateTime"`
	Updated            time.Time      `gorm:"column:updated;autoUpdateTime"`
	CreatedTimeStamp   int64          `gorm:"column:created_timestamp;"`
	UpdatedTimeStamp   int64          `gorm:"column:updated_timestamp;"`
	CreatedBy          int64          `gorm:"column:created_by"`
	UpdatedBy          int64          `gorm:"column:updated_by"`
	Archived           sql.NullTime   `gorm:"column:archived"`
	Version            int32          `gorm:"column:version"`
}

type FinanceAccountDao struct {
	ID              int64         `gorm:"column:id;primaryKey;"`
	AccountNumber   string        `gorm:"column:account_number"`
	SubAccountType  string        `gorm:"colum:sub_account_type"`
	LinkedAccountID sql.NullInt64 `gorm:"column:linked_account_id"`
	EntityID        int64         `gorm:"column:entity_id"`
	TaxID           sql.NullInt64 `gorm:"column:tax_id"`
	IsDetail        bool          `gorm:"column:is_detail"`
	ParentID        sql.NullInt64 `gorm:"column:parent_id"`
	Directory       string        `gorm:"column:directory"`
}

type FinanceEntityDao struct {
	ID                  int64  `gorm:"column:id;primaryKey;"`
	Name                string `gorm:"column:name"`
	BaseCurrency        string `gorm:"column:base_currency"`
	ShortName           string `gorm:"column:short_name"`
	FinancialStartMonth int    `gorm:"column:financial_start_month"`
	FinancialOffet      int    `gorm:"column:financial_offet"`
}

func (e FinanceEntityDao) TableName() string {
	return "finance_entity"
}

func (a AccountDao) TableName() string {
	return "account"
}

func (f FinanceAccountDao) TableName() string {
	return "finance_account"
}

func SaveAll(accounts []AccountDao, faccounts []FinanceAccountDao) {
	tx := DB.Begin()
	err := tx.Create(&accounts)
	if err.Error != nil {
		tx.Rollback()
		panic(err.Error)
	}
	err = tx.Create(&faccounts)
	if err.Error != nil {
		tx.Rollback()
		panic(err.Error)
	}
	tx.Commit()
}

func SaveAccount(account AccountDao) int64 {
	DB.Create(&account)
	return account.ID
}

func SaveFinanceAccout(financeAccount FinanceAccountDao) {
	DB.Create(&financeAccount)
}

func FindByName(name string) {
	if name == "" {
		return
	}
	// fmt.Println(name)
	account := &AccountDao{}
	DB.Where("account_name=? and archived is null", name).Find(&account)
	if account.AccountName == "" {
		fmt.Println(name)
	}
}

func FindEntityById(id int64) *FinanceEntityDao {
	entity := &FinanceEntityDao{}
	err := DB.First(&entity, id)
	if err.Error != nil {
		return nil
	}
	return entity
}
