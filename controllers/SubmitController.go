package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	// "os"
	"log"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SubmitController struct {
	beego.Controller
}

func (c *SubmitController) Get() {
	var name = c.GetString("fullName");
	var email = c.GetString("email");
	var country = c.GetString("country");

	fmt.Println("=====================");
	fmt.Println(name);
	fmt.Println(email);
	fmt.Println(country);
	fmt.Println("=====================");
	
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "New Member Contacted"
	to := mail.NewEmail("Example User","amit.iitr.cs@gmail.com")
	plainTextContent := "Here are the Details"
	htmlContent := "<strong>DETAILS :</strong>" + name +"<br>" + email +"<br> "+ country 
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	c.TplName = "submit.tpl"
}

