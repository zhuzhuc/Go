# Personal Website

A personal website built with Go, featuring a blog system and particle effects for visual appeal.

## Features

- Responsive design
- Blog publishing functionality
- User authentication
- Admin dashboard
- Particle effects for visual appeal
- RESTful API
- PostgreSQL database
- Docker deployment

## Tech Stack

- **Backend**: Go with Gin framework
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT and session-based auth
- **Frontend**: HTML, CSS, JavaScript
- **Visual Effects**: particles.js
- **Deployment**: Docker, Nginx

## Getting Started

### Prerequisites

- Go 1.20 or higher
- PostgreSQL
- Docker and Docker Compose (for containerized deployment)

### Local Development

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/personal-website.git
   cd personal-website
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Create a `.env` file based on `.env.example`:
   ```
   cp .env.example .env
   ```

4. Set up the database:
   ```
   createdb personal_website
   ```

5. Run the application:
   ```
   go run main.go
   ```

6. Visit `http://localhost:8080` in your browser.

### Docker Deployment

1. Build and start the containers:
   ```
   docker-compose up -d
   ```

2. Visit `http://localhost` in your browser.

## Project Structure

```
personal-website/
├── config/         # Configuration files
├── controllers/    # Request handlers
├── middleware/     # Middleware functions
├── models/         # Database models
├── static/         # Static assets (CSS, JS, images)
├── templates/      # HTML templates
├── utils/          # Utility functions
├── .env.example    # Environment variables example
├── Dockerfile      # Docker configuration
├── docker-compose.yml # Docker Compose configuration
├── go.mod          # Go modules
├── go.sum          # Go modules checksums
├── main.go         # Application entry point
└── README.md       # Project documentation
```

## API Endpoints

### Public Endpoints

- `GET /api/posts` - Get all published posts
- `GET /api/posts/:id` - Get a specific post

### Protected Endpoints

- `POST /api/posts` - Create a new post
- `PUT /api/posts/:id` - Update a post
- `DELETE /api/posts/:id` - Delete a post
- `GET /api/user` - Get current user information

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [particles.js](https://vincentgarreau.com/particles.js/)
