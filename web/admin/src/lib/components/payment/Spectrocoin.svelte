<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { createEventDispatcher } from 'svelte';
  import FormButton from '../form/Button.svelte';
  import FormInput from '../form/Input.svelte';
  import FormTextarea from '../form/Textarea.svelte';
  import FormToggle from '../form/Toggle.svelte';
  import { loadPaymentSettings, savePaymentSettings, togglePaymentActive } from '$lib/composables/usePaymentSettings';
  import { systemStore } from '$lib/stores/system';
  import type { SpectrocoinSettings } from '$lib/types/models';

  const dispatch = createEventDispatcher();

  let settings: SpectrocoinSettings = {
    active: false,
    merchant_id: '',
    project_id: '',
    private_key: ''
  };
  let formErrors: Record<string, string> = {};
  let unsubscribe: (() => void) | null = null;

  onMount(async () => {
    settings = await loadPaymentSettings<SpectrocoinSettings>('spectrocoin', settings);
    
    // Subscribe to store updates to keep settings.active in sync
    unsubscribe = systemStore.subscribe(store => {
      if (store.payments?.spectrocoin !== undefined) {
        settings.active = store.payments.spectrocoin;
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

    if (!settings.merchant_id || settings.merchant_id.length < 36) {
      formErrors.merchant_id = 'Merchant ID must be at least 36 characters';
      return;
    }
    if (!settings.project_id || settings.project_id.length < 36) {
      formErrors.project_id = 'Project ID must be at least 36 characters';
      return;
    }
    if (!settings.private_key || settings.private_key.length < 1500) {
      formErrors.private_key = 'Private key must be at least 1500 characters';
      return;
    }

    await savePaymentSettings('spectrocoin', settings, 'spectrocoin');
  }

  async function toggleActive() {
    const previousValue = settings.active;
    const success = await togglePaymentActive('spectrocoin', settings.active);
    
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
        <h1>Spectrocoin</h1>
      </div>
      <div class="pt-1">
        <FormToggle id="spectrocoin-active" bind:value={settings.active} disabled={Object.keys(formErrors).length > 0} on:change={toggleActive} />
      </div>
    </div>
  </div>

  <form on:submit|preventDefault={handleSubmit}>
    <div class="flow-root">
      <dl class="-my-3 mx-auto mb-0 mt-2 space-y-4 text-sm">
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
          <FormTextarea
            id="private_key"
            title="Private key"
            bind:value={settings.private_key}
            rows={15}
          />
          {#if formErrors.private_key}
            <span class="text-sm text-red-500 pl-4">{formErrors.private_key}</span>
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
          <FormButton type="button" name="Close" color="gray" on:click={close} />
        </div>
      </div>
    </div>
  </form>
</div>
