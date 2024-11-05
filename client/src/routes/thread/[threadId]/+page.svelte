<script lang="ts">
    import Thread from "$lib/components/Thread.svelte";
	import ReplyList from "$lib/components/ReplyList.svelte";

    const { data } = $props();
</script>

<div>
    <Thread threadData={data.threadPreview} />
    <ReplyList replyList={data.threadPreview.replies!} />
</div>

<style>
    div {
        width: 100%;
    }
</style>