package integrations

import (
	"errors"
	"log"
)

// Nhiệm vụ: Tích hợp thanh toán qua Stripe
// Liên kết:
// - Được gọi từ handlers hoặc services để xử lý thanh toán
// Vai trò trong flow:
// - Xử lý thanh toán từ doanh nghiệp/sinh viên (nếu có)
//
// TODO: Implement actual Stripe API integration. Currently a stub.
// Steps to implement:
// 1. Add stripe-go SDK: go get github.com/stripe/stripe-go/v74
// 2. Create a PaymentIntent: https://stripe.com/docs/api/payment_intents
// 3. Return client secret to frontend for confirmation

// PaymentClient handles payment processing
type PaymentClient struct {
	apiKey string
}

// NewPaymentClient khởi tạo PaymentClient với API key
func NewPaymentClient(apiKey string) *PaymentClient {
	return &PaymentClient{apiKey: apiKey}
}

// PaymentRequest represents a payment initiation request
type PaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// PaymentResponse represents the result of a payment operation
type PaymentResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
}

// CreatePayment processes a payment. Currently a stub that logs and returns success.
func (c *PaymentClient) CreatePayment(req PaymentRequest) (*PaymentResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("invalid payment amount")
	}
	if req.Currency == "" {
		req.Currency = "usd"
	}

	log.Printf("[Payment Mock] Processing payment: %d %s", req.Amount, req.Currency)

	return &PaymentResponse{
		TransactionID: "mock_tx_" + req.Currency,
		Status:        "completed",
		Amount:        req.Amount,
		Currency:      req.Currency,
	}, nil
}
