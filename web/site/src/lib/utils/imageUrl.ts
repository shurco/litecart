/**
 * Utility for getting product image URLs
 */

export type ImageSize = 'sm' | 'md' | 'lg'

export interface ImageData {
  name: string
  ext: string
}

/**
 * Gets product image URL
 * @param image - Image data or null
 * @param size - Image size (sm, md, lg)
 * @returns Image URL or placeholder path
 */
export function getProductImageUrl(image: ImageData | null | undefined, size: ImageSize = 'md'): string {
  if (!image) {
    return '/assets/img/noimage.png'
  }
  return `/uploads/${image.name}_${size}.${image.ext}`
}

/**
 * Gets URL of the first image from array
 * @param images - Array of images
 * @param size - Image size
 * @returns Image URL or placeholder path
 */
export function getFirstImageUrl(images: ImageData[] | null | undefined, size: ImageSize = 'md'): string {
  if (!images || images.length === 0) {
    return '/assets/img/noimage.png'
  }
  return getProductImageUrl(images[0], size)
}
