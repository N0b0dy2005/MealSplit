<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-xl font-header font-bold text-meal-gray-dark">Mahlzeit bearbeiten</h3>
        <button @click="$emit('close')" class="text-meal-gray hover:text-meal-error">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <form @submit.prevent="updateMeal">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Name field -->
          <div>
            <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-name">
              Mahlzeit Name
            </label>
            <input
                id="edit-meal-name"
                v-model="editedMeal.name"
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                type="text"
                placeholder="z.B. Pizza Abend"
                required
            />
          </div>

          <!-- Date field -->
          <div>
            <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-date">
              Datum
            </label>
            <input
                id="edit-meal-date"
                v-model="editedMeal.date"
                type="date"
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                required
            />
          </div>

          <!-- Payer field -->
          <div>
            <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-payer">
              Bezahlt von
            </label>
            <select
                id="edit-meal-payer"
                v-model="editedMeal.userId"
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                required
            >
              <option value="" disabled>Bitte wählen...</option>
              <option v-for="user in users" :key="user.id" :value="user.id">
                {{ user.name }}
              </option>
            </select>
          </div>

          <!-- Total amount field -->
          <div>
            <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-amount">
              Gesamtbetrag (€)
            </label>
            <input
                id="edit-meal-amount"
                v-model="editedMeal.totalAmount"
                type="number"
                step="0.01"
                min="0.01"
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                required
                placeholder="0.00"
            />
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
            <div v-if="editedMeal.participants.length === 0" class="text-meal-error text-sm mt-1">
              Bitte wähle mindestens einen Teilnehmer aus.
            </div>
          </div>

          <!-- Description field -->
          <div class="md:col-span-2">
            <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-description">
              Beschreibung (optional)
            </label>
            <textarea
                id="edit-meal-description"
                v-model="editedMeal.description"
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

            <!-- Receipt upload -->
            <div class="mb-4 p-4 border border-dashed border-meal-gray-light rounded-lg">
              <p class="text-sm text-meal-gray mb-2">Kassenbon hochladen</p>
              <div class="flex items-center space-x-3">
                <input
                    type="file"
                    @change="handleFileChange"
                    accept="image/*"
                    class="text-sm text-meal-gray-dark"
                />
                <button
                    @click="uploadReceipt"
                    :disabled="!selectedFile"
                    class="px-3 py-1 bg-meal-primary text-white rounded disabled:opacity-50 disabled:cursor-not-allowed text-sm"
                >
                  Hochladen
                </button>
              </div>
            </div>

            <!-- Product list -->
            <ProductList
                v-model="editedMeal.productsData"
                :users="users"
            />
          </div>
        </div>

        <div class="flex justify-end space-x-3 mt-6">
          <button
              type="button"
              @click="$emit('close')"
              class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
          >
            Abbrechen
          </button>
          <button
              type="submit"
              class="bg-meal-accent hover:bg-opacity-90 text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
              :disabled="editedMeal.participants.length === 0"
          >
            Speichern
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import ProductList from './ProductList.vue';
import { mutation } from '../../graphql';

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

const emit = defineEmits(['close', 'update']);
const selectedFile = ref(null);

// Create a deep copy of the meal to edit
const editedMeal = ref({...props.meal});

// Initialize editedMeal with the correct data structure
onMounted(() => {
  // Ensure productsData is an array
  if (!Array.isArray(editedMeal.value.productsData)) {
    editedMeal.value.productsData = [];
  }

  // Make a deep copy of participants array
  editedMeal.value.participants = [...props.meal.participants];
});

function handleFileChange(e) {
  selectedFile.value = e.target.files[0];
}

async function uploadReceipt() {
  if (!selectedFile.value) return;

  try {
    const [data, err] = await mutation.uploadReceipt(selectedFile.value);
    if (err) {
      console.error('Error uploading receipt:', err);
      return;
    }
    console.log('Receipt uploaded:', data);
    // Handle the response, possibly adding extracted items to productsData
  } catch (error) {
    console.error('Error uploading receipt:', error);
  }
}

function toggleParticipant(userId) {
  const index = editedMeal.value.participants.indexOf(userId);
  if (index === -1) {
    editedMeal.value.participants.push(userId);
  } else {
    editedMeal.value.participants.splice(index, 1);
  }
}

function isParticipant(userId) {
  return editedMeal.value.participants.includes(userId);
}

function addProduct() {
  if (!editedMeal.value.productsData) {
    editedMeal.value.productsData = [];
  }

  editedMeal.value.productsData.push({
    name: '',
    price: '',
    isSpecific: false,
    specificParticipants: [],
    showParticipants: false
  });
}

function updateMeal() {
  if (editedMeal.value.participants.length === 0) return;

  emit('update', editedMeal.value);
}
</script>