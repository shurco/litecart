export default {
  // Basic formatting settings
  semi: false,
  tabWidth: 2,
  singleQuote: true,
  printWidth: 120,
  trailingComma: 'none',
  // Plugins for SvelteKit and Tailwind CSS
  plugins: ['prettier-plugin-svelte', 'prettier-plugin-tailwindcss'],
  overrides: [
    {
      files: '*.svelte',
      options: {
        parser: 'svelte'
      }
    }
  ]
}
