<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { Editor } from '@tiptap/core'
  import StarterKit from '@tiptap/starter-kit'
  import Placeholder from '@tiptap/extension-placeholder'
  import SvgIcon from './SvgIcon.svelte'
  import { createEventDispatcher } from 'svelte'

  export let modelValue: string = ''
  export let placeholder: string = ''
  export let id: string | undefined = undefined

  const dispatch = createEventDispatcher()

  let editor: Editor | null = null
  let editorElement: HTMLElement

  const editorActions = [
    { name: 'undo', method: 'undo', icon: 'undo', stroke: 'currentColor', activeCondition: {} },
    { name: 'redo', method: 'redo', icon: 'redo', stroke: 'currentColor', activeCondition: {} },
    { name: 'bold', method: 'toggleBold', icon: 'bold', activeCondition: { type: 'bold' } },
    { name: 'italic', method: 'toggleItalic', icon: 'italic', activeCondition: { type: 'italic' } },
    { name: 'strike', method: 'toggleStrike', icon: 'strike', activeCondition: { type: 'strike' } },
    { name: 'paragraph', method: 'toggleParagraph', icon: 'paragraph', activeCondition: { type: 'paragraph' } },
    { name: 'h1', method: 'toggleHeading', icon: 'h1', activeCondition: { type: 'heading', options: { level: 1 } } },
    { name: 'h2', method: 'toggleHeading', icon: 'h2', activeCondition: { type: 'heading', options: { level: 2 } } },
    { name: 'h3', method: 'toggleHeading', icon: 'h3', activeCondition: { type: 'heading', options: { level: 3 } } },
    { name: 'bulletlist', method: 'toggleBulletList', icon: 'bulletlist', activeCondition: { type: 'bulletList' } },
    { name: 'orderedList', method: 'toggleOrderedList', icon: 'orderedlist', activeCondition: { type: 'orderedList' } },
    { name: 'blockquote', method: 'toggleBlockquote', icon: 'blockquote', activeCondition: { type: 'blockquote' } }
  ]

  interface EditorOptions {
    level?: number
    [key: string]: unknown
  }

  function performEditorAction(method: string, options?: EditorOptions) {
    if (!editor) return
    const chain = editor.chain().focus()
    if (options) {
      ;(chain[method as keyof typeof chain] as (options: EditorOptions) => typeof chain)(options).run()
    } else {
      ;(chain[method as keyof typeof chain] as () => typeof chain)().run()
    }
  }

  function canPerformEditorAction(method: string): boolean {
    if (!editor) return false
    if (method === 'undo' || method === 'redo') {
      const canMethod = editor.can().chain().focus()[method as 'undo' | 'redo']
      return !canMethod().run()
    }
    return false
  }

  function isActive(type: string, options?: EditorOptions): boolean {
    if (!editor) return false
    return editor.isActive(type, options)
  }

  $: if (editor && modelValue !== editor.getHTML()) {
    const isSame = editor.getHTML() === modelValue
    if (!isSame) {
      editor.commands.setContent(modelValue, false)
    }
  }

  onMount(() => {
    editor = new Editor({
      element: editorElement,
      extensions: [
        StarterKit,
        Placeholder.configure({
          placeholder: placeholder
        })
      ],
      content: modelValue,
      onUpdate: ({ editor }) => {
        const html = editor.getHTML()
        modelValue = html
        dispatch('update:modelValue', html)
      }
    })
  })

  onDestroy(() => {
    if (editor) {
      editor.destroy()
    }
  })
</script>

{#if editor}
  <div class="editor">
    {#each editorActions as action}
      <button
        on:click={() => performEditorAction(action.method, action.activeCondition.options)}
        disabled={canPerformEditorAction(action.method)}
        class={action.stroke
          ? ''
          : isActive(action.activeCondition.type, action.activeCondition.options)
            ? 'is-active'
            : ''}
      >
        <SvgIcon name={action.icon} stroke={action.stroke || 'currentColor'} />
      </button>
    {/each}
  </div>
{/if}

<article class="mt-5" bind:this={editorElement} {id}></article>

<style>
  @reference "tailwindcss";

  :global(.tiptap p.is-editor-empty:first-child::before) {
    @apply text-gray-400;
    content: attr(data-placeholder);
    float: left;
    height: 0;
    pointer-events: none;
  }

  :global(.ProseMirror:focus) {
    outline: none;
  }

  :global(.editor button),
  :global(.editor input),
  :global(.editor select) {
    @apply m-[0.2rem] rounded-[0.3rem] bg-slate-200 px-[0.6rem] py-[0.2rem] text-black;
  }

  :global(.editor button[disabled]),
  :global(.editor input[disabled]),
  :global(.editor select[disabled]) {
    opacity: 0.3;
  }

  :global(.is-active) {
    @apply bg-black text-white;
  }

  :global(.tiptap > * + *) {
    margin-top: 0.75em;
  }

  :global(.tiptap ul) {
    @apply list-disc pl-10;
  }

  :global(.tiptap ol) {
    @apply list-decimal pl-10;
  }

  :global(.tiptap blockquote) {
    @apply border-x-2 border-solid border-gray-500 pl-4;
  }

  :global(.tiptap blockquote ul) {
    @apply list-disc pl-4;
  }

  :global(.tiptap blockquote ol) {
    @apply list-decimal pl-4;
  }

  :global(.tiptap hr) {
    @apply my-8 border-0 border-t-2 border-gray-200;
  }
</style>
