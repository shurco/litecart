<script lang="ts">
  import SvgIcon from '../SvgIcon.svelte'

  interface Props {
    id?: string
    type?: string
    title?: string
    ico?: string
    error?: string
    value?: string
    placeholder?: string
    onfocusout?: (event: FocusEvent) => void
    oninput?: (event: Event) => void
  }

  let {
    id = 'name',
    type = 'text',
    title = 'Name',
    ico = undefined,
    error = undefined,
    value = $bindable(''),
    placeholder = '',
    onfocusout,
    oninput
  }: Props = $props()

  let computedPlaceholder = $derived(placeholder || `Enter ${id}`)
</script>

<div>
  <label for={id} class={error ? 'border-red-500' : ''}>
    <input
      {type}
      {id}
      bind:value
      class="form-input field peer"
      placeholder={computedPlaceholder}
      autocomplete="on"
      onfocusout={onfocusout}
      oninput={oninput}
    />
    {#if title}
      <span
        class="title peer-placeholder-shown:top-1/2 peer-placeholder-shown:text-sm peer-placeholder-shown:text-gray-400 peer-focus:top-0 peer-focus:text-xs peer-focus:text-gray-700"
      >
        {title}
      </span>
    {/if}
    {#if ico}
      <span class="ico">
        <SvgIcon name={ico} stroke="currentColor" className="h-5 w-5 {error ? 'text-red-500' : 'text-gray-400'}" />
      </span>
    {/if}
  </label>
  {#if error}
    <span class="pl-4 text-sm text-red-500">{error}</span>
  {/if}
</div>
