<script lang="ts">
  import { onMount } from 'svelte';
  import Main from '$lib/layouts/Main.svelte';
  import FormButton from '$lib/components/form/Button.svelte';
  import FormInput from '$lib/components/form/Input.svelte';
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers';

  interface WebhookSettings {
    url: string;
  }

  let formData: WebhookSettings = {
    url: ''
  };
  let formErrors: Record<string, string> = {};
  let loading = true;

  onMount(async () => {
    formData = await loadSettings<WebhookSettings>('webhook', formData);
    loading = false;
  });

  async function handleSubmit() {
    formErrors = {};

    if (formData.url && !/^https?:\/\/.+/.test(formData.url)) {
      formErrors.url = 'Valid URL is required (e.g., https://example.com/webhook)';
      return;
    }

    await saveSettings('webhook', formData);
  }
</script>

<svelte:component this={Main}>
  <h1 class="mb-5">Webhook Settings</h1>

  {#if loading}
    <div class="text-center py-8">Loading...</div>
  {:else}
    <form on:submit|preventDefault={handleSubmit} class="space-y-4 max-w-2xl">
      <FormInput
        id="url"
        type="url"
        title="Webhook URL"
        bind:value={formData.url}
        error={formErrors.url}
        ico="webhook"
        placeholder="https://example.com/webhook"
      />
      <div class="pt-4">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </form>
  {/if}
</svelte:component>
