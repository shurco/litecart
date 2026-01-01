<script lang="ts">
  import { goto } from '$app/navigation'
  import { base } from '$app/paths'
  import { onMount } from 'svelte'
  import Blank from '$lib/layouts/Blank.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import { apiPost, apiGet } from '$lib/utils/api'
  import { showMessage } from '$lib/utils/message'

  let email = ''
  let password = ''
  let emailError = ''
  let passwordError = ''

  function validateEmail(value) {
    if (!value) {
      return 'Email is required'
    }
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) {
      return 'Email is not valid'
    }
    return ''
  }

  function validatePassword(value) {
    if (!value) {
      return 'Password is required'
    }
    if (value.length < 6) {
      return 'Password must be at least 6 characters'
    }
    return ''
  }

  async function handleSubmit(event) {
    event.preventDefault()

    emailError = validateEmail(email)
    passwordError = validatePassword(password)

    if (emailError || passwordError) {
      return
    }

    try {
      const res = await apiPost(`/api/sign/in`, { email, password })
      if (res?.success) {
        goto(`${base}/products`)
      } else {
        showMessage(res?.result || res?.message || 'Login failed', 'connextError')
      }
    } catch (error) {
      showMessage('Network error. Please try again.', 'connextError')
    }
  }

  onMount(async () => {
    // Check if already authenticated
    try {
      const checkRes = await apiGet('/api/_/version')
      if (checkRes?.success) {
        goto(`${base}/products`)
      }
    } catch (error) {
      // Not authenticated, stay on signin page
    }
  })
</script>

<svelte:component this={Blank}>
  <div class="mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8">
    <div class="mx-auto max-w-lg text-center">
      <h1 class="text-2xl font-bold sm:text-3xl">👨‍🎨 Admin sign in</h1>
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
      <FormButton type="submit" name="Login" color="green" ico="arrow-right" />
    </form>
  </div>
</svelte:component>
