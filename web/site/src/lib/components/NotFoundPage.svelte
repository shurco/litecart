<script lang="ts">
  import { handleNavigation } from "$lib/utils/navigation";
  import { isBrowser } from "$lib/utils/browser";
  import { onDestroy } from "svelte";

  // Add error-page class to body synchronously before render
  $effect.pre(() => {
    if (isBrowser()) {
      document.body.classList.add("error-page");
    }
  });

  // Cleanup on component destroy
  onDestroy(() => {
    if (isBrowser()) {
      document.body.classList.remove("error-page");
    }
  });
</script>

<style>
  :global(body.error-page header),
  :global(body.error-page footer),
  :global(body.error-page [role="banner"]),
  :global(body.error-page [role="contentinfo"]) {
    display: none !important;
  }
</style>

<div class="min-h-screen bg-white flex items-center justify-center px-4 py-12">
  <div class="max-w-3xl w-full text-center">
    <div class="brutal-card p-12 bg-red-300 mb-8">
      <h1 class="text-8xl sm:text-9xl font-black uppercase tracking-tighter text-black mb-4">
        404
      </h1>
      <p class="text-3xl font-black uppercase tracking-wider text-black">
        NOT FOUND
      </p>
    </div>

    <div class="brutal-card p-8 mb-8">
      <p class="text-xl font-bold uppercase tracking-wide text-black mb-8">
        The page you're looking for doesn't exist or has been moved.
      </p>

      <div class="flex justify-center">
        <a
          href="/"
          onclick={(e) => handleNavigation(e, "/")}
          class="inline-block border-4 border-black bg-yellow-300 text-black px-8 py-4 font-black text-lg uppercase tracking-wider hover:shadow-[12px_12px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer"
        >
          GO TO HOME
        </a>
      </div>
    </div>

    <div class="brutal-card p-6 bg-white">
      <p class="text-sm font-bold text-black uppercase tracking-wide">
        Try checking the URL for typos, or return to the homepage to browse our products.
      </p>
    </div>
  </div>
</div>
