<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  export let prefix: string = "icon";
  export let name: string;
  export let className: string = "h-6 w-6";
  export let stroke: string | undefined = undefined;

  const dispatch = createEventDispatcher();
  
  // Check if click handler is provided
  $: hasClickHandler = $$props.onclick !== undefined;
</script>

<svg 
  aria-hidden="true" 
  class={className} 
  fill={stroke ? 'none' : 'currentColor'} 
  stroke={stroke || 'none'} 
  stroke-width={stroke ? '1.5' : '0'}
  viewBox="0 0 24 24"
  xmlns="http://www.w3.org/2000/svg"
  on:click={(e) => dispatch('click', e)}
  on:mouseenter={(e) => dispatch('mouseenter', e)}
  on:mouseleave={(e) => dispatch('mouseleave', e)}
  role={hasClickHandler ? 'button' : 'img'}
  {...(hasClickHandler ? { tabindex: 0 } : {})}
  style={hasClickHandler ? 'cursor: pointer;' : ''}
>
  <use href={`#${prefix}-${name}`} />
</svg>
