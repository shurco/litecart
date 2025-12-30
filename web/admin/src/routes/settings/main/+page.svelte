<script lang="ts">
  import { onMount } from 'svelte';
  import Main from '$lib/layouts/Main.svelte';
  import FormButton from '$lib/components/form/Button.svelte';
  import FormInput from '$lib/components/form/Input.svelte';
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers';
  import { validators, validateFields } from '$lib/utils/validation';

  interface MainSettings {
    site_name: string;
    domain: string;
    email: string;
  }

  let formData: MainSettings = {
    site_name: '',
    domain: '',
    email: ''
  };
  let formErrors: Record<string, string> = {};
  let loading = true;

  onMount(async () => {
    formData = await loadSettings<MainSettings>('main', formData);
    loading = false;
  });

  async function handleSubmit() {
    formErrors = validateFields(formData, [
      { field: 'site_name', ...validators.minLength(6, 'Site name must be at least 6 characters') },
      { field: 'domain', ...validators.required('Domain is required') },
      { field: 'email', ...validators.email('Valid email is required') }
    ]);

    if (Object.keys(formErrors).length > 0) {
      return;
    }

    await saveSettings('main', formData);
  }
</script>

<svelte:component this={Main}>
  <h1 class="mb-5">Main Settings</h1>

  {#if loading}
    <div class="text-center py-8">Loading...</div>
  {:else}
    <form on:submit|preventDefault={handleSubmit} class="space-y-4 max-w-2xl">
      <FormInput
        id="site_name"
        title="Site Name"
        bind:value={formData.site_name}
        error={formErrors.site_name}
        ico="home"
      />
      <FormInput
        id="domain"
        title="Domain"
        bind:value={formData.domain}
        error={formErrors.domain}
        ico="link"
      />
      <FormInput
        id="email"
        type="email"
        title="Email"
        bind:value={formData.email}
        error={formErrors.email}
        ico="at-symbol"
      />
      <div class="pt-4">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </form>
  {/if}
</svelte:component>
