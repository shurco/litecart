<script lang="ts">
  import SvgIcon from '../SvgIcon.svelte'
  import { createEventDispatcher } from 'svelte'

  export let color: string
  export let name: string = 'Name'
  export let ico: string | undefined = undefined
  export let type: 'button' | 'submit' | 'reset' = 'button'

  const dispatch = createEventDispatcher()

  const colors: Record<string, string[]> = {
    gray: ['bg-gray-600', 'bg-gray-500'],
    gray_lite: ['bg-gray-400', 'bg-gray-300'],
    green: ['bg-green-600', 'bg-green-500'],
    yellow: ['bg-yellow-600', 'bg-yellow-500'],
    red: ['bg-red-600', 'bg-red-500'],
    cyan: ['bg-cyan-600', 'bg-cyan-500']
  }

  $: colorClasses = color && colors[color] ? `${colors[color][0]} active:${colors[color][1]}` : ''
  $: icoClasses = ico ? 'focus:outline-none focus:ring' : ''

  function handleClick(event: MouseEvent) {
    dispatch('click', event)
  }
</script>

<button
  class="group relative inline-flex cursor-pointer items-center overflow-hidden rounded px-8 py-2 text-white {colorClasses} {icoClasses}"
  {type}
  on:click={handleClick}
  {...$$restProps}
>
  {#if ico}
    <SvgIcon
      name={ico}
      stroke="currentColor"
      className="h-4 w-4 absolute -start-full transition-all group-hover:start-4"
    />
  {/if}
  <span class="text-sm font-medium {ico ? 'transition-all group-hover:ms-2 group-hover:-me-2' : ''}">
    {name}
  </span>
</button>
