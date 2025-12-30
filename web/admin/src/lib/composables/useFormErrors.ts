import { writable } from "svelte/store";

export function useFormErrors() {
  const errors = writable<Record<string, string>>({});

  const setError = (field: string, message: string) => {
    errors.update((current) => ({
      ...current,
      [field]: message,
    }));
  };

  const clearError = (field: string) => {
    errors.update((current) => {
      const updated = { ...current };
      delete updated[field];
      return updated;
    });
  };

  const clearAll = () => {
    errors.set({});
  };

  const getErrors = () => {
    let currentErrors: Record<string, string> = {};
    const unsubscribe = errors.subscribe((errs) => {
      currentErrors = errs;
    });
    unsubscribe();
    return currentErrors;
  };

  const hasErrors = () => {
    return Object.keys(getErrors()).length > 0;
  };

  return {
    errors: { subscribe: errors.subscribe },
    setError,
    clearError,
    clearAll,
    hasErrors,
  };
}
