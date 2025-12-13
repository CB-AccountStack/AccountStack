# AccountStack Web Application

Modern React-based web application for AccountStack, featuring CloudBees Feature Management integration for dynamic feature control.

## Features

- **Dashboard**: Overview of all accounts with real-time balance information
- **Transactions**: Comprehensive transaction list with advanced filtering
- **Insights**: AI-powered financial insights and recommendations
- **Feature Flags**: Dynamic feature control using CloudBees Feature Management (Rox)
- **Responsive Design**: Mobile-first design with Tailwind CSS
- **Real-time Updates**: Automatic data refresh with TanStack Query
- **Type Safety**: Full TypeScript support

## Tech Stack

- **React 18** - UI framework
- **TypeScript** - Type safety
- **Vite** - Build tool and dev server
- **React Router** - Client-side routing
- **TanStack Query** - Data fetching and caching
- **Axios** - HTTP client
- **Tailwind CSS** - Styling
- **CloudBees FM (Rox)** - Feature flag management
- **Vitest** - Unit testing
- **Playwright** - E2E testing

## Feature Flags

The application uses CloudBees Feature Management to control features dynamically:

| Flag | Default | Description |
|------|---------|-------------|
| `ui.dashboardCardsV2` | `true` | Enhanced dashboard card design with gradients and better visuals |
| `ui.insightsV2` | `false` | New insights panel with improved layout and card-based design |
| `ui.alertsBanner` | `true` | Top banner for displaying important alerts and announcements |
| `ui.transactionsFilters` | `true` | Advanced filtering options for transactions (search, type, category, status) |
| `kill.ui.insights` | `false` | Kill switch to disable the entire insights feature |

## Getting Started

### Prerequisites

- Node.js 18+
- npm or yarn

### Installation

```bash
# Install dependencies
npm install
```

### Environment Configuration

Create a `.env` file based on `.env.example`:

```bash
cp .env.example .env
```

Configure the following variables:

```env
# API base URL (defaults to /api)
VITE_API_BASE_URL=/api

# CloudBees Feature Management API Key
# Get your key from: https://app.cloudbees.io/
VITE_ROX_API_KEY=your_api_key_here

# Environment
VITE_ENV=development
```

### Development

Start the development server:

```bash
npm run dev
```

The app will be available at `http://localhost:3000`

### Building for Production

```bash
npm run build
```

The built files will be in the `dist` directory.

### Preview Production Build

```bash
npm run preview
```

## Testing

### Unit Tests

```bash
# Run tests
npm run test

# Run tests in watch mode
npm run test:unit

# Generate coverage report
npm run test:coverage
```

### E2E Tests

```bash
npm run test:e2e
```

## Code Quality

### Linting

```bash
npm run lint
```

### Formatting

```bash
npm run format
```

## Project Structure

```
src/
├── components/          # Reusable UI components
│   ├── Layout.tsx      # Main layout with header, nav, footer
│   ├── AccountCard.tsx # Account card with V1/V2 variants
│   ├── TransactionList.tsx # Transaction list with filtering
│   ├── InsightsPanel.tsx   # Insights panel with V1/V2 variants
│   └── AlertBanner.tsx     # Alert banner component
├── pages/              # Page components
│   ├── Dashboard.tsx   # Dashboard page
│   ├── Transactions.tsx # Transactions page
│   └── Insights.tsx    # Insights page
├── features/           # Feature-specific code
│   └── flags.ts        # CloudBees FM integration
├── services/           # API and external services
│   └── api.ts          # Axios API client
├── styles/             # Global styles
│   └── index.css       # Tailwind CSS imports and custom styles
├── test/               # Test configuration
│   └── setup.ts        # Vitest setup
├── types.ts            # TypeScript type definitions
├── App.tsx             # Main app component with routing
├── main.tsx            # Application entry point
└── vite-env.d.ts       # Vite environment types
```

## API Integration

The application connects to the following API endpoints:

- `GET /api/accounts/me` - Get current user information
- `GET /api/accounts` - List all accounts
- `GET /api/transactions` - List all transactions
- `GET /api/insights` - Get financial insights

### API Proxy Configuration

The Vite dev server is configured to proxy API requests:

- `/api/accounts` → `http://api-accounts:8001`
- `/api/transactions` → `http://api-transactions:8002`
- `/api/insights` → `http://api-insights:8003`

## CloudBees Feature Management Integration

### Setup

1. Sign up for CloudBees Feature Management at https://app.cloudbees.io/
2. Create a new application
3. Get your API key
4. Add the API key to your `.env` file

### Using Feature Flags in Components

```tsx
import { useFeatureFlags } from '@/features/flags';

function MyComponent() {
  const { dashboardCardsV2, transactionsFilters } = useFeatureFlags();

  return (
    <div>
      {dashboardCardsV2 ? <EnhancedCard /> : <BasicCard />}
    </div>
  );
}
```

### Checking Flags Programmatically

```tsx
import {
  isDashboardCardsV2Enabled,
  isTransactionsFiltersEnabled
} from '@/features/flags';

if (isDashboardCardsV2Enabled()) {
  // Use V2 implementation
}
```

## Styling

The application uses Tailwind CSS with a custom brand color scheme:

- Primary Brand Color: `#0066cc`
- Color palette: `brand-{50-900}`

Custom component classes are available:
- `.card` - Base card styling
- `.btn-primary` - Primary button
- `.btn-secondary` - Secondary button
- `.badge-*` - Status badges
- `.input` - Form inputs

## Contributing

1. Follow the existing code style
2. Run linting and formatting before committing
3. Write tests for new features
4. Update documentation as needed

## License

Proprietary - All rights reserved
