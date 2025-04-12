<template>
  <div class="min-h-screen flex flex-col flex-grow bg-meal-light font-sans pb-8">
    <div class="container mx-auto px-4 mt-4 sm:mt-8">
      <!-- Neue Mahlzeit erstellen -->
      <MealForm class=""
          :users="users"
          :show-form="showAddForm"
          @toggle-form="showAddForm = !showAddForm"
          @save-meal="saveMeal"
      />
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

// Typ-Definitionen
interface Product {
  name: string;
  price: number;
  isSpecific?: boolean;
  specificParticipants?: number[];
  showParticipants?: boolean;
}

interface User {
  id: number;
  name: string;
}

interface Meal {
  id: number;
  name: string;
  date: string;
  totalAmount: number;
  userId: number;
  description?: string;
  participants: number[];
  productsData?: Product[];
}

// Data
const users = ref<User[]>([]);
const meals = ref<Meal[]>([]);
const showAddForm = ref(false);
const searchQuery = ref('');
const filterMonth = ref('');
const filterYear = ref('');
const selectedMeal = ref<Meal | null>(null);
const mealToDelete = ref<Meal | null>(null);
const editingMeal = ref<Meal | null>(null);

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
      // Parse productsData
      if (meal.productsData && typeof meal.productsData === 'string') {
        meal.productsData = JSON.parse(meal.productsData);
      } else if (meal.products && typeof meal.products === 'string') {
        meal.productsData = JSON.parse(meal.products);
      } else if (!meal.productsData) {
        meal.productsData = [];
      }

      // Ensure participants is always an array of numbers
      if (!meal.participants) {
        if (meal.userIds && typeof meal.userIds === 'string') {
          meal.participants = meal.userIds.split(',').map(id => parseInt(id));
        } else {
          meal.participants = [];
        }
      }

      // Log for debugging
      console.log(`Meal ${meal.name} has ${meal.productsData.length} products and ${meal.participants.length} participants`);
    } catch (e) {
      console.error('Error parsing meal data', e);
      meal.productsData = [];
      meal.participants = [];
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
    produkts: productsDataString
  };

  console.log("Saving meal with products:", meal);



  const [data, err] = await mutation.createMeal(meal);
  if (err) {
    console.error('Error creating meal:', err,data);
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
  console.log("Mahlzeit ausgewählt:", meal);

  // Create a deep copy with all necessary properties
  const mealCopy = {
    ...meal,
    participants: Array.isArray(meal.participants) ? [...meal.participants] : [],
    productsData: Array.isArray(meal.productsData) ? [...meal.productsData] : []
  };

  // Ensure participants is correctly set if only userIds exists
  if ((!mealCopy.participants || mealCopy.participants.length === 0) && meal.userIds) {
    try {
      if (typeof meal.userIds === 'string') {
        mealCopy.participants = meal.userIds.split(',').map(id => parseInt(id));
      }
    } catch (e) {
      console.error('Error parsing participants', e);
      mealCopy.participants = [];
    }
  }

  // Ensure productsData is correctly set
  if (!mealCopy.productsData || mealCopy.productsData.length === 0) {
    if (meal.products && typeof meal.products === 'string') {
      try {
        mealCopy.productsData = JSON.parse(meal.products);
        console.log("Parsed products for detail view:", mealCopy.productsData);
      } catch (e) {
        console.error('Error parsing products for detail view', e);
        mealCopy.productsData = [];
      }
    }
  }

  selectedMeal.value = mealCopy;
}

function editMeal(meal) {
  console.log("Mahlzeit bearbeiten:", meal);

  // Create a deep copy to avoid modifying the original
  let mealCopy = {
    ...meal,
    participants: Array.isArray(meal.participants) ? [...meal.participants] : [],
    productsData: meal.productsData ? [...meal.productsData] : []
  };

  // Ensure participants is correctly set if only userIds exists
  if ((!mealCopy.participants || mealCopy.participants.length === 0) && meal.userIds) {
    try {
      if (typeof meal.userIds === 'string') {
        mealCopy.participants = meal.userIds.split(',').map(id => parseInt(id));
      }
    } catch (e) {
      console.error('Error parsing participants', e);
      mealCopy.participants = [];
    }
  }

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
  console.log("Mahlzeit löschen:", meal);

  // Create a deep copy with all necessary properties
  const mealCopy = {
    ...meal,
    date: meal.date || new Date().toISOString().split('T')[0]
  };

  mealToDelete.value = mealCopy;
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