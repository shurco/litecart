export interface ApiResponse<T = any> {
  success: boolean
  message?: string
  result?: T
}

export interface RequestOptions extends RequestInit {
  credentials?: RequestCredentials
  method: 'GET' | 'POST' | 'PATCH' | 'DELETE'
  body?: string | FormData
  headers?: HeadersInit
}
