<script lang="ts">
  import { settingsStore } from "$lib/stores/settings";
  import { socialUrl } from "$lib/utils/socialUrl";
  import { handleNavigation } from "$lib/utils/navigation";

  let settings = $derived($settingsStore);
  let pages = $derived(settings?.pages || []);
  let socials = $derived(settings?.socials || {});
  let footerPages = $derived(pages.filter((p) => p.position === "footer"));
</script>

<footer class="bg-black text-white border-t-4 border-yellow-300">
  <div class="mx-auto max-w-screen-xl px-4 py-10 sm:px-6 lg:px-8">
    <div class="flex flex-col lg:flex-row justify-between items-start lg:items-center gap-8 mb-8">
      {#if footerPages.length > 0}
        <nav class="flex flex-wrap gap-4">
          {#each footerPages as page}
            <a
              href="/{page.slug}"
              onclick={(e) => handleNavigation(e, `/${page.slug}`)}
              class="inline-block border-2 border-transparent hover:border-yellow-300 px-4 py-2 font-black uppercase text-sm transition-all duration-200 hover:bg-yellow-300 hover:text-black cursor-pointer"
            >
              {page.name}
            </a>
          {/each}
        </nav>
      {/if}
      
      {#if Object.keys(socials).length > 0}
        <div class="flex items-center gap-3">
          <span class="text-xs font-black uppercase tracking-wider text-yellow-300 mr-2">Follow:</span>
          <ul class="flex items-center gap-2">
            {#each Object.entries(socials) as [key, value], i}
              {#if value}
                <li>
                  <a
                    href="{socialUrl[key]}{value}"
                    rel="noreferrer"
                    target="_blank"
                    class="block border-4 border-white bg-white text-black p-1 cursor-pointer hover:shadow-[6px_6px_0px_0px_rgba(255,235,59,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1"
                    aria-label={key}
                  >
                    <svg class="h-7 w-7">
                      <use href="/assets/img/socials.svg#{key}" />
                    </svg>
                  </a>
                </li>
              {/if}
            {/each}
          </ul>
        </div>
      {/if}
    </div>
    
    <div class="border-t-4 border-yellow-300 pt-6">
      <div class="flex flex-col sm:flex-row justify-between items-center gap-4">
        <p class="text-xs font-black uppercase tracking-widest text-yellow-300">
          © {new Date().getFullYear()} All Rights Reserved
        </p>
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://github.com/shurco/litecart"
          class="text-xs font-black uppercase tracking-wider text-yellow-300 hover:text-white transition-colors duration-200 cursor-pointer"
        >
          Powered by Litecart
        </a>
      </div>
    </div>
  </div>
</footer>
