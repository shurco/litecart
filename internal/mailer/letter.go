package mailer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
)

// SendTestLetter is ...
func SendTestLetter(letterName string) error {
	db := queries.DB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mailSetting, err := queries.GetSettingByGroup[models.Mail](ctx, db)
	if err != nil {
		return err
	}

	settingEmail, err := db.GetSettingByKey(ctx, "email", letterName)
	if err != nil {
		return err
	}

	letter := &models.MessageMail{
		To: settingEmail[0].Value.(string),
		Letter: models.Letter{
			Subject: "litecart test smtp settings",
			Text:    "test message",
		},
		Data: map[string]string{
			"Payment_URL":    "https://payment.com/order/1234567890",
			"Admin_Email":    "Admin Name <admin@mail.com>",
			"Site_Name":      "Site name",
			"Amount_Payment": "21.00 USD",
		},
	}

	if letterName != "smtp" {
		if err := json.Unmarshal([]byte(settingEmail[1].Value.(string)), &letter.Letter); err != nil {
			return err
		}
	}

	if err := SendMail(mailSetting, letter); err != nil {
		return err
	}

	return nil
}

// SendPrepaymentLetter is ...
func SendPrepaymentLetter(email, amountPayment, paymentURL string) error {
	db := queries.DB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	letter, err := db.CartLetterPayment(ctx, email, amountPayment, paymentURL)
	if err != nil {
		return err
	}

	mailSetting, err := queries.GetSettingByGroup[models.Mail](ctx, db)
	if err != nil {
		return err
	}

	if err := SendMail(mailSetting, letter); err != nil {
		return err
	}

	return nil
}

// SendCartLetter is ...
func SendCartLetter(cartID string) error {
	db := queries.DB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	letter, err := db.CartLetterPurchase(ctx, cartID)
	if err != nil {
		return err
	}

	mailSetting, err := queries.GetSettingByGroup[models.Mail](ctx, db)
	if err != nil {
		return err
	}

	if err := SendMail(mailSetting, letter); err != nil {
		return err
	}

	return nil
}
