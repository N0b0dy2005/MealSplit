<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
    <div class="bg-white rounded-xl shadow-xl p-4 sm:p-6 w-full max-w-md">
      <h3 class="text-lg sm:text-xl font-header font-bold text-meal-gray-dark mb-3 sm:mb-4">{{ title }}</h3>
      <p class="mb-4 sm:mb-6 text-meal-gray-dark text-sm sm:text-base">
        {{ message }}
      </p>
      
      <slot></slot>
      
      <div class="flex justify-end space-x-3 mt-4">
        <button
          @click="cancelAction"
          class="px-3 py-2 sm:px-4 sm:py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark text-sm sm:text-base hover:bg-meal-gray-light transition-colors"
        >
          Abbrechen
        </button>
        <button
          @click="confirmAction"
          class="px-3 py-2 sm:px-4 sm:py-2 bg-meal-error hover:bg-red-700 text-white text-sm sm:text-base rounded-lg transition-colors"
        >
          Bestätigen
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string;
  message: string;
  confirmAction: () => void;
  cancelAction: () => void;
}>();
</script> 