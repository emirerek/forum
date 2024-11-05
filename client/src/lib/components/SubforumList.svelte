<script lang="ts">
    import type { SubforumModel } from "$lib/models/subforum";
    import Subforum from "./Subforum.svelte";

    const { subforumList }: { subforumList: SubforumModel[] } = $props(); // MAKE TWO ROW COLUMN!
</script>

<div class="subforum-list">
    {#each subforumList as subforum, index (subforum.id)}
        <Subforum subforum={subforum} />
    {/each}
</div>

<style>
    .subforum-list {
        width: 100%;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        gap: 8px;
        list-style: none;
    }
</style>
