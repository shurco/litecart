<script lang="ts">
  import SvgIcon from '../SvgIcon.svelte';

  export let id: string = 'name';
  export let type: string = 'text';
  export let title: string = 'Name';
  export let ico: string | undefined = undefined;
  export let error: string | undefined = undefined;
  export let value: string = '';
  export let placeholder: string = '';

  $: computedPlaceholder = placeholder || `Enter ${id}`;
</script>

<div>
  <label for={id} class={error ? 'border-red-500' : ''}>
    <input
      type={type}
      {id}
      bind:value
      class="form-input field peer"
      placeholder={computedPlaceholder}
      autocomplete="on"
      on:focusout
    />
    {#if title}
      <span class="peer-placeholder-shown:top-1/2 peer-placeholder-shown:text-sm peer-focus:top-0 peer-focus:text-xs peer-placeholder-shown:text-gray-400 peer-focus:text-gray-700 title">
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
    <span class="text-sm text-red-500 pl-4">{error}</span>
  {/if}
</div>