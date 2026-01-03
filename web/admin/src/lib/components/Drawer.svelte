<script lang="ts">
  import { onDestroy } from 'svelte'

  interface Props {
    isOpen?: boolean
    maxWidth?: string
    backgroundColor?: string
    onclose?: () => void
    children?: import('svelte').Snippet
  }

  let {
    isOpen = $bindable(false),
    maxWidth = '500px',
    backgroundColor = '#fafafa',
    onclose,
    children
  }: Props = $props()

  let isVisible = $state(false)
  let isTransitioning = $state(false)
  let drawerContent: HTMLElement | undefined = $state()

  $effect(() => {
    if (isOpen) {
      if (drawerContent) {
        drawerContent.scrollTop = 0
      }
      toggleBackgroundScrolling(true)
      isVisible = true
      // Force reflow for animation
      requestAnimationFrame(() => {
        requestAnimationFrame(() => {
          // Animation will trigger via class binding
        })
      })
    } else {
      toggleBackgroundScrolling(false)
      setTimeout(() => {
        isVisible = false
      }, 200)
    }
  })

  function toggleBackgroundScrolling(enable: boolean) {
    const body = document.querySelector('body')
    if (body) {
      body.style.overflow = enable ? 'hidden' : ''
    }
  }

  function closeDrawer(event: MouseEvent) {
    if (!isTransitioning && event.target === event.currentTarget) {
      isTransitioning = true
      setTimeout(() => {
        onclose?.()
        isTransitioning = false
      }, 200)
    }
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Escape' || event.key === 'Enter' || event.key === ' ') {
      event.preventDefault()
      closeDrawer(event as unknown as MouseEvent)
    }
  }

  onDestroy(() => {
    // document.removeEventListener('click', handleClickOutside);
    toggleBackgroundScrolling(false)
  })
</script>

{#if isVisible}
  <div class="drawer">
    <div
      class="overlay fixed inset-x-0 inset-y-0 z-50 w-full bg-black transition-opacity select-none {isOpen
        ? 'opacity-50'
        : 'opacity-0'}"
      style="transition-duration: 200ms"
      role="button"
      tabindex="0"
      aria-label="Close drawer"
      onclick={closeDrawer}
      onkeydown={handleKeydown}
    ></div>

    <div
      bind:this={drawerContent}
      id="drawer_content"
      class="content {isOpen ? 'translate-x-0' : 'translate-x-full'}"
      style="max-width: {maxWidth}; transition-duration: 200ms; background-color: {backgroundColor};"
    >
      {#if children}
        {@render children()}
      {/if}
    </div>
  </div>
{/if}

<style>
  @reference "tailwindcss";

  :global(.drawer .content) {
    @apply fixed inset-y-0 right-0 z-[999] flex h-full w-full flex-col overflow-auto bg-white p-6 shadow-2xl transition-transform;
  }
</style>
