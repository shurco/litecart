<script lang="ts">
  import { onMount } from 'svelte';
  import Main from '$lib/layouts/Main.svelte';
  import FormButton from '$lib/components/form/Button.svelte';
  import FormInput from '$lib/components/form/Input.svelte';
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers';
  import { saveData } from '$lib/utils/apiHelpers';
  import { validators, validateFields } from '$lib/utils/validation';

  interface AuthSettings {
    email: string;
  }

  interface PasswordData {
    old: string;
    new: string;
  }

  let formData: AuthSettings = {
    email: ''
  };
  let passwordData: PasswordData = {
    old: '',
    new: ''
  };
  let formErrors: Record<string, string> = {};
  let passwordErrors: Record<string, string> = {};
  let loading = true;

  onMount(async () => {
    formData = await loadSettings<AuthSettings>('auth', formData);
    loading = false;
  });

  async function handleSubmit() {
    formErrors = validateFields(formData, [
      { field: 'email', ...validators.email('Valid email is required') }
    ]);

    if (Object.keys(formErrors).length > 0) {
      return;
    }

    await saveSettings('auth', formData);
  }

  async function handlePasswordSubmit() {
    passwordErrors = validateFields(passwordData, [
      { field: 'old', ...validators.required('Old password is required') },
      { field: 'new', ...validators.minLength(6, 'New password must be at least 6 characters') }
    ]);

    if (Object.keys(passwordErrors).length > 0) {
      return;
    }

    const result = await saveData<PasswordData>(
      '/api/_/settings/password',
      passwordData,
      true,
      'Password updated',
      'Failed to update password'
    );
    
    if (result !== null) {
      passwordData = { old: '', new: '' };
    }
  }
</script>

<svelte:component this={Main}>
  <h1 class="mb-5">Authentication Settings</h1>

  {#if loading}
    <div class="text-center py-8">Loading...</div>
  {:else}
    <form on:submit|preventDefault={handleSubmit} class="space-y-4 max-w-2xl">
      <h2 class="text-xl font-bold mb-4">Email</h2>
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

    <hr class="mt-5" />

    <form on:submit|preventDefault={handlePasswordSubmit} class="space-y-4 max-w-2xl">
      <h2 class="text-xl font-bold mb-4">Change Password</h2>
      <FormInput
        id="old_password"
        type="password"
        title="Old Password"
        bind:value={passwordData.old}
        error={passwordErrors.old}
        ico="finger-print"
      />
      <FormInput
        id="new_password"
        type="password"
        title="New Password"
        bind:value={passwordData.new}
        error={passwordErrors.new}
        ico="finger-print"
      />
      <div class="pt-4">
        <FormButton type="submit" name="Update Password" color="green" />
      </div>
    </form>
  {/if}
</svelte:component>
