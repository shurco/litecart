<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers'

  interface SocialSettings {
    facebook: string
    instagram: string
    twitter: string
    dribbble: string
    github: string
    youtube: string
    other: string
  }

  let formData: SocialSettings = {
    facebook: '',
    instagram: '',
    twitter: '',
    dribbble: '',
    github: '',
    youtube: '',
    other: ''
  }
  let formErrors: Record<string, string> = {}
  let loading = true

  onMount(async () => {
    formData = await loadSettings<SocialSettings>('social', formData)
    loading = false
  })

  async function handleSubmit() {
    formErrors = {}

    if (formData.other && !/^https?:\/\/.+/.test(formData.other)) {
      formErrors.other = 'Valid URL is required'
      return
    }

    await saveSettings('social', formData)
  }
</script>

<svelte:component this={Main}>
  <h1 class="mb-5">Social Settings</h1>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else}
    <form on:submit|preventDefault={handleSubmit} class="max-w-2xl space-y-4">
      <FormInput
        id="facebook"
        title="Facebook"
        bind:value={formData.facebook}
        error={formErrors.facebook}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="instagram"
        title="Instagram"
        bind:value={formData.instagram}
        error={formErrors.instagram}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="twitter"
        title="Twitter"
        bind:value={formData.twitter}
        error={formErrors.twitter}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="dribbble"
        title="Dribbble"
        bind:value={formData.dribbble}
        error={formErrors.dribbble}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="github"
        title="GitHub"
        bind:value={formData.github}
        error={formErrors.github}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="youtube"
        title="YouTube"
        bind:value={formData.youtube}
        error={formErrors.youtube}
        ico="user-group"
        placeholder="username"
      />
      <FormInput
        id="other"
        type="url"
        title="Other (URL)"
        bind:value={formData.other}
        error={formErrors.other}
        ico="link"
        placeholder="https://example.com"
      />
      <div class="pt-4">
        <FormButton type="submit" name="Save" color="green" />
      </div>
    </form>
  {/if}
</svelte:component>
