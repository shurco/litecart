/**
 * Utility for working with cart
 */

import { cartStore } from '$lib/stores/cart'
import type { Product, CartItem } from '$lib/types/models'

/**
 * Toggles product in cart (adds or removes)
 * @param product - Product to add/remove
 * @param cartItems - Current cart items to check availability
 */
export function toggleCartItem(product: Product, cartItems: CartItem[]): void {
  const inCart = cartItems.some((item) => item.id === product.id)

  if (inCart) {
    cartStore.remove(product.id)
  } else {
    const image = product.images?.[0] ? { name: product.images[0].name, ext: product.images[0].ext } : null

    const cartItem: CartItem = {
      id: product.id,
      name: product.name,
      slug: product.slug,
      amount: product.amount,
      image
    }

    cartStore.add(cartItem)
  }
}
