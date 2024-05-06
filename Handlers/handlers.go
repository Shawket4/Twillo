package Handlers

import (
	"fmt"
	"log"
	"net/http"

	"Twillo/Models"
	"Twillo/Regex"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func EchoJSON(c *gin.Context) {
	data := c.Query("message")

	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"message": "Received"})
}

func RegisterTransaction(c *gin.Context) {
	data := c.Query("message")
	variables := Regex.ParseTransactionMessage(data)
	fmt.Println(variables)
	var transaction Models.Transaction
	place := Regex.ParseTransactionMessagePlace(data)
	for i, variable := range variables {
		switch i {
		case 0:
			transaction.Amount = variable
		case 1:
			transaction.CardShortNo = variable
		case 2:
			transaction.Place = place[0]
		case 3:
			transaction.Date = variable
		case 4:
			transaction.Time = variable
		case 5:
			transaction.BalanceAvailable = variable
		}
	}
	fmt.Println(transaction)
	f, err := excelize.OpenFile("Transactions.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, _ := f.GetRows("Sheet1")
	index := len(rows)
	f.SetCellValue("Sheet1", fmt.Sprintf("A%v", index+1), transaction.Date)
	f.SetCellValue("Sheet1", fmt.Sprintf("B%v", index+1), transaction.Time)
	f.SetCellValue("Sheet1", fmt.Sprintf("C%v", index+1), transaction.Amount)
	f.SetCellValue("Sheet1", fmt.Sprintf("D%v", index+1), transaction.CardShortNo)
	f.SetCellValue("Sheet1", fmt.Sprintf("E%v", index+1), transaction.Place)
	f.SetCellValue("Sheet1", fmt.Sprintf("F%v", index+1), transaction.BalanceAvailable)
	if err := f.SaveAs("Transactions.xlsx"); err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, nil)
}

func RegisterInstapay(c *gin.Context) {
	data := c.Query("message")
	var output Models.Instapay
	variables := Regex.ParseInstapayMessage(data, "Al Ahly")
	if variables == nil {
		variables = Regex.ParseInstapayMessage(data, "SAIB")
		for i, variable := range variables {
			switch i {
			case 0:
				output.Amount = variable
			case 1:
				output.Date = variable
			}
		}
	} else {
		for i, variable := range variables {
			switch i {
			case 0:
				output.CardShortNo = variable
			case 1:
				output.Date = variable
			case 2:
				output.Time = variable
			case 3:
				output.Amount = variable
			}
		}
	}

	fmt.Println(output)
	f, err := excelize.OpenFile("Instapay.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, _ := f.GetRows("Instapay")
	index := len(rows)
	f.SetCellValue("Instapay", fmt.Sprintf("A%v", index+1), output.Date)
	f.SetCellValue("Instapay", fmt.Sprintf("B%v", index+1), output.CardShortNo)
	f.SetCellValue("Instapay", fmt.Sprintf("C%v", index+1), output.Amount)
	f.SetCellValue("Instapay", fmt.Sprintf("D%v", index+1), output.Time)
	f.SetCellValue("Instapay", fmt.Sprintf("E%v", index+1), output.Date)
	if err := f.SaveAs("Instapay.xlsx"); err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, nil)
}

func RegisterInstapayNew(c *gin.Context) {
	var message Models.MessageReceived
	if err := c.ShouldBindJSON(&message); err != nil {
		log.Println(err)
	}
	f, err := excelize.OpenFile("Instapay.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, _ := f.GetRows("Instapay")
	index := len(rows)
	f.SetCellValue("Instapay", fmt.Sprintf("A%v", index+1), message.DateTime)
	f.SetCellValue("Instapay", fmt.Sprintf("B%v", index+1), message.Card)
	f.SetCellValue("Instapay", fmt.Sprintf("C%v", index+1), message.Amount)
	f.SetCellValue("Instapay", fmt.Sprintf("D%v", index+1), message.Notes)
	f.SetCellValue("Instapay", fmt.Sprintf("E%v", index+1), message.Bank)
	if err := f.SaveAs("Instapay.xlsx"); err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, nil)
}
