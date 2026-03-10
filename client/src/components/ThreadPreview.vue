<template>
	<article class="thread-preview">
		<a class="title" :href="threadUrl" :title="threadPreview.title">
			{{ threadPreview.title }}
		</a>
		<div class="info">
			<a :href="accountUrl">
				{{ threadPreview.account.username }}
			</a>
			<span class="at">posted at</span>
			<Date v-if="threadPreview.account" :date="threadPreview.account.createdAt" />
			<Date v-else :date="threadPreview.createdAt" />
		</div>
	</article>
</template>

<script setup lang="ts">
import { computed } from "vue";
import Date from "./Date.vue";
import type { ThreadModel } from "@/models/thread";

const props = defineProps<{ threadPreview: ThreadModel }>();
const threadUrl = computed(() => `/thread/${props.threadPreview.id}`);
const accountUrl = computed(() => props.threadPreview.account ? `/account/${props.threadPreview.account.id}` : "#");
</script>

<style scoped>
.thread-preview {
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: flex-start;
	align-items: flex-start;
	background-color: var(--component);
}

.title {
	font-size: large;
	text-overflow: ellipsis;
	white-space: nowrap;
	overflow: hidden;
	padding: 8px
}

.info {
	width: 100%;
	display: flex;
	flex-direction: row;
	justify-content: flex-start;
	align-items: center;
	padding: 8px;
	background-color: var(--bg-secondary);
}

.at {
	margin: 0 4px;
}
</style>
