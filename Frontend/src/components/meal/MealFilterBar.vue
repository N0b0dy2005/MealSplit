<template>
  <div class="bg-white rounded-xl shadow-meal p-4 mb-8 flex flex-wrap gap-4 items-center">
    <div class="flex-grow">
      <input
          :value="searchQuery"
          @input="$emit('update:searchQuery', $event.target.value)"
          type="text"
          placeholder="Mahlzeit suchen..."
          class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
      />
    </div>

    <div class="flex gap-2">
      <select
          :value="filterMonth"
          @change="$emit('update:filterMonth', $event.target.value)"
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
          :value="filterYear"
          @change="$emit('update:filterYear', $event.target.value)"
          class="appearance-none border border-meal-gray-light rounded py-2 px-3 text-meal-gray-dark leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
      >
        <option value="">Alle Jahre</option>
        <option v-for="year in availableYears" :key="year" :value="year">{{ year }}</option>
      </select>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps({
  searchQuery: String,
  filterMonth: String,
  filterYear: String,
  availableYears: Array
});

defineEmits([
  'update:searchQuery',
  'update:filterMonth',
  'update:filterYear'
]);
</script>