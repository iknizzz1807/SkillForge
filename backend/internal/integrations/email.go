package integrations

// Nhiệm vụ: Tích hợp gửi email qua SMTP
// Liên kết:
// - Dùng github.com/jordan-wright/email để gửi email
// - Được gọi từ services/notification_service.go để gửi thông báo email
// Vai trò trong flow:
// - Gửi email thông báo (ví dụ: project created, new application) tới user

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailClient struct {
	// host là địa chỉ server SMTP (ví dụ: smtp.gmail.com)
	host string
	// port là cổng SMTP (ví dụ: 587)
	port string
	// username là tên đăng nhập SMTP
	username string
	// password là mật khẩu SMTP
	password string
}

// NewEmailClient khởi tạo EmailClient với thông tin cấu hình
// Input: host (string), port (string), username (string), password (string)
// Return: *EmailClient - con trỏ đến EmailClient
func NewEmailClient(host, port, username, password string) *EmailClient {
	return &EmailClient{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

// SendEmail gửi email tới người nhận
// Input: to (string), subject (string), body (string)
// Return: error (nếu có lỗi)
func (c *EmailClient) SendEmail(to, subject, body string) error {
	// Tạo email mới
	e := email.NewEmail()
	// Điền thông tin email
	e.From = c.username
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	// Gửi email qua SMTP
	err := e.Send(c.host+":"+c.port, smtp.PlainAuth("", c.username, c.password, c.host))
	if err != nil {
		return err
	}

	// Trả về nil nếu thành công
	return nil
}
