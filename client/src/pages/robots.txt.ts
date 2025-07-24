import type { APIRoute } from "astro"

const robots = `
User-Agent: *
Allow: /
`

export const GET: APIRoute = () => {
  return new Response(robots.trim(), {
    headers: {
      "Content-Type": "text/plain; charset=utf-8"
    }
  })
}