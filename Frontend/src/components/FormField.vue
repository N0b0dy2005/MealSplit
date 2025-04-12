<template>
  <div class="mb-3 sm:mb-4">
    <label class="block text-meal-gray text-xs sm:text-sm font-bold mb-1 sm:mb-2" :for="id">
      {{ label }}
    </label>
    <input
      :id="id"
      :type="type"
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      class="appearance-none border border-meal-gray-light rounded w-full py-2 px-3 text-meal-gray-dark text-sm sm:text-base leading-tight focus:outline-none focus:ring-2 focus:ring-meal-primary"
      :placeholder="placeholder"
      :required="required"
      :minlength="minlength"
    />
  </div>
</template>

<script setup lang="ts">
defineProps<{
  id: string;
  label: string;
  modelValue: string;
  type?: string;
  placeholder?: string;
  required?: boolean;
  minlength?: number;
}>();

defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();
</script> 