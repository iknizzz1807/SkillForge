package integrations

// Nhiệm vụ: Tích hợp gửi email qua SMTP
// Liên kết:
// - Dùng github.com/jordan-wright/email để gửi email
// - Được gọi từ services/notification_service.go để gửi thông báo email
// Vai trò trong flow:
// - Gửi email thông báo (ví dụ: project created, new application) tới user

import (
	gomail "gopkg.in/gomail.v2"

	"github.com/iknizzz1807/SkillForge/internal/constants"
)

type EmailClient struct {
	client *gomail.Dialer
}

func NewEmailClient(host string, port int, username string, password string) *EmailClient {
	return &EmailClient{gomail.NewDialer(host, port, username, password)}
}

func (e *EmailClient) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", constants.EmailFrom)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return e.client.DialAndSend(m)
}