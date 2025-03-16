package integrations

// Nhiệm vụ: Tích hợp WebRTC để hỗ trợ video call
// Liên kết:
// - Dùng github.com/pion/webrtc/v3 để xử lý WebRTC
// - Được gọi từ services/message_service.go để khởi tạo video call
// Vai trò trong flow:
// - Tạo kết nối video giữa sinh viên và mentor

import (
	"github.com/pion/webrtc/v3"
)

type WebRTCClient struct {
	// config là cấu hình WebRTC (ICE servers, v.v.)
	config webrtc.Configuration
}

// NewWebRTCClient khởi tạo WebRTCClient
// Return: *WebRTCClient - con trỏ đến WebRTCClient
func NewWebRTCClient() *WebRTCClient {
	// Cấu hình ICE servers (STUN/TURN servers)
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	return &WebRTCClient{config}
}

// CreateOffer tạo SDP offer cho video call
// Return: *webrtc.SessionDescription (offer), error (nếu có lỗi)
func (c *WebRTCClient) CreateOffer() (*webrtc.SessionDescription, error) {
	// Tạo peer connection
	peerConnection, err := webrtc.NewPeerConnection(c.config)
	if err != nil {
		return nil, err
	}

	// Thêm video/audio track
	_, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo)
	if err != nil {
		return nil, err
	}
	_, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio)
	if err != nil {
		return nil, err
	}

	// Tạo offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		return nil, err
	}

	// Đặt local description
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		return nil, err
	}

	// Trả về offer
	return &offer, nil
}
