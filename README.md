# AccountStack

A forkable, mono-repo reference application demonstrating modern CI/CD, feature management, governance, and SmartTests in regulated enterprise environments.

## Overview

AccountStack is a realistic **account overview portal** similar to those used by banks, insurers, utilities, and large enterprises — designed for enablement, workshops, and executive demonstrations.

**Key Features**:
- Multi-component architecture (UI + multiple APIs)
- Local-first development (works offline)
- Built-in feature management (local + CloudBees FM integration)
- Comprehensive test coverage for SmartTests demonstrations
- Fork-friendly and customizable
- CloudBees Unify CI/CD workflows

## Quick Start

Run everything locally:

```bash
docker compose up --build
```

This starts:
- React UI (port 3000)
- Accounts API (port 8001)
- Transactions API (port 8002)
- Insights API (port 8003)

Access the application at: http://localhost:3000

## Components

- **apps/web** - React UI
- **apps/api-accounts** - Accounts service
- **apps/api-transactions** - Transactions service
- **apps/api-insights** - Insights and analytics service
- **config/** - Feature flag reference documentation

## Feature Management

AccountStack uses **CloudBees Feature Management** for all feature flags with **real-time updates**.

**Key Features**:
- **Real-time flag updates** - Changes propagate instantly without reload
- **CloudBees FM SDK** - Integrated in all components
- **Demo-resilient** - Falls back to hardcoded defaults when offline
- **No local FM system needed** - Flags managed in CloudBees FM

**How it works**:
- All flags defined in CloudBees Feature Management
- Components use CloudBees FM SDK (Rox)
- Flag changes visible immediately in UI/API (no reload required)
- Perfect for live demos - toggle flags and see instant changes

**Offline mode**:
- Works without internet connection
- Uses hardcoded default values
- No degraded experience

See [Feature Flags Reference](config/README.md) for complete flag list and code examples.

## Testing

Run all tests locally:

```bash
make test
```

Or run specific test suites:

```bash
make test-unit          # ~200-300 unit tests
make test-integration   # ~80-100 integration tests
make test-e2e          # ~30-40 end-to-end tests
```

High test volume demonstrates CloudBees SmartTests impact analysis and test subsetting.

## Documentation

- [Architecture Details](docs/ARCHITECTURE.md) - Component design, tech stack, and customization guide
- [Setup Guide](docs/SETUP.md) - Detailed local and CloudBees configuration (coming soon)
- [Changelog](CHANGELOG.md) - Version history and migration guides

## Customization

AccountStack is designed to be forked and customized:

1. Fork to your organization
2. Update `config/features.yaml` with your branding
3. Application name, colors, and logos are feature-flag controlled
4. Repo and org names are auto-detected from git remote

See [Architecture Documentation](docs/ARCHITECTURE.md) for details.

## Technology Stack

- **Frontend**: React, TypeScript, Vite
- **Backend**: Go (APIs)
- **Infrastructure**: Docker, Docker Compose, Kubernetes
- **CI/CD**: CloudBees Unify
- **Feature Management**: CloudBees Feature Management (optional)
- **Testing**: Jest, Playwright, Go testing

## Positioning

**AccountStack** is for:
- Enablement and training
- Customer workshops
- Executive demonstrations
- Fast iteration and prototyping

For deep, multi-service platform demonstrations, see **SquidStack**.

## Non-Goals

- No payments or PCI complexity
- No trading functionality
- No microservice sprawl
- Simple, focused domain model

## License

MIT

## Support

For issues or questions, please open a GitHub issue or contact the CloudBees SE team.
