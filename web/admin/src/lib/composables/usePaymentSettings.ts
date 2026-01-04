import { apiGet, apiUpdate, showMessage } from '$lib/utils'
import { systemStore } from '$lib/stores/system'
import { t } from '$lib/i18n'
import type { StripeSettings, PaypalSettings, SpectrocoinSettings } from '$lib/types/models'

export async function loadPaymentSettings<T extends StripeSettings | PaypalSettings | SpectrocoinSettings>(
  endpoint: string,
  defaultSettings: T
): Promise<T> {
  try {
    const res = await apiGet<T>(`/api/_/settings/${endpoint}`)
    if (res.success && res.result) {
      return { ...defaultSettings, ...res.result }
    }
    showMessage(res.message || t('settings.failedToLoadSettings'), 'connextError')
    return defaultSettings
  } catch (error) {
    showMessage(t('common.networkError'), 'connextError')
    return defaultSettings
  }
}

export async function savePaymentSettings<T extends StripeSettings | PaypalSettings | SpectrocoinSettings>(
  endpoint: string,
  settings: T,
  providerName: 'stripe' | 'paypal' | 'spectrocoin'
): Promise<boolean> {
  try {
    const update = { [providerName]: settings }
    const res = await apiUpdate(`/api/_/settings/${endpoint}`, update)
    if (res.success) {
      showMessage(res.message || t('settings.settingsSaved'), 'connextSuccess')
      return true
    }
    showMessage(res.message || t('settings.failedToSaveSettings'), 'connextError')
    return false
  } catch (error) {
    showMessage(t('common.networkError'), 'connextError')
    return false
  }
}

export async function togglePaymentActive(
  providerName: 'stripe' | 'paypal' | 'spectrocoin',
  active: boolean
): Promise<boolean> {
  try {
    const res = await apiUpdate(`/api/_/settings/${providerName}_active`, {
      value: active
    })
    if (res.success) {
      systemStore.update((store) => ({
        ...store,
        payments: {
          ...store.payments,
          [providerName]: active
        }
      }))
      showMessage(res.message || t('common.statusUpdated'), 'connextSuccess')
      return true
    }
    showMessage(res.message || t('common.failedToUpdateStatus'), 'connextError')
    return false
  } catch (error) {
    showMessage(t('common.networkError'), 'connextError')
    return false
  }
}
