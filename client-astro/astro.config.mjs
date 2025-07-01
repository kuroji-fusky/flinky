// @ts-check
import { defineConfig } from "astro/config"

import tailwindcss from "@tailwindcss/vite"
import sitemap from "@astrojs/sitemap"
import svelte from "@astrojs/svelte"

export default defineConfig({
  vite: {
    plugins: [tailwindcss()],
  },
  output: "server",

  integrations: [
    sitemap(),
    // svelte()
  ],
})
