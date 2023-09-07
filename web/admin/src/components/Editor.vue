<template>
  <div v-if="editor">
    <button @click="editor.chain().focus().undo().run()" :disabled="!editor.can().chain().focus().undo().run()">
      <SvgIcon name="undo" />
    </button>
    <button @click="editor.chain().focus().redo().run()" :disabled="!editor.can().chain().focus().redo().run()">
      <SvgIcon name="redo" />
    </button>

    <button @click="editor.chain().focus().toggleBold().run()" :disabled="!editor.can().chain().focus().toggleBold().run()" :class="{ 'is-active': editor.isActive('bold') }">
      <SvgIcon name="bold" />
    </button>
    <button @click="editor.chain().focus().toggleItalic().run()" :disabled="!editor.can().chain().focus().toggleItalic().run()" :class="{ 'is-active': editor.isActive('italic') }">
      <SvgIcon name="italic" />
    </button>
    <button @click="editor.chain().focus().toggleStrike().run()" :disabled="!editor.can().chain().focus().toggleStrike().run()" :class="{ 'is-active': editor.isActive('strike') }">
      <SvgIcon name="strike" />
    </button>

    <button @click="editor.chain().focus().setParagraph().run()" :class="{ 'is-active': editor.isActive('paragraph') }">
      <SvgIcon name="paragraph" />
    </button>
    <button @click="editor.chain().focus().toggleHeading({ level: 1 }).run()" :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }">
      <SvgIcon name="h1" />
    </button>
    <button @click="editor.chain().focus().toggleHeading({ level: 2 }).run()" :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }">
      <SvgIcon name="h2" />
    </button>
    <button @click="editor.chain().focus().toggleHeading({ level: 3 }).run()" :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }">
      <SvgIcon name="h3" />
    </button>
    <button @click="editor.chain().focus().toggleBulletList().run()" :class="{ 'is-active': editor.isActive('bulletList') }">
      <SvgIcon name="bulletlist" />
    </button>
    <button @click="editor.chain().focus().toggleOrderedList().run()" :class="{ 'is-active': editor.isActive('orderedList') }">
      <SvgIcon name="orderedlist" />
    </button>

    <button @click="editor.chain().focus().toggleBlockquote().run()" :class="{ 'is-active': editor.isActive('blockquote') }">
      <SvgIcon name="blockquote" />
    </button>
  </div>

  <article class="mt-5">
    <EditorContent :editor="editor" />
  </article>
</template>

<script setup>
import { onBeforeUnmount, onMounted, ref, watch } from "vue";
import { EditorContent, Editor } from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import SvgIcon from "svg-icon";

const emits = defineEmits(["update:modelValue"]);
const editor = ref();
const props = defineProps({
  modelValue: {
    type: String,
    required: true,
    default: "",
  },
});

onMounted(() => {
  editor.value = new Editor({
    extensions: [StarterKit],
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
