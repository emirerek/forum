<script lang="ts">
    import type { ReplyModel } from "$lib/models/reply";
    import AccountPreview from "./AccountPreview.svelte";

    const { reply }: { reply: ReplyModel } = $props();
</script>

<article>
    <AccountPreview accountPreview={reply.account} />
    <div>{reply.content}</div>
</article>