<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers'
  import { saveData } from '$lib/utils/apiHelpers'
  import { validators, validateFields } from '$lib/utils/validation'
  import { translate } from '$lib/i18n'

  // Reactive translation function
  let t = $derived($translate)

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
    formErrors = validateFields(formData, [{ field: 'email', ...validators.email(t('settings.validEmailRequired')) }])

    if (Object.keys(formErrors).length > 0) {
      return
    }

    await saveSettings('auth', formData)
  }

  async function handlePasswordSubmit() {
    passwordErrors = validateFields(passwordData, [
      { field: 'old', ...validators.required(t('settings.oldPasswordRequired')) },
      { field: 'new', ...validators.minLength(6, t('settings.newPasswordMinLength')) }
    ])

    if (Object.keys(passwordErrors).length > 0) {
      return
    }

    const result = await saveData<PasswordData>(
      '/api/_/settings/password',
      passwordData,
      true,
      t('settings.passwordUpdated'),
      t('settings.failedToUpdatePassword')
    )

    if (result !== null) {
      passwordData = { old: '', new: '' }
    }
  }
</script>

<Main>
  <h1 class="mb-5">{t('settings.authSettings')}</h1>

  {#if loading}
    <div class="py-8 text-center">{t('common.loading')}</div>
  {:else}
    <form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="max-w-2xl space-y-4">
      <h2 class="mb-4 text-xl font-bold">{t('settings.email')}</h2>
      <FormInput
        id="email"
        type="email"
        title={t('settings.email')}
        bind:value={formData.email}
        error={formErrors.email}
        ico="at-symbol"
      />
      <div class="pt-4">
        <FormButton type="submit" name={t('common.save')} color="green" />
      </div>
    </form>

    <hr class="mt-5" />

    <form onsubmit={(e) => { e.preventDefault(); handlePasswordSubmit(); }} class="max-w-2xl space-y-4">
      <h2 class="mb-4 text-xl font-bold">{t('settings.changePassword')}</h2>
      <FormInput
        id="old_password"
        type="password"
        title={t('settings.oldPassword')}
        bind:value={passwordData.old}
        error={passwordErrors.old}
        ico="finger-print"
      />
      <FormInput
        id="new_password"
        type="password"
        title={t('settings.newPassword')}
        bind:value={passwordData.new}
        error={passwordErrors.new}
        ico="finger-print"
      />
      <div class="pt-4">
        <FormButton type="submit" name={t('settings.updatePassword')} color="green" />
      </div>
    </form>
  {/if}
</Main>
