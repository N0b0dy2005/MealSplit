// MealsPieChart.vue
<template>
  <div class="bg-white rounded-xl shadow-meal p-3">
    <div class="flex flex-col">
      <!-- Diagramm und Legende in einem Flex-Container -->
      <div class="flex">
        <!-- Diagramm - Maximierte Größe -->
        <div class="w-3/4 pr-2">
          <svg :width="size" :height="size" class="pie-chart">
            <g :transform="`translate(${size/2}, ${size/2})`">
              <!-- Diagramm-Segmente -->
              <path
                  v-for="(segment, i) in pieData"
                  :key="i"
                  :d="arc(segment)"
                  :fill="getUserColor(userExpenseData[i].id)"
                  stroke="#ffffff"
                  stroke-width="1"
                  @mouseenter="activeSegment = i"
                  @mouseleave="activeSegment = null"
                  :style="{
                  transform: activeSegment === i ? 'scale(1.03)' : 'scale(1)',
                  transformOrigin: 'center',
                  filter: activeSegment === i ? 'brightness(1.1)' : 'none',
                  transition: 'transform 0.2s ease, filter 0.2s ease'
                }"
              />

              <!-- Labels direkt im Diagramm - nur für wirklich große Segmente -->
              <g v-for="(segment, i) in pieData" :key="`label-${i}`"
                 v-show="userExpenseData[i].percentage >= 15">
                <text
                    :transform="`translate(${labelPosition(segment)})`"
                    text-anchor="middle"
                    class="text-xs font-bold select-none"
                    fill="white"
                    style="text-shadow: 0 0 3px rgba(0,0,0,0.7);"
                >
                  {{ userExpenseData[i].name.charAt(0) }}
                </text>
              </g>

              <!-- Mitte (Loch) mit Gesamt-Info -->
              <circle r="75" fill="white" />
              <text
                  text-anchor="middle"
                  dy="-15"
                  class="text-base font-semibold text-meal-gray-dark"
              >
                Gesamt
              </text>
              <text
                  text-anchor="middle"
                  dy="20"
                  class="text-2xl font-bold text-meal-primary"
              >
                {{totalMeals}}
              </text>
              <text
                  text-anchor="middle"
                  dy="45"
                  class="text-sm text-meal-gray"
              >
                Mahlzeiten
              </text>
            </g>
          </svg>
        </div>

        <!-- Kompakte Legende - Direkt neben dem Diagramm -->
        <div class="w-1/4 flex flex-col justify-center">
          <div
              v-for="(user, index) in userExpenseData"
              :key="index"
              @mouseenter="activeSegment = index"
              @mouseleave="activeSegment = null"
              class="mb-2 flex items-center rounded p-1 cursor-pointer"
              :class="{ 'bg-meal-light': activeSegment === index }"
              style="transition: background-color 0.15s ease;"
          >
            <div class="h-4 w-4 rounded-full mr-1 flex-shrink-0"
                 :style="{ backgroundColor: getUserColor(user.id) }"></div>
            <div class="truncate">
              <div class="text-xs font-medium text-meal-gray-dark truncate">{{ user.name }}</div>
              <div class="text-xs text-meal-primary">{{ formatPercentage(user.percentage) }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Details-Panel für ausgewähltes Segment - Immer unter dem Diagramm -->
      <transition name="fade">
        <div v-if="activeSegment !== null" class="mt-2 p-2 bg-meal-light rounded-lg">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="h-4 w-4 rounded-full mr-2"
                   :style="{ backgroundColor: getUserColor(userExpenseData[activeSegment].id) }"></div>
              <span class="font-medium">{{ userExpenseData[activeSegment].name }}</span>
            </div>
            <div class="text-right">
              <div class="text-meal-primary font-bold">
                {{ userExpenseData[activeSegment].amount }} Mahlzeiten
              </div>
              <div class="text-xs text-meal-gray">
                {{ formatPercentage(userExpenseData[activeSegment].percentage) }} aller Mahlzeiten
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue';
import * as d3 from 'd3';

export default {
  name: 'MealsPieChart',
  props: {
    userExpenseData: {
      type: Array,
      required: true
    },
    getUserColor: {
      type: Function,
      required: true
    },
    formatPercentage: {
      type: Function,
      required: true
    },
    title: {
      type: String,
      default: 'Mahlzeiten pro Person'
    }
  },
  setup(props) {
    const activeSegment = ref(null);
    const size = 340; // Maximierte Größe für bessere Sichtbarkeit

    // Berechne die Gesamtzahl der Mahlzeiten
    const totalMeals = computed(() => {
      return props.userExpenseData.reduce((sum, item) => sum + item.amount, 0);
    });

    // D3 pie generator
    const pie = d3.pie()
        .value(d => d.percentage)
        .sort(null);

    // D3 arc generator - größerer Radius für maximale Sichtbarkeit
    const arc = d3.arc()
        .innerRadius(75) // Größeres Loch für die Mitte
        .outerRadius(160); // Maximaler äußerer Radius

    // Transformiere Daten fürs Pie-Chart
    const pieData = computed(() => {
      return pie(props.userExpenseData);
    });

    // Hilfsfunktion zur Berechnung der Label-Position
    const labelPosition = (segment) => {
      // Position zwischen innerem und äußerem Radius für optimale Platzierung
      const midRadius = (75 + 160) / 2;
      const midAngle = (segment.startAngle + segment.endAngle) / 2;
      return [
        Math.sin(midAngle) * midRadius * 0.75,
        -Math.cos(midAngle) * midRadius * 0.75
      ];
    };

    return {
      activeSegment,
      size,
      totalMeals,
      pieData,
      arc,
      labelPosition
    };
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.pie-chart {
  overflow: visible; /* Für Hover-Effekte */
}

/* Verbesserte Animation für Hover-Effekt */
.pie-chart path {
  transform-box: fill-box;
  transform-origin: center;
}
</style>