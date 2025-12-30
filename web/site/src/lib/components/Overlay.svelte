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
    class="fixed top-0 left-0 w-full h-full bg-black/50 backdrop-blur-sm flex items-center justify-center z-40 animate-fade-in"
    role="dialog"
    aria-modal="true"
    aria-labelledby="overlay-title"
  >
    {#if error}
      <div
        role="alert"
        class="rounded border-s-4 border-red-500 bg-red-50 p-4 max-w-md mx-4 animate-slide-up"
      >
        <div class="flex items-center gap-2 text-red-800">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            class="h-5 w-5"
          >
            <path
              fill-rule="evenodd"
              d="M9.401 3.003c1.155-2 4.043-2 5.197 0l7.355 12.748c1.154 2-.29 4.5-2.599 4.5H4.645c-2.309 0-3.752-2.5-2.598-4.5L9.4 3.003zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z"
              clip-rule="evenodd"
            />
          </svg>
          <strong id="overlay-title" class="block font-medium">{error}</strong>
        </div>
        <p class="mt-2 text-sm text-red-700">
          Try again in a little while, if the error doesn't go away - contact the
          administrator.
        </p>
        {#if onClose}
          <button
            onclick={onClose}
            class="mt-4 disabled:opacity-25 disabled:bg-gray-400 cursor-pointer block rounded bg-gray-700 px-5 py-3 text-sm text-gray-100 transition hover:bg-gray-600"
          >
            Close
          </button>
        {/if}
      </div>
    {:else}
      <div class="flex flex-col items-center gap-4 animate-fade-in">
        <!-- Spinner -->
        <div class="relative w-20 h-20">
          <div class="absolute top-0 left-0 w-full h-full border-4 border-gray-200 rounded-full"></div>
          <div class="absolute top-0 left-0 w-full h-full border-4 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
        </div>
        <!-- Pulsing dots -->
        <div class="flex gap-2">
          <div class="w-3 h-3 bg-blue-600 rounded-full animate-pulse" style="animation-delay: 0s"></div>
          <div class="w-3 h-3 bg-blue-600 rounded-full animate-pulse" style="animation-delay: 0.2s"></div>
          <div class="w-3 h-3 bg-blue-600 rounded-full animate-pulse" style="animation-delay: 0.4s"></div>
        </div>
      </div>
    {/if}
  </div>
{/if}

<style>
  @keyframes fade-in {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  @keyframes slide-up {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .animate-fade-in {
    animation: fade-in 0.3s ease-out;
  }

  .animate-slide-up {
    animation: slide-up 0.4s ease-out;
  }
</style>
