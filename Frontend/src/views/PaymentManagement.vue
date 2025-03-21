<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <!-- Header -->
    <header class="bg-meal-primary text-white p-4 shadow-md mb-4 sm:mb-8">
      <div class="container mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 sm:h-8 sm:w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"/>
          </svg>
          <h1 class="text-xl sm:text-2xl font-header font-bold">Zahlungen</h1>
        </div>
        <button @click="goBackToHomeScreen"
                class="text-white hover:text-meal-accent-light transition-colors duration-200">
          <span class="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
            <span class="hidden sm:inline">Zurück zur Startseite</span>
            <span class="sm:hidden">Zurück</span>
          </span>
        </button>
      </div>
    </header>

    <div class="container mx-auto px-4">
      <!-- Neue Zahlung erstellen -->
      <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6 mb-6 sm:mb-8">
        <h2 class="text-xl sm:text-2xl font-header font-bold text-meal-gray-dark mb-4 sm:mb-6">Zahlung erfassen</h2>

        <div>
          <div class="grid grid-cols-1 gap-4 sm:gap-6">
            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="from-user">
                Wer hat bezahlt?
              </label>
              <select
                  id="from-user"
                  v-model="newPayment.fromUserId"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
              >
                <option value="" disabled selected>Bitte wählen...</option>
                <option v-for="user in users" :key="user.id" :value="user.id">
                  {{ user.name }}
                </option>
              </select>
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="to-user">
                Wer hat das Geld erhalten?
              </label>
              <select
                  id="to-user"
                  v-model="newPayment.toUserId"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
              >
                <option value="" disabled selected>Bitte wählen...</option>
                <option v-for="user in users" :key="user.id" :value="user.id">
                  {{ user.name }}
                </option>
              </select>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div>
                <label class="block text-meal-gray text-sm font-bold mb-2" for="amount">
                  Betrag (€)
                </label>
                <input
                    id="amount"
                    v-model="newPayment.amount"
                    type="number"
                    step="0.01"
                    min="0.01"
                    class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                    required
                    placeholder="0.00"
                />
              </div>

              <div>
                <label class="block text-meal-gray text-sm font-bold mb-2" for="date">
                  Datum
                </label>
                <input
                    id="date"
                    v-model="newPayment.date"
                    type="date"
                    class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                    required
                />
              </div>
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="description">
                Beschreibung (optional)
              </label>
              <textarea
                  id="description"
                  v-model="newPayment.description"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  rows="3"
                  placeholder="Wofür war diese Zahlung?"
              ></textarea>
            </div>
          </div>

          <div class="flex justify-center sm:justify-end mt-6">
            <button @click="createPayment"
                    type="submit"
                    class="bg-meal-accent hover:bg-opacity-90 text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200 w-full sm:w-auto"
            >
              Zahlung erfassen
            </button>
          </div>
        </div>
      </div>

      <!-- Zahlungsliste -->
      <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6 mb-6 sm:mb-8">
        <h2 class="text-xl sm:text-2xl font-header font-bold text-meal-gray-dark mb-4 sm:mb-6">Zahlungsübersicht</h2>

        <!-- Desktop Table (hidden on mobile) -->
        <div class="hidden sm:block overflow-x-auto">
          <table class="min-w-full bg-white">
            <thead>
            <tr class="bg-meal-light text-meal-dark text-left">
              <th class="py-3 px-4 rounded-tl-lg">Datum</th>
              <th class="py-3 px-4">Von</th>
              <th class="py-3 px-4">An</th>
              <th class="py-3 px-4">Betrag</th>
              <th class="py-3 px-4 rounded-tr-lg text-right">Aktionen</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="payment in sortedPayments" :key="payment.id"
                class="border-b border-meal-gray-light hover:bg-meal-gray-light transition-colors duration-150">
              <td class="py-3 px-4">{{ formatDate(payment.date) }}</td>
              <td class="py-3 px-4">
                <div class="flex items-center">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(payment.fromUserId) }">
                    {{ getUserInitial(payment.fromUserId) }}
                  </div>
                  {{ getUserName(payment.fromUserId) }}
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(payment.toUserId) }">
                    {{ getUserInitial(payment.toUserId) }}
                  </div>
                  {{ getUserName(payment.toUserId) }}
                </div>
              </td>
              <td class="py-3 px-4 font-bold">{{ formatCurrency(payment.amount) }}</td>
              <td class="py-3 px-4 text-right">
                <button
                    v-if="!payment.confirmed"
                    @click="confirmPayment(payment)"
                    class="text-meal-primary hover:text-meal-dark mr-2"
                    title="Bestätigen"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd"
                          d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                          clip-rule="evenodd"/>
                  </svg>
                </button>
                <button
                    @click="deletePayment(payment.id)"
                    class="text-meal-error hover:text-red-700"
                    title="Löschen"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd"
                          d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                          clip-rule="evenodd"/>
                  </svg>
                </button>
              </td>
            </tr>
            <tr v-if="payments.length === 0">
              <td colspan="6" class="py-4 px-4 text-center text-meal-gray">
                Keine Zahlungen vorhanden. Erfasse deine erste Zahlung!
              </td>
            </tr>
            </tbody>
          </table>
        </div>

        <!-- Mobile Cards (shown only on mobile) -->
        <div class="sm:hidden">
          <div v-if="payments.length === 0" class="py-4 px-4 text-center text-meal-gray">
            Keine Zahlungen vorhanden. Erfasse deine erste Zahlung!
          </div>

          <div v-else class="space-y-4">
            <div v-for="payment in sortedPayments" :key="payment.id"
                 class="border border-meal-gray-light rounded-lg p-4 hover:bg-meal-gray-light transition-colors duration-150">
              <div class="flex justify-between items-center mb-3">
                <span class="text-sm text-meal-gray">{{ formatDate(payment.date) }}</span>
                <div class="flex items-center">
                  <button
                      v-if="!payment.confirmed"
                      @click="confirmPayment(payment)"
                      class="text-meal-primary hover:text-meal-dark p-2"
                      title="Bestätigen"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd"
                            d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                            clip-rule="evenodd"/>
                    </svg>
                  </button>
                  <button
                      @click="deletePayment(payment.id)"
                      class="text-meal-error hover:text-red-700 p-2"
                      title="Löschen"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd"
                            d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                            clip-rule="evenodd"/>
                    </svg>
                  </button>
                </div>
              </div>

              <div class="flex justify-between items-center mb-3">
                <div class="flex items-center">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(payment.fromUserId) }">
                    {{ getUserInitial(payment.fromUserId) }}
                  </div>
                  <div>
                    <div class="text-xs text-meal-gray">Von</div>
                    <div class="text-sm font-medium">{{ getUserName(payment.fromUserId) }}</div>
                  </div>
                </div>

                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-meal-gray mx-1" fill="none"
                     viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/>
                </svg>

                <div class="flex items-center">
                  <div>
                    <div class="text-xs text-meal-gray text-right">An</div>
                    <div class="text-sm font-medium">{{ getUserName(payment.toUserId) }}</div>
                  </div>
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold ml-2"
                       :style="{ backgroundColor: getUserColor(payment.toUserId) }">
                    {{ getUserInitial(payment.toUserId) }}
                  </div>
                </div>
              </div>

              <div class="text-center font-bold text-lg">
                {{ formatCurrency(payment.amount) }}
              </div>

              <div v-if="payment.description" class="mt-2 text-xs text-meal-gray border-t border-meal-gray-light pt-2">
                {{ payment.description }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Schuldenübersicht -->
      <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
        <h2 class="text-xl sm:text-2xl font-header font-bold text-meal-gray-dark mb-4 sm:mb-6">Schuldenübersicht</h2>

        <div v-if="debts && debts.length > 0">
          <!-- Desktop Table (hidden on mobile) -->
          <div class="hidden sm:block overflow-x-auto">
            <table class="min-w-full bg-white">
              <thead>
              <tr class="bg-meal-light text-meal-dark text-left">
                <th class="py-3 px-4 rounded-tl-lg">Datum</th>
                <th class="py-3 px-4">Schuldner</th>
                <th class="py-3 px-4">Gläubiger</th>
                <th class="py-3 px-4">Betrag</th>
                <th class="py-3 px-4 rounded-tr-lg">Mahlzeiten</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="debt in debts" :key="debt.id"
                  class="border-b border-meal-gray-light hover:bg-meal-gray-light transition-colors duration-150">
                <td class="py-3 px-4">{{ new Date(debt.createDate).toLocaleDateString('de-DE') }}</td>
                <td class="py-3 px-4">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                         :style="{ backgroundColor: getUserColor(debt.fromUserId) }">
                      {{ getUserInitial(debt.fromUserId) }}
                    </div>
                    {{ getUserName(debt.fromUserId) }}
                  </div>
                </td>
                <td class="py-3 px-4">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                         :style="{ backgroundColor: getUserColor(debt.toUserId) }">
                      {{ getUserInitial(debt.toUserId) }}
                    </div>
                    {{ getUserName(debt.toUserId) }}
                  </div>
                </td>
                <td class="py-3 px-4 font-bold text-meal-error">{{ formatCurrency(parseFloat(debt.amount)) }}</td>
                <td class="py-3 px-4">{{ debt.mealsCount }} {{ debt.mealsCount === 1 ? 'Mahlzeit' : 'Mahlzeiten' }}</td>
              </tr>
              </tbody>
            </table>
          </div>

          <!-- Mobile Cards (shown only on mobile) -->
          <div class="sm:hidden space-y-4">
            <div v-for="debt in debts" :key="debt.id"
                 class="border border-meal-gray-light rounded-lg p-4 hover:bg-meal-gray-light transition-colors duration-150">
              <div class="flex justify-between items-center mb-3">
                <span class="text-sm text-meal-gray">{{ new Date(debt.createDate).toLocaleDateString('de-DE') }}</span>
                <span class="text-xs bg-meal-light rounded-full px-2 py-1">
                  {{ debt.mealsCount }} {{ debt.mealsCount === 1 ? 'Mahlzeit' : 'Mahlzeiten' }}
                </span>
              </div>

              <div class="flex justify-between items-center mb-3">
                <div class="flex items-center">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(debt.fromUserId) }">
                    {{ getUserInitial(debt.fromUserId) }}
                  </div>
                  <div>
                    <div class="text-xs text-meal-gray">Schuldner</div>
                    <div class="text-sm font-medium">{{ getUserName(debt.fromUserId) }}</div>
                  </div>
                </div>

                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-meal-gray mx-1" fill="none"
                     viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/>
                </svg>

                <div class="flex items-center">
                  <div>
                    <div class="text-xs text-meal-gray text-right">Gläubiger</div>
                    <div class="text-sm font-medium">{{ getUserName(debt.toUserId) }}</div>
                  </div>
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold ml-2"
                       :style="{ backgroundColor: getUserColor(debt.toUserId) }">
                    {{ getUserInitial(debt.toUserId) }}
                  </div>
                </div>
              </div>

              <div class="text-center font-bold text-lg text-meal-error">
                {{ formatCurrency(parseFloat(debt.amount)) }}
              </div>
            </div>
          </div>
        </div>

        <div v-else class="text-center py-6 text-meal-gray">
          Keine offenen Beträge vorhanden. Super!
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, ref} from "vue";
import router from "../router";
import {mutation, query} from "../graphql.ts";

// Beispiel-Benutzer (sollten aus einer Store/API kommen)
const users = ref([]);

// Zahlungen
const payments = ref([]);

// Neue Zahlung
const newPayment = ref({
  fromUserId: '',
  toUserId: '',
  amount: '',
  date: new Date().toISOString().split('T')[0],
  description: ''
});

// Farben für Benutzer-Avatare
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

// Berechnete Eigenschaften
const sortedPayments = computed(() => {
  return [...payments.value].sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime());
});

async function getDebts() {
  const [data, err] = await query.getDebts();
  if (err) {
    console.error('Error getting debts:', err);
    return;
  }
  debts.value = data;
}

// Beträge zwischen Benutzern berechnen
const debts = ref([]);

// Methoden
function formatCurrency(value) {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR'
  }).format(value);
}

function formatDate(dateString) {
  return new Date(dateString).toLocaleDateString('de-DE');
}

function getUserColor(userId) {
  return avatarColors[userId % avatarColors.length];
}

function getUserInitial(userId) {
  const user = users.value.find(u => u.id === userId);
  return user ? user.name.charAt(0).toUpperCase() : '?';
}

function getUserName(userId) {
  const user = users.value.find(u => u.id === userId);
  return user ? user.name : 'Unbekannt';
}

async function getUsers() {
  const [data, err] = await query.getUsers();
  if (err) {
    console.error('Error getting users:', err);
    return;
  }
  users.value = data;
}

async function createPayment() {
  let payment = {
    fromUserId: newPayment.value.fromUserId,
    toUserId: newPayment.value.toUserId,
    amount: parseFloat(newPayment.value.amount),
    description: newPayment.value.description,
  }

  const [data, err] = await mutation.createPayment(payment);
  if (err) {
    console.error('Error getting createPayment:', err);
    return;
  }
  await getPayments();
}

async function getPayments() {
  const [data, err] = await query.getPayments();
  if (err) {
    console.error('Error getting purchaser for current user:', err);
    return;
  }
  payments.value = data;
}

function confirmPayment(payment) {
  const index = payments.value.findIndex(p => p.id === payment.id);
  if (index !== -1) {
    payments.value[index] = {...payment, confirmed: true};
  }
}

function deletePayment(paymentId) {
  payments.value = payments.value.filter(p => p.id !== paymentId);
}

function goBackToHomeScreen() {
  router.push('/homefeed');
}

// Beim Laden der Komponente
onMounted(() => {
  // Hier könntest du Daten aus dem Backend laden
  getUsers();
  getPayments();
  getDebts();
});
</script>