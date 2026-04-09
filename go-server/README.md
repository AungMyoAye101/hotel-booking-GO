# Go Server

Backend API for the GO Hotel Booking application.

## Prerequisites

- Go 1.26+
- PostgreSQL
- `git`
- `.env` file in `go-server/`

## Makefile Commands

From `go-server/`, use:

```bash
make dev          # Start the server in development mode with air
make run-server   # Run the server directly
make build-server # Build the server executable to bin/server
make test         # Run unit tests
make tidy         # Sync module dependencies
```

## How to Run

```bash
cd go-server
go mod tidy
make run-server
```

### Build and run

```bash
cd go-server
go build -o bin/server ./cmd/server
env $(cat .env) ./bin/server
```

## API Endpoints Examples

Assumes the server is running on `http://127.0.0.1:8000`.

### Auth

- Register user

```bash
curl -X POST http://127.0.0.1:8000/api/v1/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"name":"Jane Doe","email":"jane@example.com","password":"secret123"}'
```

- Login user

```bash
curl -X POST http://127.0.0.1:8000/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"jane@example.com","password":"secret123"}'
```

### Hotels

- Get hotels

```bash
curl http://127.0.0.1:8000/api/v1/hotels
```

- Get hotel by ID

```bash
curl http://127.0.0.1:8000/api/v1/hotels/{hotelId}
```

### Rooms

- Get available rooms by hotel ID

```bash
curl http://127.0.0.1:8000/api/v1/hotels/{hotelId}/rooms/available
```

### Booking

- Create booking

```bash
curl -X POST http://127.0.0.1:8000/api/v1/bookings \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <ACCESS_TOKEN>' \
  -d '{"hotel_id":"{hotelId}","room_id":"{roomId}","user_id":"{userId}","check_in":"2026-05-01","check_out":"2026-05-05","guests":2}'
```

- Update booking

```bash
curl -X PUT http://127.0.0.1:8000/api/v1/bookings/{bookingId} \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <ACCESS_TOKEN>' \
  -d '{"check_in":"2026-05-02","check_out":"2026-05-06"}'
```

### Payment

- Create payment

```bash
curl -X POST http://127.0.0.1:8000/api/v1/payments \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <ACCESS_TOKEN>' \
  -d '{"booking_id":"{bookingId}","amount":250.00,"payment_method":"credit_card"}'
```

- Get payment by ID

```bash
curl http://127.0.0.1:8000/api/v1/payments/{paymentId} \
  -H 'Authorization: Bearer <ACCESS_TOKEN>'
```

### Receipt

- Get receipt by user ID

```bash
curl http://127.0.0.1:8000/api/v1/receipts/user/{userId} \
  -H 'Authorization: Bearer <ACCESS_TOKEN>'
```

## Environment Example

Create a `.env` file inside `go-server/` with:

```env
HOST=127.0.0.1
PORT=8000
DATABASE_URL=postgres://user:password@localhost:5432/hotel_booking_db?sslmode=disable
ACCESS_TOKEN_SECRET=your_access_token_secret
REFRESH_TOKEN_SECRET=your_refresh_token_secret
```

## Tech Stack

- Go 1.26
- Echo v4
- GORM v1.31
- PostgreSQL
- Viper config
- JWT authentication
- Google UUID

## Folder Structure

- `cmd/server/` - app entry point
- `config/` - config loader and validation
- `internal/` - business logic modules
  - `auth/`
  - `booking/`
  - `hotel/`
  - `payment/`
  - `receipt/`
  - `review/`
  - `room/`
  - `user/`
- `pkg/` - shared packages
  - `db/`
  - `middlewares/`
  - `models/`
  - `pagination/`
  - `response/`
  - `utils/`
  - `validation/`

## Core Features

- User authentication and token management
- Hotel and room listing endpoints
- Booking creation and management
- Payment creation and booking confirmation
- Receipt generation
- Hotel reviews
- Pagination and standardized API responses
