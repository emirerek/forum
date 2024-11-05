<script lang="ts">
	import type { SubforumModel } from "$lib/models/subforum";
	import ThreadPreview from "./ThreadPreview.svelte";

	const { subforum }: { subforum: SubforumModel } = $props();
	const subforumUrl = $derived("/subforum/" + subforum.id);
</script>

<div class="container">
	<div class="subforum">
		<a class="title" href={subforumUrl}>
			{subforum.title}
		</a>
		<p class="description">
			{subforum.description}
		</p>
	</div>
	<div class="thread">
		<ThreadPreview threadPreview={subforum.threads![0]} />
	</div>
</div>

<style>
	.container {
		width: 100%;
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		align-items: flex-start;
	}

	.subforum {
		width: 100%;
		padding: 8px;
		background-color: var(--component-hover);
	}

	.description {
		margin: 4px 0 0 0;
	}

	.thread {
		width: 100%;
		background-color: var(--component);
	}
</style>