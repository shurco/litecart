import { writable } from "svelte/store";

export function useLoading() {
  const loading = writable(false);

  const setLoading = (value: boolean) => {
    loading.set(value);
  };

  const withLoading = async <T>(fn: () => Promise<T>): Promise<T> => {
    setLoading(true);
    try {
      return await fn();
    } finally {
      setLoading(false);
    }
  };

  return {
    loading: { subscribe: loading.subscribe },
    setLoading,
    withLoading,
  };
}
