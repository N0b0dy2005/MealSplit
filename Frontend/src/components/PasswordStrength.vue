<template>
  <div class="mt-1">
    <div class="text-xs sm:text-sm text-meal-gray">
      Passwort-Stärke: <span :class="strengthColor">{{ strength }}</span>
    </div>
    <div class="mt-1 h-1.5 sm:h-2 w-full bg-meal-gray-light rounded-full overflow-hidden">
      <div class="h-full" :class="progressClass" :style="{ width: progressWidth }"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  password: string;
  strength: string;
  strengthColor: string;
  progressClass: string;
  progressWidth: string;
}>();
</script> 