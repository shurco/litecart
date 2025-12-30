import { writable } from 'svelte/store'
import type { Settings } from '$lib/types/models'

const CACHE_DURATION = 5 * 60 * 1000 // 5 minutes

function createSettingsStore() {
  const { subscribe, set, update } = writable<Settings | null>(null)

  return {
    subscribe,
    set,
    update,
    loadFromCache: (): Settings | null => {
      if (typeof window === 'undefined') return null

      const cached = sessionStorage.getItem('settings')
      const timestamp = sessionStorage.getItem('settings_timestamp')

      if (cached && timestamp) {
        const now = Date.now()
        const cacheTime = parseInt(timestamp, 10)

        if (now < cacheTime) {
          try {
            return JSON.parse(cached)
          } catch {
            return null
          }
        }
      }

      return null
    },
    saveToCache: (settings: Settings) => {
      if (typeof window === 'undefined') return

      const expiry = Date.now() + CACHE_DURATION
      sessionStorage.setItem('settings', JSON.stringify(settings))
      sessionStorage.setItem('settings_timestamp', expiry.toString())
    },
    clearCache: () => {
      if (typeof window === 'undefined') return

      sessionStorage.removeItem('settings')
      sessionStorage.removeItem('settings_timestamp')
    }
  }
}

export const settingsStore = createSettingsStore()
