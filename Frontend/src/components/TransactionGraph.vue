<template>
  <div class="meal-graph-container w-full bg-meal-white rounded-xl shadow-meal p-4 md:p-6">
    <!-- Header mit Titel und Controls -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 mb-6">
      <h2 class="text-xl font-bold text-meal-primary">
        {{ activeView === 'yearly' ? 'Mahlzeiten-Ausgaben pro Monat' : 'Kosten im Monat' }}
      </h2>
      
      <div class="flex items-center gap-2">
        <!-- View Selector mit besserer Optik -->
        <div class="view-selector flex bg-meal-light rounded-lg p-1">
          <button
              @click="activeView = 'yearly'"
              :class="[
              'px-3 py-1.5 text-sm font-medium rounded-md transition-all duration-200',
              activeView === 'yearly'
                ? 'bg-meal-primary text-white shadow-sm'
                : 'text-meal-gray-dark hover:bg-meal-light-dark'
            ]"
          >
            Jahresübersicht
          </button>
          <button
              @click="activeView = 'monthly'"
              :class="[
              'px-3 py-1.5 text-sm font-medium rounded-md transition-all duration-200',
              activeView === 'monthly'
                ? 'bg-meal-primary text-white shadow-sm'
                : 'text-meal-gray-dark hover:bg-meal-light-dark'
            ]"
          >
            Monatsdetail
          </button>
        </div>

        <!-- Refresh Button mit besserem Design -->
        <button
            @click="refreshData"
            class="bg-meal-light text-meal-primary hover:bg-meal-light-dark p-2 rounded-md transition-all duration-200 flex items-center"
            :disabled="isLoading"
            aria-label="Daten aktualisieren"
        >
          <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              :class="{ 'animate-spin': isLoading }"
              viewBox="0 0 20 20"
              fill="currentColor"
          >
            <path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Month Navigator (nur in Monatsansicht) mit besserem Design -->
    <div v-if="activeView === 'monthly'" class="flex justify-center items-center mb-4 bg-meal-light rounded-lg py-2 px-4">
      <button
          @click="navigateMonth(-1)"
          class="p-1.5 text-meal-primary hover:bg-white hover:shadow-sm rounded-md transition-all duration-200"
          aria-label="Vorheriger Monat"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
        </svg>
      </button>
      <span class="mx-4 font-medium text-meal-gray-dark">{{ currentMonthName }} {{ currentYear }}</span>
      <button
          @click="navigateMonth(1)"
          class="p-1.5 text-meal-primary hover:bg-white hover:shadow-sm rounded-md transition-all duration-200"
          aria-label="Nächster Monat"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
        </svg>
      </button>
    </div>

    <!-- Error Message mit besserem Design -->
    <div v-if="error" class="bg-red-50 border-l-4 border-red-400 text-red-700 p-4 rounded-md mb-4">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg class="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="ml-3">
          <p>{{ error }}</p>
        </div>
      </div>
    </div>

    <!-- Chart Container mit verbessertem Design -->
    <div class="chart-container bg-white rounded-lg border border-meal-light p-4 w-full h-[350px] md:h-[400px]">
      <!-- Loading State mit verbesserter Animation -->
      <div v-if="isLoading" class="w-full h-full flex justify-center items-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-t-2 border-meal-primary"></div>
      </div>

      <!-- Empty State mit besserem Design -->
      <div v-else-if="!hasData" class="w-full h-full flex justify-center items-center text-meal-gray-dark">
        <div class="text-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto mb-3 text-meal-gray" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <p class="text-lg font-medium">Keine Mahlzeiten gefunden</p>
          <p class="text-sm text-meal-gray mt-1">Füge neue Mahlzeiten hinzu, um Statistiken zu sehen</p>
        </div>
      </div>

      <!-- Yearly Chart - NUR Balkendiagramm -->
      <div v-else-if="activeView === 'yearly'" class="h-full">
        <apexchart
            type="bar"
            height="90%"
            :options="yearlyChartOptions"
            :series="yearlyChartSeries">
        </apexchart>
        
        <!-- Manuelle Monatsbeschriftungen unter dem Diagramm -->
        <div class="flex justify-between px-6 mt-2 text-xs text-meal-gray">
          <span v-for="(month, index) in ['Jan', 'Feb', 'Mär', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Okt', 'Nov', 'Dez']" 
                :key="index" 
                class="text-center">
            {{ month }}
          </span>
        </div>
      </div>

      <!-- Monthly Chart - NUR Liniendiagramm -->
      <apexchart
          v-else
          type="line"
          height="100%"
          :options="monthlyChartOptions"
          :series="monthlyChartSeries">
      </apexchart>
    </div>

    <!-- Summary Cards mit verbessertem Design -->
    <div v-if="hasData && !isLoading" class="mt-6 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <!-- Yearly Summary Cards -->
      <template v-if="activeView === 'yearly'">
        <div class="bg-white rounded-lg p-4 border-l-4 border-indigo-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Mahlzeiten gesamt</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">{{ totalMealsCount }}</p>
        </div>
        
        <div class="bg-white rounded-lg p-4 border-l-4 border-green-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Durchschnitt/Monat</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">{{ avgMealsPerMonth.toFixed(1) }}</p>
        </div>
        
        <div class="bg-white rounded-lg p-4 border-l-4 border-yellow-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Kosten gesamt</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">€ {{ totalMealExpenses.toFixed(2) }}</p>
        </div>
        
        <div class="bg-white rounded-lg p-4 border-l-4 border-red-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Durchschnitt/Mahlzeit</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">€ {{ avgCostPerMeal.toFixed(2) }}</p>
        </div>
      </template>

      <!-- Monthly Summary Cards -->
      <template v-else>
        <div class="bg-white rounded-lg p-4 border-l-4 border-indigo-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Mahlzeiten im Monat</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">{{ monthlyMealCount }}</p>
        </div>
        
        <div class="bg-white rounded-lg p-4 border-l-4 border-green-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Durchschnitt/Tag</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">{{ avgMealsPerDay.toFixed(1) }}</p>
        </div>
        
        <div class="bg-white rounded-lg p-4 border-l-4 border-yellow-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Ausgaben im Monat</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">€ {{ monthlyMealExpenses.toFixed(2) }}</p>
        </div>
        
        <div class="bg-white rounded-lg p-4 border-l-4 border-red-500 shadow-sm">
          <h3 class="text-sm font-medium text-meal-gray-dark mb-1">Durchschnitt/Mahlzeit</h3>
          <p class="text-2xl font-bold text-meal-gray-dark">
            € {{ monthlyMealCount > 0 ? (monthlyMealExpenses / monthlyMealCount).toFixed(2) : '0.00' }}
          </p>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue';
import VueApexCharts from 'vue3-apexcharts';
import { query } from "../graphql.ts";

export default {
  components: {
    apexchart: VueApexCharts,
  },
  setup() {
    // State-Variablen
    const activeView = ref('yearly');
    const internalMeals = ref([]);
    const currentDate = ref(new Date());
    const isLoading = ref(false);
    const error = ref(null);

    // Berechnete Eigenschaften für Datumsverarbeitung
    const currentMonthName = computed(() => {
      return new Intl.DateTimeFormat('de-DE', { month: 'long' }).format(currentDate.value);
    });

    const currentYear = computed(() => {
      return currentDate.value.getFullYear();
    });

    const hasData = computed(() => {
      return internalMeals.value.length > 0;
    });

    // Berechnungen für Jahresansicht
    const yearlyMeals = computed(() => {
      return internalMeals.value.filter(m => {
        const date = new Date(m.date);
        return date.getFullYear() === currentYear.value;
      });
    });

    const totalMealsCount = computed(() => {
      return yearlyMeals.value.length;
    });

    const totalMealExpenses = computed(() => {
      return yearlyMeals.value.reduce((sum, m) => sum + parseFloat(m.totalAmount || 0), 0);
    });

    const avgMealsPerMonth = computed(() => {
      // Anzahl der Monate mit Daten (für genaueren Durchschnitt)
      const monthsWithData = new Set(
        yearlyMeals.value.map(m => new Date(m.date).getMonth())
      ).size;
      
      return monthsWithData > 0 ? totalMealsCount.value / monthsWithData : 0;
    });

    const avgCostPerMeal = computed(() => {
      return totalMealsCount.value > 0 ? totalMealExpenses.value / totalMealsCount.value : 0;
    });

    // Berechnungen für Monatsansicht
    const monthlyMeals = computed(() => {
      return internalMeals.value.filter(m => {
        const date = new Date(m.date);
        return date.getMonth() === currentDate.value.getMonth() &&
            date.getFullYear() === currentDate.value.getFullYear();
      });
    });

    const monthlyMealExpenses = computed(() => {
      return monthlyMeals.value.reduce((sum, m) => sum + parseFloat(m.totalAmount || 0), 0);
    });

    const monthlyMealCount = computed(() => {
      return monthlyMeals.value.length;
    });

    const avgMealsPerDay = computed(() => {
      // Anzahl der Tage mit Daten im Monat
      const daysWithData = new Set(
        monthlyMeals.value.map(m => new Date(m.date).getDate())
      ).size;
      
      return daysWithData > 0 ? monthlyMealCount.value / daysWithData : 0;
    });

    // Chartdaten für Jahresansicht - Balkendiagramm für Kosten der Mahlzeiten
    const yearlyChartSeries = computed(() => {
      const mealExpenseSeries = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

      internalMeals.value.forEach(meal => {
        const date = new Date(meal.date);
        const month = date.getMonth();
        const year = date.getFullYear();

        if (year === currentYear.value) {
          mealExpenseSeries[month] += parseFloat(meal.totalAmount || 0);
        }
      });

      return [{
        name: 'Mahlzeiten Kosten',
        data: mealExpenseSeries
      }];
    });

    // Chart-Optionen für Jahresansicht - Balkendiagramm
    const yearlyChartOptions = ref({
      chart: {
        type: 'bar',
        toolbar: {
          show: false
        },
        fontFamily: 'Quicksand, sans-serif',
        background: 'transparent',
        animations: {
          enabled: true,
          easing: 'easeinout',
          speed: 800
        }
      },
      plotOptions: {
        bar: {
          distributed: true,
          borderRadius: 8,
          columnWidth: '70%',
          colors: {
            ranges: [{
              from: 0,
              to: 5000,
              color: '#FF6B6B'
            }]
          }
        }
      },
      dataLabels: {
        enabled: false
      },
      grid: {
        borderColor: '#f8f8f8',
        row: {
          colors: ['#f8f8f8', 'transparent'],
          opacity: 0.2
        }
      },
      xaxis: {
        // Verstecke die X-Achsen-Beschriftungen, wir zeigen sie manuell darunter an
        labels: {
          show: false
        },
        axisTicks: {
          show: false
        },
        axisBorder: {
          show: false
        }
      },
      yaxis: {
        title: {
          text: 'Kosten (€)',
          style: {
            fontSize: '13px',
            fontWeight: 500,
            fontFamily: 'Quicksand, sans-serif',
            color: '#FF6B6B'
          }
        },
        labels: {
          formatter: function(val) {
            return "€" + val.toFixed(0);
          },
          style: {
            fontSize: '12px',
            fontFamily: 'Quicksand, sans-serif',
            colors: '#64748b'
          }
        }
      },
      tooltip: {
        y: {
          formatter: function(val) {
            return "€" + val.toFixed(2);
          }
        }
      },
      colors: ['#FF6B6B', '#FF8585', '#FFACAC', '#FFC2C2', '#FFD8D8'],
      responsive: [
        {
          breakpoint: 480,
          options: {
            dataLabels: {
              enabled: false
            }
          }
        }
      ]
    });

    // Chartdaten für Monatsansicht - Liniendiagramm für Ausgaben
    const monthlyChartSeries = computed(() => {
      const daysInMonth = new Date(
          currentDate.value.getFullYear(),
          currentDate.value.getMonth() + 1,
          0
      ).getDate();

      const dailyExpenseData = Array(daysInMonth).fill(0);

      internalMeals.value.forEach(meal => {
        const date = new Date(meal.date);

        if (date.getMonth() === currentDate.value.getMonth() &&
            date.getFullYear() === currentDate.value.getFullYear()) {
          const day = date.getDate() - 1; // Arrays beginnen bei 0
          dailyExpenseData[day] += parseFloat(meal.totalAmount || 0);
        }
      });

      return [
        {
          name: 'Tägliche Ausgaben',
          type: 'line',
          data: dailyExpenseData
        }
      ];
    });

    // Chart-Optionen für Monatsansicht - Liniendiagramm für Ausgaben
    const monthlyChartOptions = computed(() => {
      const daysInMonth = new Date(
          currentDate.value.getFullYear(),
          currentDate.value.getMonth() + 1,
          0
      ).getDate();

      const categories = [];
      for (let i = 1; i <= daysInMonth; i++) {
        categories.push(i.toString());
      }

      return {
        chart: {
          type: 'line',
          toolbar: {
            show: false
          },
          fontFamily: 'Quicksand, sans-serif',
          background: 'transparent',
          animations: {
            enabled: true,
            easing: 'easeinout',
            speed: 800
          }
        },
        stroke: {
          curve: 'smooth',
          width: 4,
          lineCap: 'round'
        },
        fill: {
          type: 'gradient',
          gradient: {
            shade: 'light',
            type: 'vertical',
            shadeIntensity: 0.3,
            opacityFrom: 0.8,
            opacityTo: 0.3,
            stops: [0, 90, 100]
          }
        },
        grid: {
          borderColor: '#f8f8f8',
          row: {
            colors: ['#f8f8f8', 'transparent'],
            opacity: 0.2
          }
        },
        xaxis: {
          categories: categories,
          axisBorder: {
            show: false
          },
          axisTicks: {
            show: false
          },
          title: {
            text: 'Tag im Monat',
            style: {
              fontSize: '13px',
              fontWeight: 500,
              fontFamily: 'Quicksand, sans-serif',
              color: '#64748b'
            }
          },
          labels: {
            style: {
              fontSize: '12px',
              fontFamily: 'Quicksand, sans-serif',
              colors: '#64748b'
            }
          }
        },
        yaxis: {
          title: {
            text: 'Ausgaben (€)',
            style: {
              fontSize: '13px',
              fontWeight: 500,
              fontFamily: 'Quicksand, sans-serif',
              color: '#FF6B6B'
            }
          },
          labels: {
            formatter: function(val) {
              return "€" + val.toFixed(0);
            },
            style: {
              fontSize: '12px',
              fontFamily: 'Quicksand, sans-serif',
              colors: '#64748b'
            }
          }
        },
        tooltip: {
          shared: true,
          intersect: false,
          theme: 'light',
          style: {
            fontSize: '12px',
            fontFamily: 'Quicksand, sans-serif'
          },
          y: {
            formatter: function(val) {
              return "€ " + val.toFixed(2);
            }
          }
        },
        colors: ['#FF6B6B'],
        markers: {
          size: 0
        },
        responsive: [
          {
            breakpoint: 480,
            options: {
              legend: {
                position: 'bottom',
                offsetY: 7
              }
            }
          }
        ]
      };
    });

    // Methode zum Navigieren zwischen Monaten
    const navigateMonth = (direction) => {
      const newDate = new Date(currentDate.value);
      newDate.setMonth(newDate.getMonth() + direction);
      currentDate.value = newDate;
    };

    // Daten abrufen
    const getMeals = async () => {
      try {
        isLoading.value = true;
        error.value = null;

        const [data, err] = await query.getMeals();
        if (err) {
          error.value = 'Fehler beim Laden der Daten: ' + (err.message || 'Unbekannter Fehler');
          console.error('Error getting meals:', err);
          return;
        }

        // Array leeren, um Duplikate zu vermeiden
        internalMeals.value = [];

        if (Array.isArray(data)) {
          internalMeals.value = data.map(meal => ({
            id: meal.id,
            name: meal.name,
            date: meal.date,
            totalAmount: meal.totalAmount,
            userId: meal.userId,
            description: meal.description,
            createDate: meal.createDate,
            updatedDate: meal.updatedDate,
            userIds: meal.userIds,
            products: meal.products
          }));
        } else {
          error.value = 'Unerwartetes Datenformat erhalten';
        }
      } catch (e) {
        error.value = 'Unerwarteter Fehler: ' + (e.message || 'Unbekannter Fehler');
        console.error('Unexpected error:', e);
      } finally {
        isLoading.value = false;
      }
    };

    // Daten aktualisieren
    const refreshData = () => {
      getMeals();
    };

    // Daten laden, wenn Komponente geladen wird
    onMounted(() => {
      getMeals();
    });

    return {
      activeView,
      internalMeals,
      currentDate,
      currentMonthName,
      currentYear,
      yearlyChartSeries,
      yearlyChartOptions,
      monthlyChartSeries,
      monthlyChartOptions,
      navigateMonth,
      refreshData,
      isLoading,
      error,
      hasData,
      totalMealsCount,
      totalMealExpenses,
      monthlyMealExpenses,
      monthlyMealCount,
      avgMealsPerMonth,
      avgCostPerMeal,
      avgMealsPerDay
    };
  }
};
</script>

<style scoped>
.meal-graph-container {
  --bg-gradient: linear-gradient(to bottom right, #ffffff, #f8fafc);
  background-image: var(--bg-gradient);
}
</style>