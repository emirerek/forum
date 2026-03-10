<script setup lang="ts">
import { ref } from "vue";

const form = ref({
	username: "",
	email: "",
	password: "",
});
const error = ref<string | null>(null);
const loading = ref(false);

async function onSubmit() {
	error.value = null;
	loading.value = true;
	const formData = new FormData()
	formData.append("username", form.value.username)
	formData.append("email", form.value.email)
	formData.append("password", form.value.password)
	const result = await fetch("/api/auth/register", {
		method: "POST",
		body: formData
	})
	if (result.ok) {
		window.location.href = "/login";
	} else {
		try {
			const err = await result.json();
			error.value = err.message || "Registration failed.";
		} catch (e) {
			error.value = "Registration failed.";
		}
	}
	loading.value = false;
}
</script>

<template>
	<main>
		<div class="register">
			<h2>Register</h2>
			<form @submit.prevent="onSubmit" enctype="multipart/form-data">
				<label for="username">
					Username:
					<input id="username" v-model="form.username" name="username" type="text" required />
				</label>
				<label for="email">
					Email:
					<input id="email" v-model="form.email" name="email" type="email" required />
				</label>
				<label for="password">
					Password:
					<input id="password" v-model="form.password" name="password" type="password" required />
				</label>
				<div v-if="error" class="error">{{ error }}</div>
				<button type="submit" :disabled="loading">{{ loading ? 'Registering...' : 'Register' }}</button>
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

.register {
	width: min-content;
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