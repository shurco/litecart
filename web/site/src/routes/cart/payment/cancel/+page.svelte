<script lang="ts">
  import { onMount } from "svelte";
  import { cartStore } from "$lib/stores/cart";
  import { goto } from "$app/navigation";

  onMount(() => {
    cartStore.clear();
  });
</script>

<div class="grid h-screen px-4 bg-white place-content-center">
  <h1 class="text-2xl font-bold text-red-600">Payment Cancelled</h1>
  <a href="/cart" onclick={(e) => { e.preventDefault(); goto("/cart"); }} class="text-blue-500 mt-4 cursor-pointer"
    >Return to cart</a
  >
</div>
