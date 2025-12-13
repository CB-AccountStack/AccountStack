package services

import (
	"sort"
	"time"

	"github.com/CB-AccountStack/AccountStack/apps/api-transactions/internal/features"
	"github.com/CB-AccountStack/AccountStack/apps/api-transactions/internal/models"
	"github.com/CB-AccountStack/AccountStack/apps/api-transactions/internal/repository"
	"github.com/sirupsen/logrus"
)

// TransactionService handles business logic for transactions
type TransactionService struct {
	repo   *repository.Repository
	flags  *features.Flags
	logger *logrus.Logger
}

// NewTransactionService creates a new transaction service
func NewTransactionService(repo *repository.Repository, flags *features.Flags, logger *logrus.Logger) *TransactionService {
	return &TransactionService{
		repo:   repo,
		flags:  flags,
		logger: logger,
	}
}

// GetTransactionByID retrieves a transaction by ID
func (s *TransactionService) GetTransactionByID(txnID string) (*models.Transaction, error) {
	return s.repo.GetTransactionByID(txnID)
}

// GetTransactions retrieves transactions with optional filters
// Advanced filters (date range, amount range, category) are only applied if feature flag is enabled
func (s *TransactionService) GetTransactions(filters *models.TransactionFilters) ([]*models.Transaction, error) {
	// Check if advanced filters are enabled
	advancedFiltersEnabled := s.flags.IsAdvancedFiltersEnabled()

	// Create a new filter set based on feature flag
	effectiveFilters := &models.TransactionFilters{
		AccountID: filters.AccountID,
	}

	// Only apply advanced filters if the feature flag is enabled
	if advancedFiltersEnabled {
		effectiveFilters.StartDate = filters.StartDate
		effectiveFilters.EndDate = filters.EndDate
		effectiveFilters.Category = filters.Category
		effectiveFilters.MinAmount = filters.MinAmount
		effectiveFilters.MaxAmount = filters.MaxAmount

		s.logger.WithFields(logrus.Fields{
			"accountId": filters.AccountID,
			"startDate": filters.StartDate,
			"endDate":   filters.EndDate,
			"category":  filters.Category,
			"minAmount": filters.MinAmount,
			"maxAmount": filters.MaxAmount,
		}).Debug("Using advanced filters")
	} else {
		// Log that advanced filters are ignored
		if filters.StartDate != nil || filters.EndDate != nil || filters.Category != "" ||
			filters.MinAmount != nil || filters.MaxAmount != nil {
			s.logger.Info("Advanced filters requested but feature flag is disabled, only accountId filter will be applied")
		}

		s.logger.WithField("accountId", filters.AccountID).Debug("Using basic filters only")
	}

	// Get filtered transactions
	transactions := s.repo.GetTransactionsByFilter(effectiveFilters)

	// Sort by date descending (most recent first)
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date.After(transactions[j].Date)
	})

	s.logger.WithField("count", len(transactions)).Info("Retrieved transactions")

	return transactions, nil
}

// ParseDateParam parses a date string in ISO 8601 format
func ParseDateParam(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}

	// Try parsing as RFC3339 (ISO 8601)
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		// Try parsing as date only (YYYY-MM-DD)
		t, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, err
		}
	}

	return &t, nil
}
