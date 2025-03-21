<template>
  <div class="bg-white rounded-xl shadow-meal p-6 mb-8">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-header font-bold text-meal-gray-dark">Neue Mahlzeit erfassen</h2>
      <button
          @click="$emit('toggle-form')"
          class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-4 rounded-lg transition-colors duration-200"
      >
        <span class="flex items-center">
          <svg v-if="!showForm" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
          </svg>
          <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M5 10a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1z" clip-rule="evenodd" />
          </svg>
          {{ showForm ? 'Formular ausblenden' : 'Neue Mahlzeit' }}
        </span>
      </button>
    </div>

    <form v-if="showForm" @submit.prevent="saveMeal" class="border-t border-meal-gray-light pt-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- Name field -->
        <div>
          <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-name">
            Mahlzeit Name
          </label>
          <input
              id="meal-name"
              v-model="newMeal.name"
              class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
              type="text"
              placeholder="z.B. Pizza Abend"
              required
          />
        </div>

        <!-- Date field -->
        <div>
          <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-date">
            Datum
          </label>
          <input
              id="meal-date"
              v-model="newMeal.date"
              type="date"
              class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
              required
          />
        </div>

        <!-- Payer field -->
        <div>
          <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-payer">
            Bezahlt von
          </label>
          <select
              id="meal-payer"
              v-model="newMeal.userId"
              class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
              required
          >
            <option value="" disabled selected>Bitte wählen...</option>
            <option v-for="user in users" :key="user.id" :value="user.id">
              {{ user.name }}
            </option>
          </select>
        </div>

        <!-- Participants -->
        <div class="md:col-span-2">
          <label class="block text-meal-gray text-sm font-bold mb-2">
            Teilnehmer
          </label>
          <div class="flex flex-wrap gap-2 mb-2">
            <div
                v-for="user in users"
                :key="user.id"
                @click="toggleParticipant(user.id)"
                :class="[
                'flex items-center px-3 py-2 rounded-full cursor-pointer transition-colors',
                isParticipant(user.id)
                  ? 'bg-meal-primary text-white'
                  : 'bg-meal-gray-light text-meal-gray-dark hover:bg-meal-gray'
              ]"
            >
              <div class="w-6 h-6 rounded-full flex items-center justify-center mr-2"
                   :class="isParticipant(user.id) ? 'bg-white text-meal-primary' : 'bg-meal-gray text-white'">
                {{ user.name.charAt(0).toUpperCase() }}
              </div>
              {{ user.name }}
            </div>
          </div>
          <div v-if="newMeal.participants.length === 0" class="text-meal-error text-sm mt-1">
            Bitte wähle mindestens einen Teilnehmer aus.
          </div>
        </div>

        <!-- Description field -->
        <div class="md:col-span-2">
          <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-description">
            Beschreibung (optional)
          </label>
          <textarea
              id="meal-description"
              v-model="newMeal.description"
              class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
              rows="3"
              placeholder="Details zur Mahlzeit"
          ></textarea>
        </div>

        <!-- Products section -->
        <div class="md:col-span-2 mt-4">
          <div class="flex justify-between items-center mb-4">
            <label class="block text-meal-gray text-sm font-bold">
              Produkte / Rezeptkomponenten
            </label>
            <button
                type="button"
                @click="addProduct"
                class="text-meal-primary hover:text-meal-dark transition-colors duration-200 flex items-center text-sm font-bold"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
              </svg>
              Produkt hinzufügen
            </button>
          </div>

          <!-- Product list -->
          <ProductList
              v-model="newMeal.products"
              :users="users"
          />

          <!-- Total amount -->
          <div class="flex justify-between items-center py-3 px-4 bg-meal-light rounded-lg mb-4">
            <span class="font-bold text-meal-gray-dark">Berechneter Gesamtbetrag:</span>
            <span class="font-bold text-meal-primary">{{ formatCurrency(calculateTotalAmount()) }}</span>
          </div>

          <!-- Manual total option -->
          <div class="flex items-center mb-4">
            <input
                id="override-total"
                type="checkbox"
                v-model="newMeal.overrideTotal"
                class="h-4 w-4 text-meal-primary focus:ring-meal-primary border-meal-gray-light rounded"
            />
            <label for="override-total" class="ml-2 block text-sm text-meal-gray-dark">
              Gesamtbetrag manuell anpassen
            </label>
          </div>

          <div v-if="newMeal.overrideTotal" class="mb-4">
            <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-amount">
              Manueller Gesamtbetrag (€)
            </label>
            <input
                id="meal-amount"
                v-model="newMeal.manualTotalAmount"
                type="number"
                step="0.01"
                min="0.01"
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                required
                placeholder="0.00"
            />
          </div>
        </div>
      </div>

      <div class="flex justify-end mt-6">
        <button
            type="submit"
            class="bg-meal-accent hover:bg-opacity-90 text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
            :disabled="!isFormValid"
        >
          Mahlzeit erfassen
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import ProductList from './ProductList.vue';

const props = defineProps({
  users: {
    type: Array,
    required: true
  },
  showForm: {
    type: Boolean,
    default: false
  }
});

const emit = defineEmits(['toggle-form', 'save-meal']);

const newMeal = ref({
  name: '',
  date: new Date().toISOString().split('T')[0],
  userId: '',
  participants: [],
  description: '',
  products: [],
  overrideTotal: false,
  manualTotalAmount: ''
});

const isFormValid = computed(() => {
  if (newMeal.value.participants.length === 0) return false;

  // Form validation logic
  if (newMeal.value.products.length === 0) {
    return !!newMeal.value.name && !!newMeal.value.date && !!newMeal.value.userId &&
        (newMeal.value.overrideTotal ? !!newMeal.value.manualTotalAmount : true);
  }

  // Check if all products are valid
  for (const product of newMeal.value.products) {
    if (!product.name || !product.price) return false;
    if (product.isSpecific && product.specificParticipants.length === 0) return false;
  }

  return !!newMeal.value.name && !!newMeal.value.date && !!newMeal.value.userId;
});

function formatCurrency(value) {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR'
  }).format(value);
}

function toggleParticipant(userId) {
  const index = newMeal.value.participants.indexOf(userId);
  if (index === -1) {
    newMeal.value.participants.push(userId);
  } else {
    newMeal.value.participants.splice(index, 1);
  }
}

function isParticipant(userId) {
  return newMeal.value.participants.includes(userId);
}

function addProduct() {
  newMeal.value.products.push({
    name: '',
    price: '',
    isSpecific: false,
    specificParticipants: [],
    showParticipants: false
  });
}

function calculateTotalAmount() {
  if (newMeal.value.products.length === 0) {
    return 0;
  }

  return newMeal.value.products.reduce((total, product) => {
    return total + parseFloat(product.price || 0);
  }, 0);
}

function saveMeal() {
  if (!isFormValid.value) return;

  emit('save-meal', { ...newMeal.value });

  // Reset form
  newMeal.value = {
    name: '',
    date: new Date().toISOString().split('T')[0],
    userId: '',
    participants: [],
    description: '',
    products: [],
    overrideTotal: false,
    manualTotalAmount: ''
  };
}
</script>