<template>
  <div class="debt-details bg-gradient-to-br from-white to-gray-50 rounded-2xl shadow-lg p-6 border border-gray-100">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-2xl font-bold text-gray-800">Schulden Übersicht</h3>
      <div class="bg-blue-50 px-4 py-2 rounded-full" :class="{'bg-red-50 text-red-600': totalAmount > 0, 'bg-green-50 text-green-600': totalAmount <= 0}">
        <span class="font-semibold">Gesamt: {{ formatCurrency(totalAmount) }}</span>
      </div>
    </div>

    <div v-if="loading" class="flex flex-col items-center justify-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-3 border-blue-500 mb-4"></div>
      <p class="text-gray-500 font-medium">Daten werden geladen...</p>
    </div>

    <div v-else-if="error" class="bg-red-50 text-red-600 p-4 rounded-xl text-center">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 mx-auto mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <p class="font-medium">{{ error }}</p>
    </div>

    <div v-else-if="details && details.length > 0" class="space-y-6">
      <!-- Filters/Controls (Optional) -->
      <div class="flex flex-wrap gap-2 mb-4">
        <button class="px-4 py-2 rounded-full text-sm font-medium bg-gray-100 hover:bg-gray-200 transition-colors">
          Alle
        </button>
        <button class="px-4 py-2 rounded-full text-sm font-medium bg-gray-100 hover:bg-gray-200 transition-colors">
          Mahlzeiten
        </button>
        <button class="px-4 py-2 rounded-full text-sm font-medium bg-gray-100 hover:bg-gray-200 transition-colors">
          Zahlungen
        </button>
      </div>

      <!-- Table -->
      <div class="overflow-hidden rounded-xl border border-gray-200 shadow-sm">
        <div class="overflow-y-auto max-h-96">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Transaktion
              </th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Datum
              </th>
              <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Typ
              </th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Betrag
              </th>
            </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(detail, index) in details" :key="index" class="hover:bg-gray-50 transition-colors duration-150">
              <td class="px-4 py-3 whitespace-nowrap">
                <div class="flex items-center">
                  <!-- Icon based on type -->
                  <div class="flex-shrink-0 h-8 w-8 mr-3 rounded-full flex items-center justify-center" :class="{'bg-blue-100 text-blue-500': detail.type === 'meal', 'bg-green-100 text-green-500': detail.type === 'payment'}">
                    <svg v-if="detail.type === 'meal'" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                    </svg>
                    <svg v-else-if="detail.type === 'payment'" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <!-- Transaction name -->
                  <div class="text-sm font-medium text-gray-900">
                      <span v-if="detail.type === 'meal'">
                        {{ detail.meal_info?.name || 'Unbenannte Mahlzeit' }}
                      </span>
                    <span v-else-if="detail.type === 'payment'">
                        {{ detail.payment_info?.description || 'Zahlung' }}
                      </span>
                    <span v-else>Unbekannt</span>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 whitespace-nowrap">
                <div class="text-sm text-gray-500">
                  {{ formatDate(detail.type === 'meal' ? detail.meal_info?.date : detail.payment_info?.date || detail.create_date) }}
                </div>
              </td>
              <td class="px-4 py-3 whitespace-nowrap">
                  <span v-if="detail.type === 'meal'" class="px-3 py-1 inline-flex text-xs leading-5 font-semibold rounded-full bg-blue-100 text-blue-800">
                    Mahlzeit
                  </span>
                <span v-else-if="detail.type === 'payment'" class="px-3 py-1 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                    Zahlung
                  </span>
                <span v-else class="px-3 py-1 inline-flex text-xs leading-5 font-semibold rounded-full bg-gray-100 text-gray-800">
                    Unbekannt
                  </span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-right text-sm font-medium" :class="{'text-red-600': isDebt(detail), 'text-green-600': !isDebt(detail)}">
                {{ formatCurrency(parseFloat(detail.amount)) }}
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Summary information -->
      <div class="bg-gray-50 rounded-xl p-4 flex justify-between items-center">
        <div>
          <span class="text-sm text-gray-500">{{ details.length }} Transaktionen</span>
        </div>
        <div class="text-lg font-bold" :class="{'text-red-600': totalAmount > 0, 'text-green-600': totalAmount <= 0}">
          {{ formatCurrency(totalAmount) }}
        </div>
      </div>
    </div>

    <div v-else class="flex flex-col items-center justify-center py-12 text-gray-400">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
      </svg>
      <p class="text-lg font-medium">Keine Transaktionen vorhanden</p>
      <p class="text-sm text-gray-400 mt-1">Hier werden Ihre Schulden und Zahlungen angezeigt</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { query } from '../graphql.ts';

// Definiere Typen für die API-Antwort
interface MealInfo {
  id: number;
  name: string;
  description?: string;
  date: string;
  total_amount: string;
}

interface PaymentInfo {
  id: number;
  description?: string;
  date: string;
  amount: string;
}

interface DebtDetail {
  id: number;
  from_user_id: number;
  to_user_id: number;
  from_user_name: string;
  to_user_name: string;
  amount: string;
  meals_count: number;
  create_date: string;
  is_confirmed: boolean;
  type: 'meal' | 'payment' | 'unknown';
  meal_id?: number;
  payment_id?: number;
  meal_info?: MealInfo;
  payment_info?: PaymentInfo;
}

// Props
const props = defineProps<{
  fromUserId: number;
  toUserId: number;
  currentUserId?: number; // Optionales Prop für den aktuellen Benutzer
}>();

// Reaktive Daten
const details = ref<DebtDetail[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

// Berechnete Eigenschaften
const totalAmount = computed(() => {
  return details.value.reduce((sum, detail) => {
    const amount = parseFloat(detail.amount);
    // Check if the amount is being paid by the user or owed to the user
    const multiplier = isDebt(detail) ? 1 : -1;
    return sum + (amount * multiplier);
  }, 0);
});

// Hilfsfunktionen
function formatCurrency(value: number): string {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR'
  }).format(value);
}

function formatDate(dateString?: string): string {
  if (!dateString) return 'Unbekannt';
  return new Date(dateString).toLocaleDateString('de-DE');
}

// Prüfe, ob der aktuelle Benutzer Geld schuldet (fromUserId)
function isDebt(detail: DebtDetail): boolean {
  if (!props.currentUserId) return detail.from_user_id === props.fromUserId;
  return detail.from_user_id === props.currentUserId;
}

// Daten laden
async function loadDebtDetails() {
  loading.value = true;
  error.value = null;

  try {
    const [data, err] = await query.getDebtDetails(
        props.fromUserId,
        props.toUserId
    );

    if (err) {
      error.value = 'Fehler beim Laden der Details: ' + err.message;
      return;
    }

    details.value = data || [];
  } catch (e) {
    error.value = 'Fehler beim Laden der Details. Bitte versuche es später erneut.';
    console.error(e);
  } finally {
    loading.value = false;
  }
}

// Initialisierung
onMounted(() => {
  loadDebtDetails();
});
</script>