<template>
	<article class="thread-preview">
		<a class="title" :href="threadUrl" :title="threadPreview.title">
			{{ threadPreview.title }}
		</a>
		<span class="by">by</span>
		<a :href="accountUrl">
			{{ threadPreview.account.username }}
		</a>
		<span class="at">posted at</span>
		<Date :date="threadPreview.createdAt" />
	</article>
</template>

<script setup lang="ts">
import { computed } from "vue";
import Date from "./Date.vue";
import type { ThreadModel } from "@/models/thread";

const props = defineProps<{ threadPreview: ThreadModel }>();
const threadUrl = computed(() => `/thread/${props.threadPreview.id}`);
const accountUrl = computed(() => `/account/${props.threadPreview.account.id}`);
</script>

<style scoped>
.thread-preview {
	display: flex;
	flex-direction: row;
	justify-content: flex-start;
	align-items: center;
	padding: 8px;
	background-color: var(--component);
}

.title {
	text-overflow: ellipsis;
	white-space: nowrap;
	overflow: hidden;
}

.meta {
	display: flex;
	justify-content: flex-start;
	align-items: center;
	flex-wrap: wrap;
}

.by,
.at {
	margin: 0 4px;
}
</style>
