<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { createEventDispatcher } from 'svelte'
  import FormButton from '../form/Button.svelte'
  import FormInput from '../form/Input.svelte'
  import FormToggle from '../form/Toggle.svelte'
  import { loadPaymentSettings, savePaymentSettings, togglePaymentActive } from '$lib/composables/usePaymentSettings'
  import { systemStore } from '$lib/stores/system'
  import type { PaypalSettings } from '$lib/types/models'

  const dispatch = createEventDispatcher()

  let settings: PaypalSettings = {
    active: false,
    client_id: '',
    secret_key: ''
  }
  let formErrors: Record<string, string> = {}
  let unsubscribe: (() => void) | null = null

  onMount(async () => {
    settings = await loadPaymentSettings<PaypalSettings>('paypal', settings)

    // Subscribe to store updates to keep settings.active in sync
    unsubscribe = systemStore.subscribe((store) => {
      if (store.payments?.paypal !== undefined) {
        settings.active = store.payments.paypal
      }
    })
  })

  onDestroy(() => {
    if (unsubscribe) {
      unsubscribe()
    }
  })

  async function handleSubmit() {
    formErrors = {}

    if (!settings.client_id || settings.client_id.length < 80) {
      formErrors.client_id = 'Client ID must be at least 80 characters'
      return
    }
    if (!settings.secret_key || settings.secret_key.length < 80) {
      formErrors.secret_key = 'Secret key must be at least 80 characters'
      return
    }

    await savePaymentSettings('paypal', settings, 'paypal')
  }

  async function toggleActive() {
    const previousValue = settings.active
    const success = await togglePaymentActive('paypal', settings.active)

    // If request failed, revert the change
    if (!success) {
      settings.active = previousValue
    }
  }

  function close() {
    dispatch('close')
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>Paypal</h1>
      </div>
      <div class="pt-1">
        <FormToggle
          id="paypal-active"
          bind:value={settings.active}
          disabled={Object.keys(formErrors).length > 0}
          on:change={toggleActive}
        />
      </div>
    </div>
  </div>

  <form on:submit|preventDefault={handleSubmit}>
    <div class="flow-root">
      <dl class="mx-auto -my-3 mt-2 mb-0 space-y-4 text-sm">
        <FormInput
          id="client_id"
          type="text"
          title="Client ID"
          bind:value={settings.client_id}
          error={formErrors.client_id}
          ico="key"
        />
      </dl>

      <dl class="mx-auto -my-3 mt-5 mb-0 space-y-4 text-sm">
        <FormInput
          id="secret_key"
          type="text"
          title="Secret key"
          bind:value={settings.secret_key}
          error={formErrors.secret_key}
          ico="key"
        />
      </dl>
    </div>

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <FormButton type="submit" name="Save" color="green" />
        </div>
        <div class="grow"></div>
        <div class="flex-none">
          <FormButton type="button" name="Close" color="gray" on:click={close} />
        </div>
      </div>
    </div>
  </form>
</div>
