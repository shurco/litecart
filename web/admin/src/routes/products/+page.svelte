<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import Drawer from '$lib/components/Drawer.svelte'
  import ProductView from '$lib/components/product/View.svelte'
  import ProductSeo from '$lib/components/product/Seo.svelte'
  import ProductDigital from '$lib/components/product/Digital.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import FormSelect from '$lib/components/form/Select.svelte'
  import FormTextarea from '$lib/components/form/Textarea.svelte'
  import Editor from '$lib/components/Editor.svelte'
  import Upload from '$lib/components/form/Upload.svelte'
  import SvgIcon from '$lib/components/SvgIcon.svelte'
  import { loadData, saveData, deleteData, toggleActive as toggleActiveApi } from '$lib/utils/apiHelpers'
  import { costFormat, formatDate, sortByDate, confirmDelete, showMessage } from '$lib/utils'
  import { apiDelete, apiUpdate } from '$lib/utils/api'
  import { validators, validateFields } from '$lib/utils/validation'
  import type { Product } from '$lib/types/models'

  interface ProductsResponse {
    products: Product[]
    currency: string
  }

  interface DrawerProduct {
    product: Product
    index: number
    currency?: string
  }

  let products: Product[] = []
  let currency = ''
  let loading = true
  let drawerOpen = false
  let drawerMode: 'view' | 'add' | 'edit' | 'seo' | 'digital' = 'view'
  let drawerProduct: DrawerProduct | null = null
  let drawerIndex = -1

  interface ProductFormData {
    name: string
    slug: string
    brief: string
    description: string
    amount: string | number
    active: boolean
    metadata: Array<{ key: string; value: string }>
    attributes: string[]
    digital: {
      type: '' | 'file' | 'data' | 'api'
    }
  }

  let formData: ProductFormData = {
    name: '',
    slug: '',
    brief: '',
    description: '',
    amount: '0',
    active: true,
    metadata: [],
    attributes: [],
    digital: {
      type: ''
    }
  }
  let amountDisplay = '0'
  let formErrors: Record<string, string> = {}
  let productImages: Product['images'] = []
  let fullProductData: Product | null = null

  onMount(async () => {
    await loadProducts()
  })

  async function loadProducts() {
    loading = true
    const result = await loadData<ProductsResponse>('/api/_/products', 'Failed to load products')
    if (result) {
      products = sortByDate(result.products || [])
      currency = result.currency || ''
    }
    loading = false
  }

  function openView(product: Product, index: number) {
    drawerProduct = { product, index, currency }
    drawerMode = 'view'
    drawerOpen = true
  }

  async function openAdd() {
    formData = {
      name: '',
      slug: '',
      brief: '',
      description: '',
      amount: '0',
      active: true,
      metadata: [],
      attributes: [],
      digital: {
        type: ''
      }
    }
    amountDisplay = '0'
    productImages = []
    fullProductData = null
    formErrors = {}
    drawerMode = 'add'
    drawerOpen = true
  }

  async function openEdit(product: Product, index: number) {
    drawerProduct = { product, index }
    formErrors = {}
    drawerMode = 'edit'

    const result = await loadData<Product>(`/api/_/products/${product.id}`, 'Failed to load product')
    if (result) {
      fullProductData = result
      formData = {
        name: result.name || '',
        slug: result.slug || '',
        brief: result.brief || '',
        description: result.description || '',
        amount: result.amount || 0,
        active: result.active !== undefined ? result.active : true,
        metadata: result.metadata || [],
        attributes: result.attributes || [],
        digital: result.digital || { type: '' }
      }
      amountDisplay = costFormat(result.amount)
      productImages = result.images || []
      drawerOpen = true
    }
  }

  function openSeo(product: Product, index: number) {
    drawerProduct = { product, index, currency }
    drawerMode = 'seo'
    drawerOpen = true
  }

  function openDigital(product: Product, index: number) {
    drawerProduct = { product, index, currency }
    drawerMode = 'digital'
    drawerOpen = true
  }

  async function handleDigitalContentUpdate() {
    await loadProducts()
  }

  function closeDrawer() {
    if (drawerOpen) {
      drawerOpen = false
      setTimeout(() => {
        drawerProduct = null
        drawerMode = 'view'
      }, 200)
    }
  }

  function updateProductInList(product: Product) {
    const productIndex = products.findIndex((p) => p.id === product.id)
    if (productIndex !== -1) {
      products = products.map((p, i) => (i === productIndex ? product : p))
    }
    if (drawerProduct && drawerProduct.product.id === product.id) {
      drawerProduct.product = product
    }
    if (fullProductData && fullProductData.id === product.id) {
      fullProductData = product
      formData.active = product.active
    }
  }

  async function handleSubmit() {
    formErrors = validateFields(formData, [
      { field: 'name', ...validators.minLength(3, 'Name must be at least 3 characters') },
      { field: 'slug', ...validators.minLength(3, 'Slug must be at least 3 characters') }
    ])

    const amountNum = parseFloat(amountDisplay)
    if (isNaN(amountNum) || amountNum < 0) {
      formErrors.amount = 'Amount must be a positive number'
    }

    if (drawerMode === 'add' && (!formData.digital?.type || formData.digital.type.trim() === '')) {
      formErrors.digital_type = 'Digital type is required'
    }

    if (Object.keys(formErrors).length > 0) {
      return
    }

    const isUpdate = drawerMode === 'edit' && drawerProduct !== null
    const url = isUpdate ? `/api/_/products/${drawerProduct!.product.id}` : '/api/_/products'
    const submitData: Partial<Product> = {
      ...formData,
      // Цены в базе хранятся в центах, поэтому умножаем на 100 при сохранении
      amount: Math.round((parseFloat(amountDisplay) || 0) * 100)
    }

    const result = await saveData<Product>(url, submitData, isUpdate, 'Product saved', 'Failed to save product')
    if (result) {
      if (isUpdate) {
        updateProductInList(result)
      } else {
        await loadProducts()
      }
      closeDrawer()
    } else if (isUpdate && drawerProduct) {
      const updatedProduct = await loadData<Product>(
        `/api/_/products/${drawerProduct.product.id}`,
        'Failed to load product'
      )
      if (updatedProduct) {
        updateProductInList(updatedProduct)
      }
    }
  }

  async function handleDeleteProduct() {
    if (!fullProductData || !confirmDelete('product', fullProductData.name)) {
      return
    }

    const success = await deleteData(
      `/api/_/products/${fullProductData.id}`,
      'Product deleted',
      'Failed to delete product'
    )
    if (success) {
      await loadProducts()
      closeDrawer()
    }
  }

  async function toggleActiveInEdit() {
    if (!fullProductData) return

    const newActive = !fullProductData.active
    const updatedProduct = { ...fullProductData, active: newActive }
    updateProductInList(updatedProduct)

    const result = await toggleActiveApi(
      `/api/_/products/${fullProductData.id}/active`,
      'Product status updated',
      'Failed to update product'
    )

    if (result === null) {
      updateProductInList(fullProductData)
    } else {
      const serverProduct = await loadData<Product>(`/api/_/products/${fullProductData.id}`, 'Failed to load product')
      if (serverProduct) {
        updateProductInList(serverProduct)
      }
    }
  }

  function addMetadataRecord() {
    formData.metadata = [...(formData.metadata || []), { key: '', value: '' }]
  }

  function deleteMetadataRecord(index: number) {
    formData.metadata = (formData.metadata || []).filter((_, i) => i !== index)
  }

  function addAttributeRecord() {
    formData.attributes = [...(formData.attributes || []), '']
  }

  function deleteAttributeRecord(index: number) {
    formData.attributes = (formData.attributes || []).filter((_, i) => i !== index)
  }

  async function deleteProductImage(index: number) {
    if (!fullProductData || !productImages || !productImages[index]) return

    try {
      const imageId = productImages[index].id
      const res = await apiDelete(`/api/_/products/${fullProductData.id}/image/${imageId}`)
      if (res.success && productImages) {
        productImages = productImages.filter((_, i) => i !== index)
        showMessage('Image deleted', 'connextSuccess')
      } else {
        showMessage(res.message || 'Failed to delete image', 'connextError')
      }
    } catch (error) {
      showMessage('Network error', 'connextError')
    }
  }

  async function handleDelete(product: Product, index: number) {
    if (!confirmDelete('product', product.name)) {
      return
    }

    const success = await deleteData(`/api/_/products/${product.id}`, 'Product deleted', 'Failed to delete product')
    if (success) {
      await loadProducts()
    }
  }

  async function toggleActive(product: Product, index: number) {
    const originalActive = product.active
    const newActive = !product.active
    products = products.map((p, i) => (i === index ? { ...p, active: newActive } : p))

    try {
      const res = await apiUpdate(`/api/_/products/${product.id}/active`, {})
      if (!res.success) {
        products = products.map((p, i) => (i === index ? { ...p, active: originalActive } : p))
        showMessage(res.message || 'Failed to update product', 'connextError')
      } else {
        showMessage(res.message || 'Product status updated', 'connextSuccess')
      }
    } catch (error) {
      products = products.map((p, i) => (i === index ? { ...p, active: originalActive } : p))
      showMessage('Network error', 'connextError')
    }
  }

  function digitalTypeIco(type: string | undefined): string {
    if (!type) return 'cube-transparent'
    switch (type) {
      case 'file':
        return 'paper-clip'
      case 'data':
        return 'queue-list'
      default:
        return 'cube-transparent'
    }
  }

  async function updateActive(index: number) {
    await toggleActive(products[index], index)
  }

  function handleEditorUpdate(event: CustomEvent) {
    formData.description = event.detail
  }

  function handleUpload(event: CustomEvent) {
    if (event.detail.success) {
      showMessage('File uploaded', 'connextSuccess')
      if (event.detail.result && fullProductData) {
        productImages = [...(productImages || []), event.detail.result]
      }
      loadProducts()
    }
  }
</script>

<svelte:component this={Main}>
  <div class="mb-5 flex items-center justify-between">
    <h1>Products</h1>
    <FormButton name="Add Product" color="green" ico="plus" on:click={openAdd} />
  </div>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else if products.length === 0}
    <div class="py-8 text-center text-gray-500">No products found</div>
  {:else}
    <table>
      <thead>
        <tr>
          <th class="w-28"></th>
          <th>Name</th>
          <th class="w-32">Slug</th>
          <th class="w-32">Price</th>
          <th class="w-12 px-4 py-2">
            <SvgIcon name="cube" className="h-5 w-5" stroke="currentColor" />
          </th>
          <th class="w-24 px-4 py-2"></th>
        </tr>
      </thead>
      <tbody>
        {#each products as product, index}
          <tr class:opacity-30={!product.active}>
            <td>
              {#if product.images && product.images.length > 0}
                <a href="/uploads/{product.images[0].name}.{product.images[0].ext}" target="_blank">
                  <img
                    style="width: 100%; max-width: 80px"
                    src="/uploads/{product.images[0].name}_sm.{product.images[0].ext}"
                    alt={product.name}
                    loading="lazy"
                  />
                </a>
              {:else}
                <img style="width: 100%; max-width: 80px" src="/assets/img/noimage.png" alt="" loading="lazy" />
              {/if}
            </td>
            <td on:click={() => openView(product, index)}>
              <div class="font-bold">{product.name}</div>
              {#if product.brief}
                <span class="hidden text-gray-400 xl:block">{product.brief}</span>
              {/if}
            </td>
            <td>
              {#if product.active}
                <a href="/products/{product.slug}" target="_blank">{product.slug}</a>
              {:else}
                <span>{product.slug}</span>
              {/if}
            </td>
            <td on:click={() => openView(product, index)}>
              {costFormat(product.amount)}
              {currency}
            </td>
            <td class="px-4 py-2">
              {#if product.digital && product.digital.type}
                <SvgIcon
                  name={digitalTypeIco(product.digital.type)}
                  className="h-5 w-5 cursor-pointer {product.digital.filled === true ? 'text-black' : 'text-red-600'}"
                  on:click={() => openDigital(product, index)}
                  stroke="currentColor"
                />
              {/if}
            </td>
            <td class="px-4 py-2">
              <div class="flex">
                <div class="pr-3">
                  <SvgIcon
                    name="pencil-square"
                    className="h-5 w-5 cursor-pointer"
                    on:click={() => openEdit(product, index)}
                    stroke="currentColor"
                  />
                </div>
                <div class="pr-3">
                  <SvgIcon
                    name="rocket"
                    className="h-5 w-5 cursor-pointer"
                    on:click={() => openSeo(product, index)}
                    stroke="currentColor"
                  />
                </div>
                <div>
                  <SvgIcon
                    name={product.active ? 'eye' : 'eye-slash'}
                    className="h-5 w-5 cursor-pointer"
                    on:click={() => toggleActive(product, index)}
                    stroke="currentColor"
                  />
                </div>
              </div>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
</svelte:component>

{#if drawerOpen}
  <Drawer isOpen={drawerOpen} on:close={closeDrawer} maxWidth="710px">
    {#if drawerMode === 'view' && drawerProduct}
      <ProductView drawer={drawerProduct} {updateActive} on:close={closeDrawer} />
    {:else if drawerMode === 'seo' && drawerProduct}
      <ProductSeo drawer={drawerProduct} on:close={closeDrawer} />
    {:else if drawerMode === 'digital' && drawerProduct}
      <ProductDigital drawer={drawerProduct} onContentUpdate={handleDigitalContentUpdate} on:close={closeDrawer} />
    {:else}
      <div>
        <div class="pb-8">
          <div class="flex items-center">
            <div class="pr-3">
              <h1>{drawerMode === 'add' ? 'Add' : `Edit ${formData.name || ''}`}</h1>
            </div>
            {#if drawerMode === 'edit' && fullProductData}
              <div>
                <SvgIcon
                  name={fullProductData.active ? 'eye' : 'eye-slash'}
                  className="h-5 w-5 cursor-pointer"
                  on:click={toggleActiveInEdit}
                  stroke="currentColor"
                />
              </div>
            {/if}
          </div>
        </div>

        <form on:submit|preventDefault={handleSubmit}>
          <div class="flow-root">
            <dl class="mx-auto -my-3 mt-2 mb-0 space-y-4 text-sm">
              <FormInput id="name" title="Name" bind:value={formData.name} error={formErrors.name} ico="at-symbol" />
              <div class="flex flex-row">
                <div class="pr-3">
                  <FormInput
                    id="amount"
                    type="text"
                    title="Amount"
                    bind:value={amountDisplay}
                    error={formErrors.amount}
                    ico="money"
                  />
                </div>
                <div class="mt-3">{currency}</div>
              </div>

              {#if drawerMode === 'add'}
                <div class="flex">
                  <div class="grow pr-3">
                    <FormInput
                      id="slug"
                      title="Slug"
                      bind:value={formData.slug}
                      error={formErrors.slug}
                      ico="glob-alt"
                    />
                  </div>
                  <div class="grow">
                    <FormSelect
                      id="digital_type"
                      title="Digital type"
                      options={['file', 'data', 'api']}
                      bind:value={formData.digital.type}
                      error={formErrors.digital_type}
                      ico="cube"
                    />
                  </div>
                </div>
              {:else}
                <FormInput id="slug" title="Slug" bind:value={formData.slug} error={formErrors.slug} ico="glob-alt" />
              {/if}

              <hr />
              <p class="font-semibold">Metadata</p>
              {#each formData.metadata || [] as metadata, index}
                <div class="flex">
                  <div class="grow pr-3">
                    <FormInput id="mtd-key-{index}" type="text" title="Key" bind:value={metadata.key} />
                  </div>
                  <div class="grow">
                    <FormInput id="mtd-value-{index}" type="text" title="Value" bind:value={metadata.value} />
                  </div>
                  <div
                    class="flex-none cursor-pointer pt-3 pl-3"
                    role="button"
                    tabindex="0"
                    on:click={() => deleteMetadataRecord(index)}
                    on:keydown={(e) => {
                      if (e.key === 'Enter' || e.key === ' ') {
                        e.preventDefault()
                        deleteMetadataRecord(index)
                      }
                    }}
                  >
                    <SvgIcon name="trash" className="h-5 w-5" stroke="currentColor" />
                  </div>
                </div>
              {/each}
              <div class="flex">
                <div class="grow"></div>
                <div class="mt-2 flex-none">
                  <button
                    type="button"
                    class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700"
                    on:click={addMetadataRecord}
                  >
                    Add metadata record
                  </button>
                </div>
              </div>

              <hr />
              <p class="font-semibold">Attributes</p>
              {#each formData.attributes || [] as attribute, index}
                <div class="flex">
                  <div class="grow">
                    <FormInput id="atr-key-{index}" type="text" title="" bind:value={formData.attributes[index]} />
                  </div>
                  <div
                    class="flex-none cursor-pointer pt-3 pl-3"
                    role="button"
                    tabindex="0"
                    on:click={() => deleteAttributeRecord(index)}
                    on:keydown={(e) => {
                      if (e.key === 'Enter' || e.key === ' ') {
                        e.preventDefault()
                        deleteAttributeRecord(index)
                      }
                    }}
                  >
                    <SvgIcon name="trash" className="h-5 w-5" stroke="currentColor" />
                  </div>
                </div>
              {/each}
              <div class="flex">
                <div class="grow"></div>
                <div class="mt-2 flex-none">
                  <button
                    type="button"
                    class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700"
                    on:click={addAttributeRecord}
                  >
                    Add attribute record
                  </button>
                </div>
              </div>

              {#if drawerMode === 'edit' && fullProductData}
                <hr />
                <p class="font-semibold">Images</p>
                {#if productImages && productImages.length > 0}
                  <div class="grid grid-cols-4 content-start gap-4">
                    {#each productImages as image, index}
                      <div class="relative" style="width: 100%; max-width: 150px">
                        <a href="/uploads/{image.name}.{image.ext}" target="_blank">
                          <img src="/uploads/{image.name}_sm.{image.ext}" alt="" />
                        </a>
                        <div
                          role="button"
                          tabindex="0"
                          class="absolute end-4 top-4 cursor-pointer bg-white p-2"
                          on:click={() => deleteProductImage(index)}
                          on:keydown={(e) => {
                            if (e.key === 'Enter' || e.key === ' ') {
                              e.preventDefault()
                              deleteProductImage(index)
                            }
                          }}
                        >
                          <SvgIcon name="trash" className="h-5 w-5" stroke="currentColor" />
                        </div>
                      </div>
                    {/each}
                  </div>
                {/if}
                <Upload
                  section="image"
                  productId={fullProductData.id}
                  accept=".jpg,.jpeg,.png"
                  on:added={handleUpload}
                />
              {/if}

              <hr />
              <p class="font-semibold">Short description</p>
              <FormTextarea id="brief" title="Brief" bind:value={formData.brief} />

              <hr />
              <p class="font-semibold">Description</p>
              <Editor
                bind:modelValue={formData.description}
                placeholder="type description here"
                on:update:modelValue={handleEditorUpdate}
              />
            </dl>
          </div>

          <div class="pt-8">
            <div class="flex">
              <div class="flex-none">
                <FormButton type="submit" name={drawerMode === 'add' ? 'Add' : 'Save'} color="green" />
                <FormButton type="button" name="Close" color="gray" on:click={closeDrawer} />
              </div>
              <div class="grow"></div>
              {#if drawerMode === 'edit' && fullProductData}
                <div class="mt-2 flex-none">
                  <span
                    on:click={handleDeleteProduct}
                    on:keydown={(e) => {
                      if (e.key === 'Enter' || e.key === ' ') {
                        e.preventDefault()
                        handleDeleteProduct()
                      }
                    }}
                    role="button"
                    tabindex="0"
                    class="cursor-pointer text-red-700">Delete</span
                  >
                </div>
              {/if}
            </div>
          </div>
        </form>
      </div>
    {/if}
  </Drawer>
{/if}
