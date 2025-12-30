<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { createEventDispatcher } from 'svelte';
  import FormButton from '../form/Button.svelte';
  import FormInput from '../form/Input.svelte';
  import FormToggle from '../form/Toggle.svelte';
  import { loadPaymentSettings, savePaymentSettings, togglePaymentActive } from '$lib/composables/usePaymentSettings';
  import { systemStore } from '$lib/stores/system';
  import type { StripeSettings } from '$lib/types/models';

  const dispatch = createEventDispatcher();

  let settings: StripeSettings = {
    active: false,
    secret_key: ''
  };
  let formErrors: Record<string, string> = {};
  let unsubscribe: (() => void) | null = null;

  onMount(async () => {
    settings = await loadPaymentSettings<StripeSettings>('stripe', settings);
    
    // Subscribe to store updates to keep settings.active in sync
    unsubscribe = systemStore.subscribe(store => {
      if (store.payments?.stripe !== undefined) {
        settings.active = store.payments.stripe;
      }
    });
  });

  onDestroy(() => {
    if (unsubscribe) {
      unsubscribe();
    }
  });

  async function handleSubmit() {
    formErrors = {};

    if (!settings.secret_key || settings.secret_key.length < 100) {
      formErrors.secret_key = 'Secret key must be at least 100 characters';
      return;
    }

    await savePaymentSettings('stripe', settings, 'stripe');
  }

  async function toggleActive() {
    const previousValue = settings.active;
    const success = await togglePaymentActive('stripe', settings.active);
    
    // If request failed, revert the change
    if (!success) {
      settings.active = previousValue;
    }
  }

  function close() {
    dispatch('close');
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>Stripe</h1>
      </div>
      <div class="pt-1">
        <FormToggle id="stripe-active" bind:value={settings.active} disabled={Object.keys(formErrors).length > 0} on:change={toggleActive} />
      </div>
    </div>
  </div>

  <form on:submit|preventDefault={handleSubmit}>
    <div class="flow-root">
      <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
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
