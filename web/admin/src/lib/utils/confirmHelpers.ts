export function confirmAction(message: string): boolean {
  return confirm(message);
}

export function confirmDelete(entityName: string, name: string): boolean {
  return confirmAction(`Are you sure you want to delete "${name}"?`);
}
