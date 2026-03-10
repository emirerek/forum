<script setup lang="ts">
import { ref, onMounted, watchEffect } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { useFetch } from "@vueuse/core";
import type { AccountModel } from "@/models/account";
import type { SubforumModel } from "@/models/subforum";

const route = useRoute();
const router = useRouter();
const auth = useAuthStore();

const form = ref({
	title: "",
	content: "",
});

const error = ref<string | null>(null);
const loading = ref(false);
const subforums = ref<SubforumModel[]>([]);
const selectedSubforumId = ref<number | null>(null);

const { data: subforumData, isFetching } = useFetch("/api/subforum/").get().json<SubforumModel[]>();

watchEffect(() => {
	if (subforumData.value) {
		subforums.value = subforumData.value;
		if (subforums.value.length > 0 && !selectedSubforumId.value) {
			selectedSubforumId.value = subforums.value[0].id;
		}
	}
});

async function onSubmit() {
	if (!auth.authenticated) {
		error.value = "You must be logged in to create a thread.";
		return;
	}
	const user = auth.user as { accountId: number; username: string } | null;
	if (!user) {
		error.value = "User info not loaded.";
		return;
	}
	loading.value = true;
	error.value = null;
	try {
		console.log(user.accountId)
		const res = await fetch("/api/thread/", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				title: form.value.title,
				content: form.value.content,
				subforumId: selectedSubforumId.value,
				accountId: user.accountId,
			}),
			credentials: "include",
		});
		if (res.ok) {
			// If backend returns only a string, show a success message and redirect to subforum or home
			try {
				const data = await res.json();
				if (data && data.id) {
					router.push(`/thread/${data.id}`);
				} else {
					// fallback: redirect to subforum or home
					router.push("/subforum/" + selectedSubforumId.value);
				}
			} catch (e) {
				// fallback: redirect to subforum or home
				router.push("/subforum/" + selectedSubforumId.value);
			}
		} else {
			const err = await res.json();
			error.value = err.message || "Failed to create thread.";
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
		<h2>Create Thread</h2>
		<div class="create-thread">
			<form @submit.prevent="onSubmit">
				<label for="title">
					Title:
					<input id="title" v-model="form.title" name="title" type="text" required />
				</label>
				<label for="content">
					Content:
					<textarea id="content" v-model="form.content" name="content" required rows="6" />
				</label>
				<label for="subforum">
					Subforum:
					<select id="subforum" v-model="selectedSubforumId" required>
						<option v-for="subforum in subforums" :key="subforum.id" :value="subforum.id">
							{{ subforum.title }}
						</option>
					</select>
				</label>
				<div v-if="error" class="error">{{ error }}</div>
				<button type="submit" :disabled="loading">{{ loading ? 'Creating...' : 'Create Thread' }}</button>
			</form>
		</div>
	</main>
</template>

<style scoped>
main {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: flex-start;
	margin-top: 16px;
}

.create-thread {
	width: 512px;
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	align-items: center;
	margin-top: 8px;
	padding: 32px;
	background-color: var(--bg-secondary, #f5f5f5);
}

form {
	width: 100%;
	display: flex;
	flex-direction: column;
}

label {
	display: flex;
	flex-direction: column;
}

input,
textarea,
select {
	margin: 8px 0;
}

button {
	margin: 8px 0 0 0;
}

.error {
	color: red;
	margin-bottom: 8px;
}
</style>
