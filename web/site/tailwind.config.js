/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,js}", "!./**/node_modules/**"],
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
