<script lang="ts">
    import type { ThreadModel } from "$lib/models/thread";
	import AccountUsername from "./AccountUsername.svelte";
    import Date from "./Date.svelte";

    const { threadPreview }: { threadPreview: ThreadModel } = $props();
    const threadUrl = $derived("/thread/" + threadPreview.id);
</script>

<article class="thread-preview">
    <a class="title" href={threadUrl} title={threadPreview.title}>
        {threadPreview.title}
    </a>
    <div class="meta">
        <AccountUsername username={threadPreview.account.username} />
        <Date date={threadPreview.createdAt} />
    </div>
</article>

<style>
    .thread-preview {
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        padding: 8px;
    }

    .title {
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
        margin: 0 0 4px 0;
    }

    .meta {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
    }
</style>
