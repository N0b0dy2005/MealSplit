const plugin = require('tailwindcss/plugin')

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        // Original Grünes Schema
        "meal-primary-green": "#2E7D32",      // Satter Grünton (Geld/Finanzen)
        "meal-secondary-green": "#66BB6A",    // Mittleres Grün
        "meal-light-green": "#E8F5E9",        // Heller Grünton für Hintergründe
        "meal-accent-green": "#FF8F00",       // Warmes Orange (Essen-Assoziation)
        "meal-accent-light-green": "#FFE0B2",  // Helles Orange
        "meal-dark-green": "#1B5E20",         // Dunkelgrün
        "meal-error-green": "#D32F2F",        // Rot für negative Beträge
        "meal-positive-green": "#388E3C",     // Grün für positive Beträge
        "meal-gray-dark-green": "#263238",    // Dunkelgrau für Text
        "meal-gray-green": "#607D8B",         // Mittelgrau
        "meal-gray-light-green": "#ECEFF1",   // Hellgrau für Hintergründe

        // Neues Blaues Schema - helleres, kräftigeres Blau
        "meal-primary-blue": "#1E88E5",      // Helleres, kräftigeres Blau
        "meal-secondary-blue": "#42A5F5",    // Helleres mittleres Blau
        "meal-light-blue": "#E3F2FD",        // Heller Blauton für Hintergründe
        "meal-accent-blue": "#FF9800",       // Warmes Orange als Akzent
        "meal-accent-light-blue": "#FFE0B2",  // Helles Orange
        "meal-dark-blue": "#0D47A1",         // Dunkleres Blau für Kontrast
        "meal-error-blue": "#F44336",        // Rot für negative Beträge
        "meal-positive-blue": "#4CAF50",     // Grün für positive Beträge
        "meal-gray-dark-blue": "#263238",    // Dunkelgrau für Text
        "meal-gray-blue": "#546E7A",         // Mittelgrau
        "meal-gray-light-blue": "#ECEFF1",   // Hellgrau für Hintergründe

        // Aktive Farben (standardmäßig auf Grün gesetzt)
        "meal-primary": "var(--meal-primary-color, #2E7D32)",
        "meal-secondary": "var(--meal-secondary-color, #66BB6A)",
        "meal-light": "var(--meal-light-color, #E8F5E9)",
        "meal-accent": "var(--meal-accent-color, #FF8F00)",
        "meal-accent-light": "var(--meal-accent-light-color, #FFE0B2)",
        "meal-dark": "var(--meal-dark-color, #1B5E20)",
        "meal-error": "var(--meal-error-color, #D32F2F)",
        "meal-positive": "var(--meal-positive-color, #388E3C)",
        "meal-gray-dark": "var(--meal-gray-dark-color, #263238)",
        "meal-gray": "var(--meal-gray-color, #607D8B)",
        "meal-gray-light": "var(--meal-gray-light-color, #ECEFF1)",
        "meal-white": "#FFFFFF",
      },
      fontFamily: {
        header: ['"Nunito"', 'sans-serif'],
        sans: ['"Quicksand"', 'sans-serif'],
      },
      boxShadow: {
        'sm': '0 1px 3px rgba(164, 157, 147, 0.1)',
        'md': '0 4px 6px rgba(164, 157, 147, 0.1)',
        'meal': '0 4px 6px var(--meal-shadow-color, rgba(46, 125, 50, 0.1))',
      },
      borderRadius: {
        'xl': '1rem',
      },
      textShadow: {
        sm: '0 1px 2px var(--tw-shadow-color)',
        DEFAULT: '0 2px 4px var(--tw-shadow-color)',
        lg: '0 8px 16px var(--tw-shadow-color)',
      },
    },
  },
  plugins: [
    plugin(function ({ matchUtilities, theme }) {
      matchUtilities(
          {
            'text-shadow': (value) => ({
              textShadow: value,
            }),
          },
          { values: theme('textShadow') }
      )
    }),
  ],
}