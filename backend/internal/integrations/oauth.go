package integrations

// Nhiệm vụ: Tích hợp OAuth2 cho đăng nhập qua Google, GitHub, v.v.
// Liên kết:
// - Dùng golang.org/x/oauth2 để xử lý OAuth2 flow
// - Được gọi từ services/auth_service.go cho đăng nhập qua bên thứ ba
// Vai trò trong flow:
// - Xử lý xác thực OAuth2, lấy thông tin user từ provider
//
// TODO: Implement OAuth2 providers (Google, GitHub, etc.)
