<script lang="ts">
  import MainLayout from '$lib/layouts/MainLayout.svelte'
  import { page } from '$app/stores'
  import '../app.css'

  interface Props {
    children: import('svelte').Snippet
  }

  let { children }: Props = $props()

  // Check if this is an error page (has error or status >= 400)
  const isErrorPage = $derived($page.error !== null || ($page.status && $page.status >= 400))
</script>

{#if isErrorPage}
  <style>
    :global(header),
    :global(footer) {
      display: none !important;
    }
  </style>
  {@render children()}
{:else}
  <MainLayout>
    {@render children()}
  </MainLayout>
{/if}
