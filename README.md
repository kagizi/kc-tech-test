## References

- [How to Domain Driven Design (DDD) Golang](https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/)
- [go-ddd by sklinkert](https://github.com/sklinkert/go-ddd)
- [ddd-go by percybolmer](https://github.com/percybolmer/ddd-go)
## API Endpoints

### 1. Register User
```http
POST /register
Content-Type: application/json

{
    "name": "Max Verstappen",
    "phone": "1234567890"
}
```

**Success Response (200):**
```json
{
    "user_id": "01234567-89ab-cdef-0123-456789abcdef",
    "wallet_id": "01234567-89ab-cdef-0123-456789abcdef",
    "balance": "0"
}
```

**Error Response (400/500):**
```json
{
    "error": "invalid request"
}
```

### 2. Get User by Phone
```http
GET /user?phone=1234567890
```

**Success Response (200):**
```json
{
    "user_id": "01234567-89ab-cdef-0123-456789abcdef",
    "name": "Max Verstappen",
    "phone": "1234567890",
    "created_at": "2025-08-02T05:37:50Z",
    "wallet": {
        "wallet_id": "01234567-89ab-cdef-0123-456789abcdef",
        "balance": "365919",
        "created_at": "2025-08-02T05:37:50Z"
    }
}
```

**Error Response (400/404):**
```json
{
    "error": "phone required"
}
```

### 3. Deposit Money
```http
POST /deposit
Content-Type: application/json

{
    "wallet_id": "01234567-89ab-cdef-0123-456789abcdef",
    "amount": "100.50"
}
```

**Success Response (200):**
```json
{
    "message": "deposit successful",
    "balance": "150.75"
}
```

**Error Response (400):**
```json
{
    "error": "invalid wallet_id"
}
```

### 4. Withdraw Money
```http
POST /withdraw
Content-Type: application/json

{
    "wallet_id": "01234567-89ab-cdef-0123-456789abcdef",
    "amount": "50.25"
}
```

**Success Response (200):**
```json
{
    "message": "withdrawal successful",
    "balance": "100.50"
}
```

**Error Response (400):**
```json
{
    "error": "insufficient balance"
}
```

### 5. Check Balance
```http
GET /balance?wallet_id=01234567-89ab-cdef-0123-456789abcdef
```

**Success Response (200):**
```json
{
    "wallet_id": "01234567-89ab-cdef-0123-456789abcdef",
    "balance": "100.50"
}
```

**Error Response (400/404):**
```json
{
    "error": "wallet_id required"
}
```
