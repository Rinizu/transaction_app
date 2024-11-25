# Transaction App

This is a simple transaction application built using Go (Golang) and Gin-Gonic framework. It allows customers (users) to register, login/logout with cookies feature, perform transaction with merchants, check all history or desired customer history of transaction.

## Table of Contents
1. [Requirements](#requirements)
2. [Data Preparation](#data-preparation) 
3. [Environment Variables](#environment-variable) 
4. [API Endpoints](#api-endpoints)

## Requirements

- Go 1.20+
- `github.com/joho/godotenv` for reading .env config file
- `gin-gonic/gin` for routing
- `golang-jwt/jwt` for JWT-based authentication
- `x/crypto/bcrypt` for hashing sensitive-information (passwords)
- `golang.org/x/exp/rand` for randomize id value

## Data Preparation

1. Add or remove merchant data manually first on `../data/merchants.json` 

2. Install dependecies using 
```bash
go mod tidy
```

3. Make sure manually change customer amount of money (default amount = 0) after doing register customer in `../data/customers.json` to use transaction activity

## Environment Variables

1. Make sure you have a .env file with the following example (you can check on env_example file)
```bash
JWT_SECRET=your_jwt_secret

CUSTOMER_FILE=./data/customers.json
MERCHANT_FILE=./data/merchants.json
HISTORY_FILE=./data/history.json
```

2. Run the application 
```bash
go run main.go
```

## API Endpoints

1. Customer Registration
POST `localhost:8080/api/customer/register`
* Request Body:
```json
{
  "name": "customer1",
  "email": "customer1@example.com",
  "password": "customer1"
}
```

2. Customer Login
POST `localhost:8080/api/customer/login`
* Request Body:
```json
{
  "email": "customer1@example.com",
  "password": "customer1"
}
```

3. Get All Merchants
GET `localhost:8080/api/merchants`
* Response Body:
```json
{
    "merchants": [
        {
            "id": 1,
            "name": "Merchant A",
            "category": "Electronics"
        },
        {
            "id": 2,
            "name": "Merchant B",
            "category": "Clothing"
        }
    ]
}
```

4. Transaction Customer With Merchant
POST `localhost:8080/api/transaction`
* Request Body:
```json
{
  "customer_id": 1,
  "merchant_id": 1,
  "amount": 10000.0
}
```

5. Get All History of Transaction
GET `localhost:8080/api/history`
* Response Body: 
```json
{
    "data": [
        {
            "id": 88721,
            "customer_id": 6298,
            "merchant_id": 1,
            "amount": 100000,
            "timestamp": "2024-11-25T14:42:37.019886+07:00"
        },
        {
            "id": 84419,
            "customer_id": 6298,
            "merchant_id": 1,
            "amount": 100000,
            "timestamp": "2024-11-25T14:43:26.993304+07:00"
        },
        {
            "id": 89546,
            "customer_id": 6298,
            "merchant_id": 1,
            "amount": 100000,
            "timestamp": "2024-11-25T14:43:34.623315+07:00"
        }
    ],
    "message": "All transaction fetched successfully"
}
```

6. Get Transaction History of Desired Customer By ID
GET `localhost:8080/api/history/customer?customer_id=6298`
* Response Body:
```json
{
    "data": [
        {
            "id": 88721,
            "customer_id": 6298,
            "merchant_id": 1,
            "amount": 100000,
            "timestamp": "2024-11-25T14:42:37.019886+07:00"
        },
        {
            "id": 84419,
            "customer_id": 6298,
            "merchant_id": 1,
            "amount": 100000,
            "timestamp": "2024-11-25T14:43:26.993304+07:00"
        },
        {
            "id": 89546,
            "customer_id": 6298,
            "merchant_id": 1,
            "amount": 100000,
            "timestamp": "2024-11-25T14:43:34.623315+07:00"
        }
    ],
    "message": "All transaction fetched successfully"
}
```

7. Logout Customer
POST `localhost:8080/api/customer/logout`
* Response Body:
```json
{
    "message": "customer logout successfully"
}
```