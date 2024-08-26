/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./app/**/*.{ts,tsx,jsx,js}", "./components/**/*.{ts,tsx,jsx,js}"],
  theme: {
    extend: {
            fontFamily: {
                inter: ['Inter', 'sans-serif'],
                hanken: ['Hanken', 'sans-serif'],
            },
            colors: {
                'background': '#F5F5F5',
            }
        },
  },
  plugins: [],
}

