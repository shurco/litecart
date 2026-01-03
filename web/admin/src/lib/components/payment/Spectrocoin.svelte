<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import FormButton from '../form/Button.svelte'
  import FormInput from '../form/Input.svelte'
  import FormTextarea from '../form/Textarea.svelte'
  import FormToggle from '../form/Toggle.svelte'
  import { loadPaymentSettings, savePaymentSettings, togglePaymentActive } from '$lib/composables/usePaymentSettings'
  import { systemStore } from '$lib/stores/system'
  import {
    MIN_MERCHANT_ID_LENGTH,
    MIN_PROJECT_ID_LENGTH,
    MIN_PRIVATE_KEY_LENGTH,
    ERROR_MESSAGES
  } from '$lib/constants/validation'
  import type { SpectrocoinSettings } from '$lib/types/models'

  interface Props {
    onclose?: () => void
  }

  let { onclose }: Props = $props()

  let settings = $state<SpectrocoinSettings>({
    active: false,
    merchant_id: '',
    project_id: '',
    private_key: ''
  })
  let formErrors = $state<Record<string, string>>({})
  let unsubscribe: (() => void) | null = null

  onMount(async () => {
    settings = await loadPaymentSettings<SpectrocoinSettings>('spectrocoin', settings)

    unsubscribe = systemStore.subscribe((store) => {
      if (store.payments?.spectrocoin !== undefined) {
        settings.active = store.payments.spectrocoin
      }
    })
  })

  onDestroy(() => {
    unsubscribe?.()
  })

  async function handleSubmit(event: SubmitEvent) {
    event.preventDefault()
    formErrors = {}

    if (!settings.merchant_id || settings.merchant_id.length < MIN_MERCHANT_ID_LENGTH) {
      formErrors.merchant_id = ERROR_MESSAGES.MERCHANT_ID_TOO_SHORT
      return
    }
    if (!settings.project_id || settings.project_id.length < MIN_PROJECT_ID_LENGTH) {
      formErrors.project_id = ERROR_MESSAGES.PROJECT_ID_TOO_SHORT
      return
    }
    if (!settings.private_key || settings.private_key.length < MIN_PRIVATE_KEY_LENGTH) {
      formErrors.private_key = ERROR_MESSAGES.PRIVATE_KEY_TOO_SHORT
      return
    }

    await savePaymentSettings('spectrocoin', settings, 'spectrocoin')
  }

  async function handleToggleActive() {
    const previousValue = settings.active
    const success = await togglePaymentActive('spectrocoin', settings.active)

    if (!success) {
      settings.active = previousValue
    }
  }

  function close() {
    onclose?.()
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>Spectrocoin</h1>
      </div>
      <div class="pt-1">
        <FormToggle
          id="spectrocoin-active"
          bind:value={settings.active}
          disabled={Object.keys(formErrors).length > 0}
          onchange={handleToggleActive}
        />
      </div>
    </div>
  </div>

  <form onsubmit={handleSubmit}>
    <div class="flow-root">
      <dl class="mx-auto -my-3 mt-2 mb-0 space-y-4 text-sm">
        <FormInput
          id="merchant_id"
          type="text"
          title="Merchant ID"
          bind:value={settings.merchant_id}
          error={formErrors.merchant_id}
          ico="key"
        />
        <div class="mt-5">
          <FormInput
            id="project_id"
            type="text"
            title="Project ID"
            bind:value={settings.project_id}
            error={formErrors.project_id}
            ico="key"
          />
        </div>
        <div class="mt-5">
          <FormTextarea id="private_key" title="Private key" bind:value={settings.private_key} rows={15} />
          {#if formErrors.private_key}
            <span class="pl-4 text-sm text-red-500">{formErrors.private_key}</span>
          {/if}
        </div>
      </dl>
    </div>

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <FormButton type="submit" name="Save" color="green" />
        </div>
        <div class="grow"></div>
        <div class="flex-none">
          <FormButton type="button" name="Close" color="gray" onclick={close} />
        </div>
      </div>
    </div>
  </form>
</div>
