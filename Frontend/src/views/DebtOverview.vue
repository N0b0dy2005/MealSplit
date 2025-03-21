<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <!-- Header -->
    <header class="bg-meal-primary text-white p-4 shadow-md mb-8">
      <div class="container mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 sm:h-8 sm:w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
          <h1 class="text-xl sm:text-2xl font-header font-bold">
            <span class="hidden sm:inline">Schuldenübersicht</span>
            <span class="sm:hidden">Beträge</span>
          </h1>
        </div>
        <button @click="goBackToHomeScreen" class="text-white  hover:text-meal-accent-light transition-colors duration-200">
          <span class="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            <span class="hidden sm:inline">Zurück zur Startseite</span>
            <span class="sm:hidden">Zurück</span>
          </span>
        </button>
      </div>
    </header>

    <div class="container mx-auto px-4">
      <!-- Zusammenfassung -->
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 sm:gap-6 mb-6 sm:mb-8">
        <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-1 sm:mb-2">Offene Beträge</h3>
          <p class="text-2xl sm:text-3xl font-bold text-meal-error">{{ formatCurrency(totalDebt) }}</p>
          <p class="text-xs sm:text-sm text-meal-gray mt-1">Summe aller ausstehenden Beträge</p>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-1 sm:mb-2">Aktive Transaktionen</h3>
          <p class="text-2xl sm:text-3xl font-bold text-meal-primary">{{ debts.length }}</p>
          <p class="text-xs sm:text-sm text-meal-gray mt-1">Anzahl der offenen Schuldenbeziehungen</p>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-4 sm:p-6">
          <h3 class="text-base sm:text-lg font-header font-bold text-meal-gray-dark mb-1 sm:mb-2">Letzte Zahlung</h3>
          <p class="text-lg sm:text-xl font-bold text-meal-positive">{{ formatCurrency(lastPayment.amount) }}</p>
          <p class="text-xs sm:text-sm text-meal-gray mt-1">
            {{ lastPayment.from }} → {{ lastPayment.to }} ({{ formatDate(lastPayment.date) }})
          </p>
        </div>
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
                <path fill-rule="evenodd" d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd" />
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
                <path d="M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V5zM11 13a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
              </svg>
              <span class="hidden sm:inline">Beträge Matrix</span>
            </span>
          </button>
          <button
              @click="activeTab = 'optimize'"
              :class="[
              'py-2 sm:py-4 px-2 sm:px-6 font-bold text-center flex-1 text-sm sm:text-base',
              activeTab === 'optimize'
                ? 'text-meal-primary border-b-2 border-meal-primary'
                : 'text-meal-gray hover:text-meal-gray-dark'
            ]"
          >
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 sm:mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z" clip-rule="evenodd" />
              </svg>
              <span class="hidden sm:inline">Optimierte Zahlungen</span>
            </span>
          </button>
        </div>

        <!-- Beträge Liste Ansicht -->
        <div v-if="activeTab === 'list'" class="p-4 sm:p-6">
          <div v-if="debts.length === 0" class="text-center py-6 sm:py-8 text-meal-gray">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 sm:h-16 sm:w-16 mx-auto mb-4 text-meal-gray-light" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
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
                  <th class="py-3 px-4 rounded-tl-lg">Schuldner</th>
                  <th class="py-3 px-4 text-center">Schuldet</th>
                  <th class="py-3 px-4">Gläubiger</th>
                  <th class="py-3 px-4 text-center">Basierend auf</th>
                  <th class="py-3 px-4 rounded-tr-lg text-right">Aktionen</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(debt, index) in debts" :key="index" class="border-b border-meal-gray-light hover:bg-meal-gray-light transition-colors duration-150">
                  <!-- Schuldner -->
                  <td class="py-3 px-4">
                    <div class="flex items-center">
                      <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold mr-3"
                           :style="{ backgroundColor: getUserColor(debt.fromUserId) }">
                        {{ debt.from.name.charAt(0).toUpperCase() }}
                      </div>
                      {{ debt.from.name }}
                    </div>
                  </td>

                  <!-- Betrag -->
                  <td class="py-3 px-4 text-center">
                    <span class="font-bold text-meal-error">{{ formatCurrency(debt.amount) }}</span>
                  </td>

                  <!-- Gläubiger -->
                  <td class="py-3 px-4">
                    <div class="flex items-center">
                      <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold mr-3"
                           :style="{ backgroundColor: getUserColor(debt.toUserId) }">
                        {{ debt.to.name.charAt(0).toUpperCase() }}
                      </div>
                      {{ debt.to.name }}
                    </div>
                  </td>

                  <!-- Mahlzeiten -->
                  <td class="py-3 px-4 text-center">
                    <span>{{ debt.meals }} Mahlzeiten</span>
                  </td>

                  <!-- Aktionen -->
                  <td class="py-3 px-4 text-right">
                    <button
                        @click="markAsPaid(index)"
                        class="bg-meal-primary hover:bg-meal-dark text-white px-3 py-1 rounded transition-colors duration-200"
                    >
                      Als bezahlt markieren
                    </button>
                  </td>
                </tr>
                </tbody>
              </table>
            </div>

            <!-- Mobile Cards -->
            <div class="sm:hidden space-y-4">
              <div v-for="(debt, index) in debts" :key="index"
                   class="bg-white border border-meal-gray-light rounded-lg overflow-hidden">
                <div class="p-4">
                  <div class="flex justify-between items-center mb-3">
                    <!-- Schuldner info -->
                    <div class="flex items-center">
                      <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                           :style="{ backgroundColor: getUserColor(debt.fromUserId) }">
                        {{ debt.from.name.charAt(0).toUpperCase() }}
                      </div>
                      <div>
                        <div class="text-xs text-meal-gray">Schuldner</div>
                        <div class="font-medium text-sm">{{ debt.from.name }}</div>
                      </div>
                    </div>

                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-meal-gray mx-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                    </svg>

                    <!-- Gläubiger info -->
                    <div class="flex items-center">
                      <div>
                        <div class="text-xs text-meal-gray text-right">Gläubiger</div>
                        <div class="font-medium text-sm text-right">{{ debt.to.name }}</div>
                      </div>
                      <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold ml-2"
                           :style="{ backgroundColor: getUserColor(debt.toUserId) }">
                        {{ debt.to.name.charAt(0).toUpperCase() }}
                      </div>
                    </div>
                  </div>

                  <div class="flex justify-between items-center mb-3">
                    <div class="text-center flex-1">
                      <div class="text-xs text-meal-gray">Betrag</div>
                      <div class="text-lg font-bold text-meal-error">{{ formatCurrency(debt.amount) }}</div>
                    </div>
                    <div class="text-center flex-1">
                      <div class="text-xs text-meal-gray">Basis</div>
                      <div class="text-sm">{{ debt.meals }} Mahlzeiten</div>
                    </div>
                  </div>

                  <button
                      @click="markAsPaid(index)"
                      class="w-full bg-meal-primary hover:bg-meal-dark text-white px-3 py-2 rounded transition-colors duration-200 text-sm"
                  >
                    Als bezahlt markieren
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Beträge Matrix Ansicht -->
        <div v-if="activeTab === 'matrix'" class="p-4 sm:p-6">
          <!-- Matrix View (Desktop only) -->
          <div class="hidden sm:block overflow-x-auto">
            <table class="min-w-full bg-white">
              <thead>
              <tr class="bg-meal-light text-meal-dark text-left">
                <th class="py-3 px-4 rounded-tl-lg">Benutzer</th>
                <th v-for="user in users" :key="user.id" class="py-3 px-4 text-center">
                  <div class="flex flex-col items-center">
                    <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mb-1"
                         :style="{ backgroundColor: getUserColor(user.id) }">
                      {{ user.name.charAt(0).toUpperCase() }}
                    </div>
                    <span class="text-xs">{{ user.name }}</span>
                  </div>
                </th>
                <th class="py-3 px-4 rounded-tr-lg text-center">Gesamt</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="fromUser in users" :key="fromUser.id" class="border-b border-meal-gray-light hover:bg-meal-gray-light transition-colors duration-150">
                <td class="py-3 px-4">
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
                      <span v-if="getDebtAmount(fromUser.id, toUser.id) > 0" class="font-bold text-meal-error">
                        {{ formatCurrency(getDebtAmount(fromUser.id, toUser.id)) }}
                      </span>
                    <span v-else-if="getDebtAmount(toUser.id, fromUser.id) > 0" class="font-bold text-meal-positive">
                        {{ formatCurrency(getDebtAmount(toUser.id, fromUser.id)) }}
                      </span>
                    <span v-else class="text-meal-gray">-</span>
                  </template>
                  <span v-else class="text-meal-gray">x</span>
                </td>
                <td class="py-3 px-4 text-center font-bold" :class="{
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
                     class="p-2 border border-meal-gray-light rounded-lg flex justify-between items-center">
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
                  </div>
                  <span class="font-bold text-meal-error text-sm">{{ formatCurrency(debt.amount) }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="flex justify-center mt-4 text-xs sm:text-sm text-meal-gray flex-wrap gap-2">
            <div class="flex items-center mr-2 sm:mr-4">
              <div class="w-3 h-3 bg-meal-error rounded-full mr-1"></div>
              <span>Schuldet Geld</span>
            </div>
            <div class="flex items-center mr-2 sm:mr-4">
              <div class="w-3 h-3 bg-meal-positive rounded-full mr-1"></div>
              <span>Bekommt Geld</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-meal-gray-light rounded-full mr-1"></div>
              <span>Keine Verbindung</span>
            </div>
          </div>
        </div>

        <!-- Optimierte Zahlungen -->
        <div v-if="activeTab === 'optimize'" class="p-4 sm:p-6">
          <div class="bg-meal-light p-3 sm:p-4 rounded-lg mb-4 sm:mb-6">
            <p class="text-sm text-meal-gray-dark mb-0 sm:mb-2">
              Die folgenden Zahlungen sind optimiert, um die Anzahl der Transaktionen zu minimieren.
              Anstatt viele kleine Zahlungen zwischen allen Benutzern, empfehlen wir diese effizientere Lösung.
            </p>
          </div>

          <div v-if="optimizedPayments.length === 0" class="text-center py-6 sm:py-8 text-meal-gray">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 sm:h-16 sm:w-16 mx-auto mb-4 text-meal-gray-light" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <p class="text-lg sm:text-xl mb-2">Keine Zahlungen notwendig</p>
            <p>Alle Beträge sind bereits ausgeglichen.</p>
          </div>

          <div v-else class="space-y-4 sm:space-y-6">
            <div v-for="(payment, index) in optimizedPayments" :key="index"
                 class="bg-white border border-meal-gray-light rounded-lg overflow-hidden">
              <!-- Desktop Layout -->
              <div class="hidden sm:flex items-center p-4 bg-meal-light">
                <div class="w-12 h-12 rounded-full flex items-center justify-center text-white font-bold mr-3"
                     :style="{ backgroundColor: getUserColor(payment.fromUserId) }">
                  {{ payment.from.name.charAt(0).toUpperCase() }}
                </div>
                <div class="flex-grow">
                  <p class="font-bold text-meal-gray-dark">{{ payment.from.name }}</p>
                  <p class="text-sm text-meal-gray">zahlt an</p>
                </div>
                <div class="font-bold text-2xl text-meal-error mx-4">
                  {{ formatCurrency(payment.amount) }}
                </div>
                <div class="flex items-center">
                  <div class="flex-grow text-right mr-3">
                    <p class="font-bold text-meal-gray-dark">{{ payment.to.name }}</p>
                    <p class="text-sm text-meal-gray">erhält</p>
                  </div>
                  <div class="w-12 h-12 rounded-full flex items-center justify-center text-white font-bold"
                       :style="{ backgroundColor: getUserColor(payment.toUserId) }">
                    {{ payment.to.name.charAt(0).toUpperCase() }}
                  </div>
                </div>
              </div>

              <!-- Mobile Layout -->
              <div class="sm:hidden flex flex-col p-3 bg-meal-light">
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center">
                    <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold mr-2"
                         :style="{ backgroundColor: getUserColor(payment.fromUserId) }">
                      {{ payment.from.name.charAt(0).toUpperCase() }}
                    </div>
                    <div>
                      <p class="font-bold text-meal-gray-dark text-sm">{{ payment.from.name }}</p>
                      <p class="text-xs text-meal-gray">zahlt an</p>
                    </div>
                  </div>

                  <div class="flex items-center">
                    <div>
                      <p class="font-bold text-meal-gray-dark text-sm text-right">{{ payment.to.name }}</p>
                      <p class="text-xs text-meal-gray text-right">erhält</p>
                    </div>
                    <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold ml-2"
                         :style="{ backgroundColor: getUserColor(payment.toUserId) }">
                      {{ payment.to.name.charAt(0).toUpperCase() }}
                    </div>
                  </div>
                </div>

                <div class="font-bold text-xl text-meal-error text-center mb-2">
                  {{ formatCurrency(payment.amount) }}
                </div>
              </div>

              <div class="p-4 border-t border-meal-gray-light">
                <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center gap-3">
                  <button
                      @click="markOptimizedAsPaid(index)"
                      class="bg-meal-primary hover:bg-meal-dark text-white px-4 py-2 rounded transition-colors duration-200 w-full sm:w-auto text-sm sm:text-base"
                  >
                    Als bezahlt markieren
                  </button>

                  <div class="flex items-center text-sm text-meal-gray-dark">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-5 sm:w-5 mr-1" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" />
                    </svg>
                    <span>Diese Zahlung löst {{ payment.resolvesCount }} Beträge gleichzeitig</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Zahlungen Historie -->
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
                   :style="{ backgroundColor: getUserColor(payment.fromUserId) }">
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
              <span class="font-bold text-meal-positive mr-2">{{ formatCurrency(payment.amount) }}</span>
              <span
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-meal-positive text-white"
              >
                Bezahlt
              </span>
            </div>
          </div>

          <!-- Mobile Payment History -->
          <div v-for="payment in recentPayments" :key="payment.id"
               class="flex sm:hidden flex-col p-3 border-b border-meal-gray-light">
            <div class="flex items-center mb-2">
              <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                   :style="{ backgroundColor: getUserColor(payment.fromUserId) }">
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
              <span class="font-bold text-meal-positive mr-2 text-sm">{{ formatCurrency(payment.amount) }}</span>
              <span
                  class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-meal-positive text-white"
              >
                Bezahlt
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import router from "../router";
import {mutation, query} from "../graphql.ts";
import {create, get} from "axios";

// Beispiel-Benutzer
const users = ref([]);

// Beispiel-Beträge (in der Realität berechnet aus Mahlzeiten und Zahlungen)
const debts = ref([]);

// Beispiel-Zahlungshistorie
const recentPayments = ref([]);


async function getUsers() {
  const [data, err] = await query.getUsers();
  if (err) {
    console.error('Error getting users:', err);
    return;
  }
  users.value = data;
}

async function getDebts() {
  const [data, err] = await query.getDebts();
  if (err) {
    console.error('Error getting debts:', err);
    return;
  }

  // Transform debts to include user objects
  debts.value = data.map(debt => {
    const fromUser = users.value.find(user => user.id === debt.fromUserId);
    const toUser = users.value.find(user => user.id === debt.toUserId);

    return {
      ...debt,
      amount: parseFloat(debt.amount),
      meals: debt.mealsCount,
      from: {
        id: fromUser?.id,
        name: fromUser?.name
      },
      to: {
        id: toUser?.id,
        name: toUser?.name
      }
    };
  });
}

async function getRecentPayments() {
  const [data, err] = await query.getPayments();
  if (err) {
    console.error('Error getting payments:', err);
    return;
  }

  // Transform payments to include user objects
  recentPayments.value = data.map(payment => {
    const fromUser = users.value.find(user => user.id === payment.fromUserId);
    const toUser = users.value.find(user => user.id === payment.toUserId);

    return {
      ...payment,
      amount: parseFloat(payment.amount),
      from: {
        id: fromUser?.id,
        name: fromUser?.name
      },
      to: {
        id: toUser?.id,
        name: toUser?.name
      }
    };
  });
}
// UI Status
const activeTab = ref('list');

// Farben für Benutzer-Avatare
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

// Berechnete Eigenschaften
const totalDebt = computed(() => {
  return debts.value.reduce((sum, debt) => sum + debt.amount, 0);
});

const lastPayment = computed(() => {
  if (recentPayments.value.length === 0) {
    return { from: 'Niemand', to: 'Niemand', amount: 0, date: new Date().toISOString() };
  }
  return {
    from: recentPayments.value[0].from.name,
    to: recentPayments.value[0].to.name,
    amount: recentPayments.value[0].amount,
    date: recentPayments.value[0].date
  };
});
const optimizedPayments = computed(() => {
  if (users.value.length === 0 || debts.value.length === 0) {
    return [];
  }

  // Gruppiere Beträge nach Schuldner-Gläubiger-Paar
  const groupedDebts = {};

  debts.value.forEach(debt => {
    const key = `${debt.fromUserId}-${debt.toUserId}`;
    if (!groupedDebts[key]) {
      groupedDebts[key] = {
        fromUserId: debt.fromUserId,
        toUserId: debt.toUserId,
        from: { ...debt.from },
        to: { ...debt.to },
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

async function markAsPaid(index) {
  const debt = debts.value[index];
  let payment = {
    fromUserId: debt.from.id,
    toUserId: debt.to.id,
    amount: debt.amount,
    description: `Bezahlungen für` + debt.meals
  }
  const [data, err] = await mutation.createPayment(payment);
  if (err) {
    console.error('Error getting debts:', err);
    return;
  }
  reloadPage();
}


function markOptimizedAsPaid(index) {
  // Ähnlich wie markAsPaid, aber für optimierte Zahlungen
  const payment = optimizedPayments.value[index];

  // Füge eine neue Zahlung zur Historie hinzu
  recentPayments.value.unshift({
    id: recentPayments.value.length + 1,
    from: payment.from,
    to: payment.to,
    amount: payment.amount,
    date: new Date().toISOString().split('T')[0]
  });

  // Aktualisiere Beträge - in Realität würde man hier alle betroffenen Beträge aktualisieren
  // Diese einfache Implementierung entfernt einfach die erste passende Schuld
  const debtIndex = debts.value.findIndex(
      debt => debt.fromUserId === payment.fromUserId && debt.toUserId === payment.toUserId
  );

  if (debtIndex >= 0) {
    const debt = debts.value[debtIndex];
    if (debt.amount <= payment.amount) {
      debts.value.splice(debtIndex, 1);
    } else {
      debt.amount -= payment.amount;
    }
  }
}

function getDebtAmount(fromId, toId) {
  const debt = debts.value.find(d => d.fromUserId === fromId && d.toUserId === toId);
  return debt ? debt.amount : 0;
}

function getUserBalance(userId) {
  let balance = 0;

  // Beträge (negativ)
  debts.value.forEach(debt => {
    if (debt.fromUserId === userId) {
      balance -= debt.amount;
    }
  });

  // Forderungen (positiv)
  debts.value.forEach(debt => {
    if (debt.toUserId === userId) {
      balance += debt.amount;
    }
  });

  return balance;
}

function goBackToHomeScreen() {
  router.push('/homefeed');
}

async function reloadPage() {
  await getUsers();
  await getDebts();
  await getRecentPayments();
}
// Beim Laden der Komponente
onMounted(() => {
  reloadPage();
});

</script>