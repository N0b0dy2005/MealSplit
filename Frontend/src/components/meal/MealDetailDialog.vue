<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-2xl font-header font-bold text-meal-gray-dark">Details: {{ meal.name }}</h3>
        <button @click="$emit('close')" class="text-meal-gray hover:text-meal-error">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        <div>
          <p class="text-sm text-meal-gray mb-1">Datum</p>
          <p class="font-bold">{{ formatDate(meal.date) }}</p>
        </div>

        <div>
          <p class="text-sm text-meal-gray mb-1">Gesamtbetrag</p>
          <p class="font-bold text-meal-primary">{{ formatCurrency(meal.totalAmount) }}</p>
        </div>

        <div>
          <p class="text-sm text-meal-gray mb-1">Bezahlt von</p>
          <div class="flex items-center">
            <div class="w-6 h-6 rounded-full flex items-center justify-center text-white font-bold mr-2"
                 :style="{ backgroundColor: getUserColor(meal.userId) }">
              {{ getUserInitial(meal.userId) }}
            </div>
            {{ getUserName(meal.userId) }}
          </div>
        </div>

        <div>
          <p class="text-sm text-meal-gray mb-1">Anzahl Teilnehmer</p>
          <p class="font-bold">{{ meal.participants.length }} Personen</p>
        </div>
      </div>

      <div v-if="meal.description" class="mb-6">
        <p class="text-sm text-meal-gray mb-1">Beschreibung</p>
        <p class="bg-meal-light p-3 rounded">{{ meal.description }}</p>
      </div>

      <!-- Products display -->
      <div v-if="meal.productsData && meal.productsData.length > 0" class="mb-6">
        <p class="text-sm text-meal-gray mb-2">Produkte</p>
        <div class="overflow-hidden border rounded-lg">
          <table class="min-w-full divide-y divide-meal-gray-light">
            <thead class="bg-meal-light">
            <tr>
              <th scope="col" class="px-4 py-2 text-left text-xs font-medium text-meal-gray-dark uppercase tracking-wider">Produkt</th>
              <th scope="col" class="px-4 py-2 text-left text-xs font-medium text-meal-gray-dark uppercase tracking-wider">Preis</th>
              <th scope="col" class="px-4 py-2 text-left text-xs font-medium text-meal-gray-dark uppercase tracking-wider">Für</th>
            </tr>
            </thead>
            <tbody class="bg-white divide-y divide-meal-gray-light">
            <tr v-for="(product, index) in meal.productsData" :key="index">
              <td class="px-4 py-2 whitespace-nowrap text-sm font-medium text-meal-gray-dark">
                {{ product.name }}
              </td>
              <td class="px-4 py-2 whitespace-nowrap text-sm text-meal-gray-dark">
                {{ formatCurrency(product.price) }}
              </td>
              <td class="px-4 py-2 text-sm text-meal-gray-dark">
                <span v-if="!product.isSpecific">Alle Teilnehmer</span>
                <div v-else class="flex flex-wrap gap-1">
                    <span v-for="participantId in product.specificParticipants" :key="participantId"
                          class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-meal-light">
                      {{ getUserName(participantId) }}
                    </span>
                </div>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div class="mb-6">
        <p class="text-sm text-meal-gray mb-2">Kosten pro Person</p>
        <div class="bg-meal-light p-4 rounded-lg">
          <div class="grid grid-cols-1 gap-3">
            <div v-for="participantId in meal.participants" :key="participantId"
                 class="flex items-center justify-between pb-2 border-b border-white">
              <div class="flex items-center">
                <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                     :style="{ backgroundColor: getUserColor(participantId) }">
                  {{ getUserInitial(participantId) }}
                </div>
                {{ getUserName(participantId) }}
              </div>
              <div class="font-bold">
                {{ formatCurrency(calculateParticipantCost(meal, participantId)) }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="meal.userId" class="mb-6">
        <p class="text-sm text-meal-gray mb-2">Schulden gegenüber {{ getUserName(meal.userId) }}</p>
        <div class="space-y-2">
          <div v-for="participantId in meal.participants.filter(id => id !== meal.userId)" :key="participantId"
               class="flex items-center justify-between bg-meal-light p-3 rounded">
            <div class="flex items-center">
              <div class="w-6 h-6 rounded-full flex items-center justify-center text-white font-bold mr-2"
                   :style="{ backgroundColor: getUserColor(participantId) }">
                {{ getUserInitial(participantId) }}
              </div>
              {{ getUserName(participantId) }}
            </div>
            <div class="font-bold text-meal-error">
              {{ formatCurrency(calculateParticipantCost(meal, participantId)) }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  meal: {
    type: Object,
    required: true
  },
  users: {
    type: Array,
    required: true
  }
});

defineEmits(['close']);

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

function formatDate(dateString) {
  return new Date(dateString).toLocaleDateString('de-DE');
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

// Calculate costs per participant based on products
function calculateParticipantCost(meal, participantId) {
  if (!meal.productsData || meal.productsData.length === 0) {
    // If no products, divide total amount equally
    return meal.totalAmount / meal.participants.length;
  }

  let totalCost = 0;

  for (const product of meal.productsData) {
    if (!product.isSpecific) {
      // Product for all participants - divide equally
      totalCost += parseFloat(product.price) / meal.participants.length;
    } else if (product.specificParticipants.includes(participantId)) {
      // Product only for specific participants - divide among them
      totalCost += parseFloat(product.price) / product.specificParticipants.length;
    }
  }

  return totalCost;
}
</script>