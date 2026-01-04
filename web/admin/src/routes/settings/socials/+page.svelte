<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers'
  import { translate } from '$lib/i18n'

  // Reactive translation function
  let t = $derived($translate)

  interface SocialSettings {
    facebook: string
    instagram: string
    twitter: string
    dribbble: string
    github: string
    youtube: string
    other: string
  }

  let formData = $state<SocialSettings>({
    facebook: '',
    instagram: '',
    twitter: '',
    dribbble: '',
    github: '',
    youtube: '',
    other: ''
  })
  let formErrors = $state<Record<string, string>>({})
  let loading = $state(true)

  onMount(async () => {
    const loaded = await loadSettings<SocialSettings>('social', formData)
    if (loaded) {
      formData = {
        facebook: loaded.facebook || '',
        instagram: loaded.instagram || '',
        twitter: loaded.twitter || '',
        dribbble: loaded.dribbble || '',
        github: loaded.github || '',
        youtube: loaded.youtube || '',
        other: loaded.other || ''
      }
    }
    loading = false
  })

  async function handleSubmit() {
    formErrors = {}

    if (formData.other && !/^https?:\/\/.+/.test(formData.other)) {
      formErrors.other = t('settings.validUrlRequired')
      return
    }

    await saveSettings('social', formData)
  }
</script>

<Main>
  <h1 class="mb-5">{t('settings.socialSettings')}</h1>

  {#if loading}
    <div class="py-8 text-center">{t('common.loading')}</div>
  {:else}
    <form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="max-w-2xl space-y-4">
      <FormInput
        id="facebook"
        title={t('settings.facebook')}
        bind:value={formData.facebook}
        error={formErrors.facebook}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="instagram"
        title={t('settings.instagram')}
        bind:value={formData.instagram}
        error={formErrors.instagram}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="twitter"
        title={t('settings.twitter')}
        bind:value={formData.twitter}
        error={formErrors.twitter}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="dribbble"
        title={t('settings.dribbble')}
        bind:value={formData.dribbble}
        error={formErrors.dribbble}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="github"
        title={t('settings.github')}
        bind:value={formData.github}
        error={formErrors.github}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="youtube"
        title={t('settings.youtube')}
        bind:value={formData.youtube}
        error={formErrors.youtube}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="other"
        type="url"
        title={t('settings.otherUrl')}
        bind:value={formData.other}
        error={formErrors.other}
        ico="link"
        placeholder="https://example.com"
      />
      <div class="pt-4">
        <FormButton type="submit" name={t('common.save')} color="green" />
      </div>
    </form>
  {/if}
</Main>
