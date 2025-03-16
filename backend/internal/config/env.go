package config

// Nhiệm vụ: Định nghĩa struct Env để lưu trữ tất cả biến môi trường cần thiết
// Liên kết:
// - Được sử dụng trong config.go để trả về cấu hình
// - Được các file khác (app.go, integrations) truy cập để lấy thông tin
// Vai trò trong flow:
// - Là "kho chứa" cấu hình, giúp các thành phần khác truy cập dễ dàng
// - Đảm bảo tất cả biến môi trường được định nghĩa tập trung tại một nơi

// Env là struct chính chứa toàn bộ cấu hình của ứng dụng
type Env struct {
	// MongoURI: Đường dẫn kết nối đến MongoDB
	MongoURI string

	// EmailHost: Địa chỉ server SMTP để gửi email
	EmailHost string
	// EmailPort: Cổng SMTP (thường là 587 cho TLS)
	EmailPort string
	// EmailUser: Tên người dùng để đăng nhập SMTP
	EmailUser string
	// EmailPass: Mật khẩu để đăng nhập SMTP
	EmailPass string

	// AIURL: URL của dịch vụ AI (FastAPI) để gọi API
	AIURL string

	// GitHubToken: Token để gọi GitHub API (commit tracking, repository)
	GitHubToken string

	// StripeKey: API key của Stripe để xử lý thanh toán
	StripeKey string

	// StorageConfig: Cấu hình lưu trữ file (AWS S3)
	StorageConfig StorageConfig
}

// StorageConfig chứa thông tin cấu hình cho AWS S3
type StorageConfig struct {
	// AWSAccessKey: Khóa truy cập AWS
	AWSAccessKey string
	// AWSSecretKey: Khóa bí mật AWS
	AWSSecretKey string
	// AWSRegion: Vùng AWS (ví dụ: us-east-1)
	AWSRegion string
	// AWSBucket: Tên bucket S3 để lưu trữ file
	AWSBucket string
}
