import { redirect, type Handle } from '@sveltejs/kit'

const BASE_PATH = '/_'

export const handle: Handle = async ({ event, resolve }) => {
  const { url } = event

  // Allow install and signin pages without authentication
  // pathname includes base path, so we check with base
  if (url.pathname === `${BASE_PATH}/install` || url.pathname === `${BASE_PATH}/signin`) {
    return resolve(event)
  }

  // For other routes, authentication check will be done client-side
  // or via API calls in the page components
  // This is a simplified version - in production you might want
  // to check cookies/session here

  return resolve(event)
}
