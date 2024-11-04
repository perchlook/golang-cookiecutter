const plugin = require("tailwindcss/plugin");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js}"],
  theme: {
    extend: {
      fontFamily: {
        serif: ["Literata", "Georgia", "sans-serif"],
        // Add more custom font families as needed
      },
      colors: {
        primary: {
          '50': '#f9f5f3',
          '100': '#f1e9e3',
          '200': '#e2d0c6',
          '300': '#d0b1a1',
          '400': '#bc8d7b',
          '500': '#ae7361',
          '600': '#a66558',
          '700': '#864f48',
          '800': '#6d423f',
          '900': '#593835',
          '950': '#2f1b1b',
        },
      },
    },
  },
  plugins: [
    plugin(function ({ addVariant }) {
      addVariant("htmx-settling", ["&.htmx-settling", ".htmx-settling &"]);
      addVariant("htmx-request", ["&.htmx-request", ".htmx-request &"]);
      addVariant("htmx-swapping", ["&.htmx-swapping", ".htmx-swapping &"]);
      addVariant("htmx-added", ["&.htmx-added", ".htmx-added &"]);
    }),
    require("@tailwindcss/typography"),
    require("@tailwindcss/forms"),
    require("@tailwindcss/aspect-ratio"),
    require("@tailwindcss/container-queries"),
  ],
};
