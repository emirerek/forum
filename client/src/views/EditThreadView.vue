<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { useFetch } from "@vueuse/core";
import type { ThreadModel } from "@/models/thread";

const route = useRoute();
const router = useRouter();
const auth = useAuthStore();

const threadId = Number(route.params.threadId);
const form = ref({
	title: "",
	content: "",
});
const error = ref<string | null>(null);
const loading = ref(false);

// Fetch thread data on mount
onMounted(async () => {
	const res = await fetch(`/api/thread/${threadId}`);
	if (res.ok) {
		const thread: ThreadModel = await res.json();
		form.value.title = thread.title;
		form.value.content = thread.content;
	} else {
		error.value = "Failed to load thread.";
	}
});

async function onSubmit() {
	if (!auth.authenticated) {
		error.value = "You must be logged in to edit a thread.";
		return;
	}
	loading.value = true;
	error.value = null;
	try {
		const res = await fetch(`/api/thread/${threadId}`, {
			method: "PATCH",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({
				title: form.value.title,
				content: form.value.content,
			}),
			credentials: "include",
		});
		if (res.ok) {
			router.push(`/thread/${threadId}`);
		} else {
			const err = await res.json();
			error.value = err.message || "Failed to update thread.";
		}
	} catch (e: any) {
		error.value = e.message || String(e);
	} finally {
		loading.value = false;
	}
}
</script>

<template>
	<main>
		<h2>Edit Thread</h2>
		<div class="edit-thread">
			<form @submit.prevent="onSubmit">
				<label for="title">
					Title:
					<input id="title" v-model="form.title" name="title" type="text" required />
				</label>
				<label for="content">
					Content:
					<textarea id="content" v-model="form.content" name="content" required rows="6" />
				</label>
				<div v-if="error" class="error">{{ error }}</div>
				<button type="submit" :disabled="loading">{{ loading ? 'Saving...' : 'Save Changes' }}</button>
			</form>
		</div>
	</main>
</template>

<style scoped>
h2 {
	margin-bottom: 8px;
}

main {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: flex-start;
	min-height: 60vh;
}

.edit-thread {
	width: 100%;
	max-width: 600px;
	background: var(--bg-secondary, #222);
	padding: 2rem;
}

label {
	display: block;
	margin-bottom: 1rem;
}

input,
textarea,
select {
	width: 100%;
	padding: 0.5rem;
	background: var(--bg-primary, #111);
	color: var(--text-primary, #fff);
}

button {
	padding: 0.5rem 1.5rem;
}

.error {
	color: #ffb3b3;
	margin-bottom: 1rem;
}
</style>
