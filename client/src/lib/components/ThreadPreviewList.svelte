<script lang="ts">
	import type { ThreadModel } from "$lib/models/thread";
	import ThreadPreview from "./ThreadPreview.svelte";

	const { threadPreviewList }: { threadPreviewList: ThreadModel[] } = $props();
</script>

<ul>
	{#each threadPreviewList as threadPreview (threadPreview.id)}
		<li>
			<ThreadPreview threadPreview={threadPreview} />
		</li>
	{/each}
</ul>


<style>
	ul {
		list-style: none;
	}
</style>