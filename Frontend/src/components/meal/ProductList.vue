<template>
  <div class="mb-4">
    <div v-if="modelValue && modelValue.length > 0" class="border rounded-lg overflow-hidden">
      <table class="min-w-full divide-y divide-meal-gray-light">
        <thead class="bg-meal-light">
        <tr>
          <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-meal-gray-dark uppercase tracking-wider">Produkt</th>
          <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-meal-gray-dark uppercase tracking-wider">Preis</th>
          <th scope="col" class="px-3 py-2 text-left text-xs font-medium text-meal-gray-dark uppercase tracking-wider">Für</th>
          <th scope="col" class="px-3 py-2 text-right text-xs font-medium text-meal-gray-dark uppercase tracking-wider"></th>
        </tr>
        </thead>
        <tbody class="bg-white divide-y divide-meal-gray-light">
        <tr v-for="(product, index) in localProducts" :key="index" class="hover:bg-meal-light">
          <td class="px-3 py-2 whitespace-nowrap">
            <input
                v-if="editable"
                v-model="localProducts[index].name"
                @input="updateProducts"
                class="w-full border-0 bg-transparent focus:ring-0 px-0 py-0 placeholder-meal-gray"
                type="text"
                placeholder="Produktname"
                required
            />
            <span v-else>{{ product.name }}</span>
          </td>
          <td class="px-3 py-2 whitespace-nowrap">
            <input
                v-if="editable"
                v-model="localProducts[index].price"
                @input="updateProducts"
                class="w-24 border-0 bg-transparent focus:ring-0 px-0 py-0"
                type="number"
                step="0.01"
                min="0.01"
                placeholder="0.00"
                required
            />
            <span v-else>{{ formatCurrency(product.price) }}</span>
          </td>
          <td class="px-3 py-2">
            <div v-if="editable">
              <button
                  type="button"
                  @click="toggleShowParticipants(index)"
                  class="text-xs flex items-center whitespace-nowrap text-meal-primary hover:text-meal-dark"
              >
                {{ product.isSpecific ? 'Für bestimmte Personen' : 'Für alle' }}
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-1" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </button>

              <!-- Participant selection dropdown -->
              <div v-if="product.showParticipants" class="mt-2 p-2 bg-meal-light rounded-md">
                <div class="flex items-center mb-2">
                  <input
                      :id="`product-specific-${index}`"
                      type="checkbox"
                      v-model="localProducts[index].isSpecific"
                      @change="updateProducts"
                      class="h-4 w-4 text-meal-primary focus:ring-meal-primary border-meal-gray-light rounded"
                  />
                  <label :for="`product-specific-${index}`" class="ml-2 block text-sm text-meal-gray-dark">
                    Nur für bestimmte Teilnehmer
                  </label>
                </div>

                <div v-if="product.isSpecific" class="mt-2">
                  <div class="flex flex-wrap gap-1">
                    <div
                        v-for="user in users"
                        :key="`${index}-${user.id}`"
                        @click="toggleProductParticipant(index, user.id)"
                        :class="[
                          'flex items-center px-2 py-1 rounded-full cursor-pointer transition-colors text-xs',
                          isProductParticipant(index, user.id)
                            ? 'bg-meal-accent text-white'
                            : 'bg-meal-gray-light text-meal-gray-dark hover:bg-meal-gray'
                        ]"
                    >
                      {{ user.name }}
                    </div>
                  </div>
                  <div v-if="product.isSpecific && product.specificParticipants.length === 0" class="text-meal-error text-xs mt-1">
                    Bitte wähle mindestens einen Teilnehmer.
                  </div>
                </div>
              </div>
            </div>
            <div v-else>
              <span v-if="!product.isSpecific">Alle Teilnehmer</span>
              <div v-else class="flex flex-wrap gap-1">
                  <span v-for="participantId in product.specificParticipants" :key="participantId"
                        class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-meal-light">
                    {{ getUserName(participantId) }}
                  </span>
              </div>
            </div>
          </td>
          <td class="px-3 py-2 text-right">
            <button
                v-if="editable"
                type="button"
                @click="removeProduct(index)"
                class="text-meal-error hover:text-red-700"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="text-center p-4 bg-meal-light rounded-lg text-meal-gray">
      <p>Noch keine Produkte hinzugefügt.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue';

const props = defineProps({
  modelValue: {
    type: Array,
    required: true
  },
  users: {
    type: Array,
    required: true
  },
  editable: {
    type: Boolean,
    default: true
  }
});

const emit = defineEmits(['update:modelValue']);

// Create a local copy of products to avoid directly modifying props
const localProducts = ref([]);

// Initialize local products from props
onMounted(() => {
  localProducts.value = JSON.parse(JSON.stringify(props.modelValue || []));
});

// Update localProducts when modelValue changes
watch(() => props.modelValue, (newVal) => {
  if (JSON.stringify(newVal) !== JSON.stringify(localProducts.value)) {
    localProducts.value = JSON.parse(JSON.stringify(newVal || []));
  }
}, { deep: true });

// Send updates to parent component
function updateProducts() {
  emit('update:modelValue', JSON.parse(JSON.stringify(localProducts.value)));
}

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

function toggleShowParticipants(index) {
  localProducts.value[index].showParticipants = !localProducts.value[index].showParticipants;
  // No need to emit update here as it doesn't affect the data model
}

function toggleProductParticipant(productIndex, userId) {
  const product = localProducts.value[productIndex];

  if (!product.specificParticipants) {
    product.specificParticipants = [];
  }

  const index = product.specificParticipants.indexOf(userId);

  if (index === -1) {
    product.specificParticipants.push(userId);
  } else {
    product.specificParticipants.splice(index, 1);
  }

  updateProducts();
}

function isProductParticipant(productIndex, userId) {
  const product = localProducts.value[productIndex];
  return product.specificParticipants && product.specificParticipants.includes(userId);
}

function removeProduct(index) {
  localProducts.value.splice(index, 1);
  updateProducts();
}
</script>