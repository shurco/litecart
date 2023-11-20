package litepay

type Status string

const (
	NEW       Status = "new"
	UNPAID    Status = "unpaid"
	PAY       Status = "pay"
	CANCELED  Status = "canceled"
	FAILED    Status = "failed"
	PROCESSED Status = "processed"
	TEST      Status = "test"
)

type Cfg struct {
	paymentSystem PaymentSystem
	api           string   // API path
	currency      []string // support currency
	callbackURL   string
	successURL    string
	cancelURL     string
}

type LitePay interface {
	Pay(cart Cart) (*Payment, error)
	Checkout(payment *Payment, session string) (*Payment, error)
}

func New(callbackURL, successURL, cancelURL string) Cfg {
	return Cfg{
		callbackURL: callbackURL,
		successURL:  successURL,
		cancelURL:   cancelURL,
	}
}
