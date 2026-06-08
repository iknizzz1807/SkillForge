package integrations

// Nhiệm vụ: Tích hợp thanh toán qua Stripe
// Liên kết:
// - Được gọi từ handlers hoặc services để xử lý thanh toán
// Vai trò trong flow:
// - Xử lý thanh toán từ doanh nghiệp/sinh viên (nếu có)

// PaymentClient handles payment processing
type PaymentClient struct {
	apiKey string
}

// NewPaymentClient khởi tạo PaymentClient với API key
// Input: apiKey (string) - Stripe secret key
// Return: *PaymentClient - con trỏ đến PaymentClient
func NewPaymentClient(apiKey string) *PaymentClient {
	return &PaymentClient{apiKey: apiKey}
}
