PerimeterSecure - Quickstart

Requirements: Docker and Docker Compose installed.

Quickstart:
1. docker-compose up --build
2. Frontend: http://localhost:3000
3. Backend health: http://localhost:8080/health
4. AI service: http://localhost:8000/docs

Notes:
- AI service uses OpenCV to demonstrate face + motion detection (placeholder).
- Backend is written in Go using Gin.
