<script lang="ts">
  import SvgIcon from '../SvgIcon.svelte'
  import { DEFAULT_BUTTON_NAME } from '$lib/constants/ui'

  interface Props {
    color: string
    name?: string
    ico?: string
    type?: 'button' | 'submit' | 'reset'
    onclick?: (event: MouseEvent) => void
    class?: string
  }

  let {
    color,
    name = DEFAULT_BUTTON_NAME,
    ico = undefined,
    type = 'button',
    onclick,
    class: className = ''
  }: Props = $props()

  const COLOR_CLASSES: Record<string, string[]> = {
    gray: ['bg-gray-600', 'bg-gray-500'],
    gray_lite: ['bg-gray-400', 'bg-gray-300'],
    green: ['bg-green-600', 'bg-green-500'],
    yellow: ['bg-yellow-600', 'bg-yellow-500'],
    red: ['bg-red-600', 'bg-red-500'],
    cyan: ['bg-cyan-600', 'bg-cyan-500']
  }

  let colorClasses = $derived(
    color && COLOR_CLASSES[color] ? `${COLOR_CLASSES[color][0]} active:${COLOR_CLASSES[color][1]}` : ''
  )
  let icoClasses = $derived(ico ? 'focus:outline-none focus:ring' : '')
</script>

<button
  class="group relative inline-flex cursor-pointer items-center overflow-hidden rounded px-8 py-2 text-white {colorClasses} {icoClasses} {className}"
  {type}
  onclick={onclick}
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
