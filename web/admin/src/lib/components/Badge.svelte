<script lang="ts">
  import SvgIcon from './SvgIcon.svelte'

  export let ico: string | undefined = undefined
  export let svgClass: string = '-ms-0.5 me-5.5 h-5 w-5'
</script>

<span class="badge">
  {#if ico}
    <SvgIcon name={ico} stroke="currentColor" className={ico ? `-ms-0.5 me-1.5 ${svgClass}` : 'h-0 w-0'} />
  {/if}
  <p>
    <slot />
  </p>
</span>

<style>
  @reference "tailwindcss";

  :global(.badge) {
    @apply inline-flex items-center justify-center rounded-full px-2.5 py-0.5;
  }

  :global(.badge p) {
    @apply text-xs whitespace-nowrap;
  }
</style>
