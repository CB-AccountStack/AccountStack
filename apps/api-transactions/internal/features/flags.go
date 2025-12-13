package features

import (
	"sync"

	"github.com/rollout/rox-go/v5/core/model"
	"github.com/rollout/rox-go/v5/core/roxx"
	"github.com/sirupsen/logrus"
)

// Flags holds all feature flags for the application
type Flags struct {
	AdvancedFilters model.RoxFlag
	logger          *logrus.Logger
	mu              sync.RWMutex
}

var flags *Flags

// Initialize sets up feature flags with CloudBees Feature Management
func Initialize(apiKey string, logger *logrus.Logger) (*Flags, error) {
	flags = &Flags{
		logger: logger,
	}

	// Register feature flag: api.advancedFilters (default: false)
	// When false: only accountId filter is allowed
	// When true: all filters are allowed (date range, amount range, category)
	flags.AdvancedFilters = model.NewRoxFlag(false)

	// Register flags with CloudBees
	roxx.Register("api", flags)

	if apiKey != "" && apiKey != "dev-mode" {
		// Setup Rox with API key
		options := roxx.NewRoxOptions(roxx.RoxOptionsBuilder{})
		<-roxx.Setup(apiKey, options)

		logger.Info("CloudBees Feature Management initialized successfully")

		// Fetch latest feature flags
		go func() {
			roxx.Fetch()
			logger.Info("Initial feature flags fetched")
		}()
	} else {
		logger.Warn("CloudBees Feature Management API key not provided, using default flag values")
	}

	logger.WithFields(logrus.Fields{
		"advancedFilters": flags.IsAdvancedFiltersEnabled(),
	}).Info("Feature flags initialized")

	return flags, nil
}

// GetFlags returns the global flags instance
func GetFlags() *Flags {
	return flags
}

// IsAdvancedFiltersEnabled returns whether advanced filters are enabled
func (f *Flags) IsAdvancedFiltersEnabled() bool {
	if f == nil || f.AdvancedFilters == nil {
		return false
	}
	return f.AdvancedFilters.IsEnabled(nil)
}

// Shutdown gracefully shuts down the feature management system
func Shutdown() {
	if flags != nil {
		roxx.Shutdown()
		flags.logger.Info("CloudBees Feature Management shutdown complete")
	}
}
