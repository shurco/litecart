<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import Drawer from '$lib/components/Drawer.svelte'
  import CartView from '$lib/components/cart/View.svelte'
  import SvgIcon from '$lib/components/SvgIcon.svelte'
  import { loadData, handleApiCall } from '$lib/utils/apiHelpers'
  import { apiPost } from '$lib/utils'
  import { costFormat, formatDate } from '$lib/utils'
  import { STRIPE_DASHBOARD_URL } from '$lib/utils/constants'
  import type { Cart } from '$lib/types/models'

  interface DrawerCart {
    cart: Cart
  }

  let carts: Cart[] = []
  let loading = true
  let drawerOpen = false
  let drawerCart: DrawerCart | null = null

  onMount(async () => {
    await loadCarts()
  })

  async function loadCarts() {
    loading = true
    const result = await loadData<Cart[]>('/api/_/carts', 'Failed to load carts')
    if (result) {
      carts = result
    }
    loading = false
  }

  function openView(cart: Cart) {
    drawerCart = { cart }
    drawerOpen = true
  }

  function closeDrawer() {
    if (drawerOpen) {
      drawerOpen = false
      setTimeout(() => {
        drawerCart = null
      }, 200)
    }
  }

  async function sendMail(cartId: string, event: Event) {
    event.stopPropagation()
    await handleApiCall(
      () => apiPost(`/api/_/carts/${cartId}/mail`, {}),
      'Mail sent successfully',
      'Failed to send mail'
    )
  }
</script>

<svelte:component this={Main}>
  <div class="mb-5 flex items-center justify-between">
    <h1>Carts</h1>
  </div>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else if carts.length === 0}
    <div class="py-8 text-center text-gray-500">No carts found</div>
  {:else}
    <table>
      <thead>
        <tr>
          <th>Email</th>
          <th>Price</th>
          <th>Status</th>
          <th>Payment</th>
          <th class="w-48">Created</th>
          <th class="w-48">Updated</th>
          <th class="w-12"></th>
        </tr>
      </thead>
      <tbody>
        {#each carts as cart, index}
          <tr
            class:bg-green-50={cart.payment_status === 'paid'}
            class="cursor-pointer hover:bg-gray-50"
            onclick={() => openView(cart)}
          >
            <td>{cart.email || '-'}</td>
            <td>
              {costFormat(cart.amount_total)} {cart.currency || ''}
            </td>
            <td
              class={cart.payment_status === 'paid'
                ? 'text-green-600'
                : cart.payment_status === 'pending'
                  ? 'text-yellow-600'
                  : cart.payment_status === 'failed'
                    ? 'text-red-600'
                    : 'text-gray-600'}
            >
              {cart.payment_status || '-'}
            </td>
            <td>{cart.payment_system || '-'}</td>
            <td>{formatDate(cart.created)}</td>
            <td>
              {#if cart.updated}
                {formatDate(cart.updated)}
              {/if}
            </td>
            <td onclick={(e) => e.stopPropagation()}>
              {#if cart.payment_status === 'paid'}
                <SvgIcon
                  name="envelope"
                  className="h-5 w-5 cursor-pointer"
                  onclick={(e: Event) => sendMail(cart.id, e)}
                  stroke="currentColor"
                />
              {:else}
                <SvgIcon name="envelope" className="h-5 w-5 opacity-30" stroke="currentColor" />
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
</svelte:component>

{#if drawerOpen}
  <Drawer isOpen={drawerOpen} on:close={closeDrawer} maxWidth="710px">
    {#if drawerCart}
      <CartView drawer={drawerCart} on:close={closeDrawer} />
    {/if}
  </Drawer>
{/if}
