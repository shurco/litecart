export function costFormat(amount: string | number | null | undefined): string {
  if (!amount) return '0.00'
  // Prices in database are stored in cents, so divide by 100 for display
  return (parseFloat(String(amount)) / 100).toFixed(2)
}

export function formatPrice(amount: string | number | null | undefined): string {
  if (!amount || parseFloat(String(amount)) === 0) return 'free'
  // Prices in database are stored in cents, so divide by 100 for display
  return (parseFloat(String(amount)) / 100).toFixed(2)
}
