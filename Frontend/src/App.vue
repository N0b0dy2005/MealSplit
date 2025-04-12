<template>
  <div class="flex flex-col min-h-screen">
    <Navbar v-if="route.path !== '/login'" />
    <div class="flex-grow">
      <router-view/>
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from './components/Navbar.vue';
import { useRoute } from 'vue-router';

const route = useRoute();
</script>
