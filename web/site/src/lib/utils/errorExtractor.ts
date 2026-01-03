/**
 * Extracts error message from API response
 * Centralizes error message extraction logic to eliminate duplication
 */
export function extractErrorMessage(data: any): string {
  if (data.result && typeof data.result === 'string' && data.result.trim()) {
    return data.result
  }

  if (data.result && typeof data.result === 'object') {
    if (data.result.message && typeof data.result.message === 'string') {
      return data.result.message
    }
    if (data.result.error && typeof data.result.error === 'string') {
      return data.result.error
    }
  }

  if (data.message && typeof data.message === 'string') {
    return data.message
  }

  return 'Request failed'
}
