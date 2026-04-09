# GO Hotel Booking

A full-stack hotel booking system with a Next.js frontend and a Go backend API.

## Project Overview

This repository contains:

- `client/` — a Next.js 16 app built with React 19 for hotel search, booking, payments, receipts, reviews, and user management.
- `go-server/` — a Go 1.26 API server using Echo, GORM, PostgreSQL, JWT auth, and a modular service structure.

## Prerequisites

- Node.js 20+
- `pnpm` (recommended) or `npm` / `yarn`
- Go 1.26+
- PostgreSQL database
- A terminal in the repository root

## Tech Stack

### Frontend

- Next.js 16
- React 19
- TypeScript
- Tailwind CSS 4
- `@tanstack/react-query`
- Axios
- `zustand`
- `@heroui/react`
- `@react-pdf/renderer`

### Backend

- Go 1.26
- Echo v4
- GORM v1.31
- PostgreSQL
- Viper for config
- JWT auth
- UUID support

## Running the Project

### Backend

```bash
cd go-server
go mod tidy
go run ./cmd/server
```

### Frontend

```bash
cd client
pnpm install
pnpm dev
```

## Future Plan
- Add background image upload
- Add automated tests for frontend and backend
- Add CI/CD and deployment pipeline

