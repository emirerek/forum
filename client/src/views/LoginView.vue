<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

const form = ref({
	username: "",
	password: "",
});
const router = useRouter();
const error = ref<string | null>(null);
const loading = ref(false);

async function onSubmit() {
	error.value = null;
	loading.value = true;
	const formData = new FormData()
	formData.append("username", form.value.username)
	formData.append("password", form.value.password)
	const result = await fetch("/api/auth/login", {
		method: "POST",
		body: formData,
		credentials: "include"
	})
	if (result.ok) {
		router.push("/");
	} else {
		try {
			const err = await result.json();
			error.value = err.message || "Login failed.";
		} catch (e) {
			error.value = "Login failed.";
		}
	}
	loading.value = false;
}
</script>

<template>
	<main>
		<div class="login">
			<h2>Log In</h2>
			<form @submit.prevent="onSubmit" enctype="multipart/form-data">
				<label for="username">
					Username:
					<input id="username" v-model="form.username" name="username" type="text" required />
				</label>
				<label for="password">
					Password:
					<input id="password" v-model="form.password" name="password" type="password" required />
				</label>
				<div v-if="error" class="error">{{ error }}</div>
				<button type="submit" :disabled="loading">{{ loading ? 'Logging in...' : 'Log In' }}</button>
			</form>
		</div>
	</main>
</template>

<style scoped>
main {
	display: flex;
	justify-content: center;
}

h2 {
	margin-bottom: 16px;
}

.login {
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	align-items: center;
	margin-top: 8px;
	padding: 32px;
	background-color: var(--bg-secondary);
}

form {
	display: flex;
	flex-direction: column;
}

label {
	display: flex;
	flex-direction: column;
}

input {
	margin: 8px 0;
}

.error {
	color: #ffb3b3;
	margin-bottom: 1rem;
}
</style>