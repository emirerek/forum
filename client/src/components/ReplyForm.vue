<template>
	<form class="reply-form" @submit.prevent="onSubmit">
		<textarea id="reply-content" v-model="content" required rows="4" placeholder="Write your reply..."></textarea>
		<div v-if="error" class="error">{{ error }}</div>
		<button type="submit" :disabled="loading">{{ loading ? 'Posting...' : 'Post Reply' }}</button>
	</form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import type { AccountModel } from '@/models/account';

const props = defineProps<{ threadId: number }>();
const emit = defineEmits(['reply-posted']);

const content = ref('');
const error = ref<string | null>(null);
const loading = ref(false);
const auth = useAuthStore();

async function onSubmit() {
	error.value = null;
	loading.value = true;
	try {
		const user = auth.user as { accountId?: number; id?: number } | null;
		const res = await fetch(`/api/reply/`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				threadId: props.threadId,
				content: content.value,
				accountId: user?.accountId ?? user?.id,
			}),
			credentials: 'include',
		});
		if (res.ok) {
			content.value = '';
			emit('reply-posted');
		} else {
			const err = await res.json();
			error.value = err.message || 'Failed to post reply.';
		}
	} catch (e: any) {
		error.value = e.message || String(e);
	} finally {
		loading.value = false;
	}
}
</script>

<style scoped>
.reply-form {
	display: flex;
	flex-direction: column;
	gap: 8px;
	margin-top: 24px;
}

textarea {
	resize: vertical;
}

.error {
	color: red;
}
</style>
