<template>
	<header>
		<div>
			<h1><a href="/">FORUM</a></h1>
			<nav>
				<ul>
					<li v-if="auth.authenticated">
						<a href="" @click="handleLogout">Log out</a>
					</li>
					<li v-if="!auth.authenticated">
						<router-link to="/login">Login</router-link>
					</li>
					<li v-if="!auth.authenticated">
						<router-link to="/register">Register</router-link>
					</li>
				</ul>
			</nav>
		</div>
	</header>
</template>

<script setup lang="ts">
import { useAuthStore } from "../stores/auth";
import { useRouter } from "vue-router";

const auth = useAuthStore();
const router = useRouter();

async function handleLogout() {
	await auth.logout();
	router.push("/");
}
</script>

<style scoped>
header {
	width: 100%;
	height: 48px;
	position: sticky;
	top: 0;
	display: flex;
	justify-content: center;
	align-items: center;
	padding: 0 16px;
	background-color: var(--bg-secondary);
}

header h1 {
	margin: 0;
}

div {
	width: var(--app-width);
	display: flex;
	flex-direction: row;
	justify-content: space-between;
	align-items: center;
}

nav {
	display: flex;
	justify-content: flex-end;
	align-items: center;
}

nav>ul {
	display: flex;
	flex-direction: row;
	justify-content: flex-end;
	align-items: center;
	padding: 0;
	list-style-type: none;
}

ul>li {
	margin: 0 8px 0 8px;
}
</style>
