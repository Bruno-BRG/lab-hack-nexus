# Lab Hack Nexus

A website for educational content about cybersecurity and hacking.

## Project Overview

Lab Hack Nexus is a full-stack web application built with React and Go, designed to provide educational content about cybersecurity and ethical hacking. The project features a modern UI with category-based content organization and a responsive design.

## Tech Stack

### Frontend
- Vite + React
- TypeScript
- Tailwind CSS
- shadcn/ui components
- React Context for state management

### Backend
- Go
- SQLite database
- RESTful API architecture

## Project Structure

```
├── frontend/          # React frontend application
│   ├── src/
│   │   ├── components/    # React components
│   │   ├── pages/        # Page components
│   │   ├── contexts/     # React contexts
│   │   └── hooks/        # Custom React hooks
│   
├── backend/           # Go backend server
│   ├── handlers/     # API endpoint handlers
│   ├── models/       # Data models
│   ├── config/       # Configuration
│   └── data/         # SQLite database
```

## Getting Started

### Prerequisites
- Node.js (v16 or higher)
- Go (v1.16 or higher)
- Git

### Frontend Setup

1. Navigate to the frontend directory:
```sh
cd frontend
```

2. Install dependencies:
```sh
npm install
```

3. Start the development server:
```sh
npm run dev
```

The frontend will be available at http://localhost:5173

### Backend Setup

1. Navigate to the backend directory:
```sh
cd backend
```

2. Run the Go server:
```sh
go run main.go
```

The backend API will be available at http://localhost:3000

## Features

- Category-based content organization
- Create and manage educational posts
- Modern, responsive UI
- SQLite database for data persistence
- RESTful API endpoints

## API Endpoints

### Categories
- GET /api/categories - List all categories
- POST /api/categories - Create a new category
- GET /api/categories/:id - Get category details
- PUT /api/categories/:id - Update a category
- DELETE /api/categories/:id - Delete a category

### Posts
- GET /api/posts - List all posts
- POST /api/posts - Create a new post
- GET /api/posts/:id - Get post details
- PUT /api/posts/:id - Update a post
- DELETE /api/posts/:id - Delete a post

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a Pull Request

## Project Repository

**URL**: https://github.com/Bruno-BRG/lab-hack-nexus
