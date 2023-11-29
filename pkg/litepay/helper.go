package litepay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
)

func findInSlice(slice []string, value string) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}

func signMessage(message, privKey string) (string, error) {
	block, _ := pem.Decode([]byte(privKey))
	if block == nil {
		return "", errors.New("invalid private key")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("key is not a valid RSA private key")
	}

	hash := sha1.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hash[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func parseBody(r io.Reader) (map[string]any, error) {
	var data map[string]any

	body, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.New("error reading request body")
	}

	if len(body) > 0 {
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, errors.New("error decoding request body")
		}
	}

	return data, nil
}

// StatusPayment is ...
func StatusPayment(system PaymentSystem, status string) Status {
	statusBase := map[string]Status{}

	switch system {
	case STRIPE:
		statusBase = map[string]Status{
			"pay":                     PAID,
			"paid":                    PAID,
			"unpaid":                  UNPAID,
			"open":                    PROCESSED,
			"complete":                PAID,
			"expired":                 CANCELED,
			"requires_payment_method": FAILED,
			"requires_confirmation":   FAILED,
			"requires_action":         FAILED,
			"processing":              PROCESSED,
			"requires_capture":        PROCESSED,
			"canceled":                CANCELED,
			"succeeded":               PAID,
		}

	case PAYPAL:
		statusBase = map[string]Status{
			"CREATED":               PROCESSED,
			"SAVED":                 PROCESSED,
			"APPROVED":              PROCESSED,
			"VOIDED":                CANCELED,
			"COMPLETED":             PAID,
			"PAYER_ACTION_REQUIRED": PROCESSED,
		}

	case SPECTROCOIN:
		statusBase = map[string]Status{
			"1": UNPAID,    // new
			"2": PROCESSED, // pending, Payment (or part of it) was received, but still waiting for confirmation
			"3": PAID,      // paid, Order is completed
			"4": FAILED,    // failed, Some error occurred
			"5": FAILED,    // expired, Payment was not received in time
			"6": TEST,      // test, Test order
		}
	}

	statusTmp := statusBase[status]
	if statusTmp == "" {
		statusTmp = FAILED
	}

	return statusTmp
}
