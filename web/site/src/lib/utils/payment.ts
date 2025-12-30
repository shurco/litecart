/**
 * Utility for working with payment providers
 */

import type { PaymentMethods } from '$lib/types/models'

/**
 * Gets available payment providers
 * @param payments - Object with available providers
 * @returns Array of available providers
 */
export function getAvailableProviders(payments: PaymentMethods): string[] {
  const providers: string[] = []
  if (payments.stripe) providers.push('stripe')
  if (payments.paypal) providers.push('paypal')
  if (payments.spectrocoin) providers.push('spectrocoin')
  return providers
}

/**
 * Automatically selects provider if only one is available
 * @param payments - Object with available providers
 * @returns Provider name or empty string
 */
export function autoSelectProvider(payments: PaymentMethods): string {
  const providers = getAvailableProviders(payments)
  return providers.length === 1 ? providers[0] : ''
}

/**
 * Checks if any payment providers are available
 * @param payments - Object with available providers
 * @returns true if at least one provider is available
 */
export function hasPaymentProviders(payments: PaymentMethods): boolean {
  return !!(payments.stripe || payments.paypal || payments.spectrocoin)
}
