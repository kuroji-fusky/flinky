import type { Snippet } from "svelte";

export type WithSnippet<P extends Record<string, any> = {}> = P & { children?: Snippet }