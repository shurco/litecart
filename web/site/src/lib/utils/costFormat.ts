export function costFormat(cost: number): string {
  if (!cost || cost === 0) return 'free'
  return (cost / 100).toFixed(2)
}
