<script lang="ts">
	import ThreadPreviewList from "$lib/components/ThreadPreviewList.svelte";

    const { data } = $props();
</script>

<div>
    <h2>Subforum name here lol</h2>
    <ThreadPreviewList threadPreviewList={data.threadPreviewList} />
</div>

<style>
    div {
        width: 100%;
    }
</style>
