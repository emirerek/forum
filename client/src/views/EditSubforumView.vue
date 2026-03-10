<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';

const route = useRoute();
const router = useRouter();
const auth = useAuthStore();

const title = ref('');
const description = ref('');
const error = ref('');
const loading = ref(false);

onMounted(async () => {
	const res = await fetch(`/api/subforum/${route.params.subforumId}`);
	if (res.ok) {
		const data = await res.json();
		title.value = data.title;
		description.value = data.description;
	} else {
		error.value = 'Failed to load subforum.';
	}
});

async function updateSubforum() {
	error.value = '';
	loading.value = true;
	try {
		const res = await fetch(`/api/subforum/${route.params.subforumId}`, {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include',
			body: JSON.stringify({ title: title.value, description: description.value })
		});
		if (res.ok) {
			router.push('/');
		} else {
			const data = await res.json().catch(() => ({}));
			error.value = data?.error || 'Failed to update subforum.';
		}
	} catch (e) {
		error.value = String(e);
	} finally {
		loading.value = false;
	}
}
</script>

<template>
	<main class="edit-subforum">
		<h2>Edit Subforum</h2>
		<form @submit.prevent="updateSubforum">
			<label>
				Title
				<input v-model="title" required />
			</label>
			<label>
				Description
				<textarea v-model="description" required></textarea>
			</label>
			<button type="submit" :disabled="loading">Update</button>
			<div v-if="error" class="error">{{ error }}</div>
		</form>
	</main>
</template>

<style scoped>
h2 {
	margin: 0 0 16px 0;
}

.edit-subforum {
	max-width: 400px;
	margin: 2rem auto;
	padding: 2rem;
	background: var(--component);
	border-radius: 8px;
}

label {
	display: flex;
	flex-direction: column;
	margin-bottom: 1rem;
}

.error {
	color: red;
	margin-top: 1rem;
}
</style>
