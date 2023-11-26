/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./templates/**/*.{html,js,templ,go}",
    "./templates/common/**/*.{html,js,templ,go}",
  ],
  theme: {
    extend: {},
    fontFamily: {
        mitr: ["Mitr"],
    },
  },
  plugins: [
      require("daisyui"),
  ],
}

