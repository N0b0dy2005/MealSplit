<template>
  <form @submit.prevent="saveUser">
    <div class="mb-4">
      <label class="block text-meal-gray text-sm font-bold mb-2" for="user-name">
        Name
      </label>
      <input
          id="user-name"
          v-model="userData.name"
          class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
          type="text"
          placeholder="Vor- und Nachname"
          required
      >
    </div>

    <div class="mb-4">
      <label class="block text-meal-gray text-sm font-bold mb-2" for="user-email">
        Email
      </label>
      <input
          id="user-email"
          v-model="userData.email"
          class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
          type="email"
          placeholder="email@beispiel.de"
          required
      >
    </div>

    <div class="mt-4 mb-6">
      <div class="flex items-center justify-center">
        <div class="w-20 h-20 rounded-full flex items-center justify-center text-white text-2xl font-bold mr-3"
             :style="{ backgroundColor: getRandomColor() }">
          {{ userData.name ? userData.name.charAt(0).toUpperCase() : '?' }}
        </div>
      </div>
      <p class="text-center text-meal-gray text-sm mt-2">
        Avatar wird automatisch generiert
      </p>
    </div>

    <div class="flex justify-end space-x-3 mt-6">
      <button
          type="button"
          @click="emit('cancel')"
          class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
      >
        Abbrechen
      </button>
      <button
          type="submit"
          class="px-4 py-2 bg-meal-primary hover:bg-meal-dark text-white rounded-lg transition-colors duration-200"
      >
        Speichern
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";

const props = defineProps({
  user: {
    type: Object,
    required: true
  }
});

const emit = defineEmits(['save', 'cancel']);

const userData = ref({ ...props.user });

// Farben für Benutzer-Avatare
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

// Aktualisiere die userData wenn sich die Props ändern
watch(() => props.user, (newUser) => {
  userData.value = { ...newUser };
});

function saveUser() {
  emit('save', { ...userData.value });
}

function getRandomColor() {
  return avatarColors[Math.floor(Math.random() * avatarColors.length)];
}
</script>