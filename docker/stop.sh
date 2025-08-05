#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Stopping Elotus Services...${NC}"
docker-compose down

echo -e "${GREEN}All services stopped successfully!${NC}" 