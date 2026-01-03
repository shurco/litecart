<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers'
  import { saveData } from '$lib/utils/apiHelpers'
  import { validators, validateFields } from '$lib/utils/validation'

  interface AuthSettings {
    email: string
  }

  interface PasswordData {
    old: string
    new: string
  }

  let formData = $state<AuthSettings>({
    email: ''
  })
  let passwordData = $state<PasswordData>({
    old: '',
    new: ''
  })
  let formErrors = $state<Record<string, string>>({})
  let passwordErrors = $state<Record<string, string>>({})
  let loading = $state(true)

  onMount(async () => {
    formData = await loadSettings<AuthSettings>('auth', formData)
    loading = false
  })

  async function handleSubmit() {
    formErrors = validateFields(formData, [{ field: 'email', ...validators.email('Valid email is required') }])

    if (Object.keys(formErrors).length > 0) {
      return
    }

    await saveSettings('auth', formData)
  }

  async function handlePasswordSubmit() {
    passwordErrors = validateFields(passwordData, [
      { field: 'old', ...validators.required('Old password is required') },
      { field: 'new', ...validators.minLength(6, 'New password must be at least 6 characters') }
    ])

    if (Object.keys(passwordErrors).length > 0) {
      return
    }

    const result = await saveData<PasswordData>(
      '/api/_/settings/password',
      passwordData,
      true,
      'Password updated',
      'Failed to update password'
    )

    if (result !== null) {
      passwordData = { old: '', new: '' }
    }
  }
</script>

<Main>
  <h1 class="mb-5">Authentication Settings</h1>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else}
    <form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="max-w-2xl space-y-4">
      <h2 class="mb-4 text-xl font-bold">Email</h2>
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

    <form onsubmit={(e) => { e.preventDefault(); handlePasswordSubmit(); }} class="max-w-2xl space-y-4">
      <h2 class="mb-4 text-xl font-bold">Change Password</h2>
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
</Main>
