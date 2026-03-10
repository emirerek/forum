Create a new subforum page for admins only.

```vue
<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";

const auth = useAuthStore();
const router = useRouter();

const title = ref("");
const description = ref("");
const error = ref("");
const loading = ref(false);

if (!auth.user || !auth.user.isAdmin) {
	router.replace("/");
}

async function createSubforum() {
	error.value = "";
	loading.value = true;
	try {
		const res = await fetch("/api/subforum/", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			credentials: "include",
			body: JSON.stringify({ title: title.value, description: description.value })
		});
		if (res.ok) {
			router.push("/");
		} else {
			const data = await res.json().catch(() => ({}));
			error.value = data?.error || "Failed to create subforum.";
		}
	} catch (e) {
		error.value = String(e);
	} finally {
		loading.value = false;
	}
}
</script>

<template>
	<main class="create-subforum">
		<h2>Create Subforum</h2>
		<form @submit.prevent="createSubforum">
			<label>
				Title
				<input v-model="title" required />
			</label>
			<label>
				Description
				<textarea v-model="description" required></textarea>
			</label>
			<button type="submit" :disabled="loading">Create</button>
			<div v-if="error" class="error">{{ error }}</div>
		</form>
	</main>
</template>

<style scoped>
h2 {
	margin: 0 0 16px 0;
}

.create-subforum {
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
```
