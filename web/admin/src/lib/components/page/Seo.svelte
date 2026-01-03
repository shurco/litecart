<script lang="ts">
  import { onMount } from 'svelte'
  import { createEventDispatcher } from 'svelte'
  import FormButton from '../form/Button.svelte'
  import FormInput from '../form/Input.svelte'
  import FormTextarea from '../form/Textarea.svelte'
  import { loadData, saveData } from '$lib/utils/apiHelpers'
  import type { Page } from '$lib/types/models'

  interface Props {
    page: Page
    onclose?: () => void
  }

  let { page, onclose }: Props = $props()

  let seoData = $state({
    title: '',
    keywords: '',
    description: ''
  })

  // Reactively load data when page.id changes
  $effect(() => {
    if (page?.id) {
      loadPage()
    }
  })

  async function loadPage() {
    if (!page?.id) return

    const pageData = await loadData<Page>(`/api/_/pages/${page.id}`, 'Failed to load page')
    if (pageData) {
      seoData = {
        title: pageData.seo?.title || '',
        keywords: pageData.seo?.keywords || '',
        description: pageData.seo?.description || ''
      }
    }
  }

  async function handleSubmit() {
    await saveData<Page>(
      `/api/_/pages/${page.id}`,
      { seo: seoData },
      true,
      'SEO settings saved',
      'Failed to save SEO settings'
    )
  }

  function close() {
    dispatch('close')
  }
</script>

<div>
  <div class="pb-8">
    <div class="flex items-center">
      <div class="pr-3">
        <h1>SEO</h1>
      </div>
    </div>
  </div>

  <form on:submit|preventDefault={handleSubmit}>
    <div class="flow-root">
      <dl class="mx-auto -my-3 mt-2 mb-0 space-y-4 text-sm">
        <FormInput id="seo-title" title="Title" bind:value={seoData.title} ico="glob-alt" />
        <FormInput id="seo-keywords" title="Keywords" bind:value={seoData.keywords} ico="glob-alt" />
        <hr />
        <FormTextarea id="seo-description" title="Description" bind:value={seoData.description} />
      </dl>
    </div>

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <FormButton type="submit" name="Save" color="green" />
          <FormButton type="button" name="Close" color="gray" on:click={close} />
        </div>
        <div class="grow"></div>
      </div>
    </div>
  </form>
</div>
