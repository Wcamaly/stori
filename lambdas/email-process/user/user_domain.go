package user

import "encoding/json"

type UserData struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
}

func ProcessJSONData(payload interface{}) []*UserData {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		println("Error al convertir interface{} a JSON: %v", err)
		return nil
	}

	println("DATA: ", jsonData)

	var users []*UserData

	// Deserializa el JSON a []UserData
	if err := json.Unmarshal(jsonData, &users); err != nil {
		println("Error deserializando JSON: %v\n", err)
		return nil
	}

	return users
}
