import { writable } from 'svelte/store'
import type { SystemStore } from '$lib/types/system'

export const systemStore = writable<SystemStore>({
  version: {},
  payments: {}
})
