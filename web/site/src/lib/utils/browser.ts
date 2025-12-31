/**
 * Browser environment utilities
 */

/**
 * Checks if code is running in browser environment
 * @returns true if running in browser
 */
export function isBrowser(): boolean {
  return typeof window !== 'undefined'
}

/**
 * Safely gets value from localStorage
 * @param key - Storage key
 * @param defaultValue - Default value if key doesn't exist
 * @returns Stored value or default
 */
export function getLocalStorage(key: string, defaultValue: string = ''): string {
  if (!isBrowser()) return defaultValue

  try {
    return localStorage.getItem(key) || defaultValue
  } catch {
    return defaultValue
  }
}

/**
 * Safely sets value to localStorage
 * @param key - Storage key
 * @param value - Value to store
 */
export function setLocalStorage(key: string, value: string): void {
  if (!isBrowser()) return

  try {
    localStorage.setItem(key, value)
  } catch {
    // Ignore storage errors
  }
}

/**
 * Safely removes value from localStorage
 * @param key - Storage key
 */
export function removeLocalStorage(key: string): void {
  if (!isBrowser()) return

  try {
    localStorage.removeItem(key)
  } catch {
    // Ignore storage errors
  }
}
