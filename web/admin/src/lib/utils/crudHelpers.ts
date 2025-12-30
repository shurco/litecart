import { loadData, saveData, deleteData } from "./apiHelpers";

export interface CrudOperations<T> {
  load: () => Promise<T[] | null>;
  create: (data: Partial<T>) => Promise<T | null>;
  update: (id: string, data: Partial<T>) => Promise<T | null>;
  delete: (id: string) => Promise<boolean>;
  toggleActive?: (id: string) => Promise<boolean>;
}

export function createCrudOperations<T extends { id: string }>(
  baseUrl: string,
  entityName: string,
): CrudOperations<T> {
  return {
    load: async () => {
      return loadData<T[]>(baseUrl, `Failed to load ${entityName}`);
    },
    create: async (data: Partial<T>) => {
      return saveData<T, Partial<T>>(
        baseUrl,
        data,
        false,
        `${entityName} created`,
        `Failed to create ${entityName}`,
      );
    },
    update: async (id: string, data: Partial<T>) => {
      return saveData<T, Partial<T>>(
        `${baseUrl}/${id}`,
        data,
        true,
        `${entityName} updated`,
        `Failed to update ${entityName}`,
      );
    },
    delete: async (id: string) => {
      return deleteData(
        `${baseUrl}/${id}`,
        `${entityName} deleted`,
        `Failed to delete ${entityName}`,
      );
    },
    toggleActive: async (id: string) => {
      const result = await saveData<T, Record<string, never>>(
        `${baseUrl}/${id}/active`,
        {},
        true,
        `${entityName} status updated`,
        `Failed to update ${entityName} status`,
      );
      return result !== null;
    },
  };
}
