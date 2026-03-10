<template>
	<article class="reply">
		<div class="info">
			<a :href="accountUrl">
				{{ reply.account.username }}
			</a>
			<span class="at">replied to</span>
			<a class="title" :href="threadUrl" :title="threadTitle">
				{{ threadTitle }}
			</a>
			<span class="at">at</span>
			<Date :date="reply.account.createdAt" />
			<button v-if="auth.user && (auth.user.isAdmin || auth.user.accountId === reply.account.id)" class="delete"
				@click="onDelete">Delete</button>
		</div>
		<div class="content">{{ reply.content }}</div>
	</article>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import Date from "./Date.vue";
import type { ReplyModel } from "@/models/reply";
import type { ThreadModel } from "@/models/thread";
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const props = defineProps<{ reply: ReplyModel }>();
const accountUrl = computed(() => `/account/${props.reply.account.id}`);
const threadUrl = computed(() => `/thread/${props.reply.threadId}`);

const auth = useAuthStore();
const router = useRouter();

const thread = ref<ThreadModel | null>(props.reply.thread || null);
const threadTitle = ref<string>(thread.value ? thread.value.title : "");

onMounted(async () => {
	if (!thread.value) {
		const res = await fetch(`/api/thread/${props.reply.threadId}`);
		if (res.ok) {
			const data = await res.json();
			thread.value = data;
			threadTitle.value = data.title;
		}
	}
});

function onDelete() {
	if (confirm('Are you sure you want to delete this reply? This cannot be undone.')) {
		fetch(`/api/reply/${props.reply.id}`, {
			method: 'DELETE',
			credentials: 'include'
		}).then(res => {
			if (res.ok) {
				// Optionally emit an event or reload the page
				window.location.reload();
			} else {
				alert('Failed to delete reply.');
			}
		});
	}
}
</script>

<style scoped>
h2 {
	margin: 0 0 16px 0;
}

.reply {
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: flex-start;
	align-items: flex-start;
	background-color: var(--bg-secondary);
}

.content {
	padding: 16px;
}

.info {
	width: 100%;
	display: flex;
	flex-direction: row;
	justify-content: flex-start;
	align-items: center;
	padding: 8px;
	background-color: var(--component);
}

.at {
	margin: 0 4px;
}

button {
	margin-left: 8px;
}
</style>
