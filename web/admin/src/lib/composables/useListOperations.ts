interface EntityWithId {
  id: string
}

/**
 * Composable for managing list operations (update, delete, toggle)
 * Eliminates code duplication across different entity lists
 */
export function useListOperations<T extends EntityWithId>(items: T[]) {
  function updateItem(updatedItem: T): void {
    const index = items.findIndex((item) => item.id === updatedItem.id)
    if (index !== -1) {
      items[index] = updatedItem
    }
  }

  function removeItem(id: string): void {
    const index = items.findIndex((item) => item.id === id)
    if (index !== -1) {
      items.splice(index, 1)
    }
  }

  function toggleItemActive(id: string, currentActive: boolean): void {
    const index = items.findIndex((item) => item.id === id)
    if (index !== -1) {
      items[index] = { ...items[index], active: !currentActive } as T
    }
  }

  function getItem(id: string): T | undefined {
    return items.find((item) => item.id === id)
  }

  return {
    updateItem,
    removeItem,
    toggleItemActive,
    getItem
  }
}
