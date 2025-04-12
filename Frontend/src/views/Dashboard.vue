<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <div class="container mx-auto px-2 sm:px-4 mt-4 sm:mt-8">
      <!-- Übersichtskarten -->
      <div class="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-4 gap-3 sm:gap-6 mb-4 sm:mb-8">
        <StatCard 
          title="Mahlzeiten" 
          :value="stats.meals" 
          :change-value="stats.mealsChange" 
          caption="seit letztem Monat"
          show-icon
          icon-color="text-meal-primary"
          iconPath="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
        />

        <StatCard 
          title="Ausgaben" 
          :value="stats.totalSpent" 
          :change-value="stats.spentChange" 
          caption="seit letztem Monat"
          is-currency
          show-icon
          icon-color="text-meal-accent"
          iconPath="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />

        <StatCard 
          title="Benutzer" 
          :value="stats.users" 
          caption="Aktiv in diesem Monat"
          :show-change-indicator="false"
          show-icon
          icon-color="text-meal-secondary"
          iconPath="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"
        />

        <StatCard 
          title="Offene Beträge" 
          :value="stats.totalDebt" 
          caption="aktive Beträge"
          :change-value="stats.activeDebts"
          :change-is-percentage="false"
          :show-change-indicator="false"
          is-currency
          show-icon
          icon-color="text-meal-error"
          iconPath="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"
        />
      </div>

      <!-- Diagramme -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6 mb-4 sm:mb-8">
        <!-- Ausgaben pro Monat -->
        <TransactionGraph :transactions="transactions" :users="users"/>

        <!-- Umschaltbares Diagramm mit Tabs -->
        <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6 h-full flex flex-col">
          <!-- Mahlzeiten-Diagramm ohne Tabs -->
          <div  class="chart-container flex-grow py-2" style="min-height: 380px;">
            <MealsPieChart 
              :user-expense-data="userExpenseData" 
              :get-user-color="getUserColor" 
              :format-percentage="formatPercentage"
              class="h-full" 
            />
          </div>
        </div>
      </div>

      <!-- Aktuelle Aktivitäten und Beträge -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 sm:gap-6 mb-4 sm:mb-8">
        <!-- Letzte Aktivitäten -->
        <div
            class="bg-white h-[350px] sm:h-[550px] flex flex-col overflow-auto rounded-xl shadow-meal p-4 sm:p-6 lg:col-span-2">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-3 sm:mb-4">Letzte
            Aktivitäten</h3>

          <div class="space-y-3 sm:space-y-4">
            <div v-for="(activity, index) in recentActivities.slice(0, showAll ? recentActivities.length : 10)"
                 :key="index"
                 class="flex items-start p-2 sm:p-3 hover:bg-meal-light rounded-lg transition-colors duration-150">
              <div :class="`bg-${getActivityIconBg(activity.type)} p-1 sm:p-2 rounded-lg mr-2 sm:mr-4`">
                <svg v-if="activity.type === 'meal'" xmlns="http://www.w3.org/2000/svg"
                     class="h-4 w-4 sm:h-6 sm:w-6 text-white"
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
                    <div
                        class="w-5 h-5 sm:w-6 sm:h-6 rounded-full flex items-center justify-center text-white font-bold mr-1"
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
                    <div
                        class="w-5 h-5 sm:w-6 sm:h-6 rounded-full flex items-center justify-center text-white font-bold mr-1"
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
                    <div
                        class="w-5 h-5 sm:w-6 sm:h-6 rounded-full flex items-center justify-center text-white font-bold mr-1"
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

          <!--
          <button @click="showMoreActivities"
                  class="w-full mt-3 sm:mt-4 py-2 text-meal-primary hover:text-meal-dark text-sm sm:text-base font-medium transition-colors duration-200">
            Weitere Aktivitäten anzeigen
          </button>
          -->
        </div>

        <!-- Top Beträge -->
        <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-3 sm:mb-4">Top Beträge</h3>

          <div class="space-y-3 sm:space-y-4">
            <div v-for="(debt, index) in topDebts" :key="index"
                 class="flex items-center justify-between p-2 sm:p-3 border-b border-meal-gray-light">
              <div class="flex items-center">
                <div
                    class="w-8 h-8 sm:w-10 sm:h-10 rounded-full flex items-center justify-center text-white font-bold mr-2 sm:mr-3"
                    :style="{ backgroundColor: getUserColor(debt.fromId) }">
                  {{ debt.fromName.charAt(0).toUpperCase() }}
                </div>
                <span class="text-xs sm:text-sm">schuldet</span>
              </div>
              <div class="flex items-center">
                <span class="font-bold text-xs sm:text-sm text-meal-error mr-2 sm:mr-3">{{
                    formatCurrency(debt.amount)
                  }}</span>
                <div class="w-8 h-8 sm:w-10 sm:h-10 rounded-full flex items-center justify-center text-white font-bold"
                     :style="{ backgroundColor: getUserColor(debt.toId) }">
                  {{ debt.toName.charAt(0).toUpperCase() }}
                </div>
              </div>
            </div>
          </div>

          <button @click="goToDebts"
                  class="w-full mt-3 sm:mt-4 py-2 text-meal-primary hover:text-meal-dark text-sm sm:text-base font-medium transition-colors duration-200">
            Alle Beträge anzeigen
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
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none"
                   viewBox="0 0 24 24"
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
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none"
                   viewBox="0 0 24 24"
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
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6 text-white" fill="none"
                   viewBox="0 0 24 24"
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
import MealsPieChart from "../components/MealsPieChart.vue";
import StatCard from "../components/StatCard.vue";
import {query} from "../graphql.ts";

// Ausgewähltes Diagramm

// Beispiel-Benutzer

const topDebts = ref([]);

// Ausgaben pro Person
const userExpenseData = ref([]);

const users = ref([]);
const transactions = ref([]); // Für TransactionGraph
const meals = ref([]); // Für die Mahlzeiten-Analyse

async function getUsers() {
  const [data, err] = await query.getUsers();
  users.value = data;
  if (err) {
    console.error('Error getting users:', err);
    return;
  }
}

async function getMeals() {
  const [data, err] = await query.getMeals();
  if (err) {
    console.error('Error getting meals:', err);
    return;
  }
  meals.value = data;
}

async function getDashboard() {
  const [data, err] = await query.getDashboard();
  if (err) {
    console.error('Error getting dashboard:', err);
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

  // Berechne die Anzahl der Mahlzeiten pro Benutzer basierend auf den tatsächlichen Mahlzeit-Daten
  const mealCountsPerUser = new Map();
  
  // Initialisiere die Map mit allen Benutzern
  users.value.forEach(user => {
    mealCountsPerUser.set(user.id, 0);
  });
  
  // Zähle die Mahlzeiten pro Benutzer
  meals.value.forEach(meal => {
    mealCountsPerUser.set(meal.userId, (mealCountsPerUser.get(meal.userId) || 0) + 1);
  });
  
  // Berechne die Gesamtzahl der Mahlzeiten
  const totalMeals = Array.from(mealCountsPerUser.values()).reduce((sum, count) => sum + count, 0);
  
  // Erstelle userExpenseData mit den Mahlzeitendaten
  userExpenseData.value = [];
  mealCountsPerUser.forEach((mealCount, userId) => {
    // Überspringe Benutzer ohne Mahlzeiten
    if (mealCount === 0) return;
    
    // Berechne den Prozentsatz der Mahlzeiten
    const percentage = (mealCount / totalMeals) * 100;
    
    userExpenseData.value.push({
      id: userId,
      name: users.value.find(u => u.id === userId)?.name || `User ${userId}`,
      amount: mealCount, // Anzahl der Mahlzeiten
      percentage: percentage
    });
  });

  console.log("userExpenseData", userExpenseData.value);
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
const userExpenseSegments = computed(() => {
  const segments = [];
  let cumulativePercentage = 0;

  userExpenseData.value.forEach(user => {
    // Umrechnung der Prozentwerte in Winkel (360° entsprechen 100%)
    const startAngle = cumulativePercentage * 3.6;
    cumulativePercentage += user.percentage;
    const endAngle = cumulativePercentage * 3.6;

    // Umrechnung in Radianten, Start bei -90° (oben)
    const startRad = (startAngle - 90) * Math.PI / 180;
    const endRad = (endAngle - 90) * Math.PI / 180;

    // Bestimme die Start- und Endpunkte des Kreisbogens (Kreisradius 50%)
    const startX = 50 + 50 * Math.cos(startRad);
    const startY = 50 + 50 * Math.sin(startRad);
    const endX = 50 + 50 * Math.cos(endRad);
    const endY = 50 + 50 * Math.sin(endRad);

    // Beginne das Polygon beim Mittelpunkt und dem Startpunkt
    let clipPath = `polygon(50% 50%, ${startX}% ${startY}%`;

    // Verwende einen kleineren Winkel-Schritt (z. B. 10°) für alle Segmente,
    // um einen glatteren Bogen zu erzielen
    const stepAngle = 10;
    if (endAngle - startAngle > stepAngle) {
      const steps = Math.ceil((endAngle - startAngle) / stepAngle);
      for (let i = 1; i < steps; i++) {
        const angle = startAngle + ((endAngle - startAngle) * (i / steps));
        const rad = (angle - 90) * Math.PI / 180;
        const x = 50 + 50 * Math.cos(rad);
        const y = 50 + 50 * Math.sin(rad);
        clipPath += `, ${x}% ${y}%`;
      }
    }

    // Füge den Endpunkt hinzu und schließe das Polygon
    clipPath += `, ${endX}% ${endY}%)`;

    // Ausgabe zum Debuggen
    console.log(`Segment für User ${user.id} (${user.name}): `, {
      userId: user.id,
      percentage: user.percentage,
      color: getUserColor(user.id),
      clipPath
    });

    segments.push({
      userId: user.id,
      clipPath,
      color: getUserColor(user.id)
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
  await getMeals(); // Hole zuerst die Mahlzeiten
  await getDashboard(); // Dann berechne das Dashboard
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