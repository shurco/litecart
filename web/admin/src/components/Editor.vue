<template>
  <div class="editor" v-if="editor">
    <button v-for="action in editorActions" :key="action.name" @click="performEditorAction(action.method, action.activeCondition.options)"
      :disabled="canPerformEditorAction(action.method)" :class="{ 'is-active': !action.stroke && editor.isActive(action.activeCondition.type, action.activeCondition.options) }">
      <SvgIcon :name="action.icon" :stroke="action.stroke" />
    </button>
  </div>

  <article class="mt-5">
    <EditorContent :editor="editor" />
  </article>
</template>

<script setup>
import { onBeforeUnmount, onMounted, ref, watch } from "vue";
import { EditorContent, Editor } from "@tiptap/vue-3";
import Placeholder from '@tiptap/extension-placeholder'
import StarterKit from "@tiptap/starter-kit";

const emits = defineEmits(["update:modelValue"]);
const editor = ref();
const props = defineProps({
  modelValue: {
    type: String,
    required: true,
    default: "",
  },
  placeholder: String,
});

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
  { name: 'blockquote', method: 'toggleBlockquote', icon: 'blockquote', activeCondition: { type: 'blockquote' } },
]

const performEditorAction = (method, options) => {
  editor.value.chain().focus()[method](options).run();
}
const canPerformEditorAction = (method) => {
  if (method === "undo" || method === "redo") {
    return !editor.value.can().chain().focus()[method]().run();
  }
  return null
}

onMounted(() => {
  editor.value = new Editor({
    extensions: [StarterKit, Placeholder.configure({
      placeholder: props.placeholder,
    })],
    content: props.modelValue,
    onUpdate: () => {
      emits("update:modelValue", editor.value.getHTML());
    },
  });
});

onBeforeUnmount(() => {
  editor.value.destroy();
});

watch(
  () => props.modelValue,
  (value) => {
    const isSame = editor.value.getHTML() === value;
    if (isSame) {
      return;
    }
    editor.value.commands.setContent(value, false);
  },
);
</script>

<style lang="scss">
.tiptap p.is-editor-empty:first-child::before {
  color: #adb5bd;
  content: attr(data-placeholder);
  float: left;
  height: 0;
  pointer-events: none;
}

.ProseMirror:focus {
  outline: none;
}

button,
input,
select {
  @apply m-[0.2rem] rounded-[0.3rem] bg-slate-200 px-[0.6rem] py-[0.2rem] text-black;
}

button[disabled],
input[disabled],
select[disabled] {
  opacity: 0.3;
}

.is-active {
  background: black;
  color: #fff;
}

.tiptap {
  >*+* {
    margin-top: 0.75em;
  }

  ul {
    @apply list-disc pl-10;
  }

  ol {
    @apply list-decimal pl-10;
  }

  blockquote {
    @apply border-x-2 border-solid border-gray-500 pl-4;

    ul {
      @apply list-disc pl-4;
    }

    ol {
      @apply list-decimal pl-4;
    }
  }

  hr {
    border: none;
    border-top: 2px solid rgba(#0d0d0d, 0.1);
    margin: 2rem 0;
  }
}
</style>
