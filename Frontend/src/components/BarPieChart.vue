<template>
  <div class="h-full">
    <div class="flex flex-col h-full">
      <!-- Kombiniertes Diagramm (Kuchendiagramm oben, Balken unten) -->
      <div class="flex-grow chart-container relative w-full">
        <svg ref="chartRef" :width="chartWidth" :height="chartHeight" class="chart-svg absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
          <!-- Kleines Kuchendiagramm oben -->
          <g :transform="`translate(${centerX}, ${centerY - 50})`">
            <g v-for="(slice, i) in pieData" :key="`pie-${i}`"
               @mouseenter="highlightSegment(i)" 
               @mouseleave="resetHighlight()"
               :class="{ 'segment-highlight': hoveredSegment === i }"
               style="cursor: pointer;">
              <path 
                :d="arc(slice)" 
                :fill="getUserColor(chartData[i].id)" 
                stroke="#fff"
                stroke-width="0.5"
                :transform="hoveredSegment === i ? `scale(1.05)` : ''"
              />
              
              <!-- Glanzeffekt -->
              <path 
                :d="arc(slice)" 
                fill="url(#barPieGradient)" 
                :opacity="0.3"
                style="pointer-events: none;"
                :transform="hoveredSegment === i ? `scale(1.05)` : ''"
              />
            </g>
          </g>
          
          <!-- Balkendiagramm unten - nur die Top 5 anzeigen für bessere Sichtbarkeit -->
          <g :transform="`translate(${chartWidth * 0.1}, ${centerY + 10})`">
            <g v-for="(item, i) in sortedData.slice(0, 5)" :key="`bar-${i}`"
               @mouseenter="highlightSegment(getOriginalIndex(item))" 
               @mouseleave="resetHighlight()"
               style="cursor: pointer;">
              <!-- Balken-Hintergrund (als Referenz) -->
              <rect 
                :x="0" 
                :y="i * (barHeight + barGap)" 
                :width="chartWidth * 0.8" 
                :height="barHeight" 
                fill="#f0f0f0" 
                rx="3" 
                ry="3"
              />
              
              <!-- Eigentlicher Balken mit Wert -->
              <rect 
                :x="0" 
                :y="i * (barHeight + barGap)" 
                :width="barScale(item.percentage)" 
                :height="barHeight" 
                :fill="getUserColor(item.id)" 
                rx="3" 
                ry="3"
                :class="{ 'bar-highlight': hoveredSegment === getOriginalIndex(item) }"
                :transform="hoveredSegment === getOriginalIndex(item) ? 'translate(0, -2)' : ''"
              />
              
              <!-- Name links -->
              <text 
                :x="5" 
                :y="i * (barHeight + barGap) + barHeight/2 + 4" 
                font-size="10" 
                fill="#fff" 
                font-weight="bold"
              >
                {{ item.name.split(' ')[0] }}
              </text>
              
              <!-- Wert rechts -->
              <text 
                :x="barScale(item.percentage) - 5" 
                :y="i * (barHeight + barGap) + barHeight/2 + 4" 
                font-size="10" 
                fill="#fff" 
                font-weight="bold"
                text-anchor="end"
                v-if="barScale(item.percentage) > 30"
              >
                {{ formatPercentage(item.percentage) }}
              </text>
              <text 
                :x="barScale(item.percentage) + 5" 
                :y="i * (barHeight + barGap) + barHeight/2 + 4" 
                font-size="10" 
                :fill="hoveredSegment === getOriginalIndex(item) ? '#000' : '#666'" 
                text-anchor="start"
                v-else
              >
                {{ formatPercentage(item.percentage) }}
              </text>
            </g>
          </g>
          
          <!-- Definitionen für Gradienten und Filter -->
          <defs>
            <linearGradient id="barPieGradient" x1="0%" y1="0%" x2="0%" y2="100%">
              <stop offset="0%" stop-color="#fff" stop-opacity="0.7" />
              <stop offset="100%" stop-color="#fff" stop-opacity="0" />
            </linearGradient>
          </defs>
        </svg>
      </div>
      
      <!-- Zusammenfassung - kompakter -->
      <div class="mt-2 w-full bg-meal-light rounded-lg p-2 text-xs">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-2 h-2 rounded-full mr-1" 
                 :style="{ backgroundColor: getUserColor(topItem.id) }"></div>
            <span class="font-bold text-meal-gray-dark truncate max-w-[100px]">{{ topItem.name }}</span>
            <span class="ml-1 text-meal-primary">{{ formatPercentage(topItem.percentage) }}</span>
          </div>
          <div class="flex items-center">
            <div class="w-2 h-2 rounded-full mr-1" 
                 :style="{ backgroundColor: getUserColor(bottomItem.id) }"></div>
            <span class="font-bold text-meal-gray-dark truncate max-w-[100px]">{{ bottomItem.name }}</span>
            <span class="ml-1 text-meal-error">{{ formatPercentage(bottomItem.percentage) }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import * as d3 from 'd3';

const props = defineProps({
  chartData: {
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
    default: 'Diagramm'
  }
});

const chartRef = ref(null);
const hoveredSegment = ref(null);

// Chart dimensions
const chartWidth = 280;
const chartHeight = 280;
const radius = Math.min(chartWidth, chartHeight) / 4; // Kleineres Kuchendiagramm
const centerX = chartWidth / 2;
const centerY = chartHeight / 2;

// Kompaktere Balken-Parameter
const barHeight = 15;
const barGap = 3;

// Initialize D3 pie generator
const pie = d3.pie()
  .value(d => d.percentage)
  .sort(null);

// Initialize D3 arc generator
const arc = d3.arc()
  .innerRadius(radius * 0.4) // Donut-Stil
  .outerRadius(radius);

// Compute pie data whenever chartData changes
const pieData = computed(() => {
  return pie(props.chartData);
});

// Sortiere Daten für das Balkendiagramm (absteigend nach Prozentsatz)
const sortedData = computed(() => {
  return [...props.chartData].sort((a, b) => b.percentage - a.percentage);
});

// Balken-Skalierung
const barScale = (value) => {
  // Maximalwert für 100% Breite
  const maxPercentage = 100;
  // Verfügbare Breite für Balken (80% der Gesamtbreite)
  const availableWidth = chartWidth * 0.8;
  
  return (value / maxPercentage) * availableWidth;
};

// Finde den ursprünglichen Index eines sortierten Elements
const getOriginalIndex = (item) => {
  return props.chartData.findIndex(d => d.id === item.id);
};

// Highlight functions
const highlightSegment = (index) => {
  hoveredSegment.value = index;
};

const resetHighlight = () => {
  hoveredSegment.value = null;
};

// Berechne höchsten und niedrigsten Wert
const topItem = computed(() => {
  return sortedData.value[0] || { name: '', percentage: 0, id: 0 };
});

const bottomItem = computed(() => {
  return sortedData.value[sortedData.value.length - 1] || { name: '', percentage: 0, id: 0 };
});

onMounted(() => {
  // Hier könnten Animationen oder zusätzliche Initialisierungen stattfinden
});
</script>

<style scoped>
.chart-container {
  position: relative;
  margin: 0 auto;
}

.segment-highlight {
  filter: brightness(1.1);
}

.bar-highlight {
  filter: brightness(1.1);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

rect {
  transition: transform 0.2s ease, filter 0.2s ease;
}

path {
  transition: transform 0.2s ease, filter 0.2s ease;
}
</style> 