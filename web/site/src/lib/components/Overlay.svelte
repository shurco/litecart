<script lang="ts">
  interface Props {
    show: boolean;
    error?: string;
    onClose?: () => void;
  }

  let { show, error, onClose }: Props = $props();
</script>

{#if show}
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/80 flex items-center justify-center z-50"
    role="dialog"
    aria-modal="true"
    aria-labelledby="overlay-title"
  >
    {#if error}
      <div
        role="alert"
        class="border-4 border-red-500 bg-white p-8 max-w-md mx-4"
      >
        <div class="flex items-center gap-4 mb-4">
          <div class="border-4 border-red-500 bg-red-500 text-white p-3">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              class="h-8 w-8"
            >
              <path
                fill-rule="evenodd"
                d="M9.401 3.003c1.155-2 4.043-2 5.197 0l7.355 12.748c1.154 2-.29 4.5-2.599 4.5H4.645c-2.309 0-3.752-2.5-2.598-4.5L9.4 3.003zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <strong id="overlay-title" class="block font-black text-xl uppercase tracking-wider text-black">
            {error}
          </strong>
        </div>
        <p class="mb-6 text-sm font-bold text-black uppercase tracking-wide">
          Try again in a little while, if the error doesn't go away - contact the administrator.
        </p>
        {#if onClose}
          <button
            onclick={onClose}
            class="border-4 border-black bg-red-500 text-white px-6 py-3 font-black text-sm uppercase tracking-wider hover:shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer"
          >
            CLOSE
          </button>
        {/if}
      </div>
    {:else}
      <div class="flex flex-col items-center gap-6">
        <!-- Brutalist Spinner -->
        <div class="relative w-24 h-24 border-4 border-yellow-300 bg-white">
          <div class="absolute inset-0 flex items-center justify-center">
            <div class="w-16 h-16 border-4 border-black border-t-transparent animate-spin"></div>
          </div>
        </div>
        <!-- Loading Text -->
        <div class="border-4 border-black bg-yellow-300 px-8 py-4">
          <p class="text-xl font-black uppercase tracking-wider text-black">
            LOADING...
          </p>
        </div>
      </div>
    {/if}
  </div>
{/if}
