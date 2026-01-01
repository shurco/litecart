<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import SvgIcon from '$lib/components/SvgIcon.svelte'
  import { loadData, handleApiCall } from '$lib/utils/apiHelpers'
  import { apiPost } from '$lib/utils'
  import { costFormat, formatDate } from '$lib/utils'
  import type { Cart } from '$lib/types/models'

  let carts: Cart[] = []
  let loading = true

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

  async function sendMail(cartId: string) {
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
        {#each carts as cart}
          <tr class:bg-green-50={cart.payment_status === 'paid'}>
            <td>{cart.email || '-'}</td>
            <td>
              {#if cart.payment_id}
                <a href="https://dashboard.stripe.com/payments/{cart.payment_id}" target="_blank">
                  {costFormat(cart.amount_total)}
                  {cart.currency || ''}
                </a>
              {:else}
                {costFormat(cart.amount_total)} {cart.currency || ''}
              {/if}
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
            <td>
              {#if cart.payment_status === 'paid'}
                <SvgIcon
                  name="envelope"
                  className="h-5 w-5 cursor-pointer"
                  on:click={() => sendMail(cart.id)}
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
