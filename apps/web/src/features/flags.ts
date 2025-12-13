// CloudBees Feature Management (Rox) integration
import Rox, { type FetcherResults, type RoxSetupOptions } from 'rox-browser';

// Define feature flags with default values
export class FeatureFlags {
  // UI Dashboard Cards V2 - Enhanced card design with better visuals
  public dashboardCardsV2 = new Rox.Flag(true);

  // UI Insights V2 - New insights panel with improved analytics
  public insightsV2 = new Rox.Flag(false);

  // UI Alerts Banner - Top banner for important alerts
  public alertsBanner = new Rox.Flag(true);

  // UI Transactions Filters - Advanced filtering for transactions
  public transactionsFilters = new Rox.Flag(true);

  // Kill switch for insights feature
  public killInsights = new Rox.Flag(false);
}

// Create feature flags instance
export const flags = new FeatureFlags();

// Configuration for CloudBees FM
interface RoxConfig {
  apiKey?: string;
  devModeSecret?: string;
}

// Initialize Rox with the feature flags
export async function initializeFeatureFlags(config: RoxConfig = {}): Promise<void> {
  // Register the feature flags container
  Rox.register('accountstack', flags);

  // Setup Rox with configuration
  const roxConfig: RoxSetupOptions = {
    debugLevel: 'verbose',
    // You can add custom properties here for targeting
    configurationFetchedHandler: (fetcherResults: FetcherResults) => {
      console.log('[FeatureFlags] Configuration fetched:', {
        hasChanges: fetcherResults.hasChanges,
        source: fetcherResults.fetcherStatus,
      });
    },
  };

  try {
    // Setup Rox - in development, this will use default values
    // In production, provide apiKey via environment variable
    const apiKey = config.apiKey || import.meta.env.VITE_ROX_API_KEY || '';

    if (apiKey) {
      await Rox.setup(apiKey, roxConfig);
      console.log('[FeatureFlags] CloudBees FM initialized successfully');
    } else {
      console.warn(
        '[FeatureFlags] No API key provided, using default flag values. ' +
        'Set VITE_ROX_API_KEY environment variable to connect to CloudBees FM.'
      );
      // In dev mode without API key, we can still use the default values
      await Rox.setup('', roxConfig);
    }
  } catch (error) {
    console.error('[FeatureFlags] Failed to initialize CloudBees FM:', error);
    // Continue with default values if setup fails
  }
}

// Helper functions to check flag values
export function isDashboardCardsV2Enabled(): boolean {
  return flags.dashboardCardsV2.isEnabled();
}

export function isInsightsV2Enabled(): boolean {
  return flags.insightsV2.isEnabled() && !flags.killInsights.isEnabled();
}

export function isAlertsBannerEnabled(): boolean {
  return flags.alertsBanner.isEnabled();
}

export function isTransactionsFiltersEnabled(): boolean {
  return flags.transactionsFilters.isEnabled();
}

export function isInsightsKilled(): boolean {
  return flags.killInsights.isEnabled();
}

// Hook for React components to use feature flags
export function useFeatureFlags() {
  return {
    dashboardCardsV2: isDashboardCardsV2Enabled(),
    insightsV2: isInsightsV2Enabled(),
    alertsBanner: isAlertsBannerEnabled(),
    transactionsFilters: isTransactionsFiltersEnabled(),
    killInsights: isInsightsKilled(),
  };
}
