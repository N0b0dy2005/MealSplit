<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md">
      <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-4">Mahlzeit löschen</h3>
      <p class="mb-6 text-meal-gray-dark">
        Bist du sicher, dass du die Mahlzeit <span class="font-bold">{{ meal.name }}</span> vom {{ formatDate(meal.date) }} löschen möchtest?
        Diese Aktion kann nicht rückgängig gemacht werden.
      </p>
      <div class="flex justify-end space-x-3">
        <button
            @click="$emit('close')"
            class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
        >
          Abbrechen
        </button>
        <button
            @click="$emit('delete')"
            class="px-4 py-2 bg-meal-error hover:bg-red-700 text-white rounded-lg transition-colors duration-200"
        >
          Löschen
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  meal: {
    type: Object,
    required: true
  }
});

defineEmits(['close', 'delete']);

function formatDate(dateString) {
  return new Date(dateString).toLocaleDateString('de-DE');
}
</script>