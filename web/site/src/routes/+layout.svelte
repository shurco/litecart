<script lang="ts">
  import MainLayout from '$lib/layouts/MainLayout.svelte'
  import { page } from '$app/stores'
  import '../app.css'

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
  <slot />
{:else}
  <MainLayout>
    {#snippet children()}
      <slot />
    {/snippet}
  </MainLayout>
{/if}
