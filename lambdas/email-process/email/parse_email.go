package email

import (
	"bytes"
	"fmt"
	"html/template"
	"stori/email-service/transaction"
	"stori/email-service/user"
	"strconv"
	"time"
)

func makeTable(movements []*transaction.MovementResponse) string {

	result := ""
	tr := `<tr style="margin-top:1em">
		<td>
			<div class style="box-sizing: border-box;font-size: 12px;font-family: Arial, Helvetica Neue, Helvetica, sans-serif;mso-line-height-alt: 18px;color: #555;line-height: 1.5;">
				<p style="box-sizing: border-box;line-height: inherit;font-family: sans-serif;font-weight: normal;margin-bottom: 15px;margin: 0;font-size: 14px; text-align: center; mso-line-height-alt: 27px;">
					<strong style="box-sizing: border-box; font-size: 18px;">Month %s</strong> 
				</p>
			</div>
		</td>
		<td>
			<div class style="box-sizing: border-box;font-size: 12px;font-family: Arial, Helvetica Neue, Helvetica, sans-serif;mso-line-height-alt: 18px;color: #555;line-height: 1.5;">
				<p style="box-sizing: border-box;line-height: inherit;font-family: sans-serif;font-weight: normal;margin-bottom: 15px;margin: 0;font-size: 14px; text-align: center; mso-line-height-alt: 27px;">
					<strong style="box-sizing: border-box; font-size: 18px;">Number of transactions credit </strong> <span>%s</span>
				</p>
				<p style="box-sizing: border-box;line-height: inherit;font-family: sans-serif;font-weight: normal;margin-bottom: 15px;margin: 0;font-size: 14px; text-align: center; mso-line-height-alt: 27px;">
					<strong style="box-sizing: border-box; font-size: 18px;">Number of transactions debit </strong> <span>%s</span>
				</p>
			</div>
		</td>
	</tr>`
	for _, movement := range movements {
		result += fmt.Sprintf(tr, time.Month(movement.Month), strconv.Itoa(movement.Increment), strconv.Itoa(movement.Decrement))
	}
	return result
}

func ParseEmail(movements []*transaction.MovementResponse, balance *transaction.BalanceResponse, user *user.UserData) (string, error) {

	data := map[string]interface{}{
		"Name":      user.FirstName + " " + user.Surname,
		"Balance":   strconv.FormatFloat(balance.Balance, 'f', 2, 64),
		"Credit":    strconv.FormatFloat(balance.Credit, 'f', 2, 64),
		"Debit":     strconv.FormatFloat(balance.Debit, 'f', 2, 64),
		"Movements": template.HTML(makeTable(movements)),
	}

	tmpl, err := template.ParseFiles("./templates/email.html")
	if err != nil {
		fmt.Println("Error al cargar el template:", err)
		return "", err
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		fmt.Println("Error al ejecutar el template:", err.Error())
		return "", err
	}
	return tpl.String(), nil
}
