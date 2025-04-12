<template>
  <div class="min-h-screen flex flex-col bg-meal-light font-sans pb-8">
    <div class="container mx-auto px-3 mt-3 sm:px-4 sm:mt-8">
      <!-- Statusmeldung -->
      <div v-if="message" :class="`mb-4 p-3 rounded-lg text-sm ${message.type === 'success' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`">
        {{ message.text }}
      </div>

      <!-- Tabs -->
      <div class="bg-white rounded-xl shadow-meal overflow-hidden mb-6">
        <div class="flex border-b border-meal-gray-light">
          <TabButton 
            v-for="tab in tabs" 
            :key="tab.id"
            :active="activeTab === tab.id"
            :icon="tab.icon"
            :text="tab.text"
            @click="activeTab = tab.id"
          />
        </div>

        <!-- Profil-Tab -->
        <div v-if="activeTab === 'profile'" class="p-4 sm:p-6">
          <div class="flex flex-col md:flex-row">
            <!-- Profilbild -->
            <div class="md:w-1/3 flex flex-col items-center mb-6 md:mb-0">
              <div class="relative">
                <div class="w-24 h-24 sm:w-32 sm:h-32 rounded-full overflow-hidden bg-meal-light flex items-center justify-center border-4 border-meal-primary">
                  <img v-if="profileImageSrc" :src="profileImageSrc" alt="Profilbild" class="w-full h-full object-cover" />
                  <span v-else class="text-4xl font-bold text-meal-primary">
                    {{ user.name.charAt(0).toUpperCase() }}
                  </span>
                </div>
                <label for="profile-image" class="absolute bottom-0 right-0 bg-meal-accent text-white rounded-full p-2 cursor-pointer hover:bg-opacity-90 transition-colors">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M4 5a2 2 0 00-2 2v8a2 2 0 002 2h12a2 2 0 002-2V7a2 2 0 00-2-2h-1.586a1 1 0 01-.707-.293l-1.121-1.121A2 2 0 0011.172 3H8.828a2 2 0 00-1.414.586L6.293 4.707A1 1 0 015.586 5H4zm6 9a3 3 0 100-6 3 3 0 000 6z" clip-rule="evenodd" />
                  </svg>
                </label>
                <input type="file" id="profile-image" class="hidden" accept="image/*" @change="onFileSelected">
              </div>
              <button v-if="profileImageSrc" @click="removeProfileImage" class="mt-3 text-meal-error hover:underline text-sm">
                Profilbild entfernen
              </button>
            </div>

            <!-- Formular -->
            <div class="md:w-2/3 md:pl-8">
              <form @submit.prevent="saveProfile">
                <FormField id="user-name" v-model="user.name" label="Name" type="text" required />
                <FormField id="user-email" v-model="user.email" label="E-Mail" type="email" required />
                <FormField id="user-phone" v-model="user.phoneNumber" label="Telefonnummer (optional)" type="tel" />

                <div class="flex justify-end">
                  <button  type="submit" class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-6 text-base rounded-lg transition-colors">
                    Änderungen speichern
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>

        <!-- Sicherheits-Tab -->
        <div v-if="activeTab === 'security'" class="p-4 sm:p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-4">Passwort ändern</h3>

          <div >
            <PasswordField 
              id="current-password" 
              v-model="passwordData.current" 
              label="Aktuelles Passwort" 
              :show-toggle="true" 
              required 
            />
            
            <PasswordField 
              id="new-password" 
              v-model="passwordData.new" 
              label="Neues Passwort" 
              :show-toggle="true" 
              minlength="8" 
              required 
            >
              <PasswordStrength 
                :password="passwordData.new" 
                :strength="passwordStrength"
                :strength-color="passwordStrengthColor"
                :progress-class="passwordProgressBarClass"
                :progress-width="passwordProgressWidth"
              />
            </PasswordField>
            
            <FormField 
              id="confirm-password" 
              v-model="passwordData.confirm" 
              label="Neues Passwort bestätigen" 
              type="password" 
              required 
            />
            <p v-if="passwordMismatch" class="text-meal-error text-sm mt-1">
              Passwörter stimmen nicht überein
            </p>

            <div class="flex justify-end">
              <button
                  @click="saveProfile"
                class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-6 rounded-lg transition-colors"
                :disabled="passwordMismatch">
                Passwort ändern
              </button>
            </div>
          </div>

          <div class="mt-6 pt-6 border-t border-meal-gray-light">
            <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-3">Konto löschen</h3>
            <p class="text-meal-gray text-sm mb-3">
              Wenn du dein Konto löschst, werden alle deine persönlichen Daten und Einstellungen unwiderruflich gelöscht.
              Du kannst später ein neues Konto mit der gleichen E-Mail-Adresse erstellen.
            </p>
            <button
              @click="showDeleteAccountConfirmation = true"
              class="bg-meal-error hover:bg-red-700 text-white font-bold py-2 px-6 rounded-lg transition-colors"
            >
              Konto löschen
            </button>
          </div>
        </div>
      </div>

      <!-- Statistik-Karten -->
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-3 sm:gap-6 mb-4">
        <StatCard 
          v-for="stat in statisticsData" 
          :key="stat.title"
          :title="stat.title" 
          :value="stat.value" 
          :caption="stat.caption"
          :is-currency="stat.isCurrency"
          :is-percentage="stat.isPercentage"
          :value-color="stat.valueColor" 
        />
      </div>
    </div>

    <!-- Konto-Löschungs-Dialog -->
    <ConfirmDialog 
      v-if="showDeleteAccountConfirmation"
      title="Konto wirklich löschen?"
      message="Bist du sicher, dass du dein Konto löschen möchtest? Diese Aktion kann nicht rückgängig gemacht werden und alle deine Daten werden gelöscht."
      :confirm-action="confirmDeleteAccount"
      :cancel-action="cancelDeleteAccount"
    >
      <FormField 
        id="delete-password" 
        v-model="deleteAccountPassword" 
        label="Bestätige mit deinem Passwort" 
        type="password" 
        placeholder="Dein aktuelles Passwort" 
        required 
      />
    </ConfirmDialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import router from "../router";
import {query, mutation, type UpdateUserInput} from "../graphql.ts";
import StatCard from "../components/StatCard.vue";
import TabButton from "../components/TabButton.vue";
import FormField from "../components/FormField.vue";
import PasswordField from "../components/PasswordField.vue";
import PasswordStrength from "../components/PasswordStrength.vue";
import ConfirmDialog from "../components/ConfirmDialog.vue";
import type { User } from "../graphql.ts";

// Benutzerdaten mit lokalen Ergänzungen
interface ExtendedUser {
  id: number;
  name: string;
  email: string;
  phoneNumber?: string;
  profileImage?: string | null;
  debts: string;
  credits?: string;
  createDate?: string;
  updatedDate?: string;
  admin?: boolean;
}

// Benutzer und UI-Status
const activeTab = ref('profile');
const message = ref<{ type: string; text: string } | null>(null);
const previewImage = ref<string | null>(null);
const showDeleteAccountConfirmation = ref(false);
const deleteAccountPassword = ref('');

// Tab-Definitionen
const tabs = [
  { 
    id: 'profile', 
    text: 'Persönliche Daten',
    icon: '<path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />'
  },
  { 
    id: 'security', 
    text: 'Sicherheit',
    icon: '<path fill-rule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />'
  }
];

// Benutzerdaten
const user = ref<ExtendedUser>({
  id: 0,
  name: '',
  email: '',
  phoneNumber: '',
  profileImage: null,
  debts: '0'
});

// Passwort-Daten
const passwordData = ref({
  current: '',
  new: '',
  confirm: ''
});

// Statistiken
const statistics = ref({
  totalSpent: 0,
  mealCount: 0,
  averagePerMeal: 0,
  totalDebt: 0,
  participationRate: 0
});

// Computed Properties
const profileImageSrc = computed(() => previewImage.value || user.value.profileImage);

const passwordMismatch = computed(() => 
  passwordData.value.new && 
  passwordData.value.confirm && 
  passwordData.value.new !== passwordData.value.confirm
);

const passwordStrength = computed(() => {
  const password = passwordData.value.new;
  if (!password) return 'Nicht angegeben';

  let strength = 0;
  if (password.length >= 8) strength += 1;
  if (password.length >= 12) strength += 1;
  if (/\d/.test(password)) strength += 1;
  if (/[a-z]/.test(password)) strength += 1;
  if (/[A-Z]/.test(password)) strength += 1;
  if (/[^A-Za-z0-9]/.test(password)) strength += 1;

  if (strength <= 2) return 'Schwach';
  if (strength <= 4) return 'Mittel';
  return 'Stark';
});

const passwordStrengthColor = computed(() => {
  const colors: Record<string, string> = {
    'Stark': 'text-green-600',
    'Mittel': 'text-yellow-600',
    'Schwach': 'text-red-600'
  };
  return colors[passwordStrength.value] || 'text-meal-gray';
});

const passwordProgressBarClass = computed(() => {
  const colors: Record<string, string> = {
    'Stark': 'bg-green-600',
    'Mittel': 'bg-yellow-600',
    'Schwach': 'bg-red-600'
  };
  return colors[passwordStrength.value] || 'bg-meal-gray';
});

const passwordProgressWidth = computed(() => {
  const widths: Record<string, string> = {
    'Stark': '100%',
    'Mittel': '66%',
    'Schwach': '33%'
  };
  return widths[passwordStrength.value] || '0%';
});

const isPasswordWeak = computed(() => 
  passwordStrength.value === 'Schwach' && passwordData.value.new !== ''
);

const statisticsData = computed(() => [
  {
    title: 'Gesamtausgaben',
    value: statistics.value.totalSpent,
    caption: `Bei ${statistics.value.mealCount} Mahlzeiten`,
    isCurrency: true,
    valueColor: 'text-meal-primary'
  },
  {
    title: 'Durchschnitt',
    value: statistics.value.averagePerMeal,
    caption: 'Pro Mahlzeit',
    isCurrency: true,
    valueColor: 'text-meal-secondary'
  },
  {
    title: 'Offene Beträge',
    value: statistics.value.totalDebt,
    caption: 'Aktueller Nettobetrag',
    isCurrency: true,
    valueColor: statistics.value.totalDebt < 0 ? 'text-meal-error' : 'text-meal-positive'
  },
  {
    title: 'Teilnahme',
    value: statistics.value.participationRate,
    caption: 'Beteiligung an Gruppenmahlzeiten',
    isPercentage: true,
    valueColor: 'text-meal-accent'
  }
]);

// Methoden
function showMessage(type: string, text: string, duration = 3000) {
  message.value = { type, text };
  setTimeout(() => message.value = null, duration);
}

function onFileSelected(event: Event) {
  const input = event.target as HTMLInputElement;
  const file = input.files?.[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = e => {
      previewImage.value = e.target?.result as string;
    };
    reader.readAsDataURL(file);
  }
}

function removeProfileImage() {
  previewImage.value = null;
  user.value.profileImage = null;
}

async function saveProfile() {
  try {
    if (passwordData.value.new !== passwordData.value.confirm) {
      showMessage('error', 'Die Passwörter stimmen nicht überein.');
      return;
    }
    console.log("hier1" ,passwordData.value);
    const [data, err] = await mutation.updateUser({
      id: user.value.id,
      name: user.value.name,
      email: user.value.email,
      phoneNumber: user.value.phoneNumber || '',
      password: passwordData.value.current,
    } as UpdateUserInput);

    if (err) {
      console.error('Fehler beim Aktualisieren des Profils:', err);
      showMessage('error', 'Fehler beim Aktualisieren des Profils.');
      return;
    }
    showMessage('success', 'Profil erfolgreich aktualisiert!');
    showMessage('success', 'Passwort erfolgreich geändert!');
    await loadUserData();
  } catch (error) {
    console.error('Exception beim Aktualisieren des Profils:', error);
    showMessage('error', 'Fehler beim Aktualisieren des Profils.');
  }
}

function cancelDeleteAccount() {
  showDeleteAccountConfirmation.value = false;
  deleteAccountPassword.value = '';
}

function confirmDeleteAccount() {
  if (!deleteAccountPassword.value) return;
  
  try {
    // Simuliere API-Aufruf
    showDeleteAccountConfirmation.value = false;
    showMessage('success', 'Dein Konto wurde erfolgreich gelöscht. Du wirst zur Anmeldeseite weitergeleitet.');
    
    setTimeout(() => router.push('/login'), 2000);
  } catch (error) {
    showMessage('error', 'Fehler beim Löschen des Kontos.');
  }
}

// Daten vom Server laden
async function loadUserData() {
  try {
    // Benutzerdaten laden
    const [userData, err] = await query.getCurrentUser();
    if (err) {
      console.error('Fehler beim Laden der Benutzerdaten:', err);
      showMessage('error', 'Fehler beim Laden der Benutzerdaten.');
      return;
    }

    
    if (userData) {
      user.value = {
        id: userData.id,
        name: userData.name,
        email: userData.email,
        phoneNumber: userData.phoneNumber || '',
        debts: userData.debts || '0',
        credits: userData.credits,
        profileImage: null
      };
      
      console.log("Transformierte Benutzerdaten:", user.value);
    }
    
    // Mahlzeitendaten laden
    const [mealsData, mealsErr] = await query.getMealsForCurrentUser();
    if (mealsErr) {
      console.error('Fehler beim Laden der Mahlzeitendaten:', mealsErr);
      return;
    }
    
    if (mealsData) {
      // Statistiken berechnen
      const totalMeals = mealsData.length || 0;
      let totalCredits = 0;
      
      mealsData.forEach(meal => {
        if (meal.totalAmount) {
          totalCredits += parseFloat(meal.totalAmount);
        }
      });
      
      // Beteiligungsrate berechnen
      let participationRate = 0;
      if (totalMeals > 0) {
        const groupMeals = mealsData.filter(meal => {
          const userIdsList = meal.userIds ? meal.userIds.split(',').filter(id => id.trim() !== '') : [];
          return userIdsList.length > 1;
        });
        
        participationRate = groupMeals.length > 0 ? Math.round((groupMeals.length / totalMeals) * 100) : 0;
      }
      
      statistics.value = {
        totalSpent: totalCredits,
        mealCount: totalMeals,
        averagePerMeal: totalMeals > 0 ? totalCredits / totalMeals : 0,
        totalDebt: parseFloat(user.value.debts) || 0,
        participationRate
      };
    }
  } catch (error) {
    console.error('Fehler beim Laden der Daten:', error);
    showMessage('error', 'Fehler beim Laden der Daten.');
  }
}

onMounted(loadUserData);
</script>