import { apiGet, apiUpdate, showMessage } from "$lib/utils";
import { systemStore } from "$lib/stores/system";
import type {
  StripeSettings,
  PaypalSettings,
  SpectrocoinSettings,
} from "$lib/types/models";

export async function loadPaymentSettings<
  T extends StripeSettings | PaypalSettings | SpectrocoinSettings,
>(endpoint: string, defaultSettings: T): Promise<T> {
  try {
    const res = await apiGet<T>(`/api/_/settings/${endpoint}`);
    if (res.success && res.result) {
      return { ...defaultSettings, ...res.result };
    }
    showMessage(res.message || "Failed to load settings", "connextError");
    return defaultSettings;
  } catch (error) {
    showMessage("Network error", "connextError");
    return defaultSettings;
  }
}

export async function savePaymentSettings<
  T extends StripeSettings | PaypalSettings | SpectrocoinSettings,
>(
  endpoint: string,
  settings: T,
  providerName: "stripe" | "paypal" | "spectrocoin",
): Promise<boolean> {
  try {
    const update = { [providerName]: settings };
    const res = await apiUpdate(`/api/_/settings/${endpoint}`, update);
    if (res.success) {
      showMessage(res.message || "Settings saved", "connextSuccess");
      return true;
    }
    showMessage(res.message || "Failed to save settings", "connextError");
    return false;
  } catch (error) {
    showMessage("Network error", "connextError");
    return false;
  }
}

export async function togglePaymentActive(
  providerName: "stripe" | "paypal" | "spectrocoin",
  active: boolean,
): Promise<boolean> {
  try {
    const res = await apiUpdate(`/api/_/settings/${providerName}_active`, {
      value: active,
    });
    if (res.success) {
      systemStore.update((store) => ({
        ...store,
        payments: {
          ...store.payments,
          [providerName]: active,
        },
      }));
      showMessage(res.message || "Status updated", "connextSuccess");
      return true;
    }
    showMessage(res.message || "Failed to update status", "connextError");
    return false;
  } catch (error) {
    showMessage("Network error", "connextError");
    return false;
  }
}
