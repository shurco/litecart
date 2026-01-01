<script lang="ts">
  import { goto } from '$app/navigation'
  import { base } from '$app/paths'
  import { onMount } from 'svelte'
  import { browser } from '$app/environment'
  import Blank from '$lib/layouts/Blank.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import { apiPost } from '$lib/utils/api'
  import { showMessage } from '$lib/utils/message'

  let email = ''
  let password = ''
  let domain = ''
  let emailError = ''
  let passwordError = ''
  let domainError = ''

  function validateEmail(value: string) {
    if (!value) {
      return 'Email is required'
    }
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) {
      return 'Email is not valid'
    }
    return ''
  }

  function validatePassword(value: string) {
    if (!value) {
      return 'Password is required'
    }
    if (value.length < 6) {
      return 'Password must be at least 6 characters'
    }
    if (value.length > 72) {
      return 'Password must be at most 72 characters'
    }
    return ''
  }

  function validateDomain(value: string) {
    if (!value) {
      return 'Domain is required'
    }
    // Basic domain validation
    const domainRegex = /^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$/i
    if (!domainRegex.test(value) && value !== 'localhost' && !value.startsWith('localhost:')) {
      return 'Domain is not valid'
    }
    return ''
  }

  async function handleSubmit(event: Event) {
    event.preventDefault()

    emailError = validateEmail(email)
    passwordError = validatePassword(password)
    domainError = validateDomain(domain)

    if (emailError || passwordError || domainError) {
      return
    }

    try {
      const res = await apiPost(`/api/install`, { email, password, domain })
      if (res?.success) {
        showMessage('Cart installed successfully!', 'connextSuccess')
        // Redirect to signin page after successful installation
        setTimeout(() => {
          goto(`${base}/signin`)
        }, 1000)
      } else {
        showMessage(res?.result || res?.message || 'Installation failed', 'connextError')
      }
    } catch (error) {
      showMessage('Network error. Please try again.', 'connextError')
    }
  }

  onMount(() => {
    // Set default domain from current location if in browser
    if (browser) {
      const url = new URL(window.location.href)
      domain = url.origin.replace(/^https?:\/\//, '')
    }
  })
</script>

<svelte:component this={Blank}>
  <div class="mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8">
    <div class="mx-auto max-w-lg text-center">
      <h1 class="text-2xl font-bold sm:text-3xl">🛒 Install Litecart</h1>
      <p class="mt-4 text-gray-600">Configure your shopping cart</p>
    </div>
    <form on:submit|preventDefault={handleSubmit} class="mx-auto mt-8 mb-0 max-w-md space-y-4">
      <FormInput id="email" type="email" title="Email" ico="at-symbol" error={emailError} bind:value={email} />
      <FormInput
        id="password"
        type="password"
        title="Password"
        ico="finger-print"
        error={passwordError}
        bind:value={password}
      />
      <FormInput
        id="domain"
        type="text"
        title="Domain"
        ico="glob-alt"
        error={domainError}
        bind:value={domain}
        placeholder="example.com"
      />
      <FormButton type="submit" name="Install" color="green" ico="arrow-right" />
    </form>
  </div>
</svelte:component>
