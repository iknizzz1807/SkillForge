package integrations

// // Nhiệm vụ: Tích hợp thanh toán qua Stripe
// // Liên kết:
// // - Dùng github.com/stripe/stripe-go/v76 để gọi Stripe API
// // - Được gọi từ handlers hoặc services để xử lý thanh toán
// // Vai trò trong flow:
// // - Xử lý thanh toán từ doanh nghiệp/sinh viên (nếu có)

// import (
// 	"github.com/stripe/stripe-go/v76"
// 	"github.com/stripe/stripe-go/v76/paymentintent"
// )

// type PaymentClient struct {
// 	// apiKey là Stripe secret key
// 	apiKey string
// }

// // NewPaymentClient khởi tạo PaymentClient với API key
// // Input: apiKey (string) - Stripe secret key
// // Return: *PaymentClient - con trỏ đến PaymentClient
// func NewPaymentClient(apiKey string) *PaymentClient {
// 	// Thiết lập API key cho Stripe
// 	stripe.Key = apiKey
// 	return &PaymentClient{apiKey}
// }

// // CreatePaymentIntent tạo payment intent cho thanh toán
// // Input: amount (int64) - số tiền (cent), currency (string) - loại tiền (ví dụ: "usd")
// // Return: string (client secret), error (nếu có lỗi)
// func (c *PaymentClient) CreatePaymentIntent(amount int64, currency string) (string, error) {
// 	// Tạo params cho payment intent
// 	params := &stripe.PaymentIntentParams{
// 		Amount:   stripe.Int64(amount),
// 		Currency: stripe.String(currency),
// 	}

// 	// Tạo payment intent
// 	pi, err := paymentintent.New(params)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Trả về client secret để frontend xử lý
// 	return pi.ClientSecret, nil
// }
