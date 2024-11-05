<script lang="ts">
    import type { ThreadModel } from "$lib/models/thread";
	import AccountPreview from "./AccountPreview.svelte";

    const { threadData }: { threadData: ThreadModel } = $props();    
</script>

<article>
    <h3>{threadData.title}</h3>
    <div>
        <AccountPreview accountPreview={threadData.account} />
        <div>{threadData.content}</div>
    </div>
</article>