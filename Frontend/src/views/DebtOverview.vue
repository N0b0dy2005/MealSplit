<template>
  <div class="h-screen overflow-auto flex flex-col bg-meal-light font-sans pb-8">
    <div class="container mx-auto px-4 mt-8">
      <!-- Zusammenfassung -->
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 sm:gap-6 mb-6 sm:mb-8">
        <StatCard title="Offene Beträge" :value="totalDebt" caption="Summe aller ausstehenden Beträge" is-currency
                  value-color="text-meal-error"/>

        <StatCard title="Aktive Transaktionen" :value="debts.length" caption="Anzahl der offenen Schuldenbeziehungen"
                  value-color="text-meal-primary"/>

      </div>
      <!-- Hauptansicht mit Tabs -->
      <div class="bg-white rounded-xl shadow-meal overflow-hidden mb-8">
        <div class="flex border-b border-meal-gray-light">
          <button
              @click="activeTab = 'list'"
              :class="[
              'py-2 sm:py-4 px-2 sm:px-6 font-bold text-center flex-1 text-sm sm:text-base',
              activeTab === 'list'
                ? 'text-meal-primary border-b-2 border-meal-primary'
                : 'text-meal-gray hover:text-meal-gray-dark'
            ]"
          >
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 sm:mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd"
                      d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                      clip-rule="evenodd"/>
              </svg>
              <span class="hidden sm:inline">Beträge Liste</span>
            </span>
          </button>
          <button
              @click="activeTab = 'matrix'"
              :class="[
              'py-2 sm:py-4 px-2 sm:px-6 font-bold text-center flex-1 text-sm sm:text-base',
              activeTab === 'matrix'
                ? 'text-meal-primary border-b-2 border-meal-primary'
                : 'text-meal-gray hover:text-meal-gray-dark'
            ]"
          >
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 sm:mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path
                    d="M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V5zM11 13a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"/>
              </svg>
              <span class="hidden sm:inline">Beträge Matrix</span>
            </span>
          </button>
        </div>

        <!-- Tabs-Inhalte (ListTab, MatrixTab) -->
        <div v-if="activeTab === 'list'" class="p-4 sm:p-6">
          <div v-if="debts.length === 0" class="text-center py-6 sm:py-8 text-meal-gray">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 sm:h-16 sm:w-16 mx-auto mb-4 text-meal-gray-light"
                 fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            <p class="text-lg sm:text-xl mb-2">Keine offenen Beträge</p>
            <p>Alle Beträge sind bezahlt. Gut gemacht!</p>
          </div>

          <div v-else>
            <!-- Desktop Table -->
            <div class="hidden sm:block overflow-x-auto">
              <table class="min-w-full bg-white">
                <thead>
                <tr class="bg-meal-light text-meal-dark text-left">
                  <th class="py-3 px-4 rounded-tl-lg w-1/4">
                    <div class="flex items-center">
                      <div
                          class="w-6 h-6 rounded-full bg-meal-error flex items-center justify-center text-white font-bold mr-2 text-xs">
                        S
                      </div>
                      <span>Schuldner</span>
                    </div>
                  </th>
                  <th class="py-3 px-4 text-center w-1/5">
                    <div class="flex flex-col items-center">
                      <span class="text-xl font-bold text-meal-error">€</span>
                      <span>Schuldet Betrag</span>
                    </div>
                  </th>
                  <th class="py-3 px-4 w-1/4">
                    <div class="flex items-center">
                      <div
                          class="w-6 h-6 rounded-full bg-meal-positive flex items-center justify-center text-white font-bold mr-2 text-xs">
                        G
                      </div>
                      <span>Gläubiger</span>
                    </div>
                  </th>
                  <th class="py-3 px-4 text-center">Details</th>
                  <th class="py-3 px-4 rounded-tr-lg text-right">Aktionen</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(debt, index) in debts" :key="index"
                    class="border-b border-meal-gray-light hover:bg-meal-gray-light transition-colors duration-150"
                    :class="{'opacity-50': !debt.isConfirmed}">
                  <!-- Schuldner -->
                  <td class="py-3 px-4">
                    <div class="flex items-center bg-red-50 p-2 rounded-lg">
                      <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold mr-3"
                           :style="{ backgroundColor: getUserColor(debt.fromUserId) }">
                        {{ debt.from.name.charAt(0).toUpperCase() }}
                      </div>
                      <div>
                        <div class="font-bold">{{ debt.from.name }}</div>
                        <div class="text-xs text-meal-error font-medium">
                          Schuldet Geld
                          <span v-if="!debt.isConfirmed" class="text-meal-gray ml-1">(Ausstehend)</span>
                        </div>
                      </div>
                    </div>
                  </td>

                  <!-- Betrag -->
                  <td class="py-3 px-4 text-center">
                    <div class="flex flex-col items-center">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-meal-error mb-1" fill="none"
                           viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                              d="M17 8l4 4m0 0l-4 4m4-4H3"/>
                      </svg>
                      <span class="text-xl font-bold text-meal-error">{{ formatCurrency(debt.amount) }}</span>
                    </div>
                  </td>

                  <!-- Gläubiger -->
                  <td class="py-3 px-4">
                    <div class="flex items-center bg-green-50 p-2 rounded-lg">
                      <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold mr-3"
                           :style="{ backgroundColor: getUserColor(debt.toUserId) }">
                        {{ debt.to.name.charAt(0).toUpperCase() }}
                      </div>
                      <div>
                        <div class="font-bold">{{ debt.to.name }}</div>
                        <div class="text-xs text-meal-positive font-medium">Bekommt Geld</div>
                      </div>
                    </div>
                  </td>

                  <!-- Details Button statt Mahlzeiten -->
                  <td class="py-3 px-4 text-center">
                    <button 
                      @click="showDebtDetails(debt)"
                      class="bg-meal-light hover:bg-meal-gray-light text-meal-primary px-3 py-1 rounded-lg transition-colors duration-200"
                    >
                     Mahlzeiten anzeigen
                    </button>
                  </td>
                  
                  <!-- Aktionen -->
                  <td v-if="debt.isConfirmed" class="py-3 px-4 text-right">
                    <button
                        @click="markAsPaid(index)"
                        class="bg-meal-primary hover:bg-meal-dark text-white px-3 py-1 rounded transition-colors duration-200 w-48"
                    >
                      Als bezahlt markieren
                    </button>
                  </td>
                  <td v-else-if="!debt.isConfirmed && currentUser?.id == debt.toUserId " class="py-3 px-4 text-right">
                    <button
                        @click="markAsPaid(index)"
                        class="bg-meal-accent hover:bg-meal-accent-dark text-white px-3 py-1 rounded transition-colors duration-200 w-48"
                    >
                      Zahlung bestätigen
                    </button>
                  </td>
                  <td v-else class="py-3 px-4 text-right">
                    <span class="text-sm font-medium text-meal-accent">Zahlung ist ausstehend</span>
                  </td>
                </tr>
                </tbody>
              </table>
            </div>

            <!-- Mobile Cards -->
            <div class="sm:hidden space-y-4">
              <div v-for="(debt, index) in debts" :key="index"
                   class="bg-white border border-meal-gray-light rounded-lg overflow-hidden"
                   :class="{'opacity-50': !debt.isConfirmed}">
                <div class="p-4">
                  <div class="flex flex-col mb-3">
                    <div class="flex items-center bg-red-50 p-2 rounded-lg mb-2">
                      <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                           :style="{ backgroundColor: getUserColor(debt.fromUserId) }">
                        {{ debt.from.name.charAt(0).toUpperCase() }}
                      </div>
                      <div>
                        <div class="font-medium text-sm">{{ debt.from.name }}</div>
                        <div class="text-xs text-meal-error">
                          Schuldner
                          <span v-if="!debt.isConfirmed" class="text-meal-gray ml-1">(Ausstehend)</span>
                        </div>
                      </div>
                    </div>

                    <div class="flex justify-center items-center">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-meal-error" fill="none"
                           viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                              d="M19 14l-7 7m0 0l-7-7m7 7V3"/>
                      </svg>
                      <span class="mx-2 text-xl font-bold text-meal-error">{{ formatCurrency(debt.amount) }}</span>
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-meal-error" fill="none"
                           viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                              d="M19 14l-7 7m0 0l-7-7m7 7V3"/>
                      </svg>
                    </div>

                    <div class="flex items-center bg-green-50 p-2 rounded-lg mt-2">
                      <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                           :style="{ backgroundColor: getUserColor(debt.toUserId) }">
                        {{ debt.to.name.charAt(0).toUpperCase() }}
                      </div>
                      <div>
                        <div class="font-medium text-sm">{{ debt.to.name }}</div>
                        <div class="text-xs text-meal-positive">Gläubiger</div>
                      </div>
                    </div>
                  </div>

                  <div class="flex justify-between items-center mb-3">
                    <button 
                      @click="showDebtDetails(debt)"
                      class="w-full bg-meal-light hover:bg-meal-gray-light text-meal-primary px-3 py-2 rounded-lg transition-colors duration-200 text-sm"
                    >
                      {{ debt.meals }} Mahlzeiten anzeigen
                    </button>
                  </div>
                  <button
                      v-if="debt.isConfirmed"
                      @click="markAsPaid(index)"
                      class="w-full bg-meal-primary hover:bg-meal-dark text-white px-3 py-2 rounded transition-colors duration-200 text-sm"
                  >
                    Als bezahlt markieren
                  </button>
                  <button
                      v-else-if="!debt.isConfirmed && currentUser?.id == debt.toUserId"
                      @click="markAsPaid(index)"
                      class="w-full bg-meal-accent hover:bg-meal-accent-dark text-white px-3 py-2 rounded transition-colors duration-200 text-sm"
                  >
                    Zahlung bestätigen
                  </button>
                  <div v-else class="w-full text-center text-meal-accent text-sm font-medium py-2">
                    Zahlung ist ausstehend
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Matrix-Ansicht -->
        <div v-if="activeTab === 'matrix'" class="p-4 sm:p-6">
          <!-- Matrix View (Desktop only) -->
          <div class="hidden sm:block overflow-x-auto">
            <div class="bg-meal-light p-3 mb-4 rounded-lg text-sm text-meal-gray-dark">
              <p>In der Matrix unten zeigt jede <span class="font-bold text-meal-error">rote Zelle</span> an, dass die
                Person in der Zeile der Person in der Spalte Geld schuldet.</p>
              <p>Jede <span class="font-bold text-meal-positive">grüne Zelle</span> zeigt an, dass die Person in der
                Zeile von der Person in der Spalte Geld bekommt.</p>
            </div>
            <table class="min-w-full bg-white">
              <thead>
              <tr class="bg-meal-light text-meal-dark text-left">
                <th class="py-3 px-4 rounded-tl-lg border-b-2 border-r-2 border-meal-gray-light">Benutzer ↓ schuldet →
                </th>
                <th v-for="user in users" :key="user.id"
                    class="py-3 px-4 text-center border-b-2 border-meal-gray-light">
                  <div class="flex flex-col items-center">
                    <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mb-1"
                         :style="{ backgroundColor: getUserColor(user.id) }">
                      {{ user.name.charAt(0).toUpperCase() }}
                    </div>
                    <span class="text-xs">{{ user.name }}</span>
                  </div>
                </th>
                <th class="py-3 px-4 rounded-tr-lg text-center border-b-2 border-l-2 border-meal-gray-light">Bilanz</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="fromUser in users" :key="fromUser.id"
                  class="border-b border-meal-gray-light hover:bg-meal-gray-light transition-colors duration-150">
                <td class="py-3 px-4 border-r-2 border-meal-gray-light">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                         :style="{ backgroundColor: getUserColor(fromUser.id) }">
                      {{ fromUser.name.charAt(0).toUpperCase() }}
                    </div>
                    {{ fromUser.name }}
                  </div>
                </td>
                <td v-for="toUser in users" :key="toUser.id" class="py-3 px-4 text-center">
                  <template v-if="fromUser.id !== toUser.id">
                      <span v-if="getDebtAmount(fromUser.id, toUser.id) > 0"
                            class="inline-block p-2 bg-red-50 rounded-lg font-bold text-meal-error">
                        {{ formatCurrency(getDebtAmount(fromUser.id, toUser.id)) }}
                      </span>
                    <span v-else-if="getDebtAmount(toUser.id, fromUser.id) > 0"
                          class="inline-block p-2 bg-green-50 rounded-lg font-bold text-meal-positive">
                        {{ formatCurrency(getDebtAmount(toUser.id, fromUser.id)) }}
                      </span>
                    <span v-else class="text-meal-gray">-</span>
                  </template>
                  <span v-else class="text-meal-gray">-</span>
                </td>
                <td class="py-3 px-4 text-center font-bold border-l-2 border-meal-gray-light" :class="{
                    'text-meal-error': getUserBalance(fromUser.id) < 0,
                    'text-meal-positive': getUserBalance(fromUser.id) > 0,
                    'text-meal-gray': getUserBalance(fromUser.id) === 0
                  }">
                  {{ formatCurrency(getUserBalance(fromUser.id)) }}
                </td>
              </tr>
              </tbody>
            </table>
          </div>

          <!-- Mobile Friendly Balance List -->
          <div class="sm:hidden">
            <div class="text-center mb-4 text-sm text-meal-gray-dark">
              Die Matrix ist auf dem Desktop besser sichtbar. Hier ist eine Übersicht der Benutzerbilanzen:
            </div>

            <div class="space-y-3">
              <div v-for="user in users" :key="user.id"
                   class="flex items-center justify-between p-3 bg-meal-gray-light rounded-lg">
                <div class="flex items-center">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(user.id) }">
                    {{ user.name.charAt(0).toUpperCase() }}
                  </div>
                  <span class="font-medium">{{ user.name }}</span>
                </div>
                <span class="font-bold" :class="{
                  'text-meal-error': getUserBalance(user.id) < 0,
                  'text-meal-positive': getUserBalance(user.id) > 0,
                  'text-meal-gray': getUserBalance(user.id) === 0
                }">
                  {{ formatCurrency(getUserBalance(user.id)) }}
                </span>
              </div>
            </div>

            <div class="mt-4 pt-4 border-t border-meal-gray-light">
              <h4 class="font-bold text-meal-gray-dark mb-2">Aktuelle Beträge</h4>
              <div v-if="debts.length === 0" class="text-center py-3 text-meal-gray">
                Keine offenen Beträge vorhanden.
              </div>
              <div v-else class="space-y-2">

                <div v-for="(debt, index) in debts" :key="index"
                     class="p-2 border border-meal-gray-light rounded-lg flex justify-between items-center"
                     :class="{'opacity-50': !debt.isConfirmed}">
                  <div class="flex items-center text-sm">
                    <div class="w-6 h-6 rounded-full flex items-center justify-center text-white font-bold mr-1"
                         :style="{ backgroundColor: getUserColor(debt.fromUserId) }">
                      {{ debt.from.name.charAt(0).toUpperCase() }}
                    </div>
                    <span class="mx-1">→</span>
                    <div class="w-6 h-6 rounded-full flex items-center justify-center text-white font-bold"
                         :style="{ backgroundColor: getUserColor(debt.toUserId) }">
                      {{ debt.to.name.charAt(0).toUpperCase() }}
                    </div>
                    <span v-if="!debt.isConfirmed" class="text-xs text-meal-gray ml-1">(Ausstehend)</span>
                  </div>

                  <span class="font-bold text-meal-error text-sm">{{ formatCurrency(debt.amount) }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="flex justify-center mt-4 text-xs sm:text-sm text-meal-gray flex-wrap gap-2">
            <div class="flex items-center mr-2 sm:mr-4 bg-red-50 p-1 px-2 rounded">
              <div class="w-3 h-3 bg-meal-error rounded-full mr-1"></div>
              <span class="font-medium">Schuldet Geld</span>
            </div>
            <div class="flex items-center mr-2 sm:mr-4 bg-green-50 p-1 px-2 rounded">
              <div class="w-3 h-3 bg-meal-positive rounded-full mr-1"></div>
              <span class="font-medium">Bekommt Geld</span>
            </div>
            <div class="flex items-center bg-meal-gray-light p-1 px-2 rounded">
              <div class="w-3 h-3 bg-meal-gray-light border border-meal-gray rounded-full mr-1"></div>
              <span class="font-medium">Keine Verbindung</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Zahlungen Historie (ausgelagerte Komponente) -->
      <PaymentHistory class="h-[300px] flex-shrink-0 overflow-auto" ref="paymentHistoryRef"/>
      
      <!-- Schulddetails Modal -->
      <Modal v-if="showDetailsModal" 
             :show="showDetailsModal" 
             :title="'Details für Betrag zwischen ' + (selectedDebt ? selectedDebt.from.name : '') + ' und ' + (selectedDebt ? selectedDebt.to.name : '')"
             @close="closeDetailsModal">
        <DebtDetails 
          v-if="selectedDebt" 
          :from-user-id="selectedDebt.fromUserId" 
          :to-user-id="selectedDebt.toUserId" />
      </Modal>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, ref} from "vue";
import {type Debt, mutation, type PaymentInput, query, type User} from "../graphql.ts";
import PaymentHistory from "../components/PaymentHistory.vue";
import StatCard from "../components/StatCard.vue";
import Modal from "../components/Modal.vue";
import DebtDetails from "../components/DebtDetails.vue";

// Ref für die PaymentHistory-Komponente
const paymentHistoryRef = ref();

// Zusätzliche Typen
interface DebtView extends Omit<Debt, 'amount'> {
  amount: number;
  meals: number;
  from: {
    id: number;
    name: string;
  };
  to: {
    id: number;
    name: string;
  };
}

interface GroupedDebt {
  fromUserId: number;
  toUserId: number;
  from: {
    id: number;
    name: string;
  };
  to: {
    id: number;
    name: string;
  };
  amount: number;
  meals: number;
  count: number;
  resolvesCount?: number;
}

interface PaymentSummary {
  from: string;
  to: string;
  amount: number;
  date: string;
}

// Anwendungsdaten
const users = ref<User[]>([]);
const debts = ref<DebtView[]>([]);
const activeTab = ref('list');
const currentUser = ref<User | null>(null);

// Modal State
const showDetailsModal = ref(false);
const selectedDebt = ref<DebtView | null>(null);

// Statische Daten
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];


async function getCurrentUser(): Promise<User | null> {
  const [data, err] = await query.getCurrentUser();
  if (err || !data) {
    console.error('Error getting current user:', err);
    return null;
  }

  currentUser.value = data;
  return data;
}

// Details Modal Funktionen
function showDebtDetails(debt: DebtView): void {
  selectedDebt.value = debt;
  showDetailsModal.value = true;
}

function closeDetailsModal(): void {
  showDetailsModal.value = false;
  selectedDebt.value = null;
}

// Berechnete Eigenschaften
const totalDebt = computed((): number => {
  return debts.value.reduce((sum, debt) => sum + debt.amount, 0);
});

const lastPayment = computed((): PaymentSummary => {
  return {from: 'Niemand', to: 'Niemand', amount: 0, date: new Date().toISOString()};
});

const optimizedPayments = computed((): GroupedDebt[] => {
  if (users.value.length === 0 || debts.value.length === 0) {
    return [];
  }

  // Gruppiere Beträge nach Schuldner-Gläubiger-Paar
  const groupedDebts: Record<string, GroupedDebt> = {};

  debts.value.forEach(debt => {
    const key = `${debt.fromUserId}-${debt.toUserId}`;
    if (!groupedDebts[key]) {
      groupedDebts[key] = {
        fromUserId: debt.fromUserId,
        toUserId: debt.toUserId,
        from: debt.from,
        to: debt.to,
        amount: 0,
        meals: 0,
        count: 0
      };
    }
    groupedDebts[key].amount += debt.amount;
    groupedDebts[key].meals += debt.meals || 0;
    groupedDebts[key].count++;
  });

  // In Array umwandeln und resolvesCount berechnen
  const results = Object.values(groupedDebts).map(debt => ({
    ...debt,
    resolvesCount: debt.count
  }));

  // Nach Betrag sortieren (absteigend)
  results.sort((a, b) => b.amount - a.amount);

  return results;
});

// Daten laden
async function loadData(): Promise<void> {
  await getUsers();
  await getDebts();
  await getCurrentUser()
}

async function getUsers(): Promise<void> {
  const [data, err] = await query.getUsers();
  if (err || !data) {
    console.error('Error getting users:', err);
    return;
  }
  users.value = data;
}

async function getDebts(): Promise<void> {
  const [data, err] = await query.getDebts();
  if (err || !data) {
    console.error('Error getting debts:', err);
    return;
  }

  // Transform debts to include user objects
  debts.value = (data || []).map(debt => {
    const fromUser = users.value.find(user => user.id === debt.fromUserId);
    const toUser = users.value.find(user => user.id === debt.toUserId);

    return {
      ...debt,
      amount: parseFloat(debt.amount), // String zu Zahl konvertieren
      meals: debt.mealsCount,
      from: {
        id: fromUser?.id || 0,
        name: fromUser?.name || 'Unbekannt'
      },
      to: {
        id: toUser?.id || 0,
        name: toUser?.name || 'Unbekannt'
      }
    };
  });
}

// Helferfunktionen
function formatCurrency(value: number): string {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR'
  }).format(value);
}

function formatDate(dateString: string): string {
  return new Date(dateString).toLocaleDateString('de-DE');
}

function getUserColor(userId: number): string {
  return avatarColors[userId % avatarColors.length];
}

// Aktionsfunktionen
async function markAsPaid(index: number): Promise<void> {
  const debt = debts.value[index];
  const payment: PaymentInput = {
    fromUserId: debt.fromUserId,
    toUserId: debt.toUserId,
    amount: debt.amount.toString(), // Zahl zurück zu String konvertieren für API
    description: `Bezahlungen für ${debt.meals} Mahlzeiten`
  };

  const [data, err] = await mutation.createPayment(payment);
  if (err) {
    console.error('Fehler beim Erstellen der Zahlung:', err);
    return;
  }
  await loadData();

  // Zahlungshistorie aktualisieren
  if (paymentHistoryRef.value) {
    await paymentHistoryRef.value.refreshPayments();
  }
}

function getDebtAmount(fromId: number, toId: number): number {
  const debt = debts.value.find(d => d.fromUserId === fromId && d.toUserId === toId);
  return debt ? debt.amount : 0;
}

function getUserBalance(userId: number): number {
  let balance = 0;

  // Negative Beträge (was der User schuldet)
  debts.value.forEach(debt => {
    if (debt.fromUserId === userId) {
      balance -= debt.amount;
    }
  });

  // Positive Beträge (was der User bekommt)
  debts.value.forEach(debt => {
    if (debt.toUserId === userId) {
      balance += debt.amount;
    }
  });

  return balance;
}

// Initialisierung
onMounted(() => {
  loadData();
});
</script>