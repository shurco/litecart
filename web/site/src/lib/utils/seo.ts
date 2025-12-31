/**
 * Utility for working with SEO meta tags
 */

export interface SEOData {
  title?: string
  keywords?: string
  description?: string
}

import { isBrowser } from './browser'

/**
 * Updates SEO meta tags on the page
 * @param seo - SEO data
 */
export function updateSEOTags(seo: SEOData): void {
  if (!isBrowser()) return

  if (seo.title) {
    document.title = seo.title
    updateMetaTag('meta[name="title"]', 'content', seo.title)
    updateMetaTag('meta[property="og:title"]', 'content', seo.title)
  }

  if (seo.keywords) {
    updateMetaTag('meta[name="keywords"]', 'content', seo.keywords)
  }

  if (seo.description) {
    updateMetaTag('meta[name="description"]', 'content', seo.description)
    updateMetaTag('meta[property="og:description"]', 'content', seo.description)
  }
}

/**
 * Updates meta tag value
 */
function updateMetaTag(selector: string, attribute: string, value: string): void {
  const element = document.querySelector(selector)
  if (element) {
    element.setAttribute(attribute, value)
  }
}
