package process

import (
	"bytes"
	"encoding/csv"
	"stori/transaction-lambda/libs/transaction"
	"stori/transaction-lambda/libs/user"
	"strconv"
)

func ProcessData(data []byte, userService *user.UserConfig) ([]*transaction.TransactionRequest, error) {

	reader := bytes.NewReader(data)
	csvReader := csv.NewReader(reader)
	transactions := make([]*transaction.TransactionRequest, 0)
	records, err := csvReader.ReadAll()
	if err != nil {
		println("Error al leer el archivo csv: ", err.Error())
		return nil, err
	}

	// Itera sobre los registros
	for _, transactioSplit := range records {

		// postion 0 ID
		// Position 1 email user
		// Position 2 created at
		// Position 3 value

		value, err := strconv.ParseFloat(transactioSplit[3], 64)
		if err != nil {
			println("Error al convertir el valor a float64: ", err.Error())
			return nil, err
		}
		us := userService.GetUser(transactioSplit[1])
		if us == nil {
			println("Error al obtener el usuario")
		} else {
			transactions = append(transactions, transaction.NewTransactionRequest(us.ID, value, transactioSplit[2]))
		}
	}
	return transactions, nil
}
