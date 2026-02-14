# URL Shortener Frontend

A simple and elegant URL shortener frontend built with React, Vite, and styled-components.

## Features

- Clean, modern interface using a custom color palette
- URL validation and shortening
- Copy to clipboard functionality
- Responsive design
- Smooth animations and transitions

## Color Palette

The application uses the following color palette:
- **Onyx**: #0F0F0F (very dark grey, almost black)
- **Liquorice**: #1A1A1A (dark grey)
- **Twilight Grey**: #2A2A2A (medium-dark grey)
- **Deep Saffron**: #FF9933 (vibrant orange)
- **Mercury**: #F1F1F1 (very light grey)
- **Snow**: #FCFCFC (off-white)

## Tech Stack

- **React 18** - UI framework
- **Vite** - Build tool and dev server
- **styled-components** - CSS-in-JS styling
- **JavaScript** - Language

## Getting Started

### Prerequisites

- Node.js (version 14 or higher)
- npm or yarn

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd url-shortener-frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm run dev
```

4. Open your browser and navigate to `http://localhost:5173`

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run lint` - Run ESLint

## Usage

1. Enter a long URL in the input field
2. Click the "Shorten" button
3. Copy the shortened URL using the "Copy" button
4. The shortened URLs are displayed in a list below

## Project Structure

```
src/
├── components/
│   └── UrlShortener.jsx    # Main URL shortener component
├── styles/
│   ├── colors.js           # Color palette and theme
│   └── GlobalStyles.js     # Global styles
├── App.jsx                 # Main application component
└── main.jsx               # Application entry point
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Commit your changes
5. Push to the branch
6. Create a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).
