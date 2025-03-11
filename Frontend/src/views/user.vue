<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <!-- Header -->
    <header class="bg-meal-primary text-white p-4 shadow-md mb-8">
      <div class="container mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"/>
          </svg>
          <h1 class="text-2xl font-header font-bold">Benutzerverwaltung</h1>
        </div>
        <button @click="goBackToHomeScreen"
                class="text-white hover:text-meal-accent-light transition-colors duration-200">
          <span class="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
            Zurück zur Startseite
          </span>
        </button>
      </div>
    </header>

    <div class="container mx-auto px-4">
      <div class="bg-white rounded-xl shadow-meal p-6 mb-8">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-header font-bold text-meal-gray-dark">Benutzer verwalten</h2>
          <button
              @click="showAddForm = true"
              class="bg-meal-primary hover:bg-meal-dark text-white font-bold py-2 px-4 rounded-lg transition-colors duration-200"
          >
            <span class="flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd"
                      d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
                      clip-rule="evenodd"/>
              </svg>
              Neuer Benutzer
            </span>
          </button>
        </div>

        <!-- Benutzer Liste -->
        <div class="overflow-x-auto">
          <table class="min-w-full bg-white">
            <thead>
            <tr class="bg-meal-light text-meal-dark text-left">
              <th class="py-3 px-4 rounded-tl-lg">Name</th>
              <th class="py-3 px-4">Email</th>
              <th class="py-3 px-4">Offene Schulden</th>
              <th class="py-3 px-4">Ausstehende Zahlungen</th>
              <th class="py-3 px-4 rounded-tr-lg text-right">Aktionen</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="user in users" :key="user.id"
                class="border-b border-meal-gray-light hover:bg-meal-gray-light transition-colors duration-150">
              <td class="py-3 px-4">
                <div class="flex items-center">
                  <div class="w-10 h-10 rounded-full flex items-center justify-center text-white font-bold mr-3"
                       :style="{ backgroundColor: getUserColor(user.id) }">
                    {{ user.name.charAt(0).toUpperCase() }}
                  </div>
                  {{ user.name }}
                </div>
              </td>
              <td class="py-3 px-4">{{ user.email }}</td>
              <td class="py-3 px-4 "
                  :class="{'text-meal-gray-dark ': user.debts < 0, 'text-meal-error': user.debts > 0}">
                {{ formatCurrency(user.debts) }}
              </td>
              <td class="py-3 px-4"
                  :class="{'text-meal-error': user.credits < 0, 'text-meal-positive': user.credits > 0}">
                {{ formatCurrency(user.credits) }}
              </td>
              <td class="py-3 px-4 text-right">
                <button
                    @click="editUser(user)"
                    class="text-meal-primary hover:text-meal-dark mr-2"
                    title="Bearbeiten"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path
                        d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"/>
                  </svg>
                </button>
                <button
                    @click="confirmDelete(user)"
                    class="text-meal-error hover:text-red-700"
                    title="Löschen"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd"
                          d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"
                          clip-rule="evenodd"/>
                  </svg>
                </button>
              </td>
            </tr>
            <tr v-if="users.length === 0">
              <td colspan="5" class="py-4 px-4 text-center text-meal-gray">
                Keine Benutzer vorhanden. Füge deinen ersten Benutzer hinzu!
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Statistik-Karten -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div class="bg-white rounded-xl shadow-meal p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-2">Gesamt Benutzer</h3>
          <p class="text-3xl font-bold text-meal-primary">{{ users.length }}</p>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-2">Offene Schulden</h3>
          <p class="text-3xl font-bold text-meal-error">{{ formatCurrency(totalDebts) }}</p>
        </div>

        <div class="bg-white rounded-xl shadow-meal p-6">
          <h3 class="text-lg font-header font-bold text-meal-gray-dark mb-2">Ausstehende Zahlungen</h3>
          <p class="text-3xl font-bold text-meal-positive">{{ formatCurrency(totalCredits) }}</p>
        </div>
      </div>
    </div>

    <!-- Dialog für Benutzer hinzufügen -->
    <div v-if="showAddForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-header font-bold text-meal-gray-dark">Neuen Benutzer hinzufügen</h3>
          <button @click="showAddForm = false" class="text-meal-gray hover:text-meal-error">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
        <user-form
            :user="newUser"
            @save="saveNewUser"
            @cancel="showAddForm = false"
        />
      </div>
    </div>

    <!-- Dialog für Benutzer bearbeiten -->
    <div v-if="showEditForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-xl font-header font-bold text-meal-gray-dark">Benutzer bearbeiten</h3>
          <button @click="showEditForm = false" class="text-meal-gray hover:text-meal-error">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
        <user-form
            :user="currentUser"
            @save="saveEditedUser"
            @cancel="showEditForm = false"
        />
      </div>
    </div>

    <!-- Dialog für Löschen bestätigen -->
    <div v-if="showDeleteConfirm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-xl p-6 w-full max-w-md">
        <h3 class="text-xl font-header font-bold text-meal-gray-dark mb-4">Benutzer löschen</h3>
        <p class="mb-6 text-meal-gray-dark">
          Bist du sicher, dass du <span class="font-bold">{{ userToDelete?.name }}</span> löschen möchtest? Diese Aktion
          kann nicht rückgängig gemacht werden.
        </p>
        <div class="flex justify-end space-x-3">
          <button
              @click="showDeleteConfirm = false"
              class="px-4 py-2 border border-meal-gray-light rounded-lg text-meal-gray-dark hover:bg-meal-gray-light transition-colors duration-200"
          >
            Abbrechen
          </button>
          <button
              @click="deleteUser"
              class="px-4 py-2 bg-meal-error hover:bg-red-700 text-white rounded-lg transition-colors duration-200"
          >
            Löschen
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import router from "../router";
import {mutation, query} from "../graphql.ts";
import UserForm from "../components/UserForm.vue";

// Daten
const users = ref([]);

async function getUsers() {
  const [data, err] = await query.getUsers();

  console.log("data", data);
  data?.forEach((user, index) => {
    if (parseFloat(user.debts)) {
      // add to total debts but convert user.debts to float
      totalDebts.value += parseFloat(user.debts);
    }
    if (parseFloat(user.credits) > 0) {
      // add to total credits but converrt user.credits to float
      totalCredits.value += parseFloat(user.credits);

    }
  });

  console.log("totalDebts", totalDebts.value);
  console.log("totalCredits", totalCredits.value);

  if (err) {
    console.error('Error getting purchaser for current user:', err);
    return;
  }
  users.value = data;
}

const showAddForm = ref(false);
const showEditForm = ref(false);
const showDeleteConfirm = ref(false);
const currentUser = ref(null);
const userToDelete = ref(null);
const totalDebts = ref(0);
const totalCredits = ref(0);
const newUser = ref({
  name: '',
  email: '',
  debts: 0,
  credits: 0
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

function getUserColor(id) {
  return avatarColors[id % avatarColors.length];
}

function editUser(user) {
  currentUser.value = {...user};
  showEditForm.value = true;
}

function confirmDelete(user) {
  userToDelete.value = user;
}

async function saveNewUser(user) {

  const [data, err] = await mutation.createUser(user.name, user.email);
  if (err) {
    console.error('Error getting purchaser for current user:', err);
    return;
  }
  await getUsers();
  showAddForm.value = false;
}

function saveEditedUser(user) {
  const index = users.value.findIndex(u => u.id === user.id);
  if (index !== -1) {
    users.value.splice(index, 1, user);
  }
  showEditForm.value = false;
  currentUser.value = null;
}

function deleteUser() {
  if (userToDelete.value) {
    users.value = users.value.filter(user => user.id !== userToDelete.value.id);
    showDeleteConfirm.value = false;
    userToDelete.value = null;
  }
}

function goBackToHomeScreen() {
  router.push('/homefeed');
}

onMounted(() => {
  getUsers();
});
</script>