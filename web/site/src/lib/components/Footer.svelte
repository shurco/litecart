<script lang="ts">
  import { settingsStore } from "$lib/stores/settings";

  let settings = $derived($settingsStore);
  let pages = $derived(settings?.pages || []);
</script>

<footer class="bg-white">
  <div class="mx-auto max-w-screen-xl px-4 pb-8 sm:px-6 lg:px-8">
    <div class="border-t border-gray-100 pt-8 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
      <ul class="flex flex-wrap justify-center gap-4 text-xs sm:justify-start">
        {#each pages.filter((p) => p.position === "footer") as page}
          <li>
            <a
              href="/{page.slug}"
              class="text-blue-500 transition hover:opacity-75 cursor-pointer"
            >
              {page.name}
            </a>
          </li>
        {/each}
      </ul>

      <a
        target="_blank"
        rel="noopener noreferrer"
        href="https://github.com/shurco/litecart"
        class="text-gray-500 text-xs cursor-pointer sm:ml-auto"
      >
        Powered by litecart
      </a>
    </div>
  </div>
</footer>
