<script lang="ts">
  import SvgIcon from '../SvgIcon.svelte';

  export let id: string = 'name';
  export let title: string = 'Name';
  export let options: string[] | Record<string, string> = [];
  export let ico: string | undefined = undefined;
  export let error: string | undefined = undefined;
  export let value: string = '';

  $: optionList = Array.isArray(options) 
    ? options.map(opt => ({ key: opt, value: opt }))
    : Object.entries(options).map(([key, val]) => ({ key, value: val }));
</script>

<div>
  <label for={id} class={error ? 'border-red-500' : ''}>
    <select
      {id}
      bind:value
      class="form-select field peer"
    >
      <option value="" disabled>Please select</option>
      {#each optionList as option}
        <option value={option.key}>{option.value}</option>
      {/each}
    </select>
    <span class="title">{title}</span>
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