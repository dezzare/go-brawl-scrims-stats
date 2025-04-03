# go-brawl-scrims-stats

## Overview
Supercell Stats API is an over-engineered project created for studying the Go language, data structures, database connections, pointers, working with JSON, environment variables, dependency management, containers and built with scalability and modularity in mind.

## Packages
There is 3 main pkg:
- **API Client**: Collects data from Supercell's API and stores it in a PostgreSQL database.
- **Stats**: Create statistics about registered players and teams based on the collected data.
- **API Server**: Share Stats data.(TODO)

## Features
- Fetches data from Supercell's API.
- Stores and manages data efficiently using PostgreSQL.
- Exposes an API for retrieving player and team statistics.

## Installation
### Prerequisites
- Go
- Docker
- Supercell API Key

### Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/dezzare/go-brawl-scrims-stats.git
   cd go-brawl-scrims-stats
   go mod tidy
   ```
2. Create .env file with:
   ```
   APIKEY=your_api_key
   DB_USERNAME=usernamedb
   DB_PASSWORD=passdb
   DB_NAME=dbdata
   DB_HOST=dbcontainer
   DB_PORT=5432
   APP_PORT=5000
   CLIENT_BASEURL=https://api.brawlstars.com/v1
   ```
3. Run database migrations (if applicable):
   ```sh
   docker compose up --build
   ```


## Usage
### Fetching Data
The Stats periodically request to API Client for fetches data and populates the database, and create statistics.


## License
This project is licensed under the MIT License.
