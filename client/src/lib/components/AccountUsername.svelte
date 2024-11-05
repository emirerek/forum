<script lang="ts">
	import Icon from "@iconify/svelte";

    let { username }: { username: string } = $props();
</script>

<span class="username">
    <Icon icon="mdi:user" />
    <span>{username}</span>
</span>

<style>
    .username {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .username span {
        margin: 0 0 0 2px;
    }
</style>