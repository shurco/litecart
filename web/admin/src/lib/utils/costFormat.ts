export function costFormat(amount: string | number | null | undefined): string {
  if (!amount) return '0.00'
  return parseFloat(String(amount)).toFixed(2)
}
