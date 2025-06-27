<script lang="ts">
	import "../app.css";
	import { Navbar } from "$lib/components";

	const { children } = $props();
</script>

<Navbar />
{@render children()}
