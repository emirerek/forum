<script lang="ts">
    import Icon from "@iconify/svelte";
    import { formatDate } from "$lib/util/formatDate";
	import { onMount } from "svelte";

    const { date }: { date: string } = $props();
    let formattedDate = $state();
    
    onMount(() => {
        formattedDate = formatDate(date);
    })
</script>

<span class="date">
    <Icon icon="material-symbols:calendar-today" />
    <time datetime={date} title={date}>
        {formattedDate}
    </time>
</span>

<style>
    .date {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .date time {
        margin: 0 0 0 2px;
    }
</style>