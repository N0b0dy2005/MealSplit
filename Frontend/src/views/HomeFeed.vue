<template>
  <div class="min-h-screen flex-grow flex flex-col bg-meal-light font-sans">
    <main class="flex-grow">
      <!-- Hero Section -->
      <section class="bg-gradient-to-br from-meal-primary to-meal-dark text-white py-10 sm:py-12 md:py-20">
        <div class="container mx-auto px-4 text-center">
          <h2 class="text-2xl sm:text-3xl md:text-5xl font-header font-bold mb-4 sm:mb-6">Nie wieder den Überblick verlieren</h2>
          <p class="text-base sm:text-lg md:text-xl max-w-2xl mx-auto mb-6 sm:mb-8 md:mb-10 px-2">
            Behalte den Überblick über gemeinsame Mahlzeiten und wer wem wie viel Geld schuldet.
            Einfach, schnell und übersichtlich.
          </p>
          <div class="flex flex-col sm:flex-row justify-center gap-3 sm:gap-4 px-4">
            <button
                @click="goToMeal"
                class="bg-meal-accent hover:bg-opacity-90 text-white font-bold py-2 sm:py-3 px-6 sm:px-8 rounded-lg shadow-lg transition-all duration-200 transform hover:scale-105 w-full sm:w-auto mb-3 sm:mb-0"
            >
              Mahlzeit erfassen
            </button>
            <button
                @click="goToUser"
                class="bg-white text-meal-primary hover:bg-opacity-90 font-bold py-2 sm:py-3 px-6 sm:px-8 rounded-lg shadow-lg transition-all duration-200 transform hover:scale-105 w-full sm:w-auto"
            >
              Benutzer verwalten
            </button>
          </div>
        </div>
      </section>

      <!-- Features Section -->
      <section class="py-10 sm:py-12 md:py-16 container mx-auto px-4">
        <h2 class="text-xl sm:text-2xl md:text-3xl font-header font-bold text-meal-gray-dark text-center mb-8 sm:mb-10 md:mb-12">So funktioniert MealSplit</h2>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-5 sm:gap-6 md:gap-8">
          <!-- Feature Card 1 -->
          <div class="bg-white rounded-xl shadow-meal p-5 sm:p-6 transition-transform duration-300 hover:-translate-y-1">
            <div class="bg-meal-primary inline-block p-3 rounded-full mb-4">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 sm:h-6 sm:w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
            </div>
            <h3 class="text-lg sm:text-xl font-header font-bold text-meal-gray-dark mb-2">1. Mahlzeiten erfassen</h3>
            <p class="text-meal-gray text-sm sm:text-base">
              Erfasse gemeinsame Mahlzeiten mit Namen, Datum, Gesamtkosten und wähle aus, wer teilgenommen hat.
            </p>
            <button
                @click="goToMeal"
                class="mt-4 w-full bg-meal-light hover:bg-meal-primary hover:text-white text-meal-primary font-medium py-2 px-4 rounded-lg transition-colors duration-200 text-sm sm:text-base"
            >
              Zur Mahlzeitenverwaltung
            </button>
          </div>

          <!-- Feature Card 2 -->
          <div class="bg-white rounded-xl shadow-meal p-5 sm:p-6 transition-transform duration-300 hover:-translate-y-1">
            <div class="bg-meal-accent inline-block p-3 rounded-full mb-4">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 sm:h-6 sm:w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <h3 class="text-lg sm:text-xl font-header font-bold text-meal-gray-dark mb-2">2. Zahlungen verfolgen</h3>
            <p class="text-meal-gray text-sm sm:text-base">
              Dokumentiere, wer wem Geld zurückgezahlt hat, und behalte alle Transaktionen im Blick.
            </p>
            <button
                @click="goToPayments"
                class="mt-4 w-full bg-meal-light hover:bg-meal-primary hover:text-white text-meal-primary font-medium py-2 px-4 rounded-lg transition-colors duration-200 text-sm sm:text-base"
            >
              Zur Zahlungsverwaltung
            </button>
          </div>

          <!-- Feature Card 3 -->
          <div class="bg-white rounded-xl shadow-meal p-5 sm:p-6 transition-transform duration-300 hover:-translate-y-1">
            <div class="bg-meal-secondary inline-block p-3 rounded-full mb-4">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 sm:h-6 sm:w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
              </svg>
            </div>
            <h3 class="text-lg sm:text-xl font-header font-bold text-meal-gray-dark mb-2">3. Beträge im Überblick</h3>
            <p class="text-meal-gray text-sm sm:text-base">
              Alle Beträge werden automatisch berechnet. Sehe auf einen Blick, wer wem wie viel schuldet.
            </p>
            <button @click="goToDebts"
                    class="mt-4 w-full bg-meal-light hover:bg-meal-primary hover:text-white text-meal-primary font-medium py-2 px-4 rounded-lg transition-colors duration-200 text-sm sm:text-base"
            >
              Zur Schuldenübersicht
            </button>
          </div>
        </div>
      </section>

      <!-- How It Works Section -->
      <section class="bg-meal-gray-light py-10 sm:py-12 md:py-16">
        <div class="container mx-auto px-4">
          <h2 class="text-xl sm:text-2xl md:text-3xl font-header font-bold text-meal-gray-dark text-center mb-8 sm:mb-10 md:mb-12">So einfach geht's</h2>

          <div class="max-w-4xl mx-auto">
            <!-- Mobile Timeline (visible on small screens only) -->
            <div class="md:hidden space-y-8">
              <!-- Step 1 -->
              <div class="flex">
                <div class="flex-shrink-0 relative">
                  <div class="hidden md:inline h-full w-0.5 bg-meal-primary absolute left-4 top-0"></div>
                  <div class="bg-meal-primary text-white rounded-full w-8 h-8 flex items-center justify-center font-bold text-lg z-10">
                    1
                  </div>
                </div>
                <div class="ml-6">
                  <h3 class="text-lg font-header font-bold text-meal-primary mb-1">Benutzer hinzufügen</h3>
                  <p class="text-meal-gray-dark text-sm sm:text-base">
                    Füge alle Personen hinzu, die regelmäßig an gemeinsamen Mahlzeiten teilnehmen.
                  </p>
                </div>
              </div>

              <!-- Step 2 -->
              <div class="flex">
                <div class="flex-shrink-0 relative">
                  <div class="h-full hidden md:inline  w-0.5 bg-meal-primary absolute left-4 top-0"></div>
                  <div class="bg-meal-primary text-white rounded-full w-8 h-8 flex items-center justify-center font-bold text-lg z-10">
                    2
                  </div>
                </div>
                <div class="ml-6">
                  <h3 class="text-lg font-header font-bold text-meal-primary mb-1">Mahlzeit erfassen</h3>
                  <p class="text-meal-gray-dark text-sm sm:text-base">
                    Trage nach einem gemeinsamen Essen die Details ein: Wer hat teilgenommen, wer hat bezahlt und wie hoch war die Rechnung?
                  </p>
                </div>
              </div>

              <!-- Step 3 -->
              <div class="flex">
                <div class="flex-shrink-0 relative">
                  <div class="h-full hidden md:inline  w-0.5 bg-meal-primary absolute left-4 top-0"></div>
                  <div class="bg-meal-primary text-white rounded-full w-8 h-8 flex items-center justify-center font-bold text-lg z-10">
                    3
                  </div>
                </div>
                <div class="ml-6">
                  <h3 class="text-lg font-header font-bold text-meal-primary mb-1">Beträge anzeigen</h3>
                  <p class="text-meal-gray-dark text-sm sm:text-base">
                    MealSplit berechnet automatisch, wer wem wie viel Geld schuldet, basierend auf allen erfassten Mahlzeiten.
                  </p>
                </div>
              </div>

              <!-- Step 4 -->
              <div class="flex">
                <div class="flex-shrink-0">
                  <div class="bg-meal-primary text-white rounded-full w-8 h-8 flex items-center justify-center font-bold text-lg z-10">
                    4
                  </div>
                </div>
                <div class="ml-6">
                  <h3 class="text-lg font-header font-bold text-meal-primary mb-1">Zahlungen erfassen</h3>
                  <p class="text-meal-gray-dark text-sm sm:text-base">
                    Wenn jemand seine Beträge begleicht, erfasse die Zahlung, und die Schuldenübersicht wird automatisch aktualisiert.
                  </p>
                </div>
              </div>
            </div>

            <!-- Desktop Timeline (hidden on small screens) -->
            <div class="hidden md:block relative">
              <!-- Timeline line -->
              <div class="absolute left-1/2 transform -translate-x-1/2 h-full w-1 bg-meal-primary"></div>

              <!-- Step 1 -->
              <div class="relative flex flex-row mb-12">
                <div class="flex-1 text-right pr-8">
                  <h3 class="text-xl font-header font-bold text-meal-primary mb-2">Benutzer hinzufügen</h3>
                  <p class="text-meal-gray-dark">
                    Füge alle Personen hinzu, die regelmäßig an gemeinsamen Mahlzeiten teilnehmen.
                  </p>
                </div>
                <div class="absolute left-1/2 transform -translate-x-1/2 bg-meal-primary text-white rounded-full w-12 h-12 flex items-center justify-center font-bold text-xl z-10">
                  1
                </div>
                <div class="flex-1 pl-8 text-left"></div>
              </div>

              <!-- Step 2 -->
              <div class="relative flex flex-row-reverse mb-12">
                <div class="flex-1 text-left pl-8">
                  <h3 class="text-xl font-header font-bold text-meal-primary mb-2">Mahlzeit erfassen</h3>
                  <p class="text-meal-gray-dark">
                    Trage nach einem gemeinsamen Essen die Details ein: Wer hat teilgenommen, wer hat bezahlt und wie hoch war die Rechnung?
                  </p>
                </div>
                <div class="absolute left-1/2 transform -translate-x-1/2 bg-meal-primary text-white rounded-full w-12 h-12 flex items-center justify-center font-bold text-xl z-10">
                  2
                </div>
                <div class="flex-1 pr-8 text-right"></div>
              </div>

              <!-- Step 3 -->
              <div class="relative flex flex-row mb-12">
                <div class="flex-1 text-right pr-8">
                  <h3 class="text-xl font-header font-bold text-meal-primary mb-2">Beträge anzeigen</h3>
                  <p class="text-meal-gray-dark">
                    MealSplit berechnet automatisch, wer wem wie viel Geld schuldet, basierend auf allen erfassten Mahlzeiten.
                  </p>
                </div>
                <div class="absolute left-1/2 transform -translate-x-1/2 bg-meal-primary text-white rounded-full w-12 h-12 flex items-center justify-center font-bold text-xl z-10">
                  3
                </div>
                <div class="flex-1 pl-8 text-left"></div>
              </div>

              <!-- Step 4 -->
              <div class="relative flex flex-row-reverse">
                <div class="flex-1 text-left pl-8">
                  <h3 class="text-xl font-header font-bold text-meal-primary mb-2">Zahlungen erfassen</h3>
                  <p class="text-meal-gray-dark">
                    Wenn jemand seine Beträge begleicht, erfasse die Zahlung, und die Schuldenübersicht wird automatisch aktualisiert.
                  </p>
                </div>
                <div class="absolute left-1/2 transform -translate-x-1/2 bg-meal-primary text-white rounded-full w-12 h-12 flex items-center justify-center font-bold text-xl z-10">
                  4
                </div>
                <div class="flex-1 pr-8 text-right"></div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Call to Action -->
      <section class="py-10 sm:py-12 md:py-16 container mx-auto px-4 text-center">
        <h2 class="text-xl sm:text-2xl md:text-3xl font-header font-bold text-meal-gray-dark mb-3 sm:mb-4 md:mb-6">Bereit, mit MealSplit zu starten?</h2>
        <p class="text-base sm:text-lg md:text-xl text-meal-gray max-w-2xl mx-auto mb-5 sm:mb-6 md:mb-8 px-2">
          Erstelle deine erste Mahlzeit und behalte den Überblick über alle Ausgaben in deiner Gruppe.
        </p>
        <div class="flex flex-col sm:flex-row flex-wrap justify-center gap-3 sm:gap-4 md:gap-6 px-4">
          <button @click="goToUser" class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 sm:py-3 px-4 sm:px-6 rounded-lg transition-colors duration-200 w-full sm:w-auto mb-3 sm:mb-0">
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z" />
              </svg>
              Benutzer verwalten
            </span>
          </button>
          <button @click="goToMeal" class="bg-meal-accent hover:bg-opacity-90 text-white font-bold py-2 sm:py-3 px-4 sm:px-6 rounded-lg transition-colors duration-200 w-full sm:w-auto mb-3 sm:mb-0">
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z" />
              </svg>
              Mahlzeit erfassen
            </span>
          </button>
          <button @click="goToPayments" class="bg-meal-secondary hover:bg-opacity-90 text-white font-bold py-2 sm:py-3 px-4 sm:px-6 rounded-lg transition-colors duration-200 w-full sm:w-auto">
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path d="M4 4a2 2 0 00-2 2v1h16V6a2 2 0 00-2-2H4z" />
                <path fill-rule="evenodd" d="M18 9H2v5a2 2 0 002 2h12a2 2 0 002-2V9zM4 13a1 1 0 011-1h1a1 1 0 110 2H5a1 1 0 01-1-1zm5-1a1 1 0 100 2h1a1 1 0 100-2H9z" clip-rule="evenodd" />
              </svg>
              Zahlung erfassen
            </span>
          </button>
        </div>
      </section>
    </main>

    <!-- Footer -->
    <footer class="bg-meal-gray-dark text-white py-5 sm:py-6 md:py-8">
      <div class="container mx-auto px-4 text-center">
        <p class="text-xs sm:text-sm md:text-base">© 2023 MealSplit - Den Überblick bei gemeinsamen Mahlzeiten behalten</p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import router from "../router";

const meals = ref([])
const payments = ref([])
const debts = ref([])

function goToUser() {
  router.push('/user')
}

function goToPayments() {
  router.push('/payments')
}

function goToMeal() {
  router.push('/meal')
}

function goToDebts() {
  router.push('/debts')
}

function goToDashboard() {
  router.push('/dashboard')
}

function goToProfile() {
  router.push('/profile')
}
</script>