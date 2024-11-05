<script lang="ts">
    import SubforumList from "$lib/components/SubforumList.svelte";

    const { data } = $props();
</script>

<div class="index">
    <section>
        <h2>Subforums</h2>
        <SubforumList subforumList={data.subforumList} />
    </section>
</div>

<style>
    .index {
        width: 100%;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: flex-start;
        margin-top: 8px;
    }

    .index section {
        width: 100%;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-self: flex-start;
        flex: 3;
    }

    .index h2 {
        margin: 0 0 8px 0;
    }
</style>