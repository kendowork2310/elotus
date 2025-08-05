# Elotus Docker Setup

This directory contains the Docker configuration to run the Elotus application with both authentication and upload services.

## Services

- **MySQL Database**: Port 3306
- **Authentication Server**: Port 8080
- **Upload Server**: Port 8081

## Quick Start

### Option 1: Using the run script (Recommended)
```bash
cd docker
./run.sh
```

This will:
- Build and start all services
- Show service status
- Keep running until you press Ctrl+C
- Automatically stop all services when you exit

### Option 2: Manual Docker Compose
```bash
cd docker

# Build and start all services
docker-compose up --build -d

# Check service status
docker-compose ps

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

## Service URLs

Once running, you can access:
- **Authentication Server**: http://localhost:8080
- **Upload Server**: http://localhost:8081
- **MySQL Database**: localhost:3306

## Environment Variables

The services are configured with the following environment variables:
- `MYSQL_URI`: mysql:3306 (internal Docker network)
- `MYSQL_DATABASE`: elotus
- `MYSQL_USERNAME`: elotus
- `MYSQL_PASSWORD`: elotus
- `AUTHEN_SERVER_PORT`: 8080
- `UPLOAD_SERVER_PORT`: 8081
- `SERVER_MODE`: debug
- `JWT_SECRET_KEY`: toilaken
- `JWT_REFRESH_SECRET_KEY`: refresh_toilaken

## Database

The MySQL database is automatically initialized with:
- `user` table for authentication
- `upload` table for file uploads

## Troubleshooting

1. **Port conflicts**: Make sure ports 8080, 8081, and 3306 are available
2. **Build issues**: Try `docker-compose build --no-cache`
3. **Database connection**: Wait a few seconds for MySQL to fully start
4. **View logs**: Use `docker-compose logs [service-name]` to debug issues

## Cleanup

To completely remove all containers, volumes, and images:
```bash
docker-compose down -v --rmi all
``` 