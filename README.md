# SkillForge

SkillForge is a collaborative project management platform that connects students with projects through AI-powered skill matching. The platform features real-time communication, task management, and intelligent project-student pairing capabilities.

Web Dev Adventure UIT project.

### Hướng dẫn cài đặt
- Tải Docker Desktop về máy chạy trên nền WSL2. Bật Docker Desktop chạy nền, vào root của thư mục dự án chạy lệnh sau nếu chạy lần đầu
```bash
docker-compose up --build -d
```
- Nếu là những lần chạy sau, chạy lệnh sau
```bash
docker-compose up -d
```
- **Ghi chú**:
  - Luôn phải bật Docker Desktop chạy nền trước khi chạy các lệnh trên.
  - Đảm bảo cập nhật Docker Desktop lên bản mới nhất nếu có thể.
  - Nếu không cài đặt được, đọc kỹ tài liệu trên trang chủ của Docker, nhớ kỹ phải cài đặt WSL2 trên máy trước.

## 🚀 Features

### Core Functionality

- **AI-Powered Matching**: Intelligent skill-based matching between students and projects using sentence transformers
- **Real-time Communication**: WebSocket-based chat, task updates, and notifications
- **Project Management**: Kanban-style task boards with live collaboration
- **User Management**: Student and project owner profiles with skill tracking

### Real-time Features

- **Project Chat**: Group messaging within project teams
- **Live Task Updates**: Real-time Kanban board synchronization
- **Instant Notifications**: User-specific event alerts
- **WebSocket Architecture**: Room-based connection management for scalable real-time updates

## 🏗️ Architecture

### Backend (Go)

- **Framework**: Go with MongoDB for data persistence
- **Real-time**: WebSocket connections with room-based architecture
- **AI Integration**: Sentence transformer model for skill matching
- **Database**: MongoDB with collections for projects, users, messages, tasks, and notifications

### Frontend (SvelteKit)

- **Framework**: SvelteKit with TypeScript
- **Real-time**: WebSocket client integration
- **UI**: Tailwind CSS for responsive design
- **Charts**: Chart.js for data visualization

### AI Matching System

- **Model**: all-MiniLM-L6-v2 sentence transformer for semantic similarity
- **Purpose**: Maps user skills and project requirements to 384-dimensional vectors for matching
- **Training**: Fine-tuned on 1B+ sentence pairs for optimal semantic understanding

## 🛠️ Technology Stack

**Backend:**

- Go
- MongoDB
- WebSocket (gorilla/websocket)
- Sentence Transformers (Python AI service)

**Frontend:**

- SvelteKit
- TypeScript
- Tailwind CSS
- Chart.js
- WebSocket client

**AI/ML:**

- Python
- Sentence Transformers
- HuggingFace Transformers

## 📁 Project Structure

```
SkillForge/
├── backend/                 # Go backend service
│   ├── internal/
│   │   ├── handlers/       # HTTP and WebSocket handlers
│   │   ├── services/       # Business logic layer
│   │   ├── repositories/   # Data access layer
│   │   ├── models/         # Data models
│   │   └── integrations/   # External service integrations
│   └── storage/            # File storage (avatars, etc.)
├── frontend/               # SvelteKit frontend
├── ai/                     # AI matching system
│   └── matching/
│       └── matching-model/ # Sentence transformer model
└── docs/                   # Documentation
```

## 🚦 Getting Started

### Prerequisites

- Go 1.19+
- Node.js 18+
- MongoDB
- Python 3.8+ (for AI matching)

### Run all services with Docker

```bash
docker-compose up
```

### Backend Setup

```bash
cd backend
go mod download
go run main.go
```

### Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

### AI Service Setup

```bash
cd ai/matching
pip install -r requirements.txt
python app.py
```

## 🔄 Real-time Communication

The platform uses a sophisticated WebSocket architecture with three main communication channels:

1. **Chat Messages**: Project-based group conversations using room format `"message{projectID}"`
2. **Task Updates**: Live Kanban board synchronization using room format `"{projectID}"`
3. **Notifications**: User-specific alerts using room format `"notification"`

Each WebSocket connection is managed through a centralized `RealtimeClient` that handles connection lifecycle, room management, and message broadcasting.

## 🤖 AI Matching Algorithm

The matching system uses semantic similarity to pair students with relevant projects:

1. **Feature Extraction**: User skills and project requirements are converted to text embeddings
2. **Similarity Calculation**: Cosine similarity between user and project vectors
3. **Score Generation**: Rounded similarity scores for ranking matches

## 📊 Database Schema

**Key Collections:**

- `users`: Student and project owner profiles
- `projects`: Project details and requirements
- `messages`: Chat communications
- `tasks`: Kanban board items
- `notifications`: User alerts
- `project_student`: Project membership relationships

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull reques>

## 📄 License

This project is licensed under the Apache 2.0 License.

## Notes

The codebase demonstrates a modern full-stack architecture with real-time capabilities and AI integration. The backend follows clean architecture principles with clear separation between handlers, services, and repositories. The AI matching system leverages state-of-the-art sentence transformers for semantic understanding, while the real-time communication system provides scalable WebSocket management through room-based organization. The project structure supports both development and production deployment scenarios.
For better understanding of the repo, follow this link https://deepwiki.com/iknizzz1807/SkillForge
