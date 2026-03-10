<template>
	<ul>
		<li v-for="reply in replyList" :key="reply.id">
			<Reply :reply="reply" />
		</li>
	</ul>
</template>

<script setup lang="ts">
import Reply from "./Reply.vue";
import type { ReplyModel } from "@/models/reply";

const props = defineProps<{ replyList: ReplyModel[] }>();
</script>

<style scoped>
ul {
	width: 100%;
}

li {
	list-style: none;
	margin: 0 0 8px 0;
}
</style>
