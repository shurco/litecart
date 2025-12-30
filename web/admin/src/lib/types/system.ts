export interface SystemVersion {
  current_version?: string;
  new?: string;
  release_url?: string;
}

export interface SystemStore {
  version: SystemVersion;
  payments: Record<string, any>;
}
