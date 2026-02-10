# urlShorter App

A minimal URL shortener built with:

- **Backend**: AWS Serverless Application Model (SAM) + Go (AWS Lambda) + DynamoDB + API Gateway (HTTP API)
- **Frontend**: Next.js (App Router) + React + styled-components

## Repository structure

- `backend/`
  - AWS SAM application (Go Lambdas + DynamoDB + HttpApi)
- `frontend/`
  - Next.js app (styled-components, minimal UI) that calls the backend

## API

The backend exposes:

- `POST /shorten`
  - Request JSON: `{ "url": "https://example.com" }`
  - Response JSON: `{ "shortCode": "...", "shortUrl": "...", "longUrl": "..." }`
- `GET /{shortenCode}`
  - 302 redirect to the stored long URL

## Prerequisites

### Backend

- Go installed (matching your project/tooling)
- AWS SAM CLI installed
- AWS credentials configured (e.g. `aws configure`)

### Frontend

- Node.js (LTS recommended)
- npm

## Running the frontend locally

1) Install dependencies:

```bash
cd frontend
npm install
```

2) Create `frontend/.env.local`:

```env
NEXT_PUBLIC_API_URL=https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com
```

Notes:

- Do **not** add `/shorten` at the end. The app will call `/shorten` automatically.
- The API URL must be the **HTTP API base URL** (see the backend deployment outputs).

3) Start the dev server:

```bash
npm run dev
```

Open:

- http://localhost:3000

### Production build (local)

```bash
cd frontend
npm run build
npm run start
```

## Deploying the backend (AWS SAM)

From `backend/`:

```bash
sam build
sam deploy
```

This project contains `backend/samconfig.toml` with default deploy parameters.

After deployment, grab the API base URL from the stack outputs (`HttpApiUrl`) and set it as `NEXT_PUBLIC_API_URL` in the frontend.

## CORS

CORS is enabled for the HttpApi via `Globals.HttpApi.CorsConfiguration` in `backend/template.yaml`.

If you deploy the frontend to Amplify, consider restricting allowed origins to your Amplify domain instead of `*`.

## Deploying the frontend (AWS Amplify)

The frontend contains an `amplify.yml` that works with Amplify Hosting.

In Amplify, set this environment variable:

- `NEXT_PUBLIC_API_URL` = `https://...execute-api...amazonaws.com`

Then trigger a build/deploy.
