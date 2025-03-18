<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <!-- Header -->
    <header class="bg-meal-primary text-white p-4 shadow-md mb-4 sm:mb-8">
      <div class="container mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 sm:h-8 sm:w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"/>
          </svg>
          <h1 class="text-xl sm:text-2xl font-header font-bold">Dashboard</h1>
        </div>
        <button @click="goBackToHomeScreen"
                class="text-white hover:text-meal-accent-light transition-colors duration-200 text-sm sm:text-base flex items-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-5 sm:w-5 mr-1" fill="none" viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
          </svg>
          <span class="hidden xs:inline">Zurück zur Startseite</span>
          <span class="xs:hidden">Zurück</span>
        </button>
      </div>
    </header>

    <div class="container mx-auto px-2 sm:px-4">
      <!-- Übersichtskarten -->
      <div class="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-4 gap-3 sm:gap-6 mb-4 sm:mb-8">
        <div class="bg-white rounded-xl shadow-meal p-3 sm:p-6">
          <div class="flex items-center justify-between mb-2 sm:mb-4">
            <h3 class="text-sm sm:text-lg font-header font-bold text-meal-gray-dark">Mahlzeiten</h3>
            <div class="p-1 sm:p-2 bg-meal-light rounded-lg">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-meal-primary" fill="none" viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
              </svg>
            </div>
          </div>
          <p class="text-xl sm:text-3xl font-bold text-meal-gray-dark mb-1">{{ stats.meals }}</p>
          <div class="flex items-center text-xs sm:text-sm">
            <span :class="stats.mealsChange >= 0 ? 'text-meal-positive' : 'text-meal-error'" class="font-medium">
              {{ stats.mealsChange >= 0 ? '+' : '' }}{{ stats.mealsChange }}%
            </span>
            <span class="text-meal-gray ml-1">seit letztem Monat</span>
          </div>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-3 sm:p-6">
          <div class="flex items-center justify-between mb-2 sm:mb-4">
            <h3 class="text-sm sm:text-lg font-header font-bold text-meal-gray-dark">Ausgaben</h3>
            <div class="p-1 sm:p-2 bg-meal-light rounded-lg">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-meal-accent" fill="none" viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
          </div>
          <p class="text-xl sm:text-3xl font-bold text-meal-gray-dark mb-1">{{ formatCurrency(stats.totalSpent) }}</p>
          <div class="flex items-center text-xs sm:text-sm">
            <span :class="stats.spentChange >= 0 ? 'text-meal-positive' : 'text-meal-error'" class="font-medium">
              {{ stats.spentChange >= 0 ? '+' : '' }}{{ stats.spentChange }}%
            </span>
            <span class="text-meal-gray ml-1">seit letztem Monat</span>
          </div>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-3 sm:p-6">
          <div class="flex items-center justify-between mb-2 sm:mb-4">
            <h3 class="text-sm sm:text-lg font-header font-bold text-meal-gray-dark">Benutzer</h3>
            <div class="p-1 sm:p-2 bg-meal-light rounded-lg">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-meal-secondary" fill="none"
                   viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
              </svg>
            </div>
          </div>
          <p class="text-xl sm:text-3xl font-bold text-meal-gray-dark mb-1">{{ stats.users }}</p>
          <div class="flex items-center text-xs sm:text-sm">
            <span class="text-meal-positive font-medium">Aktiv</span>
            <span class="text-meal-gray ml-1">in diesem Monat</span>
          </div>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-3 sm:p-6">
          <div class="flex items-center justify-between mb-2 sm:mb-4">
            <h3 class="text-sm sm:text-lg font-header font-bold text-meal-gray-dark">Offene Schulden</h3>
            <div class="p-1 sm:p-2 bg-meal-light rounded-lg">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-meal-error" fill="none" viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
            </div>
          </div>
          <p class="text-xl sm:text-3xl font-bold text-meal-gray-dark mb-1">{{ formatCurrency(stats.totalDebt) }}</p>
          <div class="flex items-center text-xs sm:text-sm">
            <span class="text-meal-error font-medium">{{ stats.activeDebts }}</span>
            <span class="text-meal-gray ml-1">aktive Schulden</span>
          </div>
        </div>
      </div>

      <!-- Diagramme -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6 mb-4 sm:mb-8">
        <!-- Ausgaben pro Monat -->
        <TransactionGraph :transactions="transactions" :users="users"/>

        <!-- Ausgaben pro Person -->
        <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-3 sm:mb-4">Ausgaben pro Person</h3>
          <div class="min-h-60 sm:h-80 flex flex-col sm:flex-row items-center justify-center">
            <!-- Kreisdiagramm mit CSS -->
            <div class="relative w-52 h-52 sm:w-64 sm:h-64 mb-4 sm:mb-0">
              <div class="absolute inset-0 rounded-full overflow-hidden">
                <div v-for="(segment, index) in userExpenseSegments" :key="index"
                     class="absolute"
                     :style="{
                       backgroundColor: getUserColor(segment.userId),
                       top: 0,
                       left: 0,
                       width: '100%',
                       height: '100%',
                       clipPath: `polygon(50% 50%, ${segment.startX}% ${segment.startY}%, ${segment.endX}% ${segment.endY}%, ${segment.extraPoints})`,
                     }">
                </div>
              </div>
              <!-- Mittleres Loch im Kreis -->
              <div
                  class="absolute w-24 h-24 sm:w-32 sm:h-32 bg-white rounded-full top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"></div>
            </div>

            <!-- Legende -->
            <div class="sm:ml-6 space-y-1 sm:space-y-2 grid grid-cols-2 sm:block gap-1">
              <div v-for="user in userExpenseData" :key="user.id" class="flex items-center">
                <div class="w-3 h-3 sm:w-4 sm:h-4 rounded mr-1 sm:mr-2" :style="{ backgroundColor: getUserColor(user.id) }"></div>
                <span class="text-xs sm:text-sm text-meal-gray-dark whitespace-nowrap">{{ user.name }} ({{
                    formatPercentage(user.percentage)
                  }})</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Aktuelle Aktivitäten und Schulden -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 sm:gap-6 mb-4 sm:mb-8">
        <!-- Letzte Aktivitäten -->
        <div class="bg-white h-[350px] sm:h-[550px] flex flex-col overflow-auto rounded-xl shadow-meal p-4 sm:p-6 lg:col-span-2">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-3 sm:mb-4">Letzte Aktivitäten</h3>

          <div class="space-y-3 sm:space-y-4">
            <div v-for="(activity, index) in recentActivities.slice(0, showAll ? recentActivities.length : 10)" :key="index"
                 class="flex items-start p-2 sm:p-3 hover:bg-meal-light rounded-lg transition-colors duration-150">
              <div :class="`bg-${getActivityIconBg(activity.type)} p-1 sm:p-2 rounded-lg mr-2 sm:mr-4`">
                <svg v-if="activity.type === 'meal'" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white"
                     fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253"/>
                </svg>
                <svg v-else-if="activity.type === 'payment'" xmlns="http://www.w3.org/2000/svg"
                     class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"/>
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none"
                     viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
                </svg>
              </div>

              <div class="flex-grow">
                <div class="flex justify-between">
                  <h4 class="font-bold text-sm sm:text-base text-meal-gray-dark">{{ getActivityTitle(activity) }}</h4>
                  <span class="text-xs sm:text-sm text-meal-gray">{{ formatDate(activity.date) }}</span>
                </div>
                <p class="text-xs sm:text-sm text-meal-gray">{{ activity.description }}</p>

                <!-- Spezifische Darstellung je nach Typ -->
                <div v-if="activity.type === 'payment'" class="mt-1 flex items-center flex-wrap text-xs sm:text-sm">
                  <div class="flex items-center mr-2">
                    <div class="w-5 h-5 sm:w-6 sm:h-6 rounded-full flex items-center justify-center text-white font-bold mr-1"
                         :style="{ backgroundColor: getUserColor(activity.fromUserId) }">
                      {{ getUserName(activity.fromUserId).charAt(0).toUpperCase() }}
                    </div>
                    <span>{{ getUserName(activity.fromUserId) }}</span>
                  </div>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 sm:h-4 sm:w-4 text-meal-gray mx-1" fill="none"
                       viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/>
                  </svg>
                  <div class="flex items-center">
                    <div class="w-5 h-5 sm:w-6 sm:h-6 rounded-full flex items-center justify-center text-white font-bold mr-1"
                         :style="{ backgroundColor: getUserColor(activity.toUserId) }">
                      {{ getUserName(activity.toUserId).charAt(0).toUpperCase() }}
                    </div>
                    <span>{{ getUserName(activity.toUserId) }}</span>
                  </div>
                  <span class="ml-auto font-bold text-meal-positive mt-1 sm:mt-0">
                    {{ formatCurrency(Number(activity.amount)) }}
                  </span>
                </div>

                <div v-if="activity.type === 'meal'" class="mt-1 flex items-center flex-wrap text-xs sm:text-sm">
                  <div class="flex items-center">
                    <div class="w-5 h-5 sm:w-6 sm:h-6 rounded-full flex items-center justify-center text-white font-bold mr-1"
                         :style="{ backgroundColor: getUserColor(activity.userId) }">
                      {{ getUserName(activity.userId).charAt(0).toUpperCase() }}
                    </div>
                    <span>{{ getUserName(activity.userId) }} hat bezahlt</span>
                  </div>
                  <span class="ml-auto font-bold text-meal-primary mt-1 sm:mt-0">
                    {{ formatCurrency(Number(activity.amount)) }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <button @click="showMoreActivities"
                  class="w-full mt-3 sm:mt-4 py-2 text-meal-primary hover:text-meal-dark text-sm sm:text-base font-medium transition-colors duration-200">
            Weitere Aktivitäten anzeigen
          </button>
        </div>

        <!-- Top Schulden -->
        <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-3 sm:mb-4">Top Schulden</h3>

          <div class="space-y-3 sm:space-y-4">
            <div v-for="(debt, index) in topDebts" :key="index"
                 class="flex items-center justify-between p-2 sm:p-3 border-b border-meal-gray-light">
              <div class="flex items-center">
                <div class="w-8 h-8 sm:w-10 sm:h-10 rounded-full flex items-center justify-center text-white font-bold mr-2 sm:mr-3"
                     :style="{ backgroundColor: getUserColor(debt.fromId) }">
                  {{ debt.fromName.charAt(0).toUpperCase() }}
                </div>
                <span class="text-xs sm:text-sm">schuldet</span>
              </div>
              <div class="flex items-center">
                <span class="font-bold text-xs sm:text-sm text-meal-error mr-2 sm:mr-3">{{ formatCurrency(debt.amount) }}</span>
                <div class="w-8 h-8 sm:w-10 sm:h-10 rounded-full flex items-center justify-center text-white font-bold"
                     :style="{ backgroundColor: getUserColor(debt.toId) }">
                  {{ debt.toName.charAt(0).toUpperCase() }}
                </div>
              </div>
            </div>
          </div>

          <button @click="goToDebts"
                  class="w-full mt-3 sm:mt-4 py-2 text-meal-primary hover:text-meal-dark text-sm sm:text-base font-medium transition-colors duration-200">
            Alle Schulden anzeigen
          </button>
        </div>
      </div>

      <!-- Schnellzugriff -->
      <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6 mb-4 sm:mb-8">
        <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-3 sm:mb-4">Schnellzugriff</h3>

        <div class="grid grid-cols-1 sm:grid-cols-3 gap-3 sm:gap-4">
          <button
              @click="goToMeal"
              class="flex items-center p-3 sm:p-4 rounded-lg border border-meal-gray-light hover:bg-meal-light transition-colors duration-200"
          >
            <div class="p-1 sm:p-2 bg-meal-primary rounded-full mr-2 sm:mr-3">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none" viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253"/>
              </svg>
            </div>
            <div class="text-left">
              <h4 class="font-bold text-sm sm:text-base text-meal-gray-dark">Neue Mahlzeit</h4>
              <p class="text-xs sm:text-sm text-meal-gray">Mahlzeit erfassen</p>
            </div>
          </button>

          <button
              @click="goToPayments"
              class="flex items-center p-3 sm:p-4 rounded-lg border border-meal-gray-light hover:bg-meal-light transition-colors duration-200"
          >
            <div class="p-1 sm:p-2 bg-meal-accent rounded-full mr-2 sm:mr-3">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none" viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
            </div>
            <div class="text-left">
              <h4 class="font-bold text-sm sm:text-base text-meal-gray-dark">Neue Zahlung</h4>
              <p class="text-xs sm:text-sm text-meal-gray">Zahlung erfassen</p>
            </div>
          </button>

          <button
              @click="goToUser"
              class="flex items-center p-3 sm:p-4 rounded-lg border border-meal-gray-light hover:bg-meal-light transition-colors duration-200"
          >
            <div class="p-1 sm:p-2 bg-meal-secondary rounded-full mr-2 sm:mr-3">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none" viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
              </svg>
            </div>
            <div class="text-left">
              <h4 class="font-bold text-sm sm:text-base text-meal-gray-dark">Benutzer</h4>
              <p class="text-xs sm:text-sm text-meal-gray">Benutzer verwalten</p>
            </div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, ref} from "vue";
import router from "../router";
import TransactionGraph from "../components/TransactionGraph.vue";
import {query} from "../graphql.ts";

// Beispiel-Benutzer

const topDebts = ref([]);

// Ausgaben pro Person
const userExpenseData = ref([]);

const users = ref([]);

async function getUsers() {
  const [data, err] = await query.getUsers();
  users.value = data;
  if (err) {
    console.error('Error getting users:', err);
    return;
  }
}

async function getDashboard() {
  const [data, err] = await query.getDashboard();
  if (err) {
    console.error('Error getting users:', err);
    return;
  }
  stats.value = {
    meals: data?.totalMeals,
    //mealsChange: data.mealsChange,
    totalSpent: data?.totalCredits,
    // spentChange: data.spentChange,
    users: data?.totalUsers,
    totalDebt: data?.totalDebts,
    //activeDebts: data.activeDebts
  };
  data?.tobDebtsPerUser.forEach(debt => {
    let percentage = debt.amount / data?.totalDebts * 100;

    topDebts.value.push({
      fromId: debt.userId,
      fromName: users.value.find(u => u.id === debt.userId).name,
      toName: "",
      amount: debt.amount,
      percentage: percentage
    });
  });

  data?.totalCreditsPerUser.forEach(user => {
    let percentage = user.amount / data?.totalCredits * 100;
    userExpenseData.value.push({
      id: user.userId,
      name: users.value.find(u => u.id === user.userId).name,
      amount: user.amount,
      percentage: percentage
    });
  });

  console.log("userExpenseData",userExpenseData.value);

}

// Statistiken
const stats = ref({});
const showAll = ref(false);

// Monatliche Ausgabendaten
const monthlyExpenseData = ref([
  {label: 'Jan', amount: 120.50},
  {label: 'Feb', amount: 85.75},
  {label: 'Mär', amount: 135.20},
  {label: 'Apr', amount: 92.30},
  {label: 'Mai', amount: 105.60},
  {label: 'Jun', amount: 75.40},
  {label: 'Jul', amount: 110.25},
  {label: 'Aug', amount: 180.50},
  {label: 'Sep', amount: 130.80},
  {label: 'Okt', amount: 142.75},
  {label: 'Nov', amount: 95.30},
  {label: 'Dez', amount: 0}
]);

// Letzte Aktivitäten (Daten aus der DB, z. B. über GraphQL)
const recentActivities = ref([]);

async function getActivities() {
  const [data, err] = await query.getActivities();
  recentActivities.value = data;
  if (err) {
    console.error('Error getting activities:', err);
    return;
  }
}

// Top Schulden


// Berechnung für das Kreisdiagramm
const userExpenseSegments = computed(() => {
  const segments = [];
  let cumulativePercentage = 0;

  userExpenseData.value.forEach(user => {
    const startAngle = cumulativePercentage * 3.6;
    cumulativePercentage += user.percentage;
    const endAngle = cumulativePercentage * 3.6;

    const startRad = (startAngle - 90) * Math.PI / 180;
    const endRad = (endAngle - 90) * Math.PI / 180;

    const startX = 50 + 50 * Math.cos(startRad);
    const startY = 50 + 50 * Math.sin(startRad);
    const endX = 50 + 50 * Math.cos(endRad);
    const endY = 50 + 50 * Math.sin(endRad);

    let extraPoints = '';
    if (endAngle - startAngle > 180) {
      const midRad = (startRad + endRad) / 2;
      const midX = 50 + 50 * Math.cos(midRad);
      const midY = 50 + 50 * Math.sin(midRad);
      extraPoints = `${midX}% ${midY}%,`;
    }

    segments.push({
      userId: user.id,
      startX,
      startY,
      endX,
      endY,
      extraPoints
    });
  });

  return segments;
});

// Berechnete Eigenschaften
const highestMonthlyExpense = computed(() => {
  return Math.max(...monthlyExpenseData.value.map(month => month.amount));
});

// Farben für Benutzer-Avatare
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

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

function formatPercentage(value) {
  return `${value.toFixed(0)}%`;
}

function getUserColor(userId) {
  return avatarColors[userId % avatarColors.length];
}

function getUserName(userId) {
  const user = users.value.find(u => u.id === userId);
  return user ? user.name : `User ${userId}`;
}

function getActivityTitle(activity) {
  if (activity.type === 'meal') return "Mahlzeit erfasst";
  if (activity.type === 'payment') return "Zahlung erfasst";
  if (activity.type === 'user') return "Benutzer erstellt";
  return "";
}

function getActivityIconBg(type) {
  if (type === 'meal') return 'meal-primary';
  if (type === 'payment') return 'meal-accent';
  if (type === 'user') return 'meal-secondary';
  return 'meal-light';
}

function showMoreActivities() {
  showAll.value = true;
}

function goBackToHomeScreen() {
  router.push('/homefeed');
}

function goToUser() {
  router.push('/user');
}

function goToPayments() {
  router.push('/payments');
}

function goToMeal() {
  router.push('/meal');
}

function goToDebts() {
  router.push('/debts');
}


// Beim Laden der Komponente
onMounted(async () => {
  await getActivities();
  await getUsers();
  await getDashboard();
});
</script>

<style>
/* Add this to your CSS to create an extra small breakpoint */
@media (min-width: 480px) {
  .xs\:inline {
    display: inline;
  }
  .xs\:hidden {
    display: none;
  }
}
</style>