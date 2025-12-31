<script lang="ts">
  import { page } from "$app/stores";
  import { apiGet } from "$lib/utils/api";
  import type { Page } from "$lib/types/models";
  import { updateSEOTags } from "$lib/utils/seo";
  import { isBrowser } from "$lib/utils/browser";
  import NotFoundPage from "$lib/components/NotFoundPage.svelte";

  let content = $state<Page | null>(null);
  let notFound = $state(false);
  let loading = $state(true);

  $effect(() => {
    const slug = $page.params.slug;
    if (slug) {
      // Reset state when slug changes
      content = null;
      notFound = false;
      loading = true;
      loadPage(slug);
    }
  });

  async function loadPage(slug: string) {
    const res = await apiGet<Page>(`/api/pages/${slug}`);
    loading = false;
    
    if (res.success && res.result) {
      content = res.result;

      if (isBrowser() && content.seo) {
        updateSEOTags(content.seo);
      }
    } else {
      // Page not found
      notFound = true;
    }
  }
</script>

{#if loading}
  <div class="min-h-screen bg-white flex items-center justify-center">
    <div class="border-4 border-black bg-yellow-300 px-8 py-6 inline-block">
      <p class="text-xl font-black uppercase tracking-wider text-black">
        LOADING...
      </p>
    </div>
  </div>
{:else if notFound}
  <NotFoundPage />
{:else if content}
  <section class="min-h-screen bg-white py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-screen-xl mx-auto">
      <div>
        <h1 class="text-4xl sm:text-5xl font-black uppercase tracking-tighter text-black mb-8 border-b-4 border-black pb-6">
          {content.name}
        </h1>
        <div class="tiptap text-lg font-bold text-black leading-relaxed space-y-6">
          {@html content.content}
        </div>
      </div>
    </div>
  </section>
{/if}
