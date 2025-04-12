<template>
  <div class="bg-white rounded-xl shadow-meal p-3 sm:p-6">
    <div class="flex items-center justify-between mb-2 sm:mb-4" v-if="showIcon">
      <h3 class="text-sm sm:text-lg font-header font-bold text-meal-gray-dark">{{ title }}</h3>
      <div class="p-1 sm:p-2 bg-meal-light rounded-lg">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 sm:h-6 sm:w-6" :class="iconColor" fill="none"
             viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="iconPath" />
        </svg>
      </div>
    </div>
    <h3 v-else class="text-sm sm:text-lg font-header font-bold text-meal-gray-dark mb-1 sm:mb-2">{{ title }}</h3>
    
    <p class="text-xl sm:text-3xl font-bold" :class="valueColor">{{ formattedValue }}</p>
    
    <div v-if="caption && changeValue !== undefined" class="flex items-center text-xs sm:text-sm">
      <span v-if="showChangeIndicator" :class="changeValue >= 0 ? 'text-meal-positive' : 'text-meal-error'" class="font-medium">
        {{ changeValue >= 0 ? '+' : '' }}{{ changeValue }}{{ changeIsPercentage ? '%' : '' }}
      </span>
      <span class="text-meal-gray ml-1">{{ caption }}</span>
    </div>
    <p v-else-if="caption" class="text-xs sm:text-sm text-meal-gray mt-1">{{ caption }}</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

// Props Definition
const props = defineProps({
  // Grundlegende Inhalte
  title: {
    type: String,
    required: true
  },
  value: {
    type: [Number, String],
    required: true
  },
  caption: {
    type: String,
    default: ''
  },
  
  // Styling
  valueColor: {
    type: String,
    default: 'text-meal-gray-dark'
  },
  
  // Wertänderung
  changeValue: {
    type: Number,
    default: undefined
  },
  changeIsPercentage: {
    type: Boolean,
    default: true
  },
  showChangeIndicator: {
    type: Boolean,
    default: true
  },
  
  // Icon
  showIcon: {
    type: Boolean,
    default: false
  },
  iconPath: {
    type: String,
    default: 'M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253'
  },
  iconColor: {
    type: String,
    default: 'text-meal-primary'
  },
  
  // Formatierung
  isCurrency: {
    type: Boolean,
    default: false
  },
  isPercentage: {
    type: Boolean,
    default: false
  }
});

// Formatierte Wertausgabe
const formattedValue = computed(() => {
  if (props.isCurrency) {
    return formatCurrency(Number(props.value));
  } else if (props.isPercentage) {
    return `${props.value}%`;
  }
  return props.value;
});

// Hilfsfunktion zum Formatieren von Währungsbeträgen
function formatCurrency(value: number): string {
  return new Intl.NumberFormat('de-DE', {
    style: 'currency',
    currency: 'EUR'
  }).format(value);
}
</script> 