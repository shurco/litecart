/**
 * Cookie consent management utilities
 * Handles GDPR-compliant cookie consent storage and retrieval
 */

import { isBrowser, getLocalStorage, setLocalStorage } from './browser'

const COOKIE_CONSENT_KEY = 'cookie_consent'
const COOKIE_CONSENT_VERSION = '1.0'

export type CookieConsentStatus = 'accepted' | 'rejected' | null

export interface CookieConsentData {
  status: CookieConsentStatus
  timestamp: number
  version: string
}

/**
 * Gets current cookie consent status
 * @returns Cookie consent data or null if not set
 */
export function getCookieConsent(): CookieConsentData | null {
  if (!isBrowser()) return null

  try {
    const stored = getLocalStorage(COOKIE_CONSENT_KEY)
    if (!stored) return null

    const data: CookieConsentData = JSON.parse(stored)
    return data
  } catch {
    return null
  }
}

/**
 * Checks if user has given consent (accepted or rejected)
 * @returns true if consent has been given
 */
export function hasCookieConsent(): boolean {
  const consent = getCookieConsent()
  return consent !== null && consent.status !== null
}

/**
 * Sets cookie consent status
 * @param status - Consent status (accepted or rejected)
 */
export function setCookieConsent(status: 'accepted' | 'rejected'): void {
  if (!isBrowser()) return

  const data: CookieConsentData = {
    status,
    timestamp: Date.now(),
    version: COOKIE_CONSENT_VERSION
  }

  try {
    setLocalStorage(COOKIE_CONSENT_KEY, JSON.stringify(data))
  } catch {
    // Ignore storage errors
  }
}

/**
 * Clears cookie consent (allows banner to show again)
 */
export function clearCookieConsent(): void {
  if (!isBrowser()) return

  try {
    localStorage.removeItem(COOKIE_CONSENT_KEY)
  } catch {
    // Ignore storage errors
  }
}
