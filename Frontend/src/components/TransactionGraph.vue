<template>
  <div class="transaction-graph-container w-full bg-meal-white rounded-xl shadow-meal p-4 md:p-6">
    <!-- Controls Section -->
    <div class="controls flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 mb-4">
      <!-- View Selector -->
      <div class="view-selector flex">
        <button
            @click="activeView = 'yearly'"
            :class="[
            'px-3 py-2 md:px-4 md:py-2 text-sm md:text-base rounded-l-lg transition-colors',
            activeView === 'yearly'
              ? 'bg-meal-primary text-white'
              : 'bg-meal-light text-meal-gray-dark hover:bg-meal-light hover:bg-opacity-80'
          ]"
        >
          Jährlich
        </button>
        <button
            @click="activeView = 'monthly'"
            :class="[
            'px-3 py-2 md:px-4 md:py-2 text-sm md:text-base rounded-r-lg transition-colors',
            activeView === 'monthly'
              ? 'bg-meal-primary text-white'
              : 'bg-meal-light text-meal-gray-dark hover:bg-meal-light hover:bg-opacity-80'
          ]"
        >
          Monatlich
        </button>
      </div>

      <!-- Month Navigator (nur in Monatsansicht) -->
      <div v-if="activeView === 'monthly'" class="month-navigator flex items-center">
        <button
            @click="navigateMonth(-1)"
            class="p-2 text-meal-primary hover:bg-meal-light rounded-full transition-colors"
            aria-label="Vorheriger Monat"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
          </svg>
        </button>
        <span class="mx-2 font-medium">{{ currentMonthName }} {{ currentYear }}</span>
        <button
            @click="navigateMonth(1)"
            class="p-2 text-meal-primary hover:bg-meal-light rounded-full transition-colors"
            aria-label="Nächster Monat"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
          </svg>
        </button>
      </div>

      <!-- Refresh Button -->
      <button
          @click="refreshData"
          class="text-meal-primary hover:bg-meal-light p-2 rounded-full transition-colors flex items-center"
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

    <!-- Error Message -->
    <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
      <p>{{ error }}</p>
    </div>

    <!-- Chart Container -->
    <div class="chart-container w-full h-[350px] md:h-[400px]">
      <!-- Loading State -->
      <div v-if="isLoading" class="w-full h-full flex justify-center items-center">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-meal-primary"></div>
      </div>

      <!-- Empty State -->
      <div v-else-if="!hasData" class="w-full h-full flex justify-center items-center text-meal-gray">
        <div class="text-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <p>Keine Daten verfügbar</p>
        </div>
      </div>

      <!-- Yearly Chart -->
      <apexchart
          v-else-if="activeView === 'yearly'"
          type="bar"
          height="100%"
          :options="yearlyChartOptions"
          :series="yearlyChartSeries">
      </apexchart>

      <!-- Monthly Chart -->
      <apexchart
          v-else
          type="line"
          height="100%"
          :options="monthlyChartOptions"
          :series="monthlyChartSeries">
      </apexchart>
    </div>

    <!-- Summary Section -->
    <div v-if="hasData && !isLoading" class="mt-4 pt-4 border-t border-meal-gray-light">
      <!-- Yearly Summary -->
      <div v-if="activeView === 'yearly'" class="text-sm text-meal-gray-dark">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <p class="font-medium">Gesamteinnahmen:</p>
            <p class="text-meal-positive text-lg font-bold">€ {{ totalIncome.toFixed(2) }}</p>
          </div>
          <div>
            <p class="font-medium">Gesamtausgaben:</p>
            <p class="text-meal-error text-lg font-bold">€ {{ totalExpenses.toFixed(2) }}</p>
          </div>
        </div>
      </div>

      <!-- Monthly Summary -->
      <div v-else class="text-sm text-meal-gray-dark">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <p class="font-medium">Monatssaldo:</p>
            <p
                :class="[
                'text-lg font-bold',
                monthlyBalance >= 0 ? 'text-meal-positive' : 'text-meal-error'
              ]"
            >
              € {{ monthlyBalance.toFixed(2) }}
            </p>
          </div>
          <div>
            <p class="font-medium">Transaktionen:</p>
            <p class="text-lg font-bold">{{ monthlyTransactionCount }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue';
import VueApexCharts from 'vue3-apexcharts';
import {mutation, query} from "../graphql.ts";

export default {
  components: {
    apexchart: VueApexCharts,
  },
  setup() {
    // State-Variablen
    const activeView = ref('yearly');
    const internalTransactions = ref([]);
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
      return internalTransactions.value.length > 0;
    });

    // Berechnungen für Jahresansicht
    const totalIncome = computed(() => {
      return internalTransactions.value
          .filter(t => {
            const date = new Date(t.created_at);
            return date.getFullYear() === currentYear.value && t.amount > 0;
          })
          .reduce((sum, t) => sum + t.amount, 0);
    });

    const totalExpenses = computed(() => {
      return Math.abs(
          internalTransactions.value
              .filter(t => {
                const date = new Date(t.created_at);
                return date.getFullYear() === currentYear.value && t.amount < 0;
              })
              .reduce((sum, t) => sum + t.amount, 0)
      );
    });

    // Berechnungen für Monatsansicht
    const monthlyTransactions = computed(() => {
      return internalTransactions.value.filter(t => {
        const date = new Date(t.created_at);
        return date.getMonth() === currentDate.value.getMonth() &&
            date.getFullYear() === currentDate.value.getFullYear();
      });
    });

    const monthlyBalance = computed(() => {
      return monthlyTransactions.value.reduce((sum, t) => sum + t.amount, 0);
    });

    const monthlyTransactionCount = computed(() => {
      return monthlyTransactions.value.length;
    });

    // Chartdaten für Jahresansicht
    const yearlyChartSeries = computed(() => {
      const incomeSeries = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      const expenseSeries = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

      internalTransactions.value.forEach(transaction => {
        const date = new Date(transaction.created_at);
        const month = date.getMonth();
        const year = date.getFullYear();

        if (year === currentYear.value) {
          if (transaction.amount > 0) {
            incomeSeries[month] += transaction.amount;
          } else {
            expenseSeries[month] -= transaction.amount; // In positiven Wert umwandeln
          }
        }
      });

      return [
        {
          name: 'Einnahmen',
          data: incomeSeries
        },
        {
          name: 'Ausgaben',
          data: expenseSeries
        }
      ];
    });

    // Chart-Optionen für Jahresansicht
    const yearlyChartOptions = ref({
      chart: {
        type: 'bar',
        stacked: false,
        toolbar: {
          show: false
        },
        fontFamily: 'Quicksand, sans-serif',
        background: 'transparent',
      },
      plotOptions: {
        bar: {
          horizontal: false,
          columnWidth: '55%',
          borderRadius: 5,
        },
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        show: true,
        width: 2,
        colors: ['transparent']
      },
      xaxis: {
        categories: ['Jan', 'Feb', 'Mär', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Okt', 'Nov', 'Dez'],
        labels: {
          style: {
            fontSize: '12px',
            fontFamily: 'Quicksand, sans-serif',
          },
        }
      },
      yaxis: {
        title: {
          text: 'Betrag (€)',
          style: {
            fontSize: '14px',
            fontFamily: 'Quicksand, sans-serif',
          }
        },
        labels: {
          formatter: function(val) {
            return "€" + val.toFixed(0);
          },
          style: {
            fontSize: '12px',
            fontFamily: 'Quicksand, sans-serif',
          }
        }
      },
      fill: {
        opacity: 1
      },
      tooltip: {
        y: {
          formatter: function(val) {
            return "€ " + val.toFixed(2);
          }
        }
      },
      colors: ['#388E3C', '#D32F2F'],
      legend: {
        position: 'top',
        horizontalAlign: 'right',
        fontSize: '14px',
        fontFamily: 'Quicksand, sans-serif',
        markers: {
          width: 12,
          height: 12,
          radius: 12
        },
      },
      responsive: [
        {
          breakpoint: 480,
          options: {
            legend: {
              position: 'bottom',
              offsetY: 0
            }
          }
        }
      ]
    });

    // Chartdaten für Monatsansicht
    const monthlyChartSeries = computed(() => {
      const daysInMonth = new Date(
          currentDate.value.getFullYear(),
          currentDate.value.getMonth() + 1,
          0
      ).getDate();

      const dailyData = Array(daysInMonth).fill(0);
      const cumulativeData = Array(daysInMonth).fill(0);

      internalTransactions.value.forEach(transaction => {
        const date = new Date(transaction.created_at);

        if (date.getMonth() === currentDate.value.getMonth() &&
            date.getFullYear() === currentDate.value.getFullYear()) {
          const day = date.getDate() - 1; // Arrays beginnen bei 0
          dailyData[day] += transaction.amount;
        }
      });

      // Kumulativen Saldo berechnen
      let cumulative = 0;
      for (let i = 0; i < dailyData.length; i++) {
        cumulative += dailyData[i];
        cumulativeData[i] = cumulative;
      }

      return [
        {
          name: 'Tagessaldo',
          data: dailyData
        },
        {
          name: 'Kumulativer Saldo',
          data: cumulativeData
        }
      ];
    });

    // Chart-Optionen für Monatsansicht
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
        },
        stroke: {
          curve: 'smooth',
          width: [2, 3]
        },
        xaxis: {
          categories: categories,
          title: {
            text: 'Tag des Monats',
            style: {
              fontSize: '14px',
              fontFamily: 'Quicksand, sans-serif',
            }
          },
          labels: {
            style: {
              fontSize: '12px',
              fontFamily: 'Quicksand, sans-serif',
            }
          }
        },
        yaxis: {
          title: {
            text: 'Betrag (€)',
            style: {
              fontSize: '14px',
              fontFamily: 'Quicksand, sans-serif',
            }
          },
          labels: {
            formatter: function(val) {
              return "€" + val.toFixed(0);
            },
            style: {
              fontSize: '12px',
              fontFamily: 'Quicksand, sans-serif',
            }
          }
        },
        tooltip: {
          y: {
            formatter: function(val) {
              return "€ " + val.toFixed(2);
            }
          }
        },
        colors: ['#66BB6A', '#2E7D32'],
        markers: {
          size: 0,
          hover: {
            size: 5
          }
        },
        legend: {
          position: 'top',
          horizontalAlign: 'right',
          fontSize: '14px',
          fontFamily: 'Quicksand, sans-serif',
          markers: {
            width: 12,
            height: 12,
            radius: 12
          },
        },
        grid: {
          padding: {
            left: 10,
            right: 10
          }
        },
        responsive: [
          {
            breakpoint: 480,
            options: {
              legend: {
                position: 'bottom',
                offsetY: 0
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
    const getPayments = async () => {
      try {
        isLoading.value = true;
        error.value = null;

        const [data, err] = await query.getPayments();
        if (err) {
          error.value = 'Fehler beim Laden der Daten: ' + (err.message || 'Unbekannter Fehler');
          console.error('Error getting payments:', err);
          return;
        }

        // Array leeren, um Duplikate zu vermeiden
        internalTransactions.value = [];

        if (Array.isArray(data)) {
          data.forEach(payment => {
            internalTransactions.value.push({
              id: payment.id,
              from_user_id: payment.fromUserId,
              to_user_id: payment.toUserId,
              amount: parseFloat(payment.amount), // String zu Zahl konvertieren
              created_at: payment.date
            });
          });
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
      getPayments();
    };

    // Daten laden, wenn Komponente geladen wird
    onMounted(() => {
      getPayments();
    });

    return {
      activeView,
      internalTransactions,
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
      totalIncome,
      totalExpenses,
      monthlyBalance,
      monthlyTransactionCount
    };
  }
};
</script>