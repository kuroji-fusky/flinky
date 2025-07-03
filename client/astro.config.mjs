// @ts-check
import { defineConfig } from "astro/config"
import tailwindcss from "@tailwindcss/vite"
import sitemap from "@astrojs/sitemap"
import Icons from "unplugin-icons/vite"

export default defineConfig({
  output: "server",
  
  vite: {
    plugins: [
      tailwindcss(),
      Icons({
        compiler: "astro",
      }),
    ],
  },

  integrations: [sitemap()],
})
