<template>
  <button
    @click="$emit('click')"
    :class="[
      'py-3 px-4 sm:py-4 sm:px-6 font-bold text-center flex-1 text-sm sm:text-base',
      active
        ? 'text-meal-primary border-b-2 border-meal-primary'
        : 'text-meal-gray hover:text-meal-gray-dark'
    ]"
  >
    <span class="flex items-center justify-center">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-5 sm:w-5 mr-1 sm:mr-2" viewBox="0 0 20 20" fill="currentColor" v-html="icon"></svg>
      {{ text }}
    </span>
  </button>
</template>

<script setup lang="ts">
defineProps<{
  active: boolean;
  text: string;
  icon: string;
}>();

defineEmits<{
  (e: 'click'): void;
}>();
</script> 