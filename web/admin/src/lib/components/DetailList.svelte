<script lang="ts">
  export let name: string;
  export let grid: boolean = false;
</script>

<div class="flex gap-4 py-3">
  <div class="flex-none w-1/5 font-medium text-gray-900">{name}</div>
  <div class="flex-none w-4"></div>
  <div class="grow text-gray-700 {grid ? 'grid grid-cols-3 content-start gap-4' : ''}">
    <slot />
  </div>
</div>