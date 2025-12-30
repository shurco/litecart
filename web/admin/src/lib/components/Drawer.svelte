<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { createEventDispatcher } from 'svelte';

  export let isOpen: boolean = false;
  export let maxWidth: string = "500px";
  export let backgroundColor: string = "#fafafa";

  const dispatch = createEventDispatcher();
  let isVisible = false;
  let isTransitioning = false;
  let drawerContent: HTMLElement;

  $: if (isOpen) {
    if (drawerContent) {
      drawerContent.scrollTop = 0;
    }
    toggleBackgroundScrolling(true);
    isVisible = true;
    // Force reflow for animation
    requestAnimationFrame(() => {
      requestAnimationFrame(() => {
        // Animation will trigger via class binding
      });
    });
  } else {
    toggleBackgroundScrolling(false);
    setTimeout(() => {
      isVisible = false;
    }, 200);
  }

  function toggleBackgroundScrolling(enable: boolean) {
    const body = document.querySelector("body");
    if (body) {
      body.style.overflow = enable ? "hidden" : "";
    }
  }

  function closeDrawer(event: MouseEvent) {
    if (!isTransitioning && event.target === event.currentTarget) {
      isTransitioning = true;
      setTimeout(() => {
      dispatch('close');
        isTransitioning = false;
      }, 200);
    }
  }

  function handleClickOutside(event: MouseEvent) {
    // Only close if clicking on the overlay, not on drawer content
    const target = event.target as HTMLElement;
    if (drawerContent && target) {
      // Check if click was on the overlay (has overlay class)
      const overlay = target.closest('.overlay');
      if (overlay && !drawerContent.contains(target)) {
      if (!isTransitioning) {
          isTransitioning = true;
          setTimeout(() => {
        dispatch('close');
            isTransitioning = false;
          }, 200);
        }
      }
    }
  }
  
  function handleClose() {
    if (!isTransitioning) {
      isTransitioning = true;
      // Don't modify isOpen directly, let parent handle it
      setTimeout(() => {
        dispatch('close');
        isTransitioning = false;
      }, 200);
    }
  }

  onMount(() => {
    // Remove handleClickOutside as we use closeDrawer on overlay directly
    // document.addEventListener('click', handleClickOutside);
  });

  onDestroy(() => {
    // document.removeEventListener('click', handleClickOutside);
    toggleBackgroundScrolling(false);
  });
</script>

{#if isVisible}
  <div class="drawer">
    <div
      class="overlay fixed inset-x-0 inset-y-0 z-50 w-full select-none bg-black transition-opacity {isOpen ? 'opacity-50' : 'opacity-0'}"
      style="transition-duration: 200ms"
      role="button"
      tabindex="0"
      aria-label="Close drawer"
      on:click={closeDrawer}
      on:keydown={(e) => {
        if (e.key === 'Escape' || e.key === 'Enter' || e.key === ' ') {
          e.preventDefault();
          closeDrawer(e);
        }
      }}
    ></div>

    <div
      bind:this={drawerContent}
      id="drawer_content"
      class="content {isOpen ? 'translate-x-0' : 'translate-x-full'}"
      style="max-width: {maxWidth}; transition-duration: 200ms; background-color: {backgroundColor};"
    >
      <slot />
    </div>
  </div>
{/if}

<style>
  @reference "tailwindcss";
  
  :global(.drawer .content) {
    @apply fixed inset-y-0 right-0 z-[999] flex h-full w-full flex-col overflow-auto bg-white p-6 shadow-2xl transition-transform;
  }
</style>