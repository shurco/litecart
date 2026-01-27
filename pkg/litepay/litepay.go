// Package litepay provides a unified interface for working with various payment providers.
// It supports Stripe, PayPal, SpectroCoin, and a dummy provider for testing.
//
// Example usage:
//
//	pay := litepay.New(callbackURL, successURL, cancelURL)
//	stripe := pay.Stripe("sk_test_...")
//	payment, err := stripe.Pay(cart)
package litepay

// Status represents the internal payment status.
type Status string

const (
	NEW       Status = "new"       // Payment has been created
	UNPAID    Status = "unpaid"    // Payment is awaiting payment
	PAID      Status = "paid"      // Payment has been successfully completed (final)
	CANCELED  Status = "canceled"  // Payment has been canceled (final)
	FAILED    Status = "failed"    // Payment has failed (final)
	PROCESSED Status = "processed" // Payment is being processed
	TEST      Status = "test"      // Test payment
)

// Cfg holds the configuration for payment providers.
// This structure is used internally and configured via the New function.
type Cfg struct {
	paymentSystem PaymentSystem
	api           string   // API endpoint URL
	currency      []string // supported currencies
	callbackURL   string   // webhook callback URL
	successURL    string   // redirect URL on success
	cancelURL     string   // redirect URL on cancellation
}

// LitePay is the main interface that all payment providers must implement.
// It provides methods for creating payment sessions and checking payment status.
type LitePay interface {
	// Pay creates a payment session with the provider and returns payment information
	// including a URL to redirect the user to complete the payment.
	Pay(cart Cart) (*Payment, error)

	// Checkout verifies and updates the payment status with the provider.
	// The session parameter contains the provider-specific session ID.
	Checkout(payment *Payment, session string) (*Payment, error)
}

// New creates a new payment configuration with callback URLs.
//
// Parameters:
//   - callbackURL: URL for receiving webhook notifications from payment providers
//   - successURL: URL to redirect users after successful payment
//   - cancelURL: URL to redirect users if payment is canceled
//
// Example:
//
//	pay := litepay.New(
//		"https://example.com/payment/callback",
//		"https://example.com/payment/success",
//		"https://example.com/payment/cancel",
//	)
func New(callbackURL, successURL, cancelURL string) Cfg {
	return Cfg{
		callbackURL: callbackURL,
		successURL:  successURL,
		cancelURL:   cancelURL,
	}
}
