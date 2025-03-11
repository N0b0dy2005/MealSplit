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
        // Finanzen/Essen Farbschema
        "meal-primary": "#2E7D32",      // Satter Grünton (Geld/Finanzen)
        "meal-secondary": "#66BB6A",    // Mittleres Grün
        "meal-light": "#E8F5E9",        // Heller Grünton für Hintergründe
        "meal-accent": "#FF8F00",       // Warmes Orange (Essen-Assoziation)
        "meal-accent-light": "#FFE0B2",  // Helles Orange
        "meal-dark": "#1B5E20",         // Dunkelgrün
        "meal-error": "#D32F2F",        // Rot für negative Beträge
        "meal-positive": "#388E3C",     // Grün für positive Beträge

        // UI Grautöne
        "meal-gray-dark": "#263238",    // Dunkelgrau für Text
        "meal-gray": "#607D8B",         // Mittelgrau
        "meal-gray-light": "#ECEFF1",   // Hellgrau für Hintergründe
        "meal-white": "#FFFFFF",        // Weiß

        // Behalte bisherige Farbpaletten für Kompatibilität
        // Bestehende Farben...
      },
      fontFamily: {
        header: ['"Nunito"', 'sans-serif'],
        sans: ['"Quicksand"', 'sans-serif'],
      },
      boxShadow: {
        'sm': '0 1px 3px rgba(164, 157, 147, 0.1)',
        'md': '0 4px 6px rgba(164, 157, 147, 0.1)',
        'meal': '0 4px 6px rgba(46, 125, 50, 0.1)',  // Grünlicher Schatten
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