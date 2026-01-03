<script lang="ts">
  interface Props {
    name: string
    grid?: boolean
    children?: import('svelte').Snippet
  }

  let { name, grid = false, children }: Props = $props()
</script>

<div class="flex gap-4 py-3">
  <div class="w-1/5 flex-none font-medium text-gray-900">{name}</div>
  <div class="w-4 flex-none"></div>
  <div class="grow text-gray-700 {grid ? 'grid grid-cols-3 content-start gap-4' : ''}">
    {#if children}
      {@render children()}
    {/if}
  </div>
</div>
