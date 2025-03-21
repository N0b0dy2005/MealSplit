<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <!-- Header -->
    <header class="bg-meal-primary text-white p-4 shadow-md mb-8">
      <div class="container mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
          </svg>
          <h1 class="text-2xl font-header font-bold">Mahlzeiten verwalten</h1>
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
      <!-- Neue Mahlzeit erstellen -->
      <MealForm
          :users="users"
          :show-form="showAddForm"
          @toggle-form="showAddForm = !showAddForm"
          @save-meal="saveMeal"
      />

      <!-- Filter und Suche -->
      <MealFilterBar
          v-model:search-query="searchQuery"
          v-model:filter-month="filterMonth"
          v-model:filter-year="filterYear"
          :available-years="availableYears"
      />

      <MealList
          :meals="filteredMeals"
          :users="users"
          :is-filtered="!!searchQuery || !!filterMonth || !!filterYear"
          @select-meal="selectMeal"
          @edit-meal="editMeal"
          @confirm-delete="confirmDelete"
      />
    </div>

    <!-- Dialoge -->
    <MealDetailDialog
        v-if="selectedMeal"
        :meal="selectedMeal"
        :users="users"
        @close="selectedMeal = null"
    />

    <MealDeleteDialog
        v-if="mealToDelete"
        :meal="mealToDelete"
        @close="mealToDelete = null"
        @delete="deleteMeal"
    />

    <MealEditDialog
        v-if="editingMeal"
        :meal="editingMeal"
        :users="users"
        @close="editingMeal = null"
        @update="updateMeal"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import router from "../router";
import { query, mutation } from "../graphql";
import MealForm from "../components/meal/MealForm.vue";
import MealList from "../components/meal/MealList.vue";
import MealFilterBar from "../components/meal/MealFilterBar.vue";
import MealDetailDialog from "../components/meal/MealDetailDialog.vue";
import MealDeleteDialog from "../components/meal/MealDeleteDialog.vue";
import MealEditDialog from "../components/meal/MealEditDialog.vue";

// Data
const users = ref([]);
const meals = ref([]);
const showAddForm = ref(false);
const searchQuery = ref('');
const filterMonth = ref('');
const filterYear = ref('');
const selectedMeal = ref(null);
const mealToDelete = ref(null);
const editingMeal = ref(null);

// Computed properties
const filteredMeals = computed(() => {
  return meals.value
      .filter(meal => {
        // Textsuche
        if (searchQuery.value && !meal.name.toLowerCase().includes(searchQuery.value.toLowerCase())) {
          return false;
        }

        // Monatsfilter
        if (filterMonth.value !== '') {
          const mealMonth = new Date(meal.date).getMonth();
          if (parseInt(filterMonth.value) !== mealMonth) {
            return false;
          }
        }

        // Jahresfilter
        if (filterYear.value !== '') {
          const mealYear = new Date(meal.date).getFullYear();
          if (parseInt(filterYear.value) !== mealYear) {
            return false;
          }
        }

        return true;
      })
      .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime());
});

const availableYears = computed(() => {
  const years = new Set();
  meals.value.forEach(meal => {
    years.add(new Date(meal.date).getFullYear());
  });
  return Array.from(years).sort((a, b) => b - a);
});

// Methods
async function getUsers() {
  const [data, err] = await query.getUsers();
  if (err) {
    console.error('Error getting users:', err);
    return;
  }
  users.value = data;
}

async function getMeals() {
  const [data, err] = await query.getMeals();
  if (err) {
    console.error('Error getting meals:', err);
    return;
  }

  // Parse productsData from JSON string if available
  meals.value = data.map(meal => {
    try {
      if (meal.productsData && typeof meal.productsData === 'string') {
        meal.productsData = JSON.parse(meal.productsData);
      } else if (!meal.productsData) {
        meal.productsData = [];
      }
    } catch (e) {
      console.error('Error parsing productsData', e);
      meal.productsData = [];
    }
    return meal;
  });
}

async function saveMeal(newMeal) {
  let participants = newMeal.participants.map(id => id.toString()).join(",");

  // Calculate total amount
  const totalAmount = newMeal.overrideTotal
      ? parseFloat(newMeal.manualTotalAmount)
      : calculateTotalAmount(newMeal.products);

  // Prepare products data for storage - remove UI-specific properties
  const productsToSave = newMeal.products.map(product => {
    const { showParticipants, ...productToSave } = product;
    return productToSave;
  });

  // Convert to JSON string for storage
  const productsDataString = JSON.stringify(productsToSave);

  let meal = {
    name: newMeal.name,
    date: newMeal.date,
    totalAmount: totalAmount,
    userId: newMeal.userId,
    userIds: participants,
    description: newMeal.description,
   // productsData: productsDataString // Include the products data
  };

  console.log("Saving meal with products:", meal);

  const [data, err] = await mutation.createMeal(meal);
  if (err) {
    console.error('Error creating meal:', err);
    return;
  }

  await getMeals();
  showAddForm.value = false;
}

function calculateTotalAmount(products) {
  if (!products || products.length === 0) {
    return 0;
  }

  return products.reduce((total, product) => {
    return total + parseFloat(product.price || 0);
  }, 0);
}

function selectMeal(meal) {
  selectedMeal.value = meal;
}

function editMeal(meal) {
  // Create a deep copy to avoid modifying the original
  let mealCopy = {
    ...meal,
    participants: [...meal.participants],
    productsData: meal.productsData ? [...meal.productsData] : []
  };

  // Ensure productsData is an array
  if (!Array.isArray(mealCopy.productsData)) {
    mealCopy.productsData = [];
  }

  // Add UI-specific properties
  mealCopy.productsData = mealCopy.productsData.map(product => ({
    ...product,
    showParticipants: false
  }));

  editingMeal.value = mealCopy;
}

function confirmDelete(meal) {
  mealToDelete.value = meal;
}

async function deleteMeal() {
  if (!mealToDelete.value) return;

  try {
    const [data, err] = await mutation.deleteMeal(mealToDelete.value.id);
    if (err) {
      console.error('Error deleting meal:', err);
      return;
    }

    // Update local state
    meals.value = meals.value.filter(meal => meal.id !== mealToDelete.value.id);
    mealToDelete.value = null;
  } catch (error) {
    console.error('Error deleting meal:', error);
  }
}

async function updateMeal(updatedMeal) {
  try {
    let participants = updatedMeal.participants.map(id => id.toString()).join(",");

    // Remove UI-specific properties from products
    const productsToSave = updatedMeal.productsData.map(({ showParticipants, ...product }) => product);

    let meal = {
      id: updatedMeal.id,
      name: updatedMeal.name,
      date: updatedMeal.date,
      totalAmount: parseFloat(updatedMeal.totalAmount),
      userId: updatedMeal.userId,
      userIds: participants,
      description: updatedMeal.description,
      productsData: JSON.stringify(productsToSave)
    };

    const [data, err] = await mutation.updateMeal(meal);
    if (err) {
      console.error('Error updating meal:', err);
      return;
    }

    await getMeals();
    editingMeal.value = null;
  } catch (error) {
    console.error('Error updating meal:', error);
  }
}

function goBackToHomeScreen() {
  router.push('/homefeed');
}

onMounted(() => {
  getUsers();
  getMeals();
});
</script>