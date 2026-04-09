# Client App

Next.js frontend for the GO Hotel Booking system.

## Prerequisites

- Node.js 20+
- `pnpm` (recommended) or `npm` / `yarn`
- A running backend API server

## How to Run

```bash
cd client
pnpm install
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000).

### Build for Production

```bash
pnpm build
pnpm start
```

## Environment Example

This project uses a client-side Axios base URL configured in `src/hooks/axios-api.ts`.

Create a `.env.local` file for local overrides if needed:

```env
NEXT_PUBLIC_API_BASE_URL=http://localhost:8000/api/server
```

If you change the API base URL, update `src/hooks/axios-api.ts` accordingly.

## Folder Structure

- `app/` - Next.js application routes and pages
- `components/` - reusable UI components
- `hooks/` - custom React hooks and data fetching
- `lib/` - shared utilities
- `service/` - API service calls
- `stores/` - Zustand state management
- `types/` - TypeScript type definitions
- `utils/` - helper functions
- `validations/` - validation schemas
- `public/` - static assets

## Core Features

- Search and view hotels
- Room detail and booking flow
- User authentication and profile management
- Payment handling
- Receipt history and PDF download
- Hotel reviews
- Responsive UI with Heroui components
