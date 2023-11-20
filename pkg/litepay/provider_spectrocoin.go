package litepay

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type CallbackSpectrocoin struct {
	MerchantID      int     `json:"merchantId" form:"merchantId"`
	ApiID           int     `json:"apiId" form:"apiId"`
	UserID          string  `json:"userId" form:"userId"`
	MerchantApiID   string  `json:"merchantApiId" form:"merchantApiId"`
	OrderID         string  `json:"orderId" form:"orderId"`
	PayCurrency     string  `json:"payCurrency" form:"payCurrency"`
	PayAmount       float64 `json:"payAmount" form:"payAmount"`
	ReceiveCurrency string  `json:"receiveCurrency" form:"receiveCurrency"`
	ReceiveAmount   float64 `json:"receiveAmount" form:"receiveAmount"`
	ReceivedAmount  int     `json:"receivedAmount" form:"receivedAmount"`
	Description     string  `json:"description" form:"description"`
	OrderRequestID  int     `json:"orderRequestId" form:"orderRequestId"`
	Status          int     `json:"status" form:"status"`
	Sign            string  `json:"sign" form:"sign"`
}

type spectrocoin struct {
	Cfg
	merchantID string
	projectID  string
	privateKey string
}

func (c Cfg) Spectrocoin(merchantID, projectID, privateKey string) LitePay {
	c.paymentSystem = SPECTROCOIN
	c.api = "https://spectrocoin.com"
	c.currency = []string{"EUR", "USD", "GBP", "AUD", "CAD", "JPY", "CNY", "SEK"}
	return &spectrocoin{
		Cfg:        c,
		merchantID: merchantID,
		projectID:  projectID,
		privateKey: privateKey,
	}
}

func (c *spectrocoin) Pay(cart Cart) (*Payment, error) {
	var totalAmount float64
	receiveCurrency := strings.ToUpper(cart.Currency)

	if !findInSlice(c.currency, receiveCurrency) {
		return nil, errors.New("this currency is not supported")
	}

	for _, s := range cart.Items {
		totalAmount += float64(s.PriceData.UnitAmount) / 100 * float64(s.Quantity)
	}

	_receiveAmount := fmt.Sprintf("%.2f", totalAmount)
	_receiveAmount = strings.ReplaceAll(_receiveAmount, ".00", ".0")

	body := "userId=" + c.merchantID +
		"&merchantApiId=" + c.projectID +
		"&orderId=" + cart.ID +
		"&payCurrency=BTC" +
		"&payAmount=0.0" +
		"&receiveCurrency=" + receiveCurrency +
		"&receiveAmount=" + _receiveAmount +
		"&description=" +
		"&payerEmail=" +
		"&payerName=" +
		"&payerSurname=" +
		"&culture=" +
		"&callbackUrl=" + url.QueryEscape(fmt.Sprintf("%s/?payment_system=%s&cart_id=%s", c.callbackURL, c.paymentSystem, cart.ID)) +
		"&successUrl=" + url.QueryEscape(fmt.Sprintf("%s/?payment_system=%s&cart_id=%s", c.successURL, c.paymentSystem, cart.ID)) +
		"&failureUrl=" + url.QueryEscape(fmt.Sprintf("%s/?payment_system=%s&cart_id=%s", c.cancelURL, c.paymentSystem, cart.ID))

	signature, err := signMessage(body, c.privateKey)
	if err != nil {
		return nil, err
	}
	body += "&sign=" + url.QueryEscape(signature)

	resp, err := http.Post(fmt.Sprintf("%s/api/merchant/1/createOrder", c.api), "application/x-www-form-urlencoded", bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := parseBody(resp.Body)
	if err != nil {
		return nil, err
	}

	receiveAmount, _ := strconv.ParseFloat(data["receiveAmount"].(string), 64)
	checkout := &Payment{
		AmountTotal:   int(receiveAmount * 100),
		Currency:      data["receiveCurrency"].(string),
		Status:        PROCESSED,
		URL:           data["redirectUrl"].(string),
		PaymentSystem: c.paymentSystem,
	}

	return checkout, nil
}

func (c *spectrocoin) Checkout(payment *Payment, session string) (*Payment, error) {
	return nil, nil
}
