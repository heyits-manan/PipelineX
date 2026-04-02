# PipelineX

A scalable, distributed video processing pipeline built in Go. PipelineX is designed around a Redis-first architecture for metadata, job queues, and processing status updates.

---

## 🚀 Overview

PipelineX demonstrates how to design and implement a production-style video processing system. It separates concerns across services and uses asynchronous workflows to handle compute-intensive tasks like video transcoding.

### Key Capabilities

- Video registration via API
- Asynchronous job processing using Redis queues/streams
- Multi-resolution transcoding (e.g., 360p, 720p, 1080p)
- Thumbnail generation
- Scalable worker architecture
- Redis-backed metadata and status tracking

---

## 🏗 Architecture

```
Client → API Service → Redis (metadata + queue) → Worker Service → Redis (status updates)
```

### Components

- **API Service**
  - Handles client requests
  - Registers videos and metadata
  - Publishes processing jobs to Redis

- **Worker Service**
  - Consumes jobs from Redis
  - Runs video processing (FFmpeg)
  - Updates processing status back to Redis

- **Redis**
  - Stores video metadata and processing state
  - Provides queue/stream primitives for async job handling
  - Enables scalability and retries

---

## 📁 Project Structure

```
PipelineX/
├── cmd/
│   ├── api/        # API service entry point
│   ├── worker/     # Worker service entry point
│
├── internal/
│   ├── video/      # Video domain models, handlers, and service logic
│   ├── queue/      # Queue contracts and Redis-backed queue implementation
│   ├── store/      # Redis-backed video metadata/status store
│   ├── processor/  # Worker processor and transcoder abstraction
│   ├── ffmpeg/     # FFmpeg integration and command wrappers
│   ├── config/     # Shared application configuration
│
├── configs/        # Configuration files
├── deployments/    # Docker / Kubernetes configs
├── go.mod
├── README.md
```

---

## 🔄 Data Flow

1. Client sends a create/register video request to API
2. API stores initial video metadata in Redis (`uploaded`)
3. API enqueues a processing job in Redis
4. Worker consumes the job from Redis
5. Worker processes video (FFmpeg/fake transcoder)
6. Worker updates Redis with final status (`ready`/`failed`) and metadata
7. Client fetches current status from API

---

## 🧱 Tech Stack

- **Language**: Go
- **Video Processing**: FFmpeg
- **Metadata Store**: Redis
- **Queue**: Redis (Streams or Lists)
- **Containerization**: Docker
- **Orchestration (optional)**: Kubernetes

---

## ⚙️ Getting Started

### Prerequisites

- Go 1.20+
- Redis 7+
- FFmpeg installed
- Docker (optional)

---

### 1. Clone Repository

```
git clone https://github.com/heyits-manan/PipelineX.git
cd PipelineX
```

---

### 2. Start Redis (if not already running)

```
docker run --name pipelinex-redis -p 6379:6379 -d redis:7
```

---

### 3. Run API Service

```
go run cmd/api/main.go
```

---

### 4. Run Worker Service

```
go run cmd/worker/main.go
```

---

### 5. Create a Video Job

- Call API to register a video
- API will enqueue a Redis job
- Worker will automatically process it and update status in Redis

---

## 📌 Design Principles

- **Separation of concerns**: API and worker are independent
- **Asynchronous processing**: Improves scalability and reliability
- **Interface-driven design**: Enables easy swapping of components
- **Horizontal scalability**: Workers scale independently
- **Fault tolerance**: Redis-backed queue retries and decoupling

---

## ⚠️ Known Limitations (MVP)

- No authentication/authorization
- Limited error handling and retries
- No monitoring or observability yet
- Single-region deployment
- Redis is a single dependency and potential bottleneck in the current MVP

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
