<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <!-- Header -->
    <header class="bg-meal-primary text-white p-4 shadow-md mb-8">
      <div class="container mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <h1 class="text-2xl font-header font-bold">Profil bearbeiten</h1>
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
      <!-- Erfolgs-/Fehlermeldung -->
      <div v-if="message" :class="`mb-6 p-4 rounded-lg ${message.type === 'success' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`">
        {{ message.text }}
      </div>

      <!-- Hauptinhalt in Tabs -->
      <div class="bg-white rounded-xl shadow-meal overflow-hidden mb-8">
        <div class="flex border-b border-meal-gray-light">
          <button
              @click="activeTab = 'profile'"
              :class="[
              'py-4 px-6 font-bold text-center flex-1',
              activeTab === 'profile'
                ? 'text-meal-primary border-b-2 border-meal-primary'
                : 'text-meal-gray hover:text-meal-gray-dark'
            ]"
          >
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />
              </svg>
              Persönliche Daten
            </span>
          </button>
          <button
              @click="activeTab = 'security'"
              :class="[
              'py-4 px-6 font-bold text-center flex-1',
              activeTab === 'security'
                ? 'text-meal-primary border-b-2 border-meal-primary'
                : 'text-meal-gray hover:text-meal-gray-dark'
            ]"
          >
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
              </svg>
              Sicherheit
            </span>
          </button>
          <button
              @click="activeTab = 'preferences'"
              :class="[
              'py-4 px-6 font-bold text-center flex-1',
              activeTab === 'preferences'
                ? 'text-meal-primary border-b-2 border-meal-primary'
                : 'text-meal-gray hover:text-meal-gray-dark'
            ]"
          >
            <span class="flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z" clip-rule="evenodd" />
              </svg>
              Einstellungen
            </span>
          </button>
        </div>

        <!-- Tab: Persönliche Daten -->
        <div v-if="activeTab === 'profile'" class="p-6">
          <div class="flex flex-col md:flex-row">
            <!-- Profilbild-Bereich -->
            <div class="md:w-1/3 flex flex-col items-center mb-6 md:mb-0">
              <div class="relative">
                <div class="w-32 h-32 rounded-full overflow-hidden bg-meal-light flex items-center justify-center border-4 border-meal-primary">
                  <img v-if="previewImage || user.profileImage" :src="previewImage || user.profileImage" alt="Profilbild" class="w-full h-full object-cover" />
                  <span v-else class="text-5xl font-bold text-meal-primary">
                    {{ user.name.charAt(0).toUpperCase() }}
                  </span>
                </div>
                <label for="profile-image" class="absolute bottom-0 right-0 bg-meal-accent text-white rounded-full p-2 cursor-pointer hover:bg-opacity-90 transition-colors duration-200">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M4 5a2 2 0 00-2 2v8a2 2 0 002 2h12a2 2 0 002-2V7a2 2 0 00-2-2h-1.586a1 1 0 01-.707-.293l-1.121-1.121A2 2 0 0011.172 3H8.828a2 2 0 00-1.414.586L6.293 4.707A1 1 0 015.586 5H4zm6 9a3 3 0 100-6 3 3 0 000 6z" clip-rule="evenodd" />
                  </svg>
                </label>
                <input type="file" id="profile-image" class="hidden" accept="image/*" @change="onFileSelected">
              </div>
              <button v-if="previewImage || user.profileImage" @click="removeProfileImage" class="mt-4 text-meal-error hover:underline text-sm">
                Profilbild entfernen
              </button>
            </div>

            <!-- Formular -->
            <div class="md:w-2/3 md:pl-8">
              <form @submit.prevent="saveProfile">
                <div class="mb-4">
                  <label class="block text-meal-gray text-sm font-bold mb-2" for="user-name">
                    Name
                  </label>
                  <input
                      id="user-name"
                      v-model="user.name"
                      class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                      type="text"
                      required
                  >
                </div>

                <div class="mb-4">
                  <label class="block text-meal-gray text-sm font-bold mb-2" for="user-email">
                    E-Mail
                  </label>
                  <input
                      id="user-email"
                      v-model="user.email"
                      class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                      type="email"
                      required
                  >
                </div>

                <div class="mb-4">
                  <label class="block text-meal-gray text-sm font-bold mb-2" for="user-phone">
                    Telefonnummer (optional)
                  </label>
                  <input
                      id="user-phone"
                      v-model="user.phone"
                      class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                      type="tel"
                  >
                </div>

                <div class="flex justify-end">
                  <button
                      type="submit"
                      class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
                  >
                    Änderungen speichern
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>

        <!-- Tab: Sicherheit -->
        <div v-if="activeTab === 'security'" class="p-6">
          <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-6">Passwort ändern</h3>

          <form @submit.prevent="changePassword">
            <div class="mb-4">
              <label class="block text-meal-gray text-sm font-bold mb-2" for="current-password">
                Aktuelles Passwort
              </label>
              <div class="relative">
                <input
                    id="current-password"
                    v-model="passwordData.current"
                    :type="showCurrentPassword ? 'text' : 'password'"
                    class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                    required
                >
                <button
                    type="button"
                    @click="showCurrentPassword = !showCurrentPassword"
                    class="absolute inset-y-0 right-0 flex items-center px-3 text-meal-gray"
                >
                  <svg v-if="showCurrentPassword" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                    <path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clip-rule="evenodd" />
                  </svg>
                  <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z" clip-rule="evenodd" />
                    <path d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z" />
                  </svg>
                </button>
              </div>
            </div>

            <div class="mb-4">
              <label class="block text-meal-gray text-sm font-bold mb-2" for="new-password">
                Neues Passwort
              </label>
              <div class="relative">
                <input
                    id="new-password"
                    v-model="passwordData.new"
                    :type="showNewPassword ? 'text' : 'password'"
                    class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                    minlength="8"
                    required
                >
                <button
                    type="button"
                    @click="showNewPassword = !showNewPassword"
                    class="absolute inset-y-0 right-0 flex items-center px-3 text-meal-gray"
                >
                  <svg v-if="showNewPassword" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                    <path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clip-rule="evenodd" />
                  </svg>
                  <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z" clip-rule="evenodd" />
                    <path d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z" />
                  </svg>
                </button>
              </div>
              <div class="mt-1">
                <div class="text-sm text-meal-gray">Passwort-Stärke: <span :class="passwordStrengthColor">{{ passwordStrength }}</span></div>
                <div class="mt-1 h-2 w-full bg-meal-gray-light rounded-full overflow-hidden">
                  <div class="h-full" :class="passwordProgressBarClass" :style="{ width: passwordProgressWidth }"></div>
                </div>
              </div>
            </div>

            <div class="mb-6">
              <label class="block text-meal-gray text-sm font-bold mb-2" for="confirm-password">
                Neues Passwort bestätigen
              </label>
              <input
                  id="confirm-password"
                  v-model="passwordData.confirm"
                  type="password"
                  class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                  required
              >
              <p v-if="passwordMismatch" class="text-meal-error text-sm mt-1">
                Passwörter stimmen nicht überein
              </p>
            </div>

            <div class="flex justify-end">
              <button
                  type="submit"
                  class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
                  :disabled="passwordMismatch || isPasswordWeak"
              >
                Passwort ändern
              </button>
            </div>
          </form>

          <div class="mt-8 pt-8 border-t border-meal-gray-light">
            <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-4">Konto löschen</h3>
            <p class="text-meal-gray mb-4">
              Wenn du dein Konto löschst, werden alle deine persönlichen Daten und Einstellungen unwiderruflich gelöscht.
              Du kannst später ein neues Konto mit der gleichen E-Mail-Adresse erstellen.
            </p>
            <button
                @click="showDeleteAccountConfirmation = true"
                class="bg-meal-error hover:bg-red-700 text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
            >
              Konto löschen
            </button>
          </div>
        </div>

        <!-- Tab: Einstellungen -->
        <div v-if="activeTab === 'preferences'" class="p-6">
          <div class="mb-8">
            <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-4">Benachrichtigungen</h3>

            <div class="space-y-4">
              <div class="flex items-center justify-between py-2 border-b border-meal-gray-light">
                <div>
                  <p class="font-bold text-meal-gray-dark">E-Mail-Benachrichtigungen</p>
                  <p class="text-sm text-meal-gray">Benachrichtigungen per E-Mail erhalten</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="preferences.emailNotifications" class="sr-only peer">
                  <div class="w-11 h-6 bg-meal-gray-light rounded-full peer peer-checked:bg-meal-primary peer-checked:after:translate-x-full after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all"></div>
                </label>
              </div>

              <div class="flex items-center justify-between py-2 border-b border-meal-gray-light">
                <div>
                  <p class="font-bold text-meal-gray-dark">Neue Mahlzeiten</p>
                  <p class="text-sm text-meal-gray">Benachrichtigung wenn du zu einer Mahlzeit hinzugefügt wurdest</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="preferences.newMealNotifications" class="sr-only peer">
                  <div class="w-11 h-6 bg-meal-gray-light rounded-full peer peer-checked:bg-meal-primary peer-checked:after:translate-x-full after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all"></div>
                </label>
              </div>

              <div class="flex items-center justify-between py-2 border-b border-meal-gray-light">
                <div>
                  <p class="font-bold text-meal-gray-dark">Neue Zahlungen</p>
                  <p class="text-sm text-meal-gray">Benachrichtigung bei Zahlungen an dich</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="preferences.paymentNotifications" class="sr-only peer">
                  <div class="w-11 h-6 bg-meal-gray-light rounded-full peer peer-checked:bg-meal-primary peer-checked:after:translate-x-full after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all"></div>
                </label>
              </div>
            </div>
          </div>

          <div class="mb-8">
            <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-4">Anzeige</h3>

            <div class="space-y-4">
              <div class="flex items-center justify-between py-2 border-b border-meal-gray-light">
                <div>
                  <p class="font-bold text-meal-gray-dark">Dark Mode</p>
                  <p class="text-sm text-meal-gray">Dunkles Farbschema für die Benutzeroberfläche</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input type="checkbox" v-model="preferences.darkMode" class="sr-only peer">
                  <div class="w-11 h-6 bg-meal-gray-light rounded-full peer peer-checked:bg-meal-primary peer-checked:after:translate-x-full after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all"></div>
                </label>
              </div>

              <div class="py-2 border-b border-meal-gray-light">
                <p class="font-bold text-meal-gray-dark mb-2">Währung</p>
                <select
                    v-model="preferences.currency"
                    class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                >
                  <option value="EUR">Euro (€)</option>
                  <option value="USD">US Dollar ($)</option>
                  <option value="GBP">Britisches Pfund (£)</option>
                  <option value="CHF">Schweizer Franken (CHF)</option>
                </select>
              </div>

              <div class="py-2 border-b border-meal-gray-light">
                <p class="font-bold text-meal-gray-dark mb-2">Sprache</p>
                <select
                    v-model="preferences.language"
                    class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                >
                  <option value="de">Deutsch</option>
                  <option value="en">Englisch</option>
                  <option value="fr">Französisch</option>
                  <option value="es">Spanisch</option>
                </select>
              </div>
            </div>
          </div>

          <div class="flex justify-end">
            <button
                @click="savePreferences"
                class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-6 rounded-lg transition-colors duration-200"
            >
              Einstellungen speichern
            </button>
          </div>
        </div>
      </div>

      <!-- Statistik-Karten -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-xl shadow-meal p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-2">Gesamtausgaben</h3>
          <p class="text-3xl font-bold text-meal-primary">{{ formatCurrency(statistics.totalSpent) }}</p>
          <p class="text-sm text-meal-gray mt-1">Bei {{ statistics.mealCount }} Mahlzeiten</p>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-2">Durchschnitt</h3>
          <p class="text-3xl font-bold text-meal-secondary">{{ formatCurrency(statistics.averagePerMeal) }}</p>
          <p class="text-sm text-meal-gray mt-1">Pro Mahlzeit</p>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-2">Offene Schulden</h3>
          <p class="text-3xl font-bold" :class="statistics.totalDebt < 0 ? 'text-meal-error' : 'text-meal-positive'">
            {{ formatCurrency(statistics.totalDebt) }}
          </p>
          <p class="text-sm text-meal-gray mt-1">Aktueller Nettobetrag</p>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-2">Teilnahme</h3>
          <p class="text-3xl font-bold text-meal-accent">{{ statistics.participationRate }}%</p>
          <p class="text-sm text-meal-gray mt-1">Beteiligung an Gruppenmahlzeiten</p>
        </div>
      </div>
    </div>

    <!-- Dialog für Kontolöschung-Bestätigung -->
    <div v-if="showDeleteAccountConfirmation" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md">
        <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-4">Konto wirklich löschen?</h3>
        <p class="mb-6 text-meal-gray-dark">
          Bist du sicher, dass du dein Konto löschen möchtest? Diese Aktion kann nicht rückgängig gemacht werden und alle deine Daten werden gelöscht.
        </p>
        <div class="mb-4">
          <label class="block text-meal-gray text-sm font-bold mb-2" for="delete-password">
            Bestätige mit deinem Passwort
          </label>
          <input
              id="delete-password"
              v-model="deleteAccountPassword"
              type="password"
              class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
              placeholder="Dein aktuelles Passwort"
              required
          >
        </div>
        <div class="flex justify-end space-x-3">
          <button
              @click="cancelDeleteAccount"
              class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
          >
            Abbrechen
          </button>
          <button
              @click="confirmDeleteAccount"
              class="px-4 py-2 bg-meal-error hover:bg-red-700 text-white rounded-lg transition-colors duration-200"
              :disabled="!deleteAccountPassword"
          >
            Konto endgültig löschen
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import router from "../router";

// Aktiver Tab
const activeTab = ref('profile');

// Benutzerdaten
const user = ref({
  id: 1,
  name: 'Max Mustermann',
  email: 'max@example.com',
  phone: '+49 123 4567890',
  profileImage: null
});

// Passwort-Änderungsdaten
const passwordData = ref({
  current: '',
  new: '',
  confirm: ''
});
const showCurrentPassword = ref(false);
const showNewPassword = ref(false);

// Einstellungen
const preferences = ref({
  emailNotifications: true,
  newMealNotifications: true,
  paymentNotifications: true,
  darkMode: false,
  currency: 'EUR',
  language: 'de'
});

// UI-Status
const message = ref(null);
const previewImage = ref(null);
const showDeleteAccountConfirmation = ref(false);
const deleteAccountPassword = ref('');

// Beispiel-Statistiken
const statistics = ref({
  totalSpent: 420.75,
  mealCount: 28,
  averagePerMeal: 15.03,
  totalDebt: -32.50,
  participationRate: 87
});

// Berechnete Eigenschaften
const passwordMismatch = computed(() => {
  return passwordData.value.new &&
      passwordData.value.confirm &&
      passwordData.value.new !== passwordData.value.confirm;
});

const passwordStrength = computed(() => {
  const password = passwordData.value.new;
  if (!password) return 'Nicht angegeben';

  // Einfache Passwort-Stärke-Berechnung
  let strength = 0;

  // Mindestlänge
  if (password.length >= 8) strength += 1;
  if (password.length >= 12) strength += 1;

  // Enthält Zahlen
  if (/\d/.test(password)) strength += 1;

  // Enthält Kleinbuchstaben
  if (/[a-z]/.test(password)) strength += 1;

  // Enthält Großbuchstaben
  if (/[A-Z]/.test(password)) strength += 1;

  // Enthält Sonderzeichen
  if (/[^A-Za-z0-9]/.test(password)) strength += 1;

  if (strength <= 2) return 'Schwach';
  if (strength <= 4) return 'Mittel';
  return 'Stark';
});

const passwordStrengthColor = computed(() => {
  switch (passwordStrength.value) {
    case 'Stark': return 'text-green-600';
    case 'Mittel': return 'text-yellow-600';
    case 'Schwach': return 'text-red-600';
    default: return 'text-meal-gray';
  }
});

const passwordProgressBarClass = computed(() => {
  switch (passwordStrength.value) {
    case 'Stark': return 'bg-green-600';
    case 'Mittel': return 'bg-yellow-600';
    case 'Schwach': return 'bg-red-600';
    default: return 'bg-meal-gray';
  }
});

const passwordProgressWidth = computed(() => {
  switch (passwordStrength.value) {
    case 'Stark': return '100%';
    case 'Mittel': return '66%';
    case 'Schwach': return '33%';
    default: return '0%';
  }
});

const isPasswordWeak = computed(() => {
  return passwordStrength.value === 'Schwach' && passwordData.value.new !== '';
});

// Methoden
function formatCurrency(value) {
  const currency = preferences.value.currency;
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: currency
  }).format(value);
}

function onFileSelected(event) {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = e => {
      previewImage.value = e.target.result;
    };
    reader.readAsDataURL(file);
  }
}

function removeProfileImage() {
  previewImage.value = null;
  user.value.profileImage = null;
  // In einem echten System würde man hier auch die Datei auf dem Server löschen
}

function saveProfile() {
  // Hier würdest du die Daten an deinen Server senden
  // Für jetzt simulieren wir eine erfolgreiche Anfrage
  setTimeout(() => {
    // Wenn Profilbild geändert wurde, aktualisiere es
    if (previewImage.value) {
      user.value.profileImage = previewImage.value;
      previewImage.value = null;
    }

    message.value = {
      type: 'success',
      text: 'Profil erfolgreich aktualisiert!'
    };

    // Nachricht nach 3 Sekunden ausblenden
    setTimeout(() => {
      message.value = null;
    }, 3000);
  }, 1000);
}

function changePassword() {
  if (passwordMismatch.value || isPasswordWeak.value) return;

  // Hier würdest du die Passwortänderung an deinen Server senden
  // Für jetzt simulieren wir eine erfolgreiche Anfrage
  setTimeout(() => {
    message.value = {
      type: 'success',
      text: 'Passwort erfolgreich geändert!'
    };

    // Formulardaten zurücksetzen
    passwordData.value = {
      current: '',
      new: '',
      confirm: ''
    };

    // Nachricht nach 3 Sekunden ausblenden
    setTimeout(() => {
      message.value = null;
    }, 3000);
  }, 1000);
}

function savePreferences() {
  // Hier würdest du die Einstellungen an deinen Server senden
  // Für jetzt simulieren wir eine erfolgreiche Anfrage
  setTimeout(() => {
    message.value = {
      type: 'success',
      text: 'Einstellungen erfolgreich gespeichert!'
    };

    // Nachricht nach 3 Sekunden ausblenden
    setTimeout(() => {
      message.value = null;
    }, 3000);
  }, 1000);
}

function cancelDeleteAccount() {
  showDeleteAccountConfirmation.value = false;
  deleteAccountPassword.value = '';
}

function confirmDeleteAccount() {
  if (!deleteAccountPassword.value) return;

  // Hier würdest du die Kontolöschung an deinen Server senden
  // Für jetzt simulieren wir eine erfolgreiche Anfrage
  setTimeout(() => {
    showDeleteAccountConfirmation.value = false;
    message.value = {
      type: 'success',
      text: 'Dein Konto wurde erfolgreich gelöscht. Du wirst zur Anmeldeseite weitergeleitet.'
    };

    // Nach 2 Sekunden zur Login-Seite weiterleiten
    setTimeout(() => {
      router.push('/login');
    }, 2000);
  }, 1000);
}

function goBackToHomeScreen() {
  router.push('/homefeed');
}

// Beim Laden der Komponente
onMounted(() => {
  // Hier würdest du die Benutzerdaten vom Server laden
});
</script>