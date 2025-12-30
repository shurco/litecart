<script lang="ts">
  interface Props {
    name?: string;
    color?: "green" | "yellow" | "red" | "blue";
    ico?: "cart" | "trash" | "plus" | "row";
    type?: "button" | "submit";
    disabled?: boolean;
    onclick?: () => void;
  }

  let {
    name = "Name",
    color = "blue",
    ico,
    type = "button",
    disabled = false,
    onclick,
  }: Props = $props();

  const colorClasses: Record<string, string> = {
    green: "bg-green-600 active:bg-green-500",
    yellow: "bg-yellow-600 active:bg-yellow-500",
    red: "bg-red-600 active:bg-red-500",
    blue: "bg-blue-600 active:bg-blue-500",
  };

  const icoMap: Record<string, string> = {
    row: "arrow-right",
    cart: "cart",
    trash: "trash",
    plus: "plus",
  };
</script>

<button
  {type}
  {disabled}
  onclick={() => onclick?.()}
  class="group relative inline-flex items-center overflow-hidden rounded px-8 py-3 text-white focus:outline-none focus:ring {colorClasses[color]}"
  class:opacity-50={disabled}
  class:cursor-not-allowed={disabled}
  class:cursor-pointer={!disabled}
>
  {#if ico}
    <span class="absolute -start-full transition-all group-hover:start-4">
      <svg class="h-4 w-4">
        <use href="/assets/img/sprite.svg#{icoMap[ico]}" />
      </svg>
    </span>
  {/if}

  <span
    class="text-sm font-medium"
    class:transition-all={!!ico}
    class:group-hover:ms-3={!!ico}
    class:group-hover:-me-3={!!ico}
  >
    {name}
  </span>
</button>
