<script lang="ts">
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";

  $effect(() => {
    if ($page.status === 404) {
      // Handle 404
    }
  });
</script>

<div class="grid h-screen px-4 bg-white place-content-center">
  <h1 class="tracking-widest text-gray-500">404 | Not Found</h1>
  <a href="/" onclick={(e) => { e.preventDefault(); goto("/"); }} class="text-blue-500 mt-4 cursor-pointer"
    >Go to home</a
  >
</div>
