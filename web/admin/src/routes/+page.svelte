<script lang="ts">
  import { goto } from '$app/navigation';
  import { base } from '$app/paths';
  import { onMount } from 'svelte';
  import Main from '$lib/layouts/Main.svelte';

  onMount(() => {
    goto(`${base}/products`);
  });
</script>

<svelte:component this={Main}>
  <div>Redirecting...</div>
</svelte:component>