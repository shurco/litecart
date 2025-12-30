export function costFormat(cost: number): string {
  return cost ? (cost / 100).toFixed(2) : '0.00'
}
