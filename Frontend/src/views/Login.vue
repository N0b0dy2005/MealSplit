<template>
  <div class="min-h-screen flex flex-grow flex-col bg-meal-light">
    <!-- Header mit Logo -->
    <header class="bg-meal-primary text-white py-4">
      <div class="container mx-auto px-4 text-center">
        <div class="flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
          <h1 class="text-3xl font-header font-bold">MealSplit</h1>
        </div>
      </div>
    </header>

    <!-- Login-Bereich -->
    <main class="flex-grow flex items-center justify-center p-4">
      <div class="max-w-md w-full">
        <!-- Login-Karte -->
        <div class="bg-white rounded-xl shadow-meal p-8">
          <div class="text-center mb-8">
            <h2 class="text-2xl font-header font-bold text-meal-gray-dark">Willkommen zurück!</h2>
            <p class="text-meal-gray mt-2">Melde dich an, um deine gemeinsamen Mahlzeiten zu verwalten</p>
          </div>

          <!-- Login-Formular -->
          <form @submit.prevent="onClickLogin">
            <!-- Fehlermeldung, falls vorhanden -->
            <div v-if="errorMessage" class="mb-6 p-3 bg-red-100 border border-red-300 text-red-600 rounded">
              {{ errorMessage }}
            </div>

            <!-- E-Mail Feld -->
            <div class="mb-6">
              <label class="block text-meal-gray text-sm font-bold mb-2" for="email">
                E-Mail
              </label>
              <input
                  id="email"
                  v-model="inputEmail"
                  type="email"
                  placeholder="deine@email.de"
                  required
                  class="appearance-none border border-meal-gray-light rounded w-full py-3 px-4 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
              >
            </div>

            <!-- Passwort Feld -->
            <div class="mb-6">
              <div class="flex justify-between items-center mb-2">
                <label class="block text-meal-gray text-sm font-bold" for="password">
                  Passwort
                </label>
                <button
                    @click.prevent="showForgotPassword = true"
                    type="button"
                    class="text-sm text-meal-primary hover:text-meal-dark transition-colors duration-200"
                >
                  Passwort vergessen?
                </button>
              </div>
              <div class="relative">
                <input
                    id="password"
                    v-model="inputPassword"
                    :type="showPassword ? 'text' : 'password'"
                    placeholder="Dein Passwort"
                    required
                    class="appearance-none border border-meal-gray-light rounded w-full py-3 px-4 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
                >
                <button
                    type="button"
                    @click="showPassword = !showPassword"
                    class="absolute inset-y-0 right-0 flex items-center px-3 text-meal-gray"
                >
                  <svg v-if="showPassword" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
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

            <!-- Angemeldet bleiben -->
            <div class="mb-6">
              <label class="flex items-center">
                <input type="checkbox" v-model="rememberMe" class="h-4 w-4 text-meal-primary border-meal-gray-light rounded focus:ring-meal-primary">
                <span class="ml-2 text-meal-gray">Angemeldet bleiben</span>
              </label>
            </div>

            <!-- Login-Button -->
            <button
                type="submit"
                class="w-full bg-meal-primary hover:bg-meal-dark text-white font-bold py-3 px-4 rounded-lg transition-colors duration-200 flex items-center justify-center"
            >
              {{ isLoading ? 'Anmelden...' : 'Anmelden' }}
            </button>
          </form>

          <!-- Trenner -->
          <div class="my-6 flex items-center">
            <div class="flex-grow h-px bg-meal-gray-light"></div>
            <span class="px-4 text-sm text-meal-gray">oder</span>
            <div class="flex-grow h-px bg-meal-gray-light"></div>
          </div>

          <!-- Registrieren Button -->
          <button
              @click="showRegisterForm = true"
              class="w-full bg-white border border-meal-primary text-meal-primary hover:bg-meal-light font-bold py-3 px-4 rounded-lg transition-colors duration-200"
          >
            Konto erstellen
          </button>

          <!-- Demo-Login -->
          <div class="mt-6 text-center">
            <button
                @click="handleDemoLogin"
                class="text-meal-primary hover:text-meal-dark transition-colors duration-200"
            >
              Demo-Version ohne Anmeldung starten
            </button>
          </div>
        </div>

        <!-- App-Infos -->
        <div class="mt-8 text-center text-meal-gray text-sm">
          <p>MealSplit hilft dir, gemeinsame Mahlzeiten und Schulden zu verwalten</p>
          <p class="mt-2">© 2023 MealSplit · <a href="#" class="text-meal-primary hover:underline">Datenschutz</a> · <a href="#" class="text-meal-primary hover:underline">AGB</a></p>
        </div>
      </div>
    </main>

    <!-- Passwort vergessen Dialog -->
    <div v-if="showForgotPassword" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-header font-bold text-meal-gray-dark">Passwort zurücksetzen</h3>
          <button @click="showForgotPassword = false" class="text-meal-gray hover:text-meal-error">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <p class="text-meal-gray mb-4">
          Gib deine E-Mail-Adresse ein und wir senden dir einen Link zum Zurücksetzen deines Passworts.
        </p>

        <form @submit.prevent="handlePasswordReset">
          <div class="mb-4">
            <label class="block text-meal-gray text-sm font-bold mb-2" for="reset-email">
              E-Mail
            </label>
            <input
                id="reset-email"
                v-model="resetEmail"
                type="email"
                placeholder="deine@email.de"
                required
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
            >
          </div>

          <div class="flex justify-end space-x-3">
            <button
                type="button"
                @click="showForgotPassword = false"
                class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
            >
              Abbrechen
            </button>
            <button
                type="submit"
                class="px-4 py-2 bg-meal-primary hover:bg-meal-dark text-white rounded-lg transition-colors duration-200"
            >
              Link senden
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Registrierungs Dialog -->
    <div v-if="showRegisterForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md max-h-[90vh] overflow-y-auto">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-header font-bold text-meal-gray-dark">Konto erstellen</h3>
          <button @click="showRegisterForm = false" class="text-meal-gray hover:text-meal-error">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <form @submit.prevent="handleRegister">
          <div class="mb-4">
            <label class="block text-meal-gray text-sm font-bold mb-2" for="register-name">
              Name
            </label>
            <input
                id="register-name"
                v-model="registerName"
                type="text"
                placeholder="Dein vollständiger Name"
                required
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
            >
          </div>

          <div class="mb-4">
            <label class="block text-meal-gray text-sm font-bold mb-2" for="register-email">
              E-Mail
            </label>
            <input
                id="register-email"
                v-model="registerEmail"
                type="email"
                placeholder="deine@email.de"
                required
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
            >
          </div>

          <div class="mb-4">
            <label class="block text-meal-gray text-sm font-bold mb-2" for="register-password">
              Passwort
            </label>
            <input
                id="register-password"
                v-model="registerPassword"
                type="password"
                placeholder="Mindestens 8 Zeichen"
                required
                minlength="8"
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
            >
          </div>

          <div class="mb-6">
            <label class="block text-meal-gray text-sm font-bold mb-2" for="register-password-confirm">
              Passwort bestätigen
            </label>
            <input
                id="register-password-confirm"
                v-model="registerPasswordConfirm"
                type="password"
                placeholder="Passwort wiederholen"
                required
                minlength="8"
                class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
            >
            <p v-if="passwordMismatch" class="text-meal-error text-sm mt-1">
              Passwörter stimmen nicht überein
            </p>
          </div>

          <div class="mb-6">
            <label class="flex items-start">
              <input type="checkbox" v-model="termsAccepted" required class="h-4 w-4 mt-1 text-meal-primary border-meal-gray-light rounded focus:ring-meal-primary">
              <span class="ml-2 text-meal-gray text-sm">
                Ich akzeptiere die <a href="#" class="text-meal-primary hover:underline">Nutzungsbedingungen</a> und die <a href="#" class="text-meal-primary hover:underline">Datenschutzerklärung</a>
              </span>
            </label>
          </div>

          <div class="flex justify-end space-x-3">
            <button
                type="button"
                @click="showRegisterForm = false"
                class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
            >
              Abbrechen
            </button>
            <button
                type="submit"
                :disabled="!canRegister || isRegistering"
                class="px-4 py-2 bg-meal-primary hover:bg-meal-dark text-white rounded-lg transition-colors duration-200"
            >
              <span v-if="isRegistering">Registrierung...</span>
              <span v-else>Registrieren</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import axios from "axios";
import router from "../router";

const inputEmail = ref("felix@r2e2.de");
const inputPassword = ref("test");
const showPassword = ref(false);
const rememberMe = ref(false);
const isLoading = ref(false);
const errorMessage = ref("");

// Passwort vergessen Modal
const showForgotPassword = ref(false);
const resetEmail = ref("");

// Registrierungs-Modal
const showRegisterForm = ref(false);
const registerName = ref("");
const registerEmail = ref("");
const registerPassword = ref("");
const registerPasswordConfirm = ref("");
const termsAccepted = ref(false);
const isRegistering = ref(false);

// Berechnete Eigenschaften
const passwordMismatch = computed(() => {
  return registerPassword.value &&
      registerPasswordConfirm.value &&
      registerPassword.value !== registerPasswordConfirm.value;
});

const canRegister = computed(() => {
  return registerName.value &&
      registerEmail.value &&
      registerPassword.value &&
      registerPassword.value === registerPasswordConfirm.value &&
      termsAccepted.value;
});

// Login-Funktion, die eine Anfrage an /api/login sendet
const onClickLogin = () => {
  isLoading.value = true;
  errorMessage.value = "";

  const formData = new URLSearchParams();

  formData.append('email', inputEmail.value.toLowerCase());
  formData.append('password', inputPassword.value.toLowerCase());
  axios.post('/api/login', formData, { withCredentials: true })
      .then(response => {
        errorMessage.value = "";
        router.push('/homefeed');
      })
      .catch(error => {
        errorMessage.value = error.response && error.response.data ? error.response.data : "Login fehlgeschlagen";
      })
      .finally(() => {
        isLoading.value = false;
      });
};

function handleDemoLogin() {
  router.push('/homefeed');
}

function handlePasswordReset() {
  alert(`Passwort-Reset-Link wurde an ${resetEmail.value} gesendet.`);
  showForgotPassword.value = false;
  resetEmail.value = "";
}

function handleRegister() {
  if (!canRegister.value) return;

  isRegistering.value = true;

  // Simulierte Registrierungsanfrage
  setTimeout(() => {
    alert(`Registrierung erfolgreich! Willkommen ${registerName.value}!`);
    showRegisterForm.value = false;
    isRegistering.value = false;

    // Zurücksetzen des Formulars
    registerName.value = "";
    registerEmail.value = "";
    registerPassword.value = "";
    registerPasswordConfirm.value = "";
    termsAccepted.value = false;

    // Optional: Automatisch anmelden und zur Startseite wechseln
    router.push('/homefeed');
  }, 1500);
}
</script>
