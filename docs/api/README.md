# Loan Service API Documentation

## Overview
Loan Service API provides customer management and authentication functionality for a loan management system built with Clean Architecture principles.

## Base URL
```
http://localhost:8080
```

## Authentication
The API currently uses a mock authentication system with `X-Customer-ID` header for protected endpoints. In production, this will be replaced with JWT token authentication.

## API Endpoints

### Health Check
- **GET** `/` - Health check endpoint

### Authentication Endpoints

#### Register Customer
- **POST** `/api/auth/register`
- **Description**: Register a new customer account
- **Request Body**:
```json
{
  "email": "john.doe@example.com",
  "password": "password123", 
  "name": "John Doe",
  "phone": "0812345678"
}
```
- **Success Response** (201):
```json
{
  "success": true,
  "message": "Registration successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "customer": {
      "id": 1,
      "email": "john.doe@example.com",
      "name": "John Doe",
      "phone": "0812345678",
      "address": "",
      "credit_score": 0,
      "is_verified": false,
      "created_at": "2025-07-13T10:05:05.324Z",
      "updated_at": "2025-07-13T10:05:05.324Z"
    }
  }
}
```

#### Login Customer
- **POST** `/api/auth/login`
- **Description**: Authenticate customer and get token
- **Request Body**:
```json
{
  "email": "john.doe@example.com",
  "password": "password123"
}
```
- **Success Response** (200):
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "customer": {
      "id": 1,
      "email": "john.doe@example.com",
      "name": "John Doe",
      "phone": "0812345678",
      "address": "",
      "credit_score": 0,
      "is_verified": false,
      "created_at": "2025-07-13T10:05:05.324Z",
      "updated_at": "2025-07-13T10:05:05.324Z"
    }
  }
}
```

### Customer Management Endpoints
*All endpoints require `X-Customer-ID` header for authentication*

#### Get Customer Profile
- **GET** `/api/customers/profile`
- **Headers**: `X-Customer-ID: {customer_id}`
- **Success Response** (200):
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "id": 1,
    "email": "john.doe@example.com",
    "name": "John Doe",
    "phone": "0812345678",
    "address": "123 Main Street, Bangkok",
    "credit_score": 750,
    "is_verified": true,
    "created_at": "2025-07-13T10:05:05.324Z",
    "updated_at": "2025-07-13T10:05:05.324Z"
  }
}
```

#### Update Customer Profile
- **PUT** `/api/customers/profile`
- **Headers**: `X-Customer-ID: {customer_id}`
- **Request Body**:
```json
{
  "name": "John Doe Updated",
  "phone": "0887654321",
  "address": "123 Main Street, Bangkok"
}
```
- **Success Response** (200):
```json
{
  "success": true,
  "message": "Profile updated successfully",
  "data": {
    "id": 1,
    "email": "john.doe@example.com",
    "name": "John Doe Updated",
    "phone": "0887654321",
    "address": "123 Main Street, Bangkok",
    "credit_score": 750,
    "is_verified": true,
    "created_at": "2025-07-13T10:05:05.324Z",
    "updated_at": "2025-07-13T10:05:05.324Z"
  }
}
```

#### Change Password
- **PUT** `/api/customers/password`
- **Headers**: `X-Customer-ID: {customer_id}`
- **Request Body**:
```json
{
  "current_password": "password123",
  "new_password": "newpassword456"
}
```
- **Success Response** (200):
```json
{
  "success": true,
  "message": "Password changed successfully"
}
```

### Customer Verification Endpoints
*All endpoints require `X-Customer-ID` header for authentication*

#### Verify Phone Number
- **POST** `/api/customers/verify-phone`
- **Headers**: `X-Customer-ID: {customer_id}`
- **Request Body**:
```json
{
  "phone": "0887654321",
  "otp": "123456"
}
```
- **Success Response** (200):
```json
{
  "success": true,
  "message": "Phone verified successfully"
}
```

#### Verify Identity
- **POST** `/api/customers/verify-identity`
- **Headers**: `X-Customer-ID: {customer_id}`
- **Request Body**:
```json
{
  "id_card": "1234567890123",
  "first_name": "John",
  "last_name": "Doe",
  "birth_date": "1990-01-01"
}
```
- **Success Response** (200):
```json
{
  "success": true,
  "message": "Identity verified successfully"
}
```

#### Get Credit Score
- **GET** `/api/customers/credit-score`
- **Headers**: `X-Customer-ID: {customer_id}`
- **Success Response** (200):
```json
{
  "success": true,
  "message": "Credit score retrieved successfully",
  "data": {
    "customer_id": 1,
    "credit_score": 750,
    "score_grade": "A",
    "last_updated": "2025-07-13T10:05:05.324Z"
  }
}
```

## Error Responses

### Validation Error (400)
```json
{
  "success": false,
  "message": "Validation failed",
  "error": "Field validation errors..."
}
```

### Unauthorized (401)
```json
{
  "success": false,
  "message": "Unauthorized",
  "error": "Invalid token"
}
```

### Not Found (404)
```json
{
  "success": false,
  "message": "Profile not found",
  "error": "Customer not found"
}
```

### Internal Server Error (500)
```json
{
  "success": false,
  "message": "Internal server error",
  "error": "Error details..."
}
```

## Data Validation Rules

### Customer Registration
- **email**: Must be valid email format
- **password**: Minimum 6 characters
- **name**: Minimum 2 characters
- **phone**: Thai phone number format (08xxxxxxxx or 09xxxxxxxx)

### Profile Update
- **name**: Minimum 2 characters
- **phone**: Thai phone number format
- **address**: Optional string

### Phone Verification
- **phone**: Thai phone number format
- **otp**: Exactly 6 digits

### Identity Verification
- **id_card**: 13 digits (Thai ID card format)
- **first_name**: Required string
- **last_name**: Required string
- **birth_date**: Date format (YYYY-MM-DD)

## Credit Score Grades
- **A**: 800-900 (Excellent)
- **B**: 700-799 (Good)
- **C**: 600-699 (Fair)
- **D**: Below 600 (Poor)

## Testing
Use the provided test script or Postman collection:
```bash
# Run automated tests
./test_api.sh

# Import Postman collection
loan-service.postman_collection.json
```