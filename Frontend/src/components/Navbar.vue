<template>
  <!-- Header -->
  <header class="bg-meal-primary text-white p-4 shadow-md">
    <div class="container mx-auto flex items-center justify-between">
      <div class="flex items-center space-x-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 sm:h-8 sm:w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <h1 class="text-xl sm:text-2xl font-header font-bold">{{ currentTitle }}</h1>
      </div>

      <!-- Theme Switcher for Desktop -->
      <div class="hidden md:flex space-x-2 mr-4">
        <button
            @click="switchToGreen"
            class="flex items-center px-3 py-1 rounded-lg text-sm font-medium transition-colors"
            :class="currentTheme === 'green' ? 'bg-meal-accent text-white' : 'bg-meal-dark bg-opacity-30 text-white'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" />
          </svg>
          <span>Grün</span>
        </button>
        <button
            @click="switchToBlue"
            class="flex items-center px-3 py-1 rounded-lg text-sm font-medium transition-colors"
            :class="currentTheme === 'blue' ? 'bg-meal-accent text-white' : 'bg-meal-dark bg-opacity-30 text-white'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z" />
          </svg>
          <span>Blau</span>
        </button>
      </div>

      <!-- Mobile Menu Button -->
      <button
          @click="toggleMobileMenu"
          class="block md:hidden text-white focus:outline-none"
          aria-label="Menu"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path v-if="!mobileMenuOpen" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>

      <!-- Desktop Navigation -->
      <nav class="hidden md:block">
        <ul class="flex space-x-6">
          <li><p @click="goToHome" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer" :class="{ 'text-meal-accent-light font-bold': isActive('/homefeed') }">Startseite</p></li>
          <li><p @click="goToUser" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer" :class="{ 'text-meal-accent-light font-bold': isActive('/user') }">Benutzer</p></li>
          <li><p @click="goToDashboard" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer" :class="{ 'text-meal-accent-light font-bold': isActive('/dashboard') }">Dashboard</p></li>
          <li><p @click="goToMeal" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer" :class="{ 'text-meal-accent-light font-bold': isActive('/meal') }">Mahlzeiten</p></li>
          <li><p @click="goToPayments" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer" :class="{ 'text-meal-accent-light font-bold': isActive('/payments') }">Zahlungen</p></li>
          <li><p @click="goToDebts" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer" :class="{ 'text-meal-accent-light font-bold': isActive('/debts') }">Beträge</p></li>
          <li><p @click="goToProfile" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer" :class="{ 'text-meal-accent-light font-bold': isActive('/profile') }">Profile</p></li>
          <li><p @click="logOut" class="hover:text-meal-accent-light transition-colors duration-200 cursor-pointer">Abmelden</p></li>
        </ul>
      </nav>
    </div>

    <!-- Mobile Navigation Menu -->
    <div
        v-if="mobileMenuOpen"
        class="md:hidden pt-4 pb-2 px-2 transition-all duration-300 ease-in-out"
    >
      <ul class="space-y-2">
        <li><p @click="navigateMobile(goToHome)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center" :class="{ 'bg-meal-accent': isActive('/homefeed') }">Startseite</p></li>
        <li><p @click="navigateMobile(goToUser)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center" :class="{ 'bg-meal-accent': isActive('/user') }">Benutzer</p></li>
        <li><p @click="navigateMobile(goToDashboard)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center" :class="{ 'bg-meal-accent': isActive('/dashboard') }">Dashboard</p></li>
        <li><p @click="navigateMobile(goToMeal)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center" :class="{ 'bg-meal-accent': isActive('/meal') }">Mahlzeiten</p></li>
        <li><p @click="navigateMobile(goToPayments)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center" :class="{ 'bg-meal-accent': isActive('/payments') }">Zahlungen</p></li>
        <li><p @click="navigateMobile(goToDebts)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center" :class="{ 'bg-meal-accent': isActive('/debts') }">Beträge</p></li>
        <li><p @click="navigateMobile(goToProfile)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center" :class="{ 'bg-meal-accent': isActive('/profile') }">Profile</p></li>

        <!-- Theme Switcher for Mobile -->
        <li class="flex space-x-2">
          <button
              @click="navigateMobile(switchToGreen)"
              class="flex-1 flex items-center justify-center py-2 px-3 rounded transition-colors duration-200 cursor-pointer"
              :class="currentTheme === 'green' ? 'bg-meal-accent text-white' : 'bg-meal-dark bg-opacity-30 text-white'"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" />
            </svg>
            <span>Grün</span>
          </button>
          <button
              @click="navigateMobile(switchToBlue)"
              class="flex-1 flex items-center justify-center py-2 px-3 rounded transition-colors duration-200 cursor-pointer"
              :class="currentTheme === 'blue' ? 'bg-meal-accent text-white' : 'bg-meal-dark bg-opacity-30 text-white'"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z" />
            </svg>
            <span>Blau</span>
          </button>
        </li>
        <li><p @click="navigateMobile(logOut)" class="block py-2 px-3 rounded hover:bg-meal-dark transition-colors duration-200 cursor-pointer text-center">Abmelden</p></li>
      </ul>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import router from "../router";
import { useRoute } from 'vue-router';
import {mutation} from "../graphql.ts";

const route = useRoute();
const mobileMenuOpen = ref(false);
const currentTheme = ref(localStorage.getItem('mealTheme') || 'green');

// Dynamischer Titel basierend auf der aktuellen Route
const currentTitle = computed(() => {
  switch (route.path) {
    case '/':
      return 'MealSplit';
    case '/user':
      return 'Benutzer';
    case '/dashboard':
      return 'Dashboard';
    case '/meal':
      return 'Mahlzeiten';
    case '/payments':
      return 'Zahlungen';
    case '/debts':
      return 'Beträge';
    case '/profile':
      return 'Profil';
    default:
      return 'MealSplit';
  }
});

onMounted(() => {
  applyTheme();
});

function switchToGreen() {
  currentTheme.value = 'green';
  localStorage.setItem('mealTheme', currentTheme.value);
  applyTheme();
}

function switchToBlue() {
  currentTheme.value = 'blue';
  localStorage.setItem('mealTheme', currentTheme.value);
  applyTheme();
}

function applyTheme() {
  const root = document.documentElement;

  if (currentTheme.value === 'green') {
    // Grünes Farbschema anwenden
    root.style.setProperty('--meal-primary-color', '#2E7D32');
    root.style.setProperty('--meal-secondary-color', '#66BB6A');
    root.style.setProperty('--meal-light-color', '#E8F5E9');
    root.style.setProperty('--meal-accent-color', '#FF8F00');
    root.style.setProperty('--meal-accent-light-color', '#FFE0B2');
    root.style.setProperty('--meal-dark-color', '#1B5E20');
    root.style.setProperty('--meal-error-color', '#D32F2F');
    root.style.setProperty('--meal-positive-color', '#388E3C');
    root.style.setProperty('--meal-gray-dark-color', '#263238');
    root.style.setProperty('--meal-gray-color', '#607D8B');
    root.style.setProperty('--meal-gray-light-color', '#ECEFF1');
    root.style.setProperty('--meal-shadow-color', 'rgba(46, 125, 50, 0.1)');
  } else {
    // Blaues Farbschema anwenden
    root.style.setProperty('--meal-primary-color', '#1E88E5');
    root.style.setProperty('--meal-secondary-color', '#42A5F5');
    root.style.setProperty('--meal-light-color', '#E3F2FD');
    root.style.setProperty('--meal-accent-color', '#FF9800');
    root.style.setProperty('--meal-accent-light-color', '#FFE0B2');
    root.style.setProperty('--meal-dark-color', '#0D47A1');
    root.style.setProperty('--meal-error-color', '#F44336');
    root.style.setProperty('--meal-positive-color', '#4CAF50');
    root.style.setProperty('--meal-gray-dark-color', '#263238');
    root.style.setProperty('--meal-gray-color', '#546E7A');
    root.style.setProperty('--meal-gray-light-color', '#ECEFF1');
    root.style.setProperty('--meal-shadow-color', 'rgba(30, 136, 229, 0.1)');
  }
}

function goToHome() {
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

function goToDashboard() {
  router.push('/dashboard');
}

function goToProfile() {
  router.push('/profile');
}

async function logOut() {
  const [data, err] = await mutation.logout()
  if (err) {
    console.error(err,data);
    return;
  }

  router.push('/login');
}

function toggleMobileMenu() {
  mobileMenuOpen.value = !mobileMenuOpen.value;
}

function navigateMobile(navigationFunction: () => void) {
  mobileMenuOpen.value = false;
  navigationFunction();
}

function isActive(path: string): boolean {
  if (path === '/homefeed' && route.path === '/') {
    return true;
  }
  return route.path === path;
}
</script>