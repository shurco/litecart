package litepay

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findInSlice(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	cases := []struct {
		slice  []string
		value  string
		result bool
	}{
		{slice, "banana", true},
		{slice, "grape", false},
		{[]string{}, "apple", false},
	}

	for _, tt := range cases {
		found := findInSlice(tt.slice, tt.value)
		assert.Equal(t, tt.result, found)
	}
}

func Test_signMessage(t *testing.T) {
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	privKey1Bytes := x509.MarshalPKCS1PrivateKey(privKey)
	privKey1Pem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKey1Bytes,
	})
	privKey8Bytes, _ := x509.MarshalPKCS8PrivateKey(privKey)
	privKey8Pem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privKey8Bytes,
	})

	cases := []struct {
		message string
		privKey string
		err     error
	}{
		{"Hello, World!", string(privKey8Pem), nil},
		{"Hello, World!", string(privKey1Pem), errors.New("x509: failed to parse private key (use ParsePKCS1PrivateKey instead for this key format)")},
		{"Hello, World!", "", errors.New("invalid private key")},
	}

	for _, tt := range cases {
		_, err := signMessage(tt.message, tt.privKey)
		assert.Equal(t, tt.err, err)
	}
}

func Test_parseBody(t *testing.T) {
	cases := []struct {
		body     string
		expected map[string]any
		err      error
	}{
		{
			body:     `{"name": "John", "age": 30}`,
			expected: map[string]any{"name": "John", "age": float64(30)},
			err:      nil,
		},
		{
			body:     `{"key": "value"}`,
			expected: map[string]any{"key": "value"},
			err:      nil,
		},
		{
			body:     `{`,
			expected: nil,
			err:      errors.New("error decoding request body"),
		},
	}

	for _, tt := range cases {
		r := strings.NewReader(tt.body)
		data, err := parseBody(r)
		assert.Equal(t, tt.err, err)
		assert.Equal(t, tt.expected, data)
	}
}

func Test_statusPayment(t *testing.T) {
	cases := []struct {
		status   string
		expected Status
	}{
		{"pay", PAY},
		{"paid", PAY},
		{"unpaid", UNPAID},
		{"open", PROCESSED},
		{"complete", PAY},
		{"expired", CANCELED},
		{"requires_payment_method", FAILED},
		{"requires_confirmation", FAILED},
		{"requires_action", FAILED},
		{"processing", PROCESSED},
		{"requires_capture", PROCESSED},
		{"canceled", CANCELED},
		{"succeeded", PAY},
		{"2", PROCESSED},
		{"3", PAY},
		{"4", FAILED},
		{"5", FAILED},
		{"6", TEST},
		{"", FAILED},
	}

	for _, tt := range cases {
		result := StatusPayment(tt.status)
		assert.Equal(t, tt.expected, result)
	}
}
