<template>
  <div v-if="show" class="fixed inset-0 z-50 overflow-auto bg-black bg-opacity-50 flex items-center justify-center p-4">
    <div class="bg-white rounded-lg shadow-xl w-full max-w-md md:max-w-lg lg:max-w-xl relative">
      <div class="flex justify-between items-center p-4 border-b">
        <h3 class="text-lg font-bold">{{ title }}</h3>
        <button 
          @click="close" 
          class="text-meal-gray hover:text-meal-dark transition-colors"
          aria-label="Schließen">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <div class="p-4">
        <slot></slot>
      </div>
      <div v-if="$slots.footer" class="p-4 border-t">
        <slot name="footer"></slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

// Props
const props = defineProps<{
  title: string;
  show: boolean;
}>();

// Emits
const emit = defineEmits<{
  (e: 'close'): void;
}>();

// Methoden
function close() {
  emit('close');
}

// Document body overflow handling to prevent scrolling when modal is open
watch(() => props.show, (value) => {
  if (value) {
    document.body.style.overflow = 'hidden';
  } else {
    document.body.style.overflow = '';
  }
});
</script> 