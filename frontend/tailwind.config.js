/** @type {import('tailwindcss').Config} */
export default {
  content: [
        "./index.html",
        "./src/**/*.{js,jsx}"
    ],
  theme: {
    extend: {
            colors: {
                'wc-green': '#99FFC8',
                'wc-blue': '#C4EFFD',
                'brand-black': '#1F1F1F',
                'brand-background': '#F5F5F5'
            },
            fontFamily: {
                inter: ['inter'],
                hanken: ['hanken'],
            }
        },
  },
  plugins: [],
}

