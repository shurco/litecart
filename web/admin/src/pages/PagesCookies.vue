<template>
  <div class="grid grid-cols-1 pt-4">

    <div class="flex pt-4">
      <div class="flex-none">
        <FormButton type="submit" name="Save" color="green" class="mr-3" @click="updatePage" />
      </div>
      <div class="grow"></div>
    </div>

  </div>
</template>

<script setup>
import FormButton from '@/components/form/Button.vue'
import FormTextarea from '@/components/form/Textarea.vue'

const updatePage = () => {
  console.log("save")
}
</script>
