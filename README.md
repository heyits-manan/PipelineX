# PipelineX

A scalable, distributed video processing pipeline built in Go. PipelineX is designed to handle video uploads, asynchronous processing (transcoding, thumbnail generation), and efficient delivery using a decoupled, queue-driven architecture.

---

## 🚀 Overview

PipelineX demonstrates how to design and implement a production-style video processing system. It separates concerns across services and uses asynchronous workflows to handle compute-intensive tasks like video transcoding.

### Key Capabilities

- Video upload via pre-signed URLs
- Asynchronous job processing using queues
- Multi-resolution transcoding (e.g., 360p, 720p, 1080p)
- Thumbnail generation
- Scalable worker architecture
- Pluggable storage and queue backends

---

## 🏗 Architecture

```
Client → API Service → Storage → Queue → Worker Service → Output Storage → CDN
```

### Components

- **API Service**
  - Handles client requests
  - Generates upload URLs
  - Publishes processing jobs

- **Worker Service**
  - Consumes jobs from queue
  - Runs video processing (FFmpeg)
  - Uploads processed outputs

- **Storage Layer**
  - Stores raw and processed video assets

- **Queue**
  - Decouples upload from processing
  - Enables scalability and retries

- **Database**
  - Tracks video metadata and processing status

---

## 📁 Project Structure

```
PipelineX/
├── cmd/
│   ├── api/        # API service entry point
│   ├── worker/     # Worker service entry point
│
├── internal/
│   ├── api/        # HTTP handlers and services
│   ├── worker/     # Worker and processing logic
│   ├── storage/    # Storage abstraction (S3/local)
│   ├── queue/      # Queue abstraction (SQS/Kafka)
│   ├── ffmpeg/     # Video processing logic
│   ├── db/         # Database access layer
│   ├── models/     # Shared data models
│
├── configs/        # Configuration files
├── deployments/    # Docker / Kubernetes configs
├── go.mod
├── README.md
```

---

## 🔄 Data Flow

1. Client requests upload URL from API
2. Client uploads video directly to storage
3. API enqueues a processing job
4. Worker consumes job from queue
5. Worker downloads video and processes it
6. Processed outputs are uploaded to storage
7. Database is updated with status and metadata
8. CDN serves processed video

---

## 🧱 Tech Stack

- **Language**: Go
- **Video Processing**: FFmpeg
- **Storage**: S3-compatible object storage
- **Queue**: SQS / Kafka (pluggable)
- **Database**: PostgreSQL
- **Containerization**: Docker
- **Orchestration (optional)**: Kubernetes

---

## ⚙️ Getting Started

### Prerequisites

- Go 1.20+
- FFmpeg installed
- Docker (optional)

---

### 1. Clone Repository

```
git clone https://github.com/heyits-manan/PipelineX.git
cd PipelineX
```

---

### 2. Run API Service

```
go run cmd/api/main.go
```

---

### 3. Run Worker Service

```
go run cmd/worker/main.go
```

---

### 4. Upload a Video

- Call API to generate upload URL
- Upload video using the URL
- Worker will automatically process it

---

## 📌 Design Principles

- **Separation of concerns**: API and worker are independent
- **Asynchronous processing**: Improves scalability and reliability
- **Interface-driven design**: Enables easy swapping of components
- **Horizontal scalability**: Workers scale independently
- **Fault tolerance**: Queue-based retries and decoupling

---

## ⚠️ Known Limitations (MVP)

- No authentication/authorization
- Limited error handling and retries
- No monitoring or observability yet
- Single-region deployment

---

## 🔮 Future Improvements

- Adaptive bitrate streaming (HLS/DASH)
- GPU-accelerated transcoding
- Distributed tracing and metrics
- Multi-region deployment
- AI-based video analysis (captions, moderation)

---

## 🤝 Contributing

Contributions are welcome. Feel free to open issues or submit pull requests.

---

## 📄 License

MIT License
