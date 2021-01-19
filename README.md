# Wiki App

[![Build Status](https://travis-ci.org/Sungmin-Kim-0508/Wiki.svg?branch=master)](https://travis-ci.org/Sungmin-Kim-0508/Wiki)

## Backend Part (Part 1)

### Run the application

```bash
export GO_ENV=development

# Run the application
go run main.go

# APIs
# Step 1: Check if articles are empty
curl http://localhost:9090/articles/

# Step 2: Single article is empty
curl http://localhost:9090/articles/rest_api

# Step 3: Add Article
curl -X PUT http://localhost:9090/articles/wiki -d 'A wiki is a knowledge base website'

# Step 4: Update Article
curl -X PUT http://localhost:9090/articles/wiki -d 'A wiki is the best website'

# Step 5: Get a article named 'wiki'
curl http://localhost:9090/articles/wiki

# Step 6: Get article list
curl http://localhost:9090/articles/
```

### Build the application

```bash
# Create .exe file
go build -o main.exe
```

### Test the application

```bash
go test
```

## Frontend Part (Part 2)

### Run the project
```bash
# Make sure the server is running

## Development Environment
# Set environment varibles
export REACT_APP_DEV_ENV=development

npm run start # or
yarn run start

### Production Environment
npm run start # or
yarn run start
```

### Test the project
```bash
npm run test
```

### Build the project for the production
```bash
npm run build # or
yarn run build
```

## How to dockerize the wiki app (Part 3)

```bash
docker-compose build
docker-compose up
```