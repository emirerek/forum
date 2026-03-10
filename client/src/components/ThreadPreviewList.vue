<template>
	<ul>
		<li v-for="threadPreview in threadPreviewList" :key="threadPreview.id">
			<ThreadPreview :threadPreview="threadPreview" />
		</li>
	</ul>
</template>

<script setup lang="ts">
import ThreadPreview from "./ThreadPreview.vue";
import type { ThreadModel } from "@/models/thread";

const props = defineProps<{ threadPreviewList: ThreadModel[] }>();
</script>

<style scoped>
ul {
	width: 100%;
	list-style: none;
}

li {
	margin: 0 0 8px 0;
}
</style>
