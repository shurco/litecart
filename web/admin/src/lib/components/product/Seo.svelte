<script lang="ts">
  import { onMount } from 'svelte'
  import FormButton from '../form/Button.svelte'
  import FormInput from '../form/Input.svelte'
  import FormTextarea from '../form/Textarea.svelte'
  import { loadData, saveData } from '$lib/utils/apiHelpers'
  import type { Product } from '$lib/types/models'
  import { translate } from '$lib/i18n'

  // Reactive translation function
  let t = $derived($translate)

  interface DrawerProduct {
    product: Product
    index: number
    currency?: string
  }

  interface Props {
    drawer: DrawerProduct
    onclose?: () => void
  }

  let { drawer, onclose }: Props = $props()

  let seoData = $state({
    title: '',
    keywords: '',
    description: ''
  })

  onMount(async () => {
    await loadProduct()
  })

  async function loadProduct() {
    const product = await loadData<Product>(`/api/_/products/${drawer.product.id}`, t('products.failedToLoadProduct'))
    if (product) {
      seoData = {
        title: product.seo?.title || '',
        keywords: product.seo?.keywords || '',
        description: product.seo?.description || ''
      }
    }
  }

  async function handleSubmit(event: SubmitEvent) {
    event.preventDefault()
    await saveData<Product>(
      `/api/_/products/${drawer.product.id}`,
      { seo: seoData },
      true,
      t('products.failedToSave'),
      t('products.failedToSave')
    )
  }

  function close() {
    onclose?.()
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

  <form onsubmit={handleSubmit}>
    <div class="flow-root">
      <dl class="mx-auto -my-3 mt-2 mb-0 space-y-4 text-sm">
        <FormInput id="seo-title" title={t('pages.seoTitle')} bind:value={seoData.title} ico="glob-alt" />
        <FormInput id="seo-keywords" title={t('pages.seoKeywords')} bind:value={seoData.keywords} ico="glob-alt" />
        <hr />
        <FormTextarea id="seo-description" title={t('pages.seoDescription')} bind:value={seoData.description} />
      </dl>
    </div>

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <FormButton type="submit" name={t('common.save')} color="green" />
          <FormButton type="button" name={t('common.close')} color="gray" onclick={close} />
        </div>
        <div class="grow"></div>
      </div>
    </div>
  </form>
</div>
