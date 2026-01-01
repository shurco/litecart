import { loadData, saveData } from './apiHelpers'
import { apiUpdate } from './api'

export async function loadSettings<T>(endpoint: string, defaultData: T): Promise<T> {
  const result = await loadData<T>(`/api/_/settings/${endpoint}`, 'Failed to load settings')
  return result || defaultData
}

export async function saveSettings<T>(endpoint: string, data: T, successMessage = 'Settings saved'): Promise<boolean> {
  const result = await saveData<T>(`/api/_/settings/${endpoint}`, data, true, successMessage, 'Failed to save settings')
  return result !== null
}
