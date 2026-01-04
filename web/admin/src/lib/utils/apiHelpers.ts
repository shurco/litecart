import { apiGet, apiPost, apiUpdate, apiDelete, showMessage } from './index'
import type { ApiResponse } from '$lib/types/api'

export async function handleApiCall<T>(
  apiCall: () => Promise<ApiResponse<T>>,
  successMessage?: string,
  errorMessage?: string
): Promise<T | null> {
  try {
    const res = await apiCall()
    if (res.success) {
      if (successMessage) {
        showMessage(successMessage || res.message, 'connextSuccess')
      }
      return res.result || null
    } else {
      showMessage(res.message || errorMessage || 'Operation failed', 'connextError')
      return null
    }
  } catch (error) {
    showMessage(errorMessage || 'Network error', 'connextError')
    return null
  }
}

export async function loadData<T>(url: string, errorMessage = 'Failed to load data'): Promise<T | null> {
  return handleApiCall(() => apiGet<T>(url), undefined, errorMessage)
}

export async function saveData<T, D = Partial<T>>(
  url: string,
  data: D,
  isUpdate: boolean,
  successMessage = 'Data saved',
  errorMessage = 'Failed to save data'
): Promise<T | null> {
  const apiCall = isUpdate ? () => apiUpdate<T>(url, data) : () => apiPost<T>(url, data)
  return handleApiCall(apiCall, successMessage, errorMessage)
}

export async function deleteData(
  url: string,
  successMessage = 'Deleted successfully',
  errorMessage = 'Failed to delete'
): Promise<boolean> {
  const result = await handleApiCall(() => apiDelete(url), successMessage, errorMessage)
  return result !== null
}

export async function toggleActive<T = any>(
  url: string,
  successMessage = 'Status updated',
  errorMessage = 'Failed to update status'
): Promise<T | null> {
  try {
    const res = await apiUpdate<T>(url, {})
    if (res.success) {
      if (successMessage) {
        showMessage(successMessage || res.message, 'connextSuccess')
      }
      return res.result || null
    } else {
      showMessage(res.message || errorMessage, 'connextError')
      return null
    }
  } catch (error) {
    showMessage(errorMessage || 'Network error', 'connextError')
    return null
  }
}
