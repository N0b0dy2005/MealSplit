<template>
  <div class="bg-white rounded-xl shadow-meal p-6 mb-8">
    <h2 class="text-2xl font-header font-bold text-meal-gray-dark mb-6">Alle Mahlzeiten</h2>

    <div v-if="meals.length === 0" class="text-center py-8 text-meal-gray">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto mb-4 text-meal-gray-light" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
      </svg>
      <p class="text-xl mb-2">Keine Mahlzeiten gefunden</p>
      <p>{{ isFiltered ? 'Passe die Filter an, um Ergebnisse zu sehen.' : 'Erfasse deine erste gemeinsame Mahlzeit!' }}</p>
    </div>

    <div v-else class="space-y-4">
      <div v-for="meal in meals" :key="meal.id" class="border border-meal-gray-light rounded-lg overflow-hidden hover:shadow-md transition-shadow duration-200">
        <div class="flex flex-col md:flex-row">
          <!-- Datum Spalte -->
          <div class="bg-meal-light p-4 md:w-24 flex flex-col items-center justify-center text-center">
            <div class="text-2xl font-bold text-meal-primary">{{ new Date(meal.date).getDate() }}</div>
            <div class="text-sm text-meal-gray-dark">{{ getMonthName(new Date(meal.date).getMonth()) }}</div>
            <div class="text-sm text-meal-gray">{{ new Date(meal.date).getFullYear() }}</div>
          </div>

          <!-- Hauptinformationen -->
          <div class="p-4 flex-grow">
            <div class="flex justify-between items-start mb-2">
              <h3 class="text-xl font-header font-bold text-meal-gray-dark">{{ meal.name }}</h3>
              <span class="font-bold text-meal-primary">{{ formatCurrency(meal.totalAmount) }}</span>
            </div>

            <p v-if="meal.description" class="text-meal-gray mb-3 line-clamp-2">{{ meal.description }}</p>

            <!-- Produkte anzeigen, wenn vorhanden -->
            <div v-if="meal.productsData && meal.productsData.length > 0" class="mt-2 mb-3">
              <div class="text-sm text-meal-gray font-bold mb-1">Produkte:</div>
              <div class="flex flex-wrap gap-1">
                <span v-for="(product, idx) in meal.productsData" :key="idx"
                      class="bg-meal-light px-2 py-1 text-xs rounded-lg">
                  {{ product.name }} ({{ formatCurrency(product.price) }})
                </span>
              </div>
            </div>

            <div class="flex items-center text-sm text-meal-gray mb-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path d="M8 9a3 3 0 100-6 3 3 0 000 6zM8 11a6 6 0 016 6H2a6 6 0 016-6zM16 7a1 1 0 10-2 0v1h-1a1 1 0 100 2h1v1a1 1 0 102 0v-1h1a1 1 0 100-2h-1V7z" />
              </svg>
              <span>Bezahlt von {{ getUserName(meal.userId) }}</span>
            </div>

            <div class="flex flex-wrap gap-1">
              <div v-for="participantId in meal.participants" :key="participantId"
                   class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-meal-light">
                <div class="w-4 h-4 rounded-full flex items-center justify-center text-white font-bold mr-1"
                     :style="{ backgroundColor: getUserColor(participantId) }">
                  {{ getUserInitial(participantId) }}
                </div>
                {{ getUserName(participantId) }}
              </div>
            </div>
          </div>

          <!-- Aktionen -->
          <div class="bg-meal-light p-4 flex md:flex-col justify-center space-x-4 md:space-x-0 md:space-y-4">
            <button @click="$emit('select-meal', meal)" class="text-meal-primary hover:text-meal-dark" title="Details anzeigen">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
            </button>
            <button @click="$emit('edit-meal', meal)" class="text-meal-accent hover:text-opacity-80" title="Bearbeiten">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
            </button>
            <button @click="$emit('confirm-delete', meal)" class="text-meal-error hover:text-red-700" title="Löschen">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  meals: {
    type: Array,
    required: true
  },
  users: {
    type: Array,
    required: true
  },
  isFiltered: {
    type: Boolean,
    default: false
  }
});

defineEmits(['select-meal', 'edit-meal', 'confirm-delete']);

// Avatar colors
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

function formatCurrency(value) {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR'
  }).format(value);
}

function getUserName(userId) {
  const user = props.users.find(u => u.id === userId);
  return user ? user.name : 'Unbekannt';
}

function getUserInitial(userId) {
  const user = props.users.find(u => u.id === userId);
  return user ? user.name.charAt(0).toUpperCase() : '?';
}

function getUserColor(userId) {
  return avatarColors[userId % avatarColors.length];
}

function getMonthName(monthIndex) {
  const months = ['Jan', 'Feb', 'Mär', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Okt', 'Nov', 'Dez'];
  return months[monthIndex];
}
</script>