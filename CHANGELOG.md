# Changelog

All notable changes to AccountStack will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- React UI implementation
- Go API services (accounts, transactions, insights)
- Local feature management system
- Docker Compose setup
- CloudBees Unify workflows
- Comprehensive test suite (300+ tests)
- CloudBees FM integration

---

## [0.1.0] - 2025-12-13

### Added
- Initial repository structure
- Documentation framework:
  - `README.md` - Project overview and quick start
  - `docs/ARCHITECTURE.md` - Detailed technical architecture
  - `CHANGELOG.md` - Version history (this file)
- Directory structure for mono-repo:
  - `apps/web/` - React UI (placeholder)
  - `apps/api-accounts/` - Accounts API (placeholder)
  - `apps/api-transactions/` - Transactions API (placeholder)
  - `apps/api-insights/` - Insights API (placeholder)
  - `shared/config/` - Feature flags and configuration
  - `tests/` - Test suites organization
  - `.cloudbees/workflows/` - CI/CD workflows (placeholder)
  - `data/seed/` - Demo data (placeholder)

### Architecture Decisions
- **Local-first design**: All functionality works offline by default
- **Multi-component architecture**: Separate deployable services (UI + 3 APIs)
- **Feature management**: Built-in flag system with optional CloudBees FM integration
- **Fork-friendly**: Auto-detection of org/repo names, no hardcoded references
- **Test-rich**: Target 300-440 tests for SmartTests demonstrations

### Design Principles
1. Local-first (works offline)
2. Demo-ready (impressive out of the box)
3. Fork-friendly (easy customization)
4. Test-rich (high volume, clear impact mapping)
5. Production-like (realistic architecture)
6. Simple but not trivial

---

## Version History

### Version Numbering

AccountStack uses semantic versioning:
- **MAJOR**: Incompatible API or architecture changes
- **MINOR**: New features, backwards-compatible
- **PATCH**: Bug fixes, documentation updates

### Release Cadence

- **Development**: Continuous integration to `main` branch
- **Releases**: Tagged when significant milestones achieved
- **Hotfixes**: As needed for critical issues

---

## Migration Guides

### Migrating to Future Versions

Migration guides will be added here as breaking changes are introduced.

---

## Links

- [Repository](https://github.com/CB-AccountStack/AccountStack)
- [Architecture Documentation](docs/ARCHITECTURE.md)
- [Setup Guide](docs/SETUP.md) (coming soon)
- [CloudBees Documentation](https://docs.cloudbees.com/)

---

## Contributors

AccountStack is maintained by the CloudBees SE team.

Special thanks to all contributors who help improve AccountStack for the community.
