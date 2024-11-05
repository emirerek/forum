<script lang="ts">
	import "../app.css";
	import Footer from '$lib/components/Footer.svelte';
	import Header from '$lib/components/Header.svelte';
	export const ssr = false;

    const { children } = $props();
</script>

<Header />
<main>
	{@render children()}
</main>
<Footer />

<style>
	main {
		width: var(--app-width);
	}
</style>