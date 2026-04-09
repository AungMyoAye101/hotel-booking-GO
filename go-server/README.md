# Go Server

Backend API for the GO Hotel Booking application.

## Prerequisites

- Go 1.26+
- PostgreSQL
- `GIT`
- `.env` file in `go-server/`

## How to Run

```bash
cd go-server
go mod tidy
go run ./cmd/server
```

### Build and run

```bash
go build -o hotel-booking ./cmd/server
env $(cat .env) ./hotel-booking
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
