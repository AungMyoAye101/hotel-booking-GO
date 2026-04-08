// tailwind.config.js
const { heroui } = require("@heroui/theme");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./node_modules/@heroui/theme/dist/components/(avatar|button|date-input|input|progress|skeleton|ripple|spinner|form).js"
],
  theme: {
    extend: {
      colors: {
        'brands': "#333"
      }
    }
  },
  darkMode: "class",
  plugins: [heroui({
    themes: {
      light: {
        colors: {
          brand: "#6d28d9"
        }
      },
      dark: {
        colors: {
          brand: "#8b5cf6;"
        }
      }
    },

  })],
};