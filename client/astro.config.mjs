// @ts-check
import * as dotenv from "dotenv"
import path from "node:path"
import { defineConfig } from "astro/config"
import tailwindcss from "@tailwindcss/vite"
import sitemap from "@astrojs/sitemap"
import svelte from "@astrojs/svelte"
import mdx from "@astrojs/mdx"
import { default as unpluginIcons } from "unplugin-icons/vite"

dotenv.config({ path: path.resolve(process.cwd(), "../.env") })

export default defineConfig({
  output: "server",

  prefetch: {
    prefetchAll: true
  },

  integrations: [sitemap(), mdx(), svelte()],

  site: process.env.FRONTEND_SITE_URL,

  redirects: {},

  vite: {
    plugins: [
      tailwindcss(),
      unpluginIcons({
        compiler: "svelte",
      }),
    ],
  },

})
