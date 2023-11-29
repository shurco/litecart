package litepay

type PaymentSystem string

const (
	STRIPE      PaymentSystem = "stripe"
	PAYPAL      PaymentSystem = "paypal"
	SPECTROCOIN PaymentSystem = "spectrocoin"
)
