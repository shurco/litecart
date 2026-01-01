<script lang="ts">
  import { onMount } from 'svelte'
  import Main from '$lib/layouts/Main.svelte'
  import Drawer from '$lib/components/Drawer.svelte'
  import PageSeo from '$lib/components/page/Seo.svelte'
  import FormButton from '$lib/components/form/Button.svelte'
  import FormInput from '$lib/components/form/Input.svelte'
  import FormSelect from '$lib/components/form/Select.svelte'
  import Editor from '$lib/components/Editor.svelte'
  import SvgIcon from '$lib/components/SvgIcon.svelte'
  import { loadData, saveData, deleteData, toggleActive as toggleActiveApi } from '$lib/utils/apiHelpers'
  import { formatDate, confirmDelete } from '$lib/utils'
  import { validators, validateFields } from '$lib/utils/validation'
  import type { Page } from '$lib/types/models'

  let pages: Page[] = []
  let loading = true
  let drawerOpen = false
  let drawerMode: 'add' | 'edit' | 'seo' | 'view' = 'add'
  let drawerPage: Page | null = null

  let formData: Omit<Page, 'id' | 'created' | 'updated' | 'seo'> = {
    name: '',
    slug: '',
    position: 'header',
    content: '',
    active: true
  }

  const positionOptions = ['header', 'footer']
  let formErrors: Record<string, string> = {}

  onMount(async () => {
    await loadPages()
  })

  async function loadPages() {
    loading = true
    const result = await loadData<Page[]>('/api/_/pages', 'Failed to load pages')
    if (result) {
      pages = result
    }
    loading = false
  }

  function openAdd() {
    formData = {
      name: '',
      slug: '',
      position: 'header',
      content: '',
      active: true
    }
    formErrors = {}
    drawerPage = null
    drawerMode = 'add'
    drawerOpen = true
  }

  async function openEdit(page: Page) {
    // Load full page data including content
    const fullPage = await loadData<Page>(`/api/_/pages/${page.id}`, 'Failed to load page')
    if (!fullPage) {
      return
    }

    formData = {
      name: fullPage.name || '',
      slug: fullPage.slug || '',
      position: fullPage.position || 'header',
      content: fullPage.content ?? '',
      active: fullPage.active !== undefined ? fullPage.active : true
    }
    drawerPage = fullPage
    formErrors = {}
    drawerMode = 'edit'
    drawerOpen = true
  }

  function closeDrawer() {
    if (drawerOpen) {
      drawerOpen = false
      setTimeout(() => {
        drawerPage = null
        drawerMode = 'add'
      }, 200)
    }
  }

  async function handleSubmit() {
    formErrors = validateFields(formData, [
      { field: 'name', ...validators.minLength(3, 'Name must be at least 3 characters') },
      { field: 'slug', ...validators.minLength(3, 'Slug must be at least 3 characters') }
    ])

    if (Object.keys(formErrors).length > 0) {
      return
    }

    const isUpdate = drawerMode === 'edit' && drawerPage !== null
    const url = isUpdate && drawerPage ? `/api/_/pages/${drawerPage.id}` : '/api/_/pages'

    const result = await saveData<Page>(url, formData, isUpdate, 'Page saved', 'Failed to save page')
    if (result) {
      if (isUpdate && drawerPage) {
        // Find and update the specific page in the list reactively
        const pageId = drawerPage.id
        const pageIndex = pages.findIndex((p) => p.id === pageId)
        if (pageIndex !== -1) {
          pages = pages.map((p, i) => (i === pageIndex ? { ...result } : p))
        }
      } else {
        // For new pages, reload the list
        await loadPages()
      }
      closeDrawer()
    }
  }

  async function handleDelete(page: Page) {
    if (!confirmDelete('page', page.name)) {
      return
    }

    const success = await deleteData(`/api/_/pages/${page.id}`, 'Page deleted', 'Failed to delete page')
    if (success) {
      await loadPages()
    }
  }

  async function toggleActive(page: Page, index: number) {
    // Optimistically update UI first
    const newActive = !page.active
    pages = pages.map((p, i) => (i === index ? { ...p, active: newActive } : p))

    // Then make API call
    const updatedPage = await toggleActiveApi<Page>(
      `/api/_/pages/${page.id}/active`,
      'Page status updated',
      'Failed to update page'
    )

    // If API call failed, revert the change
    if (!updatedPage) {
      pages = pages.map((p, i) => (i === index ? { ...p, active: page.active } : p))
    } else {
      // Update with the response from server including updated timestamp
      pages = pages.map((p, i) => (i === index ? { ...updatedPage } : p))
    }
  }

  function handleEditorUpdate(event: CustomEvent<string>) {
    formData.content = event.detail
  }

  function openSeo(page: Page) {
    drawerPage = page
    drawerMode = 'seo'
    drawerOpen = true
  }
</script>

<svelte:component this={Main}>
  <div class="mb-5 flex items-center justify-between">
    <h1>Pages</h1>
    <FormButton name="Add Page" color="green" ico="plus" on:click={openAdd} />
  </div>

  {#if loading}
    <div class="py-8 text-center">Loading...</div>
  {:else if pages.length === 0}
    <div class="py-8 text-center text-gray-500">No pages found</div>
  {:else}
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th class="w-32">Position</th>
          <th class="w-32">Slug</th>
          <th class="w-48">Created</th>
          <th class="w-48">Updated</th>
          <th class="w-24 px-4 py-2"></th>
        </tr>
      </thead>
      <tbody>
        {#each pages as page, index}
          <tr class:opacity-30={!page.active}>
            <td>{page.name}</td>
            <td>{page.position || '-'}</td>
            <td>
              <a href="/{page.slug}" target="_blank">{page.slug}</a>
            </td>
            <td>
              {formatDate(page.created)}
            </td>
            <td>
              {#if page.updated}
                {formatDate(page.updated)}
              {/if}
            </td>
            <td class="px-4 py-2">
              <div class="flex">
                <div class="pr-3">
                  <SvgIcon
                    name="pencil-square"
                    className="h-5 w-5 cursor-pointer"
                    on:click={() => openEdit(page)}
                    stroke="currentColor"
                  />
                </div>
                <div class="pr-3">
                  <SvgIcon
                    name="rocket"
                    className="h-5 w-5 cursor-pointer"
                    on:click={() => openSeo(page)}
                    stroke="currentColor"
                  />
                </div>
                <div>
                  <SvgIcon
                    name={page.active ? 'eye' : 'eye-slash'}
                    className="h-5 w-5 cursor-pointer"
                    on:click={() => toggleActive(page, index)}
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
  <Drawer isOpen={drawerOpen} on:close={() => closeDrawer()} maxWidth="710px">
    {#if drawerMode === 'seo' && drawerPage}
      <PageSeo page={drawerPage} on:close={closeDrawer} />
    {:else}
      <div class="pb-8">
        <div class="flex items-center">
          <div class="pr-3">
            <h1>{drawerMode === 'add' ? 'New page' : 'Page setup'}</h1>
          </div>
        </div>
      </div>

      <form on:submit|preventDefault={handleSubmit}>
        <div class="flow-root">
          <dl class="mx-auto -my-3 mt-4 mb-0 space-y-4 text-sm">
            <FormInput id="name" title="name" bind:value={formData.name} error={formErrors.name} ico="at-symbol" />
            <div class="flex">
              <div class="pr-3">
                <FormSelect
                  id="position"
                  title="Position"
                  options={positionOptions}
                  bind:value={formData.position}
                  error={formErrors.position}
                />
              </div>
              <div>
                <FormInput id="slug" title="Slug" bind:value={formData.slug} error={formErrors.slug} ico="glob-alt" />
              </div>
            </div>

            <hr />
            <p class="font-semibold">Content</p>
            <Editor
              bind:modelValue={formData.content}
              placeholder="type content here"
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
            {#if drawerMode === 'edit' && drawerPage}
              <div class="mt-4 flex-none">
                <span
                  role="button"
                  tabindex="0"
                  on:click={() => drawerPage && handleDelete(drawerPage)}
                  on:keydown={(e) => {
                    if (e.key === 'Enter' || e.key === ' ') {
                      e.preventDefault()
                      drawerPage && handleDelete(drawerPage)
                    }
                  }}
                  class="cursor-pointer text-red-700"
                >
                  Delete
                </span>
              </div>
            {/if}
          </div>
        </div>
      </form>
    {/if}
  </Drawer>
{/if}
