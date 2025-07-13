#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

BASE_URL="http://localhost:8080"

echo -e "${BLUE}🚀 Quick API Test - Loan Service${NC}"
echo -e "${BLUE}=================================${NC}"

# Generate unique email for testing
TIMESTAMP=$(date +%s)
EMAIL="user${TIMESTAMP}@example.com"

echo -e "\n${YELLOW}1. Testing Health Check${NC}"
curl -s $BASE_URL/ | jq .

echo -e "\n${YELLOW}2. Testing Registration${NC}"
REGISTER_RESPONSE=$(curl -s -X POST $BASE_URL/api/auth/register \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\",
    \"password\": \"password123\",
    \"name\": \"Test User $TIMESTAMP\",
    \"phone\": \"08${TIMESTAMP:6:8}\"
  }")

echo "$REGISTER_RESPONSE" | jq .

# Extract customer ID
CUSTOMER_ID=$(echo "$REGISTER_RESPONSE" | jq -r '.data.customer.id // 1')

echo -e "\n${YELLOW}3. Testing Login${NC}"
curl -s -X POST $BASE_URL/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\",
    \"password\": \"password123\"
  }" | jq .

echo -e "\n${YELLOW}4. Testing Get Profile${NC}"
curl -s -X GET $BASE_URL/api/customers/profile \
  -H "X-Customer-ID: $CUSTOMER_ID" | jq .

echo -e "\n${YELLOW}5. Testing Update Profile${NC}"
curl -s -X PUT $BASE_URL/api/customers/profile \
  -H "Content-Type: application/json" \
  -H "X-Customer-ID: $CUSTOMER_ID" \
  -d '{
    "name": "Updated User Name",
    "phone": "0899999999",
    "address": "123 Updated Street, Bangkok"
  }' | jq .

echo -e "\n${YELLOW}6. Testing Phone Verification${NC}"
curl -s -X POST $BASE_URL/api/customers/verify-phone \
  -H "Content-Type: application/json" \
  -H "X-Customer-ID: $CUSTOMER_ID" \
  -d '{
    "phone": "0899999999",
    "otp": "123456"
  }' | jq .

echo -e "\n${YELLOW}7. Testing Identity Verification${NC}"
curl -s -X POST $BASE_URL/api/customers/verify-identity \
  -H "Content-Type: application/json" \
  -H "X-Customer-ID: $CUSTOMER_ID" \
  -d '{
    "id_card": "1357924680123",
    "first_name": "Test",
    "last_name": "User",
    "birth_date": "1990-05-15"
  }' | jq .

echo -e "\n${YELLOW}8. Testing Credit Score${NC}"
curl -s -X GET $BASE_URL/api/customers/credit-score \
  -H "X-Customer-ID: $CUSTOMER_ID" | jq .

echo -e "\n${YELLOW}9. Testing Password Change${NC}"
curl -s -X PUT $BASE_URL/api/customers/password \
  -H "Content-Type: application/json" \
  -H "X-Customer-ID: $CUSTOMER_ID" \
  -d '{
    "current_password": "password123",
    "new_password": "newpassword456"
  }' | jq .

echo -e "\n${YELLOW}10. Testing Final Profile Check${NC}"
curl -s -X GET $BASE_URL/api/customers/profile \
  -H "X-Customer-ID: $CUSTOMER_ID" | jq .

echo -e "\n${GREEN}✅ Quick Test Complete!${NC}"
echo -e "Customer ID: $CUSTOMER_ID"
echo -e "Email: $EMAIL"