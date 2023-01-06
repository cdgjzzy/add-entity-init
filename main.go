package main

import (
	"add-entity-init/db"
	"fmt"
	"time"
)

func main() {
	entity_id := 0
	fmt.Print("Please input the new entity id:")
	fmt.Scanf("%d", &entity_id)
	if entity_id == 0 {
		panic("entity_id can not be zero.")
	}
	entity := db.FindEntityById(int64(entity_id))
	if entity == nil {
		panic("entity is not exist.")
	}

	// assets
	account := &db.AccountDao{
		AccountType:        "ASSET",
		OpeningBalance:     "0.00",
		BaseOpeningBalance: "0.00",
		AccountClass:       0,
		AccountName:        "Assets",
		Created:            time.Now(),
		Updated:            time.Now(),
		CreatedTimeStamp:   time.Now().UnixMilli(),
		UpdatedTimeStamp:   time.Now().UnixMilli(),
		CreatedBy:          1,
		UpdatedBy:          1,
		Version:            1,
	}
	id := db.SaveAccount(*account)
	financeAccount := &db.FinanceAccountDao{
		ID:             id,
		AccountNumber:  "1-0000",
		SubAccountType: "Asset",
		EntityID:       int64(entity_id),
		IsDetail:       false,
		Directory:      "",
	}
	db.SaveFinanceAccout(*financeAccount)

	// LIABILITY

	account.ID = 0
	account.AccountType = "LIABILITY"
	account.AccountName = "Liabilities"
	id = db.SaveAccount(*account)
	financeAccount.ID = id
	financeAccount.AccountNumber = "2-0000"
	financeAccount.SubAccountType = "Liability"
	db.SaveFinanceAccout(*financeAccount)

	// EQUITY
	account.ID = 0
	account.AccountType = "EQUITY"
	account.AccountName = "Equity"
	id = db.SaveAccount(*account)
	financeAccount.ID = id
	financeAccount.AccountNumber = "3-0000"
	financeAccount.SubAccountType = "Equity"
	db.SaveFinanceAccout(*financeAccount)

	// INCOME
	account.ID = 0
	account.AccountType = "INCOME"
	account.AccountName = "Income"
	id = db.SaveAccount(*account)
	financeAccount.ID = id
	financeAccount.AccountNumber = "4-0000"
	financeAccount.SubAccountType = "Income"
	db.SaveFinanceAccout(*financeAccount)

	// COST_OF_SALES

	account.ID = 0
	account.AccountType = "COST_OF_SALES"
	account.AccountName = "Cost of Sales"
	id = db.SaveAccount(*account)
	financeAccount.ID = id
	financeAccount.AccountNumber = "5-0000"
	financeAccount.SubAccountType = "Cost of Sales"
	db.SaveFinanceAccout(*financeAccount)

	// EXPENSE

	account.ID = 0
	account.AccountType = "EXPENSE"
	account.AccountName = "Expenses"
	id = db.SaveAccount(*account)
	financeAccount.ID = id
	financeAccount.AccountNumber = "6-0000"
	financeAccount.SubAccountType = "Expense"
	db.SaveFinanceAccout(*financeAccount)

	// OTHER_INCOME

	account.ID = 0
	account.AccountType = "OTHER_INCOME"
	account.AccountName = "Other Income"
	id = db.SaveAccount(*account)
	financeAccount.ID = id
	financeAccount.AccountNumber = "8-0000"
	financeAccount.SubAccountType = "Other Income"
	db.SaveFinanceAccout(*financeAccount)

	// OTHER_EXPENSE
	account.ID = 0
	account.AccountType = "OTHER_EXPENSE"
	account.AccountName = "Other Expenses"
	id = db.SaveAccount(*account)
	financeAccount.ID = id
	financeAccount.AccountNumber = "9-0000"
	financeAccount.SubAccountType = "Other Expense"
	db.SaveFinanceAccout(*financeAccount)

}
