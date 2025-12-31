<script lang="ts">
  import { onMount } from "svelte";
  import { cartStore } from "$lib/stores/cart";
  import { handleNavigation } from "$lib/utils/navigation";

  onMount(() => {
    cartStore.clear();
  });
</script>

<div class="min-h-screen bg-white flex items-center justify-center px-4">
  <div class="max-w-2xl w-full">
    <div class="brutal-card p-12 text-center bg-red-300">
      <h1 class="text-4xl sm:text-5xl font-black uppercase tracking-tighter text-black mb-8">
        PAYMENT CANCELLED
      </h1>
      <p class="text-xl font-bold uppercase tracking-wider text-black mb-8">
        Your payment has been cancelled. You can return to your cart to try again.
      </p>
      <a 
        href="/cart" 
        onclick={(e) => handleNavigation(e, "/cart")} 
        class="inline-block border-4 border-black bg-yellow-300 text-black px-8 py-4 font-black text-lg uppercase tracking-wider hover:shadow-[12px_12px_0px_0px_rgba(0,0,0,1)] transition-all duration-200 hover:-translate-x-1 hover:-translate-y-1 cursor-pointer"
      >
        RETURN TO CART
      </a>
    </div>
  </div>
</div>
