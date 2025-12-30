<script lang="ts">
  import Header from "$lib/components/Header.svelte";
  import Footer from "$lib/components/Footer.svelte";
  import Overlay from "$lib/components/Overlay.svelte";
  import { settingsStore } from "$lib/stores/settings";
  import { apiGet } from "$lib/utils/api";
  import { updateSEOTags } from "$lib/utils/seo";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";

  interface Props {
    children: import("svelte").Snippet;
  }

  let { children }: Props = $props();
  let showOverlay = $state(false);
  let error = $state<string | undefined>(undefined);

  onMount(async () => {
    if (!browser) return;

    let cached = settingsStore.loadFromCache();
    if (!cached) {
      showOverlay = true;
      const res = await apiGet("/api/settings");
      if (res.success && res.result) {
        settingsStore.set(res.result);
        settingsStore.saveToCache(res.result);

        // Update meta tags
        if (res.result.main?.site_name) {
          updateSEOTags({ title: res.result.main.site_name });
        }
      } else {
        error = res.message || "Failed to load settings";
      }
      showOverlay = false;
    } else {
      settingsStore.set(cached);
    }
  });

  function closeOverlay() {
    showOverlay = false;
    error = undefined;
  }
</script>

<div>
  <Header />
  <main>
    {@render children()}
  </main>
  <Footer />
  <Overlay show={showOverlay} {error} onClose={closeOverlay} />
</div>
