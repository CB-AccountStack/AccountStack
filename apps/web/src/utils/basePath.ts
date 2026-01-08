/**
 * Runtime base path detection utility
 *
 * Detects the deployment base path from the browser's URL at runtime.
 * This allows the same Docker image to be deployed to multiple environments
 * with different URL paths without rebuilding.
 *
 * Examples:
 * - https://domain.com/CB-AccountStack/account-stack-dev/ → "/CB-AccountStack/account-stack-dev"
 * - https://domain.com/CB-AccountStack/account-stack-preprod/ → "/CB-AccountStack/account-stack-preprod"
 * - http://localhost:3000/ → "/"
 */

/**
 * Detects the base path from the current URL at runtime.
 * Matches patterns like /{org}/{env} (e.g., /CB-AccountStack/account-stack-dev)
 * Returns "/" for local development or root deployments.
 */
export function detectBasePath(): string {
  if (typeof window === 'undefined') {
    return '/';
  }

  const pathname = window.location.pathname;

  // Match pattern: /{org}/{env}/... or just /{org}/{env}
  // org and env can contain letters, numbers, hyphens, underscores
  const match = pathname.match(/^\/[a-zA-Z0-9_-]+\/[a-zA-Z0-9_-]+/);

  if (match) {
    console.log('[BasePath] Detected base path:', match[0]);
    return match[0]; // Returns something like "/CB-AccountStack/account-stack-dev"
  }

  console.log('[BasePath] Using root path for local/root deployment');
  return '/'; // Local development or root deployment
}

/**
 * Gets the API base URL using the detected base path.
 *
 * For local development (basePath="/"):
 *   Returns "api" (relative path, proxied by Vite dev server)
 *
 * For deployed environments (basePath="/{org}/{env}"):
 *   Returns "/{org}/{env}/api" (absolute path, proxied by nginx)
 */
export function getApiBaseUrl(): string {
  const basePath = detectBasePath();

  if (basePath === '/') {
    return 'api'; // Relative path for local development
  }

  const apiUrl = `${basePath}/api`;
  console.log('[BasePath] API base URL:', apiUrl);
  return apiUrl; // Absolute path for deployed environments
}
