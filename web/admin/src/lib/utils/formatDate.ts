export function formatDate(dateValue: string | number | null | undefined): string {
  if (!dateValue) return ''
  try {
    // Handle Unix timestamp (number in seconds)
    let date: Date
    if (typeof dateValue === 'number') {
      // Convert seconds to milliseconds for Date constructor
      date = new Date(dateValue * 1000)
    } else {
      date = new Date(dateValue)
    }
    return date.toLocaleDateString('ru-RU', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    })
  } catch (error) {
    return String(dateValue)
  }
}
