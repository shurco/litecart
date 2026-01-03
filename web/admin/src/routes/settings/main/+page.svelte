<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import { loadSettings, saveSettings } from '$lib/utils/settingsHelpers'
  import { validators, validateFields } from '$lib/utils/validation'

  interface MainSettings {
    site_name: string
    domain: string
    email: string
  }

  let formData = $state<MainSettings>({
    site_name: '',
    domain: '',
    email: ''
  })
  let formErrors = $state<Record<string, string>>({})
  let loading = $state(true)

  onMount(async () => {
    const loaded = await loadSettings<MainSettings>('main', formData)
    if (loaded) {
      formData = loaded
    }
    loading = false
  })

  async function handleSubmit(event: SubmitEvent) {
    event.preventDefault()
    formErrors = validateFields(formData, [
      { field: 'site_name', ...validators.minLength(6, 'Site name must be at least 6 characters') },
      { field: 'domain', ...validators.required('Domain is required') },
      { field: 'email', ...validators.email('Valid email is required') }
    ])

    if (Object.keys(formErrors).length > 0) {
      return
    }

    await saveSettings('main', formData)
  }
</script>

<Main>
  <h1 class="mb-5">Main Settings</h1>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else}
    <form onsubmit={handleSubmit} class="max-w-2xl space-y-4">
      <FormInput
        id="site_name"
        title="Site Name"
        bind:value={formData.site_name}
        error={formErrors.site_name}
        ico="home"
      />
      <FormInput id="domain" title="Domain" bind:value={formData.domain} error={formErrors.domain} ico="link" />
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
</Main>
