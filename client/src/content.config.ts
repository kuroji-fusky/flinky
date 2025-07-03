import { glob } from "astro/loaders"
import { z, defineCollection } from "astro:content"

const docs = defineCollection({
  loader: glob({
    pattern: "**/[^_]*.mdx",
    base: "./src/content/",
  }),
  schema: z.object({
    title: z.string(),
    description: z.string(),
  }),
})

export { docs }
