<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <!-- Header -->
    <header class="bg-meal-primary text-white p-4 shadow-md mb-8">
      <div class="container mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
          </svg>
          <h1 class="text-2xl font-header font-bold">Mahlzeiten verwalten</h1>
        </div>
        <button @click="goBackToHomeScreen" class="text-white hover:text-meal-accent-light transition-colors duration-200">
          <span class="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
            Zurück zur Startseite
          </span>
        </button>
      </div>
    </header>

    <div class="container mx-auto px-4">
      <!-- Neue Mahlzeit erstellen -->
      <div class="bg-white rounded-xl shadow-meal p-6 mb-8">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-header font-bold text-meal-gray-dark">Neue Mahlzeit erfassen</h2>
          <button
              @click="showAddForm = !showAddForm"
              class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-4 rounded-lg transition-colors duration-200"
          >
            <span class="flex items-center">
              <svg v-if="!showAddForm" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M5 10a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1z" clip-rule="evenodd" />
              </svg>
              {{ showAddForm ? 'Formular ausblenden' : 'Neue Mahlzeit' }}
            </span>
          </button>
        </div>

        <form v-if="showAddForm" @submit.prevent="saveMeal" class="border-t border-meal-gray-light pt-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-name">
                Mahlzeit Name
              </label>
              <input
                  id="meal-name"
                  v-model="newMeal.name"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  type="text"
                  placeholder="z.B. Pizza Abend"
                  required
              />
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-date">
                Datum
              </label>
              <input
                  id="meal-date"
                  v-model="newMeal.date"
                  type="date"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
              />
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-amount">
                Gesamtbetrag (€)
              </label>
              <input
                  id="meal-amount"
                  v-model="newMeal.totalAmount"
                  type="number"
                  step="0.01"
                  min="0.01"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
                  placeholder="0.00"
              />
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-payer">
                Bezahlt von
              </label>
              <select
                  id="meal-payer"
                  v-model="newMeal.userId"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
              >
                <option value="" disabled selected>Bitte wählen...</option>
                <option v-for="user in users" :key="user.id" :value="user.id">
                  {{ user.name }}
                </option>
              </select>
            </div>

            <div class="md:col-span-2">
              <label class="block text-meal-gray text-sm font-bold mb-2">
                Teilnehmer
              </label>
              <div class="flex flex-wrap gap-2 mb-2">
                <div
                    v-for="user in users"
                    :key="user.id"
                    @click="toggleParticipant(user.id)"
                    :class="[
                    'flex items-center px-3 py-2 rounded-full cursor-pointer transition-colors',
                    isParticipant(user.id)
                      ? 'bg-meal-primary text-white'
                      : 'bg-meal-gray-light text-meal-gray-dark hover:bg-meal-gray'
                  ]"
                >
                  <div class="w-6 h-6 rounded-full flex items-center justify-center mr-2"
                       :class="isParticipant(user.id) ? 'bg-white text-meal-primary' : 'bg-meal-gray text-white'">
                    {{ user.name.charAt(0).toUpperCase() }}
                  </div>
                  {{ user.name }}
                </div>
              </div>
              <div v-if="newMeal.participants.length === 0" class="text-meal-error text-sm mt-1">
                Bitte wähle mindestens einen Teilnehmer aus.
              </div>
            </div>

            <div class="md:col-span-2">
              <label class="block text-meal-gray text-sm font-bold mb-2" for="meal-description">
                Beschreibung (optional)
              </label>
              <textarea
                  id="meal-description"
                  v-model="newMeal.description"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  rows="3"
                  placeholder="Details zur Mahlzeit"
              ></textarea>
            </div>
          </div>

          <div class="flex justify-end mt-6">
            <button
                type="submit"
                class="bg-meal-accent hover:bg-opacity-90 text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
                :disabled="newMeal.participants.length === 0"
            >
              Mahlzeit erfassen
            </button>
          </div>
        </form>
      </div>

      <!-- Filter und Suche -->
      <div class="bg-white rounded-xl shadow-meal p-4 mb-8 flex flex-wrap gap-4 items-center">
        <div class="flex-grow">
          <input
              v-model="searchQuery"
              type="text"
              placeholder="Mahlzeit suchen..."
              class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
          />
        </div>

        <div class="flex gap-2">
          <select
              v-model="filterMonth"
              class="appearance-none border border-meal-gray-light rounded py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
          >
            <option value="">Alle Monate</option>
            <option value="0">Januar</option>
            <option value="1">Februar</option>
            <option value="2">März</option>
            <option value="3">April</option>
            <option value="4">Mai</option>
            <option value="5">Juni</option>
            <option value="6">Juli</option>
            <option value="7">August</option>
            <option value="8">September</option>
            <option value="9">Oktober</option>
            <option value="10">November</option>
            <option value="11">Dezember</option>
          </select>

          <select
              v-model="filterYear"
              class="appearance-none border border-meal-gray-light rounded py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
          >
            <option value="">Alle Jahre</option>
            <option v-for="year in availableYears" :key="year" :value="year">{{ year }}</option>
          </select>
        </div>
      </div>

      <!-- Mahlzeiten-Liste -->
      <div class="bg-white rounded-xl shadow-meal p-6 mb-8">
        <h2 class="text-2xl font-header font-bold text-meal-gray-dark mb-6">Alle Mahlzeiten</h2>

        <div v-if="filteredMeals.length === 0" class="text-center py-8 text-meal-gray">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto mb-4 text-meal-gray-light" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
          </svg>
          <p class="text-xl mb-2">Keine Mahlzeiten gefunden</p>
          <p>{{ searchQuery || filterMonth || filterYear ? 'Passe die Filter an, um Ergebnisse zu sehen.' : 'Erfasse deine erste gemeinsame Mahlzeit!' }}</p>
        </div>

        <div v-else class="space-y-4">
          <div v-for="meal in filteredMeals" :key="meal.id" class="border border-meal-gray-light rounded-lg overflow-hidden hover:shadow-md transition-shadow duration-200">
            <div class="flex flex-col md:flex-row">
              <!-- Datum Spalte -->
              <div class="bg-meal-light p-4 md:w-24 flex flex-col items-center justify-center text-center">
                <div class="text-2xl font-bold text-meal-primary">{{ new Date(meal.date).getDate() }}</div>
                <div class="text-sm text-meal-gray-dark">{{ getMonthName(new Date(meal.date).getMonth()) }}</div>
                <div class="text-sm text-meal-gray">{{ new Date(meal.date).getFullYear() }}</div>
              </div>

              <!-- Hauptinformationen -->
              <div class="p-4 flex-grow">
                <div class="flex justify-between items-start mb-2">
                  <h3 class="text-xl font-header font-bold text-meal-gray-dark">{{ meal.name }}</h3>
                  <span class="font-bold text-meal-primary">{{ formatCurrency(meal.totalAmount) }}</span>
                </div>

                <p v-if="meal.description" class="text-meal-gray mb-3 line-clamp-2">{{ meal.description }}</p>

                <div class="flex items-center text-sm text-meal-gray mb-2">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M8 9a3 3 0 100-6 3 3 0 000 6zM8 11a6 6 0 016 6H2a6 6 0 016-6zM16 7a1 1 0 10-2 0v1h-1a1 1 0 100 2h1v1a1 1 0 102 0v-1h1a1 1 0 100-2h-1V7z" />
                  </svg>
                  <span>Bezahlt von {{ getUserName(meal.userId) }}</span>
                </div>

                <div class="flex flex-wrap gap-1">
                  <div v-for="participantId in meal.participants" :key="participantId"
                       class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-meal-light">
                    <div class="w-4 h-4 rounded-full flex items-center justify-center text-white font-bold mr-1"
                         :style="{ backgroundColor: getUserColor(participantId) }">
                      {{ getUserInitial(participantId) }}
                    </div>
                    {{ getUserName(participantId) }}
                  </div>
                </div>
              </div>

              <!-- Aktionen -->
              <div class="bg-meal-light p-4 flex md:flex-col justify-center space-x-4 md:space-x-0 md:space-y-4">
                <button @click="selectMeal(meal)" class="text-meal-primary hover:text-meal-dark" title="Details anzeigen">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
                <button @click="editMeal(meal)" class="text-meal-accent hover:text-opacity-80" title="Bearbeiten">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button @click="confirmDelete(meal)" class="text-meal-error hover:text-red-700" title="Löschen">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Dialog für Mahlzeit-Details -->
    <div v-if="selectedMeal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-2xl font-header font-bold text-meal-gray-dark">Details: {{ selectedMeal.name }}</h3>
          <button @click="selectedMeal = null" class="text-meal-gray hover:text-meal-error">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
          <div>
            <p class="text-sm text-meal-gray mb-1">Datum</p>
            <p class="font-bold">{{ formatDate(selectedMeal.date) }}</p>
          </div>

          <div>
            <p class="text-sm text-meal-gray mb-1">Gesamtbetrag</p>
            <p class="font-bold text-meal-primary">{{ formatCurrency(selectedMeal.totalAmount) }}</p>
          </div>

          <div>
            <p class="text-sm text-meal-gray mb-1">Bezahlt von</p>
            <div class="flex items-center">
              <div class="w-6 h-6 rounded-full flex items-center justify-center text-white font-bold mr-2"
                   :style="{ backgroundColor: getUserColor(selectedMeal.userId) }">
                {{ getUserInitial(selectedMeal.userId) }}
              </div>
              {{ getUserName(selectedMeal.userId) }}
            </div>
          </div>

          <div>
            <p class="text-sm text-meal-gray mb-1">Anzahl Teilnehmer</p>
            <p class="font-bold">{{ selectedMeal.participants.length }} Personen</p>
          </div>
        </div>

        <div v-if="selectedMeal.description" class="mb-6">
          <p class="text-sm text-meal-gray mb-1">Beschreibung</p>
          <p class="bg-meal-light p-3 rounded">{{ selectedMeal.description }}</p>
        </div>

        <div class="mb-6">
          <p class="text-sm text-meal-gray mb-2">Kosten pro Person</p>
          <div class="bg-meal-light p-4 rounded-lg">
            <div class="grid grid-cols-1 gap-3">
              <div v-for="participantId in selectedMeal.participants" :key="participantId"
                   class="flex items-center justify-between pb-2 border-b border-white">
                <div class="flex items-center">
                  <div class="w-8 h-8 rounded-full flex items-center justify-center text-white font-bold mr-2"
                       :style="{ backgroundColor: getUserColor(participantId) }">
                    {{ getUserInitial(participantId) }}
                  </div>
                  {{ getUserName(participantId) }}
                </div>
                <div class="font-bold">
                  {{ formatCurrency(selectedMeal.totalAmount / selectedMeal.participants.length) }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="selectedMeal.userId" class="mb-6">
          <p class="text-sm text-meal-gray mb-2">Schulden gegenüber {{ getUserName(selectedMeal.userId) }}</p>
          <div class="space-y-2">
            <div v-for="participantId in selectedMeal.participants.filter(id => id !== selectedMeal.userId)" :key="participantId"
                 class="flex items-center justify-between bg-meal-light p-3 rounded">
              <div class="flex items-center">
                <div class="w-6 h-6 rounded-full flex items-center justify-center text-white font-bold mr-2"
                     :style="{ backgroundColor: getUserColor(participantId) }">
                  {{ getUserInitial(participantId) }}
                </div>
                {{ getUserName(participantId) }}
              </div>
              <div class="font-bold text-meal-error">
                {{ formatCurrency(selectedMeal.totalAmount / selectedMeal.participants.length) }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Dialog zum Löschen bestätigen -->
    <div v-if="mealToDelete" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md">
        <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-4">Mahlzeit löschen</h3>
        <p class="mb-6 text-meal-gray-dark">
          Bist du sicher, dass du die Mahlzeit <span class="font-bold">{{ mealToDelete.name }}</span> vom {{ formatDate(mealToDelete.date) }} löschen möchtest? Diese Aktion kann nicht rückgängig gemacht werden.
        </p>
        <div class="flex justify-end space-x-3">
          <button
              @click="mealToDelete = null"
              class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
          >
            Abbrechen
          </button>
          <button
              @click="deleteMeal"
              class="px-4 py-2 bg-meal-error hover:bg-red-700 text-white rounded-lg transition-colors duration-200"
          >
            Löschen
          </button>
        </div>
      </div>
    </div>

    <!-- Dialog zum Bearbeiten -->
    <div v-if="editingMeal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-header font-bold text-meal-gray-dark">Mahlzeit bearbeiten</h3>
          <button @click="editingMeal = null" class="text-meal-gray hover:text-meal-error">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <form @submit.prevent="updateMeal">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-name">
                Mahlzeit Name
              </label>
              <input
                  id="edit-meal-name"
                  v-model="editingMeal.name"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  type="text"
                  placeholder="z.B. Pizza Abend"
                  required
              />
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-date">
                Datum
              </label>
              <input
                  id="edit-meal-date"
                  v-model="editingMeal.date"
                  type="date"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
              />
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-amount">
                Gesamtbetrag (€)
              </label>
              <input
                  id="edit-meal-amount"
                  v-model="editingMeal.totalAmount"
                  type="number"
                  step="0.01"
                  min="0.01"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
                  placeholder="0.00"
              />
            </div>

            <div>
              <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-payer">
                Bezahlt von
              </label>
              <select
                  id="edit-meal-payer"
                  v-model="editingMeal.userId"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
              >
                <option value="" disabled>Bitte wählen...</option>
                <option v-for="user in users" :key="user.id" :value="user.id">
                  {{ user.name }}
                </option>
              </select>
            </div>

            <div class="md:col-span-2">
              <label class="block text-meal-gray text-sm font-bold mb-2">
                Teilnehmer
              </label>
              <div class="flex flex-wrap gap-2 mb-2">
                <div
                    v-for="user in users"
                    :key="user.id"
                    @click="toggleEditParticipant(user.id)"
                    :class="[
                    'flex items-center px-3 py-2 rounded-full cursor-pointer transition-colors',
                    isEditParticipant(user.id)
                      ? 'bg-meal-primary text-white'
                      : 'bg-meal-gray-light text-meal-gray-dark hover:bg-meal-gray'
                  ]"
                >
                  <div class="w-6 h-6 rounded-full flex items-center justify-center mr-2"
                       :class="isEditParticipant(user.id) ? 'bg-white text-meal-primary' : 'bg-meal-gray text-white'">
                    {{ user.name.charAt(0).toUpperCase() }}
                  </div>
                  {{ user.name }}
                </div>
              </div>
              <div v-if="editingMeal.participants.length === 0" class="text-meal-error text-sm mt-1">
                Bitte wähle mindestens einen Teilnehmer aus.
              </div>
            </div>

            <div class="md:col-span-2">
              <label class="block text-meal-gray text-sm font-bold mb-2" for="edit-meal-description">
                Beschreibung (optional)
              </label>
              <textarea
                  id="edit-meal-description"
                  v-model="editingMeal.description"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  rows="3"
                  placeholder="Details zur Mahlzeit"
              ></textarea>
            </div>
          </div>

          <div class="flex justify-end space-x-3 mt-6">
            <button
                type="button"
                @click="editingMeal = null"
                class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
            >
              Abbrechen
            </button>
            <button
                type="submit"
                class="bg-meal-accent hover:bg-opacity-90 text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
                :disabled="editingMeal.participants.length === 0"
            >
              Speichern
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import router from "../router";
import {query} from "../graphql.ts";
import {mutation} from "../graphql.ts";

// Beispiel-Benutzer (sollten aus einer Store/API kommen)
const users = ref([]);
const meals = ref([]);

async function getUsers() {
  const [data, err] = await query.getUsers();
  if (err) {
    console.error('Error getting purchaser for current user:', err);
    return;
  }
  users.value = data;
}

async function getMeals() {
    const [data, err] = await query.getMeals();
  if (err) {
    console.error('Error getting purchaser for current user:', err);
    return;
  }
  meals.value = data;
}
// UI-Status
const showAddForm = ref(false);
const searchQuery = ref('');
const filterMonth = ref('');
const filterYear = ref('');
const selectedMeal = ref(null);
const mealToDelete = ref(null);
const editingMeal = ref(null);

// Neue Mahlzeit
const newMeal = ref({
  name: '',
  date: new Date().toISOString().split('T')[0],
  totalAmount: '',
  userId: '',
  participants: [],
  description: ''
});

// Farben für Benutzer-Avatare
const avatarColors = [
  '#2E7D32', '#00796B', '#0277BD', '#1565C0', '#4527A0',
  '#6A1B9A', '#AD1457', '#C62828', '#EF6C00', '#FF8F00'
];

// Berechnete Eigenschaften
const filteredMeals = computed(() => {
  return meals.value
      .filter(meal => {
        // Textsuche
        if (searchQuery.value && !meal.name.toLowerCase().includes(searchQuery.value.toLowerCase())) {
          return false;
        }

        // Monatsfilter
        if (filterMonth.value !== '') {
          const mealMonth = new Date(meal.date).getMonth();
          if (parseInt(filterMonth.value) !== mealMonth) {
            return false;
          }
        }

        // Jahresfilter
        if (filterYear.value !== '') {
          const mealYear = new Date(meal.date).getFullYear();
          if (parseInt(filterYear.value) !== mealYear) {
            return false;
          }
        }

        return true;
      })
      .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime());
});

const availableYears = computed(() => {
  const years = new Set();
  meals.value.forEach(meal => {
    years.add(new Date(meal.date).getFullYear());
  });
  return Array.from(years).sort((a, b) => b - a); // Absteigende Sortierung
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

function getUserInitial(userId) {
  const user = users.value.find(u => u.id === userId);
  return user ? user.name.charAt(0).toUpperCase() : '?';
}

function getUserName(userId) {
  const user = users.value.find(u => u.id === userId);
  return user ? user.name : 'Unbekannt';
}

function getMonthName(monthIndex) {
  const months = ['Jan', 'Feb', 'Mär', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Okt', 'Nov', 'Dez'];
  return months[monthIndex];
}

function toggleParticipant(userId) {
  const index = newMeal.value.participants.indexOf(userId);
  if (index === -1) {
    newMeal.value.participants.push(userId);
  } else {
    newMeal.value.participants.splice(index, 1);
  }
}

function isParticipant(userId) {
  return newMeal.value.participants.includes(userId);
}

function toggleEditParticipant(userId) {
  if (!editingMeal.value) return;

  const index = editingMeal.value.participants.indexOf(userId);
  if (index === -1) {
    editingMeal.value.participants.push(userId);
  } else {
    editingMeal.value.participants.splice(index, 1);
  }
}

function isEditParticipant(userId) {
  return editingMeal.value && editingMeal.value.participants.includes(userId);
}

async function saveMeal() {
  if (newMeal.value.participants.length === 0) {
    return; // Mindestens ein Teilnehmer muss ausgewählt sein
  }
  let participants = newMeal.value.participants.map(id => id.toString()).join(",");
  let meal = {
    name: newMeal.value.name,
    date: newMeal.value.date,
    totalAmount: parseFloat(newMeal.value.totalAmount),
    userId: newMeal.value.userId,
    userIds: participants,
    description: newMeal.value.description
  }
  // reset formular values
  newMeal.value = {
    name: '',
    date: new Date().toISOString().split('T')[0],
    totalAmount: '',
    userId: '',
    participants: [],
    description: ''
  }


   const [data, err] = await mutation.createMeal(meal);
  if (err) {
    console.error('Error getting purchaser for current user:', err);
    return;
  }
  await getMeals()
}

function selectMeal(meal) {
  selectedMeal.value = meal;
}

function confirmDelete(meal) {
  mealToDelete.value = meal;
}

function deleteMeal() {
  if (mealToDelete.value) {
    meals.value = meals.value.filter(meal => meal.id !== mealToDelete.value.id);
    mealToDelete.value = null;
  }
}

function editMeal(meal) {
  // Kopie erstellen, um Originalwerte nicht direkt zu ändern
  editingMeal.value = { ...meal, participants: [...meal.participants] };
}

function updateMeal() {
  if (editingMeal.value.participants.length === 0) {
    return; // Mindestens ein Teilnehmer muss ausgewählt sein
  }

  const index = meals.value.findIndex(m => m.id === editingMeal.value.id);

  if (index !== -1) {
    // Aktualisiere mit den bearbeiteten Werten
    meals.value[index] = {
      ...editingMeal.value,
      totalAmount: parseFloat(editingMeal.value.totalAmount)
    };
  }

  editingMeal.value = null;
}

function goBackToHomeScreen() {
  router.push('/homefeed');
}

// Beim Laden der Komponente
onMounted(() => {
  getUsers();
  getMeals();
});
</script>