# URL Shortener App

A modern URL shortener application built with serverless architecture and a sleek dark-themed frontend.

## Architecture

- **Backend**: AWS Serverless Application Model (SAM) + Go (AWS Lambda) + DynamoDB + API Gateway (HTTP API)
- **Frontend**: React + Vite + styled-components (dark theme design)
- **CI/CD**: GitHub Actions for automatic backend deployment

## Repository Structure

```
urlshorter-app/
├── backend/
│   ├── cmd/
│   │   ├── shorten/     # Lambda function for URL shortening
│   │   └── redirect/   # Lambda function for URL redirection
│   ├── template.yaml    # SAM template for AWS resources
│   ├── samconfig.toml   # SAM configuration
│   └── go.mod         # Go dependencies
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   └── UrlShortener.jsx
│   │   ├── styles/
│   │   │   ├── colors.js
│   │   │   └── GlobalStyles.js
│   │   └── App.jsx
│   ├── package.json
│   └── vite.config.js
└── .github/
    └── workflows/
        └── deploy-backend.yml
```

## API Endpoints

The backend exposes a REST API:

### `POST /shorten`
Shortens a URL and returns the shortened version.

**Request:**
```json
{
  "url": "https://example.com"
}
```

**Response:**
```json
{
  "shortCode": "abc123",
  "shortUrl": "https://your-api.com/abc123",
  "longUrl": "https://example.com"
}
```

### `GET /{shortCode}`
Redirects to the original URL associated with the short code.

**Response:** 302 redirect to the stored long URL

## Prerequisites

### Backend Development
- Go 1.23+
- AWS SAM CLI
- AWS credentials configured (`aws configure`)
- AWS CLI

### Frontend Development
- Node.js 14+ (LTS recommended)
- npm or yarn

## Local Development

### Frontend

1. **Install dependencies:**
```bash
cd frontend
npm install
```

2. **Configure API URL:**
Create a `.env.local` file in the frontend directory:
```env
VITE_API_URL=http://localhost:3001  # For local backend testing
# OR
VITE_API_URL=https://your-deployed-api.com  # For production API
```

3. **Start development server:**
```bash
npm run dev
```

4. **Access the application:**
Open http://localhost:5173 in your browser

### Backend (Local Testing)

1. **Build the SAM application:**
```bash
cd backend
sam build
```

2. **Run locally:**
```bash
sam local start-api
```

The API will be available at http://localhost:3000

## Deployment

### Backend (Production)

The backend is automatically deployed via GitHub Actions when changes are pushed to the `main` branch.

**Manual Deployment:**
```bash
cd backend
sam build
sam deploy
```

**Required GitHub Secrets:**
- `AWS_ACCESS_KEY_ID`: AWS access key
- `AWS_SECRET_ACCESS_KEY`: AWS secret key
- `AWS_DEFAULT_REGION`: AWS region (e.g., `us-east-2`)

### Frontend (Production)

#### Option 1: Vercel
1. Connect your repository to Vercel
2. Set environment variable: `VITE_API_URL`
3. Deploy automatically on push

#### Option 2: Netlify
1. Connect your repository to Netlify
2. Set environment variable: `VITE_API_URL`
3. Deploy automatically on push

#### Option 3: AWS S3 + CloudFront
1. Build the frontend:
```bash
cd frontend
npm run build
```

2. Upload `dist/` folder to S3
3. Configure CloudFront distribution

## Features

### Frontend
- ✅ Modern dark theme design
- ✅ Responsive layout
- ✅ URL validation
- ✅ Copy to clipboard functionality
- ✅ Real-time URL shortening
- ✅ Smooth animations and transitions

### Backend
- ✅ Serverless architecture
- ✅ Auto-scaling with AWS Lambda
- ✅ Fast response with DynamoDB
- ✅ URL validation and sanitization
- ✅ Click tracking (hits counter)
- ✅ CORS enabled
- ✅ Automatic CI/CD deployment

## Color Palette

The application uses a sophisticated dark theme:
- **Onyx**: `#0F0F0F` (main background)
- **Liquorice**: `#1A1A1A` (card backgrounds)
- **Twilight Grey**: `#2A2A2A` (borders and inputs)
- **Deep Saffron**: `#FF9933` (accent color, buttons, links)
- **Mercury**: `#F1F1F1` (secondary text)
- **Snow**: `#FCFCFC` (primary text)

## Security Considerations

- CORS is configured for cross-origin requests
- URL validation prevents malicious inputs
- AWS IAM roles follow least privilege principle
- No sensitive data is exposed in the frontend

## Monitoring and Logging

- AWS CloudWatch for Lambda logs
- AWS X-Ray for request tracing (if enabled)
- GitHub Actions for deployment monitoring

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes
4. Commit your changes: `git commit -m 'Add amazing feature'`
5. Push to the branch: `git push origin feature/amazing-feature`
6. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Troubleshooting

### Common Issues

**Frontend build fails:**
```bash
rm -rf node_modules package-lock.json
npm install
```

**SAM deployment fails:**
- Check AWS credentials: `aws configure list`
- Verify region in `samconfig.toml`
- Check IAM permissions

**CORS errors:**
- Verify API URL is correctly set in frontend
- Check CORS configuration in `backend/template.yaml`

### Getting Help

- Check the [GitHub Issues](https://github.com/NaderFarhat/urlShorter/issues) for known problems
- Create a new issue for bugs or feature requests
- Review AWS CloudWatch logs for backend errors
