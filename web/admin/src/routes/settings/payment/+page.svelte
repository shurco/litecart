<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import Main from '$lib/layouts/Main.svelte';
  import Drawer from '$lib/components/Drawer.svelte';
  import Stripe from '$lib/components/payment/Stripe.svelte';
  import Paypal from '$lib/components/payment/Paypal.svelte';
  import Spectrocoin from '$lib/components/payment/Spectrocoin.svelte';
  import FormButton from '$lib/components/form/Button.svelte';
  import FormSelect from '$lib/components/form/Select.svelte';
  import { systemStore } from '$lib/stores/system';
  import { loadSettings as loadSettingsHelper, saveSettings } from '$lib/utils/settingsHelpers';
  import { loadData } from '$lib/utils/apiHelpers';
  import type { PaymentSettings } from '$lib/types/models';

  let drawerOpen = false;
  let drawerMode: 'stripe' | 'paypal' | 'spectrocoin' | null = null;
  let payments: Record<string, boolean> = {};
  let payment: PaymentSettings = {
    currency: ''
  };
  let formErrors: Record<string, string> = {};

  const currencyOptions = ['EUR', 'USD', 'JPY', 'GBP', 'AUD', 'CAD', 'CHF', 'CNY', 'SEK'];

  let unsubscribe: (() => void) | null = null;

  onMount(async () => {
    await loadPaymentSettings();
    
    // Subscribe to store updates only on client side
    unsubscribe = systemStore.subscribe(store => {
      payments = store.payments || {};
    });
  });

  onDestroy(() => {
    if (unsubscribe) {
      unsubscribe();
    }
  });

  async function loadPaymentSettings() {
    const paymentProviders = await loadData<Record<string, boolean>>('/api/cart/payment', 'Failed to load payment settings');
    if (paymentProviders) {
      payments = paymentProviders;
      systemStore.update(store => ({
        ...store,
        payments: payments
      }));
    }
    
    const paymentSettings = await loadSettingsHelper<PaymentSettings>('payment', payment);
    payment.currency = paymentSettings.currency;
  }

  async function handleCurrencySubmit() {
    formErrors = {};

    if (!payment.currency) {
      formErrors.currency = 'Currency is required';
      return;
    }

    if (!currencyOptions.includes(payment.currency)) {
      formErrors.currency = 'Currency must be one of: ' + currencyOptions.join(', ');
      return;
    }

    await saveSettings('payment', payment, 'Currency saved');
  }

  function openDrawer(mode: 'stripe' | 'paypal' | 'spectrocoin') {
    drawerMode = mode;
    drawerOpen = true;
  }

  function closeDrawer() {
    drawerOpen = false;
    setTimeout(() => {
      drawerMode = null;
    }, 200);
  }
</script>

<svelte:component this={Main}>
  <div class="pb-10">
    <header class="mb-4">
      <h1>Payment</h1>
    </header>

    <form on:submit|preventDefault={handleCurrencySubmit} class="max-w-2xl">
      <FormSelect
        id="currency"
        title="Currency"
        options={currencyOptions}
        bind:value={payment.currency}
        error={formErrors.currency}
        ico="money"
      />
      <div class="pt-5">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </form>
    <hr class="mt-5" />

    <div class="mt-5">
      <h2 class="mb-5">Payment providers</h2>
      <div class="flex">
      <div 
        class="cursor-pointer rounded p-2 {payments.stripe ? 'bg-green-200' : 'bg-gray-200'}" 
        on:click={() => openDrawer('stripe')}
        role="button"
        tabindex="0"
        on:keydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); openDrawer('stripe'); } }}
      >
        Stripe
      </div>
      <div 
        class="cursor-pointer rounded p-2 ml-5 {payments.paypal ? 'bg-green-200' : 'bg-gray-200'}" 
        on:click={() => openDrawer('paypal')}
        role="button"
        tabindex="0"
        on:keydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); openDrawer('paypal'); } }}
      >
        Paypal
      </div>
      <div 
        class="cursor-pointer rounded p-2 ml-5 {payments.spectrocoin ? 'bg-green-200' : 'bg-gray-200'}" 
        on:click={() => openDrawer('spectrocoin')}
        role="button"
        tabindex="0"
        on:keydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); openDrawer('spectrocoin'); } }}
      >
        Spectrocoin
      </div>
    </div>
    </div>
  </div>
</svelte:component>

{#if drawerOpen}
  <Drawer isOpen={drawerOpen} on:close={closeDrawer} maxWidth="725px">
    {#if drawerMode === 'stripe'}
      <Stripe on:close={closeDrawer} />
    {:else if drawerMode === 'paypal'}
      <Paypal on:close={closeDrawer} />
    {:else if drawerMode === 'spectrocoin'}
      <Spectrocoin on:close={closeDrawer} />
    {/if}
  </Drawer>
{/if}

