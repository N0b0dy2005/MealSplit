<template>
  <div class="h-full">
    <div class="flex flex-col h-full">
      <!-- 3D Kuchendiagramm -->
      <div class="flex-grow pie-chart-container relative w-full">
        <svg ref="chartRef" :width="chartWidth" :height="chartHeight" class="pie-chart-svg absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
          <!-- Schatten unter dem Diagramm -->
          <ellipse 
            :cx="centerX" 
            :cy="centerY + 20" 
            :rx="radius * 0.9" 
            :ry="radius * 0.25" 
            fill="rgba(0,0,0,0.1)" 
            filter="blur(7px)"
          />
          
          <!-- 3D-Ebenen (von unten nach oben) -->
          <g :transform="`translate(${centerX}, ${centerY})`" class="pie-3d-container">
            <!-- Schatten für 3D-Effekt -->
            <g v-for="(slice, i) in pieData" :key="`shadow-${i}`">
              <path 
                :d="createShadowPath(slice)" 
                :fill="darkenColor(getUserColor(chartData[i].id))" 
                :opacity="0.2"
              />
            </g>
            
            <!-- Hauptsegmente mit 3D-Effekt -->
            <g v-for="(slice, i) in pieData" :key="`segment-${i}`"
               @mouseenter="highlightSegment(i)" 
               @mouseleave="resetHighlight()"
               :class="{ 'segment-highlight': hoveredSegment === i }"
               style="cursor: pointer;">
              <path 
                :d="arc(slice)" 
                :fill="getUserColor(chartData[i].id)" 
                :transform="`rotate(${rotationAngle}) translate(0, ${hoveredSegment === i ? -8 : 0})`"
                stroke="#fff"
                stroke-width="0.5"
              />
              
              <!-- Glanzeffekt -->
              <path 
                :d="arc(slice)" 
                fill="url(#pieGlossGradient)" 
                :opacity="0.3"
                :transform="`rotate(${rotationAngle}) translate(0, ${hoveredSegment === i ? -8 : 0})`"
                style="pointer-events: none;"
              />
              
              <!-- Label -->
              <g v-if="chartData[i].percentage > 10"
                 :transform="`rotate(${rotationAngle}) translate(${arc.centroid(slice)[0] * 0.8}, ${arc.centroid(slice)[1] * 0.8 + (hoveredSegment === i ? -8 : 0)})`">
                <text 
                  text-anchor="middle"
                  fill="#fff"
                  font-weight="bold"
                  font-size="12"
                  style="pointer-events: none;"
                >
                  {{ chartData[i].name.split(' ')[0] }}
                </text>
              </g>
            </g>
          </g>
          
          <!-- Definitionen für Gradienten und Filter -->
          <defs>
            <radialGradient id="pieGlossGradient" cx="50%" cy="30%" r="60%" fx="50%" fy="20%">
              <stop offset="0%" stop-color="#fff" stop-opacity="0.7" />
              <stop offset="100%" stop-color="#fff" stop-opacity="0" />
            </radialGradient>
          </defs>
        </svg>
      </div>
      
      <!-- Legende - Kompakter -->
      <div class="mt-2 grid grid-cols-2 sm:grid-cols-3 gap-1 text-xs">
        <div 
          v-for="(item, index) in chartData" 
          :key="index"
          @mouseenter="highlightSegment(index)"
          @mouseleave="resetHighlight()"
          class="flex items-center p-1 rounded-lg transition-colors duration-150"
          :class="{ 'bg-meal-light': hoveredSegment === index }"
          style="cursor: pointer;"
        >
          <div class="w-2 h-2 sm:w-3 sm:h-3 rounded-full mr-1" 
               :style="{ backgroundColor: getUserColor(item.id) }"></div>
          <div class="flex flex-col">
            <span class="text-xs text-meal-gray-dark truncate max-w-[80px]">{{ item.name }}</span>
            <span class="text-xs text-meal-gray">{{ formatPercentage(item.percentage) }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue';
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
const rotationAngle = ref(-5); // Winkel weniger extrem für bessere Sichtbarkeit

// Chart dimensions
const chartWidth = 280;
const chartHeight = 280;
const radius = Math.min(chartWidth, chartHeight) / 3;
const centerX = chartWidth / 2;
const centerY = chartHeight / 2;

// Initialize D3 pie generator
const pie = d3.pie()
  .value(d => d.percentage)
  .sort(null);

// Initialize D3 arc generator
const arc = d3.arc()
  .innerRadius(radius * 0.2) // Kleines Loch in der Mitte für Donut-Effekt
  .outerRadius(radius);

// Compute pie data whenever chartData changes
const pieData = computed(() => {
  return pie(props.chartData);
});

// Function to create shadow path for 3D effect
const createShadowPath = (slice) => {
  const depth = 10; // Tiefe des 3D-Effekts
  const startAngle = slice.startAngle;
  const endAngle = slice.endAngle;
  
  // Berechne die Punkte für den 3D-Schatten
  const innerRadius = radius * 0.2;
  const outerRadius = radius;
  
  const startInnerX = innerRadius * Math.cos(startAngle);
  const startInnerY = innerRadius * Math.sin(startAngle);
  const startOuterX = outerRadius * Math.cos(startAngle);
  const startOuterY = outerRadius * Math.sin(startAngle);
  
  const endInnerX = innerRadius * Math.cos(endAngle);
  const endInnerY = innerRadius * Math.sin(endAngle);
  const endOuterX = outerRadius * Math.cos(endAngle);
  const endOuterY = outerRadius * Math.sin(endAngle);
  
  // Erstelle den Pfad für den Schatten
  return `
    M ${startInnerX} ${startInnerY + depth}
    L ${startOuterX} ${startOuterY + depth}
    A ${outerRadius} ${outerRadius} 0 ${(endAngle - startAngle > Math.PI) ? 1 : 0} 1 ${endOuterX} ${endOuterY + depth}
    L ${endInnerX} ${endInnerY + depth}
    A ${innerRadius} ${innerRadius} 0 ${(endAngle - startAngle > Math.PI) ? 1 : 0} 0 ${startInnerX} ${startInnerY + depth}
    Z
  `;
};

// Function to darken a color for shadow effect
const darkenColor = (color) => {
  // Einfache Implementierung - für komplexere Farbanpassungen könnte
  // eine Bibliothek wie 'color' verwendet werden
  return color;
};

// Highlight functions
const highlightSegment = (index) => {
  hoveredSegment.value = index;
};

const resetHighlight = () => {
  hoveredSegment.value = null;
};

// Optional: Animation
onMounted(() => {
  // Langsame Rotation für Demonstration
  /*
  const rotationInterval = setInterval(() => {
    rotationAngle.value = (rotationAngle.value + 0.2) % 360;
  }, 100);
  */
});
</script>

<style scoped>
.pie-chart-container {
  position: relative;
  margin: 0 auto;
}

.pie-3d-container {
  transform-style: preserve-3d;
  backface-visibility: hidden;
}

.segment-highlight {
  filter: brightness(1.1);
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style> 