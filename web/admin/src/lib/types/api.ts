export interface ApiResponse<T = any> {
  success: boolean
  message?: string
  result?: T
}

export interface ApiError {
  success: false
  message: string
  result?: any
}

export type RequestMethod = 'GET' | 'POST' | 'PATCH' | 'DELETE'

export interface RequestOptions extends RequestInit {
  credentials?: RequestCredentials
  method: RequestMethod
  body?: string | FormData
  headers?: HeadersInit
}
