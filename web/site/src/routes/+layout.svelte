<script lang="ts">
  import MainLayout from "$lib/layouts/MainLayout.svelte";
  import "../app.css";
</script>

<MainLayout>
  {#snippet children()}
    <slot />
  {/snippet}
</MainLayout>
