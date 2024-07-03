# GO / MongoDB Demo Application

Welcome to the GO / MongoDB Demo Application! This project is a work in progress, created as a learning exercise to explore the Go programming language.

## Getting Started

To get the application up and running, follow these steps:

1. **Start the MongoDB Database**:
   ```bash
   docker compose up -d
   ```

2. **Run the Application**:
   ```bash
   go run .
   ```

## Available APIs

You can interact with the application using the following API endpoints. Make sure to replace the example IDs with the actual generated IDs from your database.

### Create a Product

**Endpoint**: `POST localhost:8080/products`

**Headers**:
```
Accept: application/json
```

**Body**:
```json
{
  "name": "iPhone 15",
  "description": "Fancy phone",
  "price": 599.00
}
```

### Update a Product

**Endpoint**: `PATCH localhost:8080/products/{id}`

**Headers**:
```
Accept: application/json
```

**Body**:
```json
{
  "description": "Updated description for iPhone 15"
}
```

### Get a Product

**Endpoint**: `GET localhost:8080/products/{id}`

Example:
```plaintext
GET localhost:8080/products/e7fd161d-fb50-43fe-98c3-96e4b3b3329d
```

Feel free to explore and modify the application as you learn more about Go and MongoDB!