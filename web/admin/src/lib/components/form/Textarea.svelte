<script lang="ts">
  import SvgIcon from '../SvgIcon.svelte'

  interface Props {
    id?: string
    title?: string
    ico?: string
    error?: string
    value?: string
    placeholder?: string
    rows?: number
    onfocusout?: (event: FocusEvent) => void
  }

  let {
    id = 'name',
    title = 'Name',
    ico = undefined,
    error = undefined,
    value = $bindable(''),
    placeholder = '',
    rows = 4,
    onfocusout
  }: Props = $props()

  let computedPlaceholder = $derived(placeholder || `Enter ${id}`)
</script>

<div>
  <label for={id} class={error ? 'border-red-500' : ''}>
    <textarea {id} bind:value {rows} class="form-textarea field peer" placeholder={computedPlaceholder} onfocusout={onfocusout}
    ></textarea>
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
    <span class="error">{error}</span>
  {/if}
</div>
