export function costFormat(amount: string | number | null | undefined): string {
  if (!amount) return '0.00'
  // Цены в базе хранятся в центах, поэтому делим на 100 для отображения
  return (parseFloat(String(amount)) / 100).toFixed(2)
}
