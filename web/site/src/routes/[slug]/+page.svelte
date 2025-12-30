<script lang="ts">
  import { page } from "$app/stores";
  import { apiGet } from "$lib/utils/api";
  import type { Page } from "$lib/types/models";
  import { updateSEOTags } from "$lib/utils/seo";
  import { browser } from "$app/environment";

  let content = $state<Page | null>(null);

  $effect(() => {
    if ($page.params.slug) {
      loadPage($page.params.slug);
    }
  });

  async function loadPage(slug: string) {
    const res = await apiGet<Page>(`/api/pages/${slug}`);
    if (res.success && res.result) {
      content = res.result;

      if (browser && content.seo) {
        updateSEOTags(content.seo);
      }
    }
  }
</script>

{#if content}
  <section>
    <div class="tiptap max-w-screen-xl px-4 py-8 mx-auto sm:px-6 sm:py-12 lg:px-8">
      <h1><strong>{content.name}</strong></h1>
      <div class="mt-10 border-t border-gray-100"></div>
      <div class="mt-10">{@html content.content}</div>
    </div>
  </section>
{/if}
