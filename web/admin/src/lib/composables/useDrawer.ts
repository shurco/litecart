import { writable } from "svelte/store";

export function useDrawer<T extends string = string>(defaultMode: T) {
  const isOpen = writable(false);
  const mode = writable<T | null>(null);

  const open = (newMode: T) => {
    mode.set(newMode);
    isOpen.set(true);
  };

  const close = () => {
    isOpen.set(false);
    setTimeout(() => {
      mode.set(null);
    }, 200);
  };

  return {
    isOpen: { subscribe: isOpen.subscribe },
    mode: { subscribe: mode.subscribe },
    open,
    close,
  };
}
