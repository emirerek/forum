<script lang="ts">
    import type { ReplyModel } from "$lib/models/reply";
    import Reply from "./Reply.svelte";

    const { replyList }: { replyList: ReplyModel[] } = $props();
</script>

<ul>
    {#each replyList as reply (reply.id)}
        <li>
            <Reply reply={reply} />
        </li>
    {/each}
</ul>