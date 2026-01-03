<script lang="ts">
  interface Props {
    prefix?: string
    name: string
    className?: string
    stroke?: string
    onclick?: (event: MouseEvent) => void
    onmouseenter?: (event: MouseEvent) => void
    onmouseleave?: (event: MouseEvent) => void
  }

  let {
    prefix = 'icon',
    name,
    className = 'h-6 w-6',
    stroke = undefined,
    onclick,
    onmouseenter,
    onmouseleave
  }: Props = $props()

  let hasClickHandler = $derived(onclick !== undefined)
</script>

<svg
  aria-hidden="true"
  class={className}
  fill={stroke ? 'none' : 'currentColor'}
  stroke={stroke || 'none'}
  stroke-width={stroke ? '1.5' : '0'}
  viewBox="0 0 24 24"
  xmlns="http://www.w3.org/2000/svg"
  onclick={onclick}
  onmouseenter={onmouseenter}
  onmouseleave={onmouseleave}
  role={hasClickHandler ? 'button' : 'img'}
  {...(hasClickHandler ? { tabindex: 0 } : {})}
  style={hasClickHandler ? 'cursor: pointer;' : ''}
>
  <use href={`#${prefix}-${name}`} />
</svg>
