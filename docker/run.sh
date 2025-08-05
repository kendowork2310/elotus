#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting Elotus Services with Docker...${NC}"

# Function to cleanup on exit
cleanup() {
    echo -e "${YELLOW}Stopping services...${NC}"
    docker-compose down
    echo -e "${GREEN}Services stopped.${NC}"
    exit 0
}

# Set up signal handlers
trap cleanup SIGINT SIGTERM

# Build and start services
echo -e "${YELLOW}Building and starting services...${NC}"
docker-compose up --build -d

# Wait for services to be ready
echo -e "${YELLOW}Waiting for services to be ready...${NC}"
sleep 10

# Check if services are running
echo -e "${YELLOW}Checking service status...${NC}"
docker-compose ps

echo -e "${GREEN}Services are running!${NC}"
echo -e "${GREEN}Authentication Server: http://localhost:8080${NC}"
echo -e "${GREEN}Upload Server: http://localhost:8081${NC}"
echo -e "${GREEN}MySQL Database: localhost:3306${NC}"
echo -e "${YELLOW}Press Ctrl+C to stop all services${NC}"

# Keep the script running
while true; do
    sleep 1
done 