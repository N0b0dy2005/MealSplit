<template>
  <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
    <h2 class="text-xl sm:text-2xl font-header font-bold text-meal-gray-dark mb-4 sm:mb-6">Letzte Zahlungen</h2>

    <div v-if="recentPayments.length === 0" class="text-center py-6 sm:py-8 text-meal-gray">
      <p>Noch keine Zahlungen erfasst.</p>
    </div>

    <div v-else class="space-y-3 sm:space-y-4">
      <!-- Desktop Payment History -->
      <div v-for="payment in recentPayments" :key="payment.id"
           class="hidden sm:flex items-center justify-between p-3 border-b border-meal-gray-light">
        <div class="flex items-center">
          <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold mr-3"
               :style="{ backgroundColor: getUserColor(payment.from.id) }">
            {{ payment.from.name.charAt(0).toUpperCase() }}
          </div>
          <div>
            <span class="font-bold">{{ payment.from.name }}</span>
            <span class="mx-2">→</span>
            <span class="font-bold">{{ payment.to.name }}</span>
            <p class="text-sm text-meal-gray">{{ formatDate(payment.date) }}</p>
          </div>
        </div>
        <div class="flex items-center">
          <span class="font-bold" :class="payment.isConfirmed ? 'text-meal-positive' : 'text-meal-error'">
            {{ formatCurrency(payment.amount) }}
          </span>
          <span
              class="ml-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
              :class="payment.isConfirmed 
                ? 'bg-meal-positive text-white' 
                : 'bg-yellow-500 text-white'"
          >
            {{ payment.isConfirmed ? 'Bezahlt' : 'Ausstehend' }}
          </span>
        </div>
      </div>

      <!-- Mobile Payment History -->
      <div v-for="payment in recentPayments" :key="payment.id"
           class="flex sm:hidden flex-col p-3 border-b border-meal-gray-light">
        <div class="flex items-center mb-2">
          <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
               :style="{ backgroundColor: getUserColor(payment.from.id) }">
            {{ payment.from.name.charAt(0).toUpperCase() }}
          </div>
          <div>
            <div class="flex items-center">
              <span class="font-medium text-sm">{{ payment.from.name }}</span>
              <span class="mx-1 text-meal-gray">→</span>
              <span class="font-medium text-sm">{{ payment.to.name }}</span>
            </div>
            <p class="text-xs text-meal-gray">{{ formatDate(payment.date) }}</p>
          </div>
        </div>
        <div class="flex items-center justify-end">
          <span class="font-bold text-sm" :class="payment.isConfirmed ? 'text-meal-positive' : 'text-meal-error'">
            {{ formatCurrency(payment.amount) }}
          </span>
          <span
              class="ml-2 inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium"
              :class="payment.isConfirmed 
                ? 'bg-meal-positive text-white' 
                : 'bg-yellow-500 text-white'"
          >
            {{ payment.isConfirmed ? 'Bezahlt' : 'Ausstehend' }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { query, type Payment, type User } from '../graphql.ts';

// Typdefinition für die aufbereitete Zahlung
interface PaymentView {
  id: number;
  fromUserId: number;
  toUserId: number;
  amount: number;
  date: string;
  description: string;
  isConfirmed: boolean;
  from: {
    id: number;
    name: string;
  };
  to: {
    id: number;
    name: string;
  };
}

// Zustandsvariablen
const recentPayments = ref<PaymentView[]>([]);
const users = ref<User[]>([]);

// Avatar-Farben
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

// Öffentliche Methode zum Aktualisieren der Zahlungen
const refreshPayments = async (): Promise<void> => {
  await loadData();
};

// Daten laden
async function loadData(): Promise<void> {
  await getUsers();
  await getRecentPayments();
}

async function getUsers(): Promise<void> {
  const [data, err] = await query.getUsers();
  if (err || !data) {
    console.error('Fehler beim Laden der Benutzer:', err);
    return;
  }
  users.value = data;
}

async function getRecentPayments(): Promise<void> {
  const [data, err] = await query.getPayments();
  if (err || !data) {
    console.error('Fehler beim Laden der Zahlungen:', err);
    return;
  }
  
  // Zahlungen mit Benutzerdaten anreichern
  recentPayments.value = (data || []).map(payment => {
    const fromUser = users.value.find(user => user.id === payment.fromUserId);
    const toUser = users.value.find(user => user.id === payment.toUserId);

    return {
      ...payment,
      amount: parseFloat(payment.amount),
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

// Komponente initialisieren
onMounted(() => {
  loadData();
});

// Exportiere die Methode für die übergeordnete Komponente
defineExpose({
  refreshPayments
});
</script> 