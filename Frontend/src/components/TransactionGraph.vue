<template>
  <div class="bg-white rounded-xl shadow-meal p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-header font-bold text-meal-gray-dark">Transaktionsübersicht</h3>

      <div class="flex items-center space-x-2">
        <!-- Toggle zwischen Jahres- und Monatsansicht -->
        <button
            @click="viewType = 'year'"
            :class="[
            'px-3 py-1 text-sm rounded-md transition-colors',
            viewType === 'year'
              ? 'bg-meal-primary text-white'
              : 'bg-meal-light text-meal-gray-dark hover:bg-meal-gray-light'
          ]"
        >
          Jahresübersicht
        </button>
        <button
            @click="viewType = 'month'"
            :class="[
            'px-3 py-1 text-sm rounded-md transition-colors',
            viewType === 'month'
              ? 'bg-meal-primary text-white'
              : 'bg-meal-light text-meal-gray-dark hover:bg-meal-gray-light'
          ]"
        >
          Monatsansicht
        </button>

        <!-- Chart-Typen nur für Jahresansicht -->
        <div v-if="viewType === 'year'" class="flex items-center ml-2 space-x-1">
          <button
              @click="viewMode = 'line'"
              :class="[
              'p-1 rounded transition-colors',
              viewMode === 'line' ? 'bg-meal-primary text-white' : 'text-meal-gray hover:bg-meal-light'
            ]"
              title="Linien-Ansicht"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z" />
            </svg>
          </button>
          <button
              @click="viewMode = 'bar'"
              :class="[
              'p-1 rounded transition-colors',
              viewMode === 'bar' ? 'bg-meal-primary text-white' : 'text-meal-gray hover:bg-meal-light'
            ]"
              title="Balken-Ansicht"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </button>
          <button
              @click="viewMode = 'area'"
              :class="[
              'p-1 rounded transition-colors',
              viewMode === 'area' ? 'bg-meal-primary text-white' : 'text-meal-gray hover:bg-meal-light'
            ]"
              title="Bereichs-Ansicht"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </button>
        </div>

        <!-- Monatsauswahl nur für Monatsansicht -->
        <div v-if="viewType === 'month'" class="flex items-center">
          <button @click="prevMonth" class="p-1 text-meal-gray hover:text-meal-primary transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>

          <div class="px-3 font-medium text-meal-gray-dark">
            {{ monthNames[selectedMonth] }} {{ currentYear }}
          </div>

          <button @click="nextMonth" class="p-1 text-meal-gray hover:text-meal-primary transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <div class="relative h-80">
      <!-- JAHRESANSICHT -->
      <div v-if="viewType === 'year'" class="h-full">
        <!-- Y-Achse mit Beträgen -->
        <div class="absolute top-0 left-0 h-full flex flex-col justify-between pr-2 text-right">
          <div v-for="(value, index) in yAxisLabels" :key="index" class="text-xs text-meal-gray">
            {{ formatCurrency(value) }}
          </div>
        </div>

        <!-- Hauptgraph-Bereich -->
        <div class="ml-16 h-full relative">
          <!-- Horizontale Hilfslinien -->
          <div v-for="(value, index) in yAxisLabels" :key="index"
               class="absolute w-full border-t border-dashed border-meal-gray-light"
               :style="{ bottom: `${(index / (yAxisLabels.length - 1)) * 100}%` }">
          </div>

          <!-- Graph-Container mit Clipping für Animation -->
          <div class="h-full w-full overflow-hidden">
            <div class="h-full w-full relative">
              <!-- Balken-Chart -->
              <div v-if="viewMode === 'bar'" class="h-full w-full flex items-end justify-between">
                <div v-for="(point, index) in monthlyData" :key="index"
                     class="group relative h-full flex-1 mx-1 flex flex-col justify-end">
                  <div class="w-full bg-gradient-to-t from-meal-primary to-meal-accent rounded-t-md transition-all duration-500 ease-out"
                       :style="{
                        height: `${(point.value / maxValue) * 100}%`,
                        opacity: hoveredMonth === null || hoveredMonth === index ? '1' : '0.3'
                      }"
                       @mouseenter="hoveredMonth = index"
                       @mouseleave="hoveredMonth = null">
                  </div>
                  <!-- Tooltip -->
                  <div v-if="hoveredMonth === index"
                       class="absolute bottom-full mb-2 left-1/2 transform -translate-x-1/2
                              bg-meal-gray-dark text-white text-xs py-1 px-2 rounded z-10 whitespace-nowrap">
                    {{ formatCurrency(point.value) }}
                  </div>
                </div>
              </div>

              <!-- Linien-Chart -->
              <svg v-else-if="viewMode === 'line'" class="h-full w-full">
                <!-- Linie -->
                <path :d="linePath"
                      fill="none"
                      stroke="url(#lineGradient)"
                      stroke-width="3"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      class="transition-all duration-700 ease-out" />

                <!-- Punkte -->
                <g v-for="(point, index) in graphPoints" :key="index">
                  <circle
                      :cx="point.x"
                      :cy="point.y"
                      r="5"
                      class="fill-white stroke-meal-primary stroke-2 transition-all duration-300 hover:r-6"
                      @mouseenter="hoveredMonth = index"
                      @mouseleave="hoveredMonth = null" />

                  <!-- Highlight-Punkt, der bei Hover erscheint -->
                  <circle
                      v-if="hoveredMonth === index"
                      :cx="point.x"
                      :cy="point.y"
                      r="8"
                      class="fill-meal-primary opacity-30 animate-ping" />

                  <!-- Tooltip -->
                  <g v-if="hoveredMonth === index">
                    <rect
                        :x="point.x - 35"
                        :y="point.y - 35"
                        width="70"
                        height="25"
                        rx="4"
                        class="fill-meal-gray-dark" />
                    <text
                        :x="point.x"
                        :y="point.y - 20"
                        text-anchor="middle"
                        class="fill-white text-xs">
                      {{ formatCurrency(monthlyData[index].value) }}
                    </text>
                  </g>
                </g>

                <!-- Gradient für die Linie -->
                <defs>
                  <linearGradient id="lineGradient" x1="0%" y1="0%" x2="100%" y2="0%">
                    <stop offset="0%" stop-color="#4F46E5" />
                    <stop offset="100%" stop-color="#EC4899" />
                  </linearGradient>
                </defs>
              </svg>

              <!-- Flächen-Chart -->
              <svg v-else class="h-full w-full">
                <!-- Gefüllter Bereich unter der Linie -->
                <path :d="`${linePath} L ${width},${height} L 0,${height} Z`"
                      fill="url(#areaGradient)"
                      class="transition-all duration-700 ease-out opacity-70" />

                <!-- Linie oben -->
                <path :d="linePath"
                      fill="none"
                      stroke="url(#lineGradient)"
                      stroke-width="3"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      class="transition-all duration-700 ease-out" />

                <!-- Punkte -->
                <g v-for="(point, index) in graphPoints" :key="index">
                  <circle
                      :cx="point.x"
                      :cy="point.y"
                      r="4"
                      class="fill-white stroke-meal-primary stroke-2 transition-all duration-300 hover:r-6"
                      @mouseenter="hoveredMonth = index"
                      @mouseleave="hoveredMonth = null" />

                  <!-- Highlight-Punkt -->
                  <circle
                      v-if="hoveredMonth === index"
                      :cx="point.x"
                      :cy="point.y"
                      r="8"
                      class="fill-meal-primary opacity-30 animate-ping" />

                  <!-- Tooltip -->
                  <g v-if="hoveredMonth === index">
                    <rect
                        :x="point.x - 35"
                        :y="point.y - 35"
                        width="70"
                        height="25"
                        rx="4"
                        class="fill-meal-gray-dark" />
                    <text
                        :x="point.x"
                        :y="point.y - 20"
                        text-anchor="middle"
                        class="fill-white text-xs">
                      {{ formatCurrency(monthlyData[index].value) }}
                    </text>
                  </g>
                </g>

                <!-- Gradienten -->
                <defs>
                  <linearGradient id="lineGradient" x1="0%" y1="0%" x2="100%" y2="0%">
                    <stop offset="0%" stop-color="#4F46E5" />
                    <stop offset="100%" stop-color="#EC4899" />
                  </linearGradient>
                  <linearGradient id="areaGradient" x1="0%" y1="0%" x2="0%" y2="100%">
                    <stop offset="0%" stop-color="#4F46E5" stop-opacity="0.8" />
                    <stop offset="100%" stop-color="#4F46E5" stop-opacity="0.1" />
                  </linearGradient>
                </defs>
              </svg>
            </div>
          </div>

          <!-- X-Achse mit Monaten -->
          <div class="absolute bottom-0 left-0 w-full flex justify-between mt-2">
            <div v-for="(month, index) in monthlyData" :key="index"
                 :class="[
                  'text-xs text-center transition-colors py-1 px-1 cursor-pointer',
                  hoveredMonth === index ? 'text-meal-primary font-bold' : 'text-meal-gray'
                ]"
                 @mouseenter="hoveredMonth = index"
                 @mouseleave="hoveredMonth = null"
                 @click="switchToMonth(index)">
              {{ month.label }}
            </div>
          </div>
        </div>

        <!-- Highlights der letzten Transaktion -->
        <div class="absolute top-2 right-2 bg-meal-light bg-opacity-80 backdrop-blur-sm rounded-lg p-3 shadow-sm max-w-xs hidden md:block">
          <div class="text-sm font-medium text-meal-gray-dark mb-1">
            Letzte Transaktion
          </div>
          <div class="flex items-center">
            <div class="w-7 h-7 rounded-full flex items-center justify-center text-white text-xs font-bold mr-2"
                 :style="{ backgroundColor: getUserColor(lastTransaction.from_user_id) }">
              {{ getUserInitial(lastTransaction.from_user_id) }}
            </div>
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-meal-gray mx-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
            </svg>
            <div class="w-7 h-7 rounded-full flex items-center justify-center text-white text-xs font-bold mr-2"
                 :style="{ backgroundColor: getUserColor(lastTransaction.to_user_id) }">
              {{ getUserInitial(lastTransaction.to_user_id) }}
            </div>
            <span class="font-bold text-meal-primary ml-auto">
              {{ formatCurrency(lastTransaction.amount) }}
            </span>
          </div>
          <div class="text-xs text-meal-gray mt-1">
            {{ formatDate(lastTransaction.created_at) }}
          </div>
        </div>
      </div>

      <!-- MONATSANSICHT -->
      <div v-else class="h-full">
        <!-- Wenn keine Daten für den Monat vorhanden sind -->
        <div v-if="monthlyTransactions.length === 0" class="h-full flex flex-col items-center justify-center text-meal-gray">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mb-2 opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          <p class="mb-1">Keine Transaktionen im {{ monthNames[selectedMonth] }}.</p>
          <button @click="selectedMonth = currentMonth" class="text-meal-primary hover:text-meal-dark transition-colors">
            Zum aktuellen Monat wechseln
          </button>
        </div>

        <!-- Wenn Daten für den Monat vorhanden sind -->
        <div v-else class="h-full">
          <!-- Tagesansicht - Kalenderstil -->
          <div class="grid grid-cols-7 gap-1 h-full">
            <!-- Tagesüberschriften Mo-So -->
            <div v-for="day in ['Mo', 'Di', 'Mi', 'Do', 'Fr', 'Sa', 'So']" :key="day"
                 class="text-xs text-center py-1 text-meal-gray font-medium">
              {{ day }}
            </div>

            <!-- Leere Tage vor dem ersten Tag des Monats -->
            <div v-for="_ in firstDayOffset" :key="'empty-' + _" class="rounded-md bg-meal-light opacity-50"></div>

            <!-- Tage des Monats -->
            <div v-for="day in daysInMonth" :key="day"
                 :class="[
                  'relative rounded-md border transition-all p-1 flex flex-col',
                  hasTransactionOnDay(day) ? 'border-meal-primary bg-meal-light' : 'border-meal-gray-light hover:border-meal-gray',
                  isToday(day) ? 'ring-1 ring-meal-primary' : ''
                ]"
                 @click="selectDay(day)">
              <!-- Tagesnummer -->
              <div :class="[
                  'text-xs font-medium self-end mb-1',
                  isToday(day) ? 'text-meal-primary' : 'text-meal-gray-dark'
                ]">
                {{ day }}
              </div>

              <!-- Transaktionsanzeige, wenn vorhanden -->
              <div v-if="getTransactionsForDay(day).length > 0" class="flex-grow flex flex-col justify-end">
                <div v-for="(tx, idx) in getTransactionsForDay(day).slice(0, 2)" :key="idx"
                     class="flex items-center justify-between p-0.5 text-xxs rounded mb-0.5"
                     :class="selectedDay === day ? 'bg-meal-primary text-white' : 'bg-white'">
                  <div class="flex items-center">
                    <div class="w-3 h-3 rounded-full flex items-center justify-center"
                         :style="{ backgroundColor: getUserColor(tx.from_user_id) }"></div>
                    <div class="w-1"></div>
                    <div class="w-3 h-3 rounded-full flex items-center justify-center"
                         :style="{ backgroundColor: getUserColor(tx.to_user_id) }"></div>
                  </div>
                  <span>{{ formatShortCurrency(tx.amount) }}</span>
                </div>

                <!-- Mehr Transaktionen als angezeigt -->
                <div v-if="getTransactionsForDay(day).length > 2"
                     :class="[
                      'text-xxs text-center rounded-sm py-0.5',
                      selectedDay === day ? 'bg-meal-primary-dark text-white' : 'bg-meal-gray-light text-meal-gray'
                    ]">
                  +{{ getTransactionsForDay(day).length - 2 }} mehr
                </div>
              </div>
            </div>
          </div>

          <!-- Detailansicht des ausgewählten Tages -->
          <div v-if="selectedDay && getTransactionsForDay(selectedDay).length > 0"
               class="absolute inset-0 bg-white bg-opacity-95 backdrop-blur-sm rounded-lg p-4 flex flex-col transform transition-transform duration-300"
               :class="showDayDetails ? 'scale-100' : 'scale-95 opacity-0 pointer-events-none'">
            <div class="flex justify-between items-center mb-4">
              <h4 class="font-bold text-meal-gray-dark">
                {{ selectedDay }}. {{ monthNames[selectedMonth] }} {{ currentYear }}
              </h4>
              <button @click="showDayDetails = false" class="text-meal-gray hover:text-meal-gray-dark transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>

            <div class="overflow-y-auto flex-grow">
              <div v-for="(tx, idx) in getTransactionsForDay(selectedDay)" :key="idx"
                   class="flex items-center justify-between p-3 border-b border-meal-gray-light last:border-0">
                <div class="flex items-center">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white text-xs font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(tx.from_user_id) }">
                    {{ getUserInitial(tx.from_user_id) }}
                  </div>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-meal-gray mx-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                  </svg>
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white text-xs font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(tx.to_user_id) }">
                    {{ getUserInitial(tx.to_user_id) }}
                  </div>
                </div>
                <div class="text-right">
                  <div class="font-bold text-meal-primary">{{ formatCurrency(tx.amount) }}</div>
                  <div class="text-xs text-meal-gray">{{ formatTime(tx.created_at) }}</div>
                </div>
              </div>
            </div>

            <div class="mt-4 border-t border-meal-gray-light pt-3 flex justify-between items-center">
              <div class="text-sm text-meal-gray">
                <span class="font-medium">{{ getTransactionsForDay(selectedDay).length }}</span> Transaktionen
              </div>
              <div class="text-sm font-bold text-meal-primary">
                Summe: {{ formatCurrency(getDayTotal(selectedDay)) }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, defineProps, watch } from 'vue';

const props = defineProps({
  transactions: {
    type: Array,
    default: () => []
  },
  users: {
    type: Array,
    default: () => []
  }
});

// Demo-Benutzer, falls keine übergeben werden
const internalUsers = ref([
  { id: 1, name: 'Max Mustermann', email: 'max@example.com' },
  { id: 2, name: 'Lisa Müller', email: 'lisa@example.com' },
  { id: 3, name: 'Jonas Schmidt', email: 'jonas@example.com' },
  { id: 4, name: 'Sarah Weber', email: 'sarah@example.com' }
]);

// Demo-Transaktionen, falls keine übergeben werden
const internalTransactions = ref([
  { id: 1, from_user_id: 1, to_user_id: 2, amount: 42.50, meals_count: 2, created_at: '2023-01-15T12:30:00Z' },
  { id: 2, from_user_id: 3, to_user_id: 1, amount: 25.80, meals_count: 1, created_at: '2023-02-20T18:45:00Z' },
  { id: 3, from_user_id: 2, to_user_id: 4, amount: 68.30, meals_count: 3, created_at: '2023-03-05T14:15:00Z' },
  { id: 4, from_user_id: 4, to_user_id: 3, amount: 31.75, meals_count: 1, created_at: '2023-04-12T19:20:00Z' },
  { id: 5, from_user_id: 1, to_user_id: 3, amount: 55.20, meals_count: 2, created_at: '2023-05-28T13:10:00Z' },
  { id: 6, from_user_id: 2, to_user_id: 1, amount: 40.90, meals_count: 1, created_at: '2023-06-15T20:30:00Z' },
  { id: 7, from_user_id: 3, to_user_id: 4, amount: 74.60, meals_count: 3, created_at: '2023-07-08T11:45:00Z' },
  { id: 8, from_user_id: 4, to_user_id: 2, amount: 28.15, meals_count: 1, created_at: '2023-08-19T16:25:00Z' },
  { id: 9, from_user_id: 1, to_user_id: 4, amount: 63.85, meals_count: 2, created_at: '2023-09-02T17:40:00Z' },
  { id: 10, from_user_id: 3, to_user_id: 2, amount: 36.70, meals_count: 1, created_at: '2023-10-10T15:50:00Z' },
  // Zusätzliche Einträge für die Monatsansicht
  { id: 11, from_user_id: 2, to_user_id: 3, amount: 82.40, meals_count: 3, created_at: '2023-11-05T13:20:00Z' },
  { id: 12, from_user_id: 4, to_user_id: 1, amount: 45.25, meals_count: 2, created_at: '2023-11-08T19:10:00Z' },
  { id: 13, from_user_id: 1, to_user_id: 3, amount: 33.75, meals_count: 1, created_at: '2023-11-12T14:30:00Z' },
  { id: 14, from_user_id: 3, to_user_id: 2, amount: 56.20, meals_count: 2, created_at: '2023-11-15T17:45:00Z' },
  { id: 15, from_user_id: 2, to_user_id: 4, amount: 29.90, meals_count: 1, created_at: '2023-11-18T12:20:00Z' },
  { id: 16, from_user_id: 4, to_user_id: 1, amount: 64.30, meals_count: 3, created_at: '2023-11-22T20:15:00Z' },
  { id: 17, from_user_id: 1, to_user_id: 2, amount: 38.50, meals_count: 1, created_at: '2023-11-25T11:40:00Z' },
  { id: 18, from_user_id: 3, to_user_id: 4, amount: 77.10, meals_count: 3, created_at: '2023-11-28T16:55:00Z' }
]);

// Verwende die übergebenen oder die internen Daten
const effectiveTransactions = computed(() => {
  return props.transactions.length > 0 ? props.transactions : internalTransactions.value;
});

const effectiveUsers = computed(() => {
  return props.users.length > 0 ? props.users : internalUsers.value;
});

// ---- ALLGEMEINE EINSTELLUNGEN ----
const viewType = ref('year'); // 'year' oder 'month'
const viewMode = ref('line');  // 'line', 'bar', oder 'area'
const currentDate = new Date();
const currentYear = currentDate.getFullYear();
const currentMonth = currentDate.getMonth();
const today = currentDate.getDate();

const monthNames = [
  'Januar', 'Februar', 'März', 'April', 'Mai', 'Juni',
  'Juli', 'August', 'September', 'Oktober', 'November', 'Dezember'
];

// ---- JAHRESANSICHT ----
const hoveredMonth = ref(null);
const width = 980;  // Für SVG-Berechnung
const height = 300; // Für SVG-Berechnung

// Letzte Transaktion für das Highlight-Panel
const lastTransaction = computed(() => {
  const sorted = [...effectiveTransactions.value].sort((a, b) =>
      new Date(b.created_at) - new Date(a.created_at)
  );
  return sorted[0] || {};
});

// Gruppiere Transaktionen nach Monat und berechne die Summen
const monthlyData = computed(() => {
  const months = [
    'Jan', 'Feb', 'März', 'Apr', 'Mai', 'Jun',
    'Jul', 'Aug', 'Sep', 'Okt', 'Nov', 'Dez'
  ];

  // Berechne die monatlichen Summen
  const monthlySums = Array(12).fill(0);

  effectiveTransactions.value.forEach(transaction => {
    const date = new Date(transaction.created_at);
    const month = date.getMonth();
    monthlySums[month] += transaction.amount;
  });

  // Erstelle formatierte Daten für jeden Monat
  return months.map((label, index) => ({
    label,
    value: monthlySums[index]
  }));
});

// Maximaler Wert für die Skalierung
const maxValue = computed(() => {
  const values = monthlyData.value.map(m => m.value);
  const max = Math.max(...values);
  return max > 0 ? max : 100; // Fallback-Wert, falls alle Werte 0 sind
});

// Y-Achsen-Beschriftungen
const yAxisLabels = computed(() => {
  const stepCount = 5;
  const labels = [];
  const step = maxValue.value / (stepCount - 1);

  for (let i = 0; i < stepCount; i++) {
    labels.unshift(i * step); // Füge von unten nach oben hinzu
  }

  return labels;
});

// Berechne die Punkte für den SVG-Pfad
const graphPoints = computed(() => {
  const padding = 20; // Abstand zum Rand
  const availableWidth = width - 2 * padding;
  const availableHeight = height - 2 * padding;
  const segmentWidth = availableWidth / (monthlyData.value.length - 1);

  return monthlyData.value.map((month, index) => {
    const x = padding + index * segmentWidth;
    // Y ist umgekehrt, da SVG-Koordinaten von oben nach unten laufen
    const y = height - padding - (month.value / maxValue.value) * availableHeight;
    return { x, y };
  });
});

// Erstelle den SVG-Pfad für die Linie
const linePath = computed(() => {
  if (graphPoints.value.length === 0) return '';

  // Erstelle einen Pfad vom ersten bis zum letzten Punkt
  let path = `M ${graphPoints.value[0].x},${graphPoints.value[0].y}`;

  // Füge für jeden weiteren Punkt einen Kurvensegment hinzu
  for (let i = 1; i < graphPoints.value.length; i++) {
    const point = graphPoints.value[i];
    const prevPoint = graphPoints.value[i-1];

    // Erstelle eine sanfte Kurve mit kubischem Bézier
    const controlX1 = prevPoint.x + (point.x - prevPoint.x) / 3;
    const controlY1 = prevPoint.y;
    const controlX2 = point.x - (point.x - prevPoint.x) / 3;
    const controlY2 = point.y;

    path += ` C ${controlX1},${controlY1} ${controlX2},${controlY2} ${point.x},${point.y}`;
  }

  return path;
});

// ---- MONATSANSICHT ----
const selectedMonth = ref(currentMonth);
const selectedDay = ref(null);
const showDayDetails = ref(false);

// Berechne die Anzahl der Tage im ausgewählten Monat
const daysInMonth = computed(() => {
  return new Date(currentYear, selectedMonth.value + 1, 0).getDate();
});

// Berechne den Wochentag des ersten Tags des Monats (0 = Sonntag, 1 = Montag, ..., 6 = Samstag)
const firstDayOffset = computed(() => {
  const firstDay = new Date(currentYear, selectedMonth.value, 1).getDay();
  // Konvertiere von Sonntag (0) - Samstag (6) zu Montag (0) - Sonntag (6)
  return firstDay === 0 ? 6 : firstDay - 1;
});

// Filtere Transaktionen für den ausgewählten Monat
const monthlyTransactions = computed(() => {
  return effectiveTransactions.value.filter(tx => {
    const txDate = new Date(tx.created_at);
    return txDate.getMonth() === selectedMonth.value && txDate.getFullYear() === currentYear;
  });
});

// Funktionen für die Monatsansicht
function prevMonth() {
  if (selectedMonth.value > 0) {
    selectedMonth.value--;
  } else {
    // Hier könnte man eigentlich ins Vorjahr wechseln, aber wir halten es einfach
  }
  selectedDay.value = null;
  showDayDetails.value = false;
}

function nextMonth() {
  if (selectedMonth.value < 11) {
    selectedMonth.value++;
  } else {
    // Hier könnte man eigentlich ins nächste Jahr wechseln
  }
  selectedDay.value = null;
  showDayDetails.value = false;
}

function selectDay(day) {
  if (getTransactionsForDay(day).length > 0) {
    selectedDay.value = day;
    showDayDetails.value = true;
  }
}

function getTransactionsForDay(day) {
  return monthlyTransactions.value.filter(tx => {
    const txDate = new Date(tx.created_at);
    return txDate.getDate() === day;
  });
}

function hasTransactionOnDay(day) {
  return getTransactionsForDay(day).length > 0;
}

function isToday(day) {
  return selectedMonth.value === currentMonth && day === today;
}

function getDayTotal(day) {
  return getTransactionsForDay(day).reduce((sum, tx) => sum + tx.amount, 0);
}

// Zur Monatsansicht wechseln und einen bestimmten Monat auswählen
function switchToMonth(monthIndex) {
  selectedMonth.value = monthIndex;
  viewType.value = 'month';
  selectedDay.value = null;
  showDayDetails.value = false;
}

// Farben für Benutzer-Avatare
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

// Hilfsfunktionen
function formatCurrency(value) {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR',
    maximumFractionDigits: 0 // Keine Dezimalstellen für die Y-Achse
  }).format(value);
}

function formatShortCurrency(value) {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR',
    maximumFractionDigits: 0
  }).format(value);
}

function formatDate(dateString) {
  if (!dateString) return '';
  const date = new Date(dateString);
  return date.toLocaleDateString('de-DE');
}

function formatTime(dateString) {
  if (!dateString) return '';
  const date = new Date(dateString);
  return date.toLocaleTimeString('de-DE', { hour: '2-digit', minute: '2-digit' });
}

function getUserColor(userId) {
  if (!userId) return '#CCCCCC';
  return avatarColors[userId % avatarColors.length];
}

function getUserInitial(userId) {
  if (!userId) return '?';
  const user = effectiveUsers.value.find(u => u.id === userId);
  return user ? user.name.charAt(0) : '?';
}

// Wenn der ausgewählte Tag geändert wird, automatisch Details anzeigen
watch(selectedDay, (newDay) => {
  if (newDay !== null) {
    showDayDetails.value = true;
  }
});
</script>

<style scoped>
/* Für Animationen */
@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
  100% {
    transform: scale(1);
  }
}

.animate-ping {
  animation: pulse 1.5s infinite;
}

/* Für extra kleine Texte */
.text-xxs {
  font-size: 0.65rem;
  line-height: 0.9rem;
}
</style>