/**
 * Navigation utilities
 */

import { goto } from '$app/navigation'

/**
 * Handles navigation with preventDefault
 * @param e - Mouse event
 * @param path - Path to navigate to
 */
export function handleNavigation(e: MouseEvent, path: string): void {
  e.preventDefault()
  goto(path)
}
