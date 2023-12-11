package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/litepay"
)

type Event string

const (
	PAYMENT_INITIATION Event = "payment_initiation"
	PAYMENT_CALLBACK   Event = "payment_callback"
	PAYMENT_SUCCESS    Event = "payment_success"
	PAYMENT_CANCEL     Event = "payment_cancel"
	PAYMENT_ERROR      Event = "payment_error"
)

type Payment struct {
	Event     Event `json:"event"`
	TimeStamp int64 `json:"timestamp"`
	Data      Data  `json:"data"`
}

type Data struct {
	CartID        string                `json:"cart_id,omitempty"`
	PaymentSystem litepay.PaymentSystem `json:"payment_system"`
	PaymentStatus litepay.Status        `json:"payment_status"`
	TotalAmount   int                   `json:"total_amount,omitempty"`
	Currency      string                `json:"currency,omitempty"`
	CartItems     []litepay.Item        `json:"cart_items,omitempty"`
}

// SendPaymentHook is ...
func SendPaymentHook(resData *Payment) error {
	db := queries.DB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	webhookSetting, err := queries.GetSettingByGroup[models.Webhook](ctx, db)
	if err != nil {
		return err
	}

	if webhookSetting.Url != "" {
		jsonData, err := json.Marshal(resData)
		if err != nil {
			return err
		}

		errCh := make(chan error)
		go func() {
			defer close(errCh)

			res, err := Send(webhookSetting.Url, jsonData)
			if err != nil {
				errCh <- err
				return
			}
			if res.StatusCode != 200 {
				errCh <- fmt.Errorf("An issue has been identified with the payment webhook URL. Please verify that it responds with a status code of 200.")
				return
			}
		}()

		if err := <-errCh; err != nil {
			return err
		}
	}

	return nil
}
