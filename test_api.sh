#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

BASE_URL="http://localhost:8080"
TOKEN=""
CUSTOMER_ID=""

echo -e "${BLUE}🚀 Testing Loan Service API Endpoints${NC}"
echo -e "${BLUE}====================================${NC}"

# Function to make HTTP requests and show results
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local headers=$4
    local description=$5
    
    echo -e "\n${YELLOW}Testing: $description${NC}"
    echo -e "${BLUE}$method $endpoint${NC}"
    
    if [ "$method" = "GET" ]; then
        if [ -n "$headers" ]; then
            response=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X $method "$BASE_URL$endpoint" -H "$headers")
        else
            response=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X $method "$BASE_URL$endpoint")
        fi
    else
        if [ -n "$headers" ]; then
            response=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X $method "$BASE_URL$endpoint" \
                -H "Content-Type: application/json" \
                -H "$headers" \
                -d "$data")
        else
            response=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X $method "$BASE_URL$endpoint" \
                -H "Content-Type: application/json" \
                -d "$data")
        fi
    fi
    
    http_code=$(echo "$response" | grep "HTTP_CODE:" | cut -d: -f2)
    response_body=$(echo "$response" | sed '/HTTP_CODE:/d')
    
    if [ "$http_code" -ge 200 ] && [ "$http_code" -lt 300 ]; then
        echo -e "${GREEN}✅ Success (HTTP $http_code)${NC}"
    else
        echo -e "${RED}❌ Failed (HTTP $http_code)${NC}"
    fi
    
    echo "$response_body" | jq . 2>/dev/null || echo "$response_body"
    echo -e "${BLUE}----------------------------------------${NC}"
}

# 1. Health Check
test_endpoint "GET" "/" "" "" "Health Check"

# 2. Customer Registration
echo -e "\n${BLUE}📝 Testing Customer Registration${NC}"
register_data='{
    "email": "john.doe@example.com",
    "password": "password123",
    "name": "John Doe",
    "phone": "0812345678"
}'

register_response=$(curl -s -X POST "$BASE_URL/api/auth/register" \
    -H "Content-Type: application/json" \
    -d "$register_data")

echo "Registration Response:"
echo "$register_response" | jq .

# Extract token and customer ID
TOKEN=$(echo "$register_response" | jq -r '.data.token // empty')
CUSTOMER_ID=$(echo "$register_response" | jq -r '.data.customer.id // empty')

if [ -n "$TOKEN" ]; then
    echo -e "${GREEN}✅ Registration successful. Token extracted.${NC}"
else
    echo -e "${RED}❌ Registration failed. Using mock data for testing.${NC}"
    TOKEN="mock-token"
    CUSTOMER_ID="1"
fi

# 3. Customer Login
echo -e "\n${BLUE}🔑 Testing Customer Login${NC}"
login_data='{
    "email": "john.doe@example.com",
    "password": "password123"
}'

test_endpoint "POST" "/api/auth/login" "$login_data" "" "Customer Login"

# 4. Get Customer Profile
echo -e "\n${BLUE}👤 Testing Customer Profile${NC}"
test_endpoint "GET" "/api/customers/profile" "" "X-Customer-ID: $CUSTOMER_ID" "Get Customer Profile"

# 5. Update Customer Profile
echo -e "\n${BLUE}✏️ Testing Profile Update${NC}"
update_data='{
    "name": "John Doe Updated",
    "phone": "0887654321",
    "address": "123 Main Street, Bangkok"
}'

test_endpoint "PUT" "/api/customers/profile" "$update_data" "X-Customer-ID: $CUSTOMER_ID" "Update Customer Profile"

# 6. Change Password
echo -e "\n${BLUE}🔒 Testing Password Change${NC}"
password_data='{
    "current_password": "password123",
    "new_password": "newpassword456"
}'

test_endpoint "PUT" "/api/customers/password" "$password_data" "X-Customer-ID: $CUSTOMER_ID" "Change Password"

# 7. Verify Phone
echo -e "\n${BLUE}📱 Testing Phone Verification${NC}"
phone_data='{
    "phone": "0887654321",
    "otp": "123456"
}'

test_endpoint "POST" "/api/customers/verify-phone" "$phone_data" "X-Customer-ID: $CUSTOMER_ID" "Verify Phone"

# 8. Verify Identity
echo -e "\n${BLUE}🆔 Testing Identity Verification${NC}"
identity_data='{
    "id_card": "1234567890123",
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-01-01"
}'

test_endpoint "POST" "/api/customers/verify-identity" "$identity_data" "X-Customer-ID: $CUSTOMER_ID" "Verify Identity"

# 9. Get Credit Score
echo -e "\n${BLUE}💯 Testing Credit Score${NC}"
test_endpoint "GET" "/api/customers/credit-score" "" "X-Customer-ID: $CUSTOMER_ID" "Get Credit Score"

# Test with different customer for login
echo -e "\n${BLUE}🔄 Testing with Different User${NC}"
register_data2='{
    "email": "jane.smith@example.com",
    "password": "password456",
    "name": "Jane Smith",
    "phone": "0823456789"
}'

test_endpoint "POST" "/api/auth/register" "$register_data2" "" "Register Second User"

login_data2='{
    "email": "jane.smith@example.com",
    "password": "password456"
}'

test_endpoint "POST" "/api/auth/login" "$login_data2" "" "Login Second User"

# Error Cases
echo -e "\n${BLUE}❌ Testing Error Cases${NC}"

# Invalid registration
invalid_register='{
    "email": "invalid-email",
    "password": "123",
    "name": "",
    "phone": "123"
}'

test_endpoint "POST" "/api/auth/register" "$invalid_register" "" "Invalid Registration Data"

# Invalid login
invalid_login='{
    "email": "nonexistent@example.com",
    "password": "wrongpassword"
}'

test_endpoint "POST" "/api/auth/login" "$invalid_login" "" "Invalid Login Credentials"

# Unauthorized access
test_endpoint "GET" "/api/customers/profile" "" "X-Customer-ID: 999" "Unauthorized Profile Access"

echo -e "\n${GREEN}🎉 API Testing Complete!${NC}"
echo -e "${BLUE}====================================${NC}"