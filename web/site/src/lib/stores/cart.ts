import { writable } from 'svelte/store'
import type { CartItem } from '$lib/types/models'

const CART_STORAGE_KEY = 'cart'

function createCartStore() {
  const loadFromStorage = (): CartItem[] => {
    if (typeof window === 'undefined') return []

    try {
      const stored = localStorage.getItem(CART_STORAGE_KEY)
      return stored ? JSON.parse(stored) : []
    } catch {
      return []
    }
  }

  const saveToStorage = (items: CartItem[]) => {
    if (typeof window === 'undefined') return

    localStorage.setItem(CART_STORAGE_KEY, JSON.stringify(items))
  }

  const { subscribe, set, update } = writable<CartItem[]>(loadFromStorage())

  return {
    subscribe,
    add: (item: CartItem) => {
      update((items) => {
        if (items.find((i) => i.id === item.id)) {
          return items
        }
        const newItems = [...items, item]
        saveToStorage(newItems)
        return newItems
      })
    },
    remove: (id: string) => {
      update((items) => {
        const newItems = items.filter((item) => item.id !== id)
        saveToStorage(newItems)
        return newItems
      })
    },
    clear: () => {
      set([])
      saveToStorage([])
    }
  }
}

export const cartStore = createCartStore()
