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
  safelist: [
        'group-focus-within:!text-red-500',
        'group-focus-within:!border-red-500',
        'group-focus-within:!border-red-500',
        'border-red-500',
        'border-gray-400',
        'group-hover:border-red-500',
        'group-hover:border-gray-700',
        'group-focus-within:!text-green-600',
        'group-focus-within:!border-green-600',
        'group-focus-within:!border-green-600',
    ],
  plugins: [],
}

