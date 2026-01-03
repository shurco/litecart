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
  import Pagination from '$lib/components/Pagination.svelte'
  import { loadData, saveData, deleteData, toggleActive as toggleActiveApi } from '$lib/utils/apiHelpers'
  import { formatDate, confirmDelete } from '$lib/utils'
  import { validators, validateFields } from '$lib/utils/validation'
  import { MIN_NAME_LENGTH, MIN_SLUG_LENGTH, ERROR_MESSAGES } from '$lib/constants/validation'
  import type { Page } from '$lib/types/models'

  interface PagesResponse {
    pages: Page[]
    total: number
    page: number
    limit: number
  }

  import { DEFAULT_PAGE_SIZE } from '$lib/constants/pagination'
  import { DRAWER_CLOSE_DELAY_MS } from '$lib/constants/ui'

  let pages = $state<Page[]>([])
  let loading = $state(true)
  let drawerOpen = $state(false)
  let drawerMode = $state<'add' | 'edit' | 'seo' | 'view'>('add')
  let drawerPage = $state<Page | null>(null)
  let currentPage = $state(1)
  let limit = $state(DEFAULT_PAGE_SIZE)
  let total = $state(0)

  let formData = $state<Omit<Page, 'id' | 'created' | 'updated' | 'seo'>>({
    name: '',
    slug: '',
    position: 'header',
    content: '',
    active: true
  })

  const positionOptions = ['header', 'footer']
  let formErrors = $state<Record<string, string>>({})

  onMount(async () => {
    await loadPages()
  })

  async function loadPages(page = currentPage) {
    loading = true
    currentPage = page
    const result = await loadData<PagesResponse>(
      `/api/_/pages?page=${page}&limit=${limit}`,
      'Failed to load pages'
    )
    if (result) {
      pages = result.pages || []
      total = result.total || 0
    }
    loading = false
  }

  function handlePageChange(page: number) {
    loadPages(page)
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
      }, DRAWER_CLOSE_DELAY_MS)
    }
  }

  async function handleSubmit(event: SubmitEvent) {
    event.preventDefault()
    formErrors = validateFields(formData, [
      { field: 'name', ...validators.minLength(MIN_NAME_LENGTH, ERROR_MESSAGES.NAME_TOO_SHORT) },
      { field: 'slug', ...validators.minLength(MIN_SLUG_LENGTH, ERROR_MESSAGES.SLUG_TOO_SHORT) }
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
    const originalActive = page.active
    const newActive = !page.active
    
    // Optimistic update - update directly instead of map
    pages[index] = { ...pages[index], active: newActive }

    // Then make API call
    const updatedPage = await toggleActiveApi<Page>(
      `/api/_/pages/${page.id}/active`,
      'Page status updated',
      'Failed to update page'
    )

    // If API call failed, revert the change
    if (!updatedPage) {
      pages[index] = { ...pages[index], active: originalActive }
    } else {
      // Update with the response from server including updated timestamp
      pages[index] = updatedPage
    }
  }

  function handleEditorUpdate(value: string) {
    formData.content = value
  }

  function openSeo(page: Page) {
    drawerPage = page
    drawerMode = 'seo'
    drawerOpen = true
  }
</script>

<Main>
  <div class="mb-5 flex items-center justify-between">
    <h1>Pages</h1>
    <FormButton name="Add Page" color="green" ico="plus" onclick={openAdd} />
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
        {#each pages as page, index (page.id)}
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
                    onclick={() => openEdit(page)}
                    stroke="currentColor"
                  />
                </div>
                <div class="pr-3">
                  <SvgIcon
                    name="rocket"
                    className="h-5 w-5 cursor-pointer"
                    onclick={() => openSeo(page)}
                    stroke="currentColor"
                  />
                </div>
                <div>
                  <SvgIcon
                    name={page.active ? 'eye' : 'eye-slash'}
                    className="h-5 w-5 cursor-pointer"
                    onclick={() => toggleActive(page, index)}
                    stroke="currentColor"
                  />
                </div>
              </div>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>

    {#if total > 0}
      <Pagination
        currentPage={currentPage}
        totalPages={Math.ceil(total / limit)}
        onPageChange={handlePageChange}
      />
    {/if}
    {/if}
</Main>

{#if drawerOpen}
  <Drawer isOpen={drawerOpen} onclose={closeDrawer} maxWidth="710px">
    {#if drawerMode === 'seo' && drawerPage}
      <PageSeo page={drawerPage} onclose={closeDrawer} />
    {:else}
      <div class="pb-8">
        <div class="flex items-center">
          <div class="pr-3">
            <h1>{drawerMode === 'add' ? 'New page' : 'Page setup'}</h1>
          </div>
        </div>
      </div>

      <form onsubmit={handleSubmit}>
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
              onupdateModelValue={handleEditorUpdate}
            />
          </dl>
        </div>

        <div class="pt-8">
          <div class="flex">
            <div class="flex-none">
              <FormButton type="submit" name={drawerMode === 'add' ? 'Add' : 'Save'} color="green" />
              <FormButton type="button" name="Close" color="gray" onclick={closeDrawer} />
            </div>
            <div class="grow"></div>
            {#if drawerMode === 'edit' && drawerPage}
              <div class="mt-4 flex-none">
                <span
                  role="button"
                  tabindex="0"
                  onclick={() => drawerPage && handleDelete(drawerPage)}
                  onkeydown={(e) => {
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
