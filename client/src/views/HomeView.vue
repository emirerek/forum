<script setup lang="ts">
import { ref, watchEffect } from "vue";
import { useFetch } from "@vueuse/core";
import SubforumList from "@/components/SubforumList.vue";
import type { SubforumModel } from "@/models/subforum";
import { useAuthStore } from "@/stores/auth";

const auth = useAuthStore();
const subforumList = ref<SubforumModel[]>([]);
const { data, isFetching } = useFetch("/api/subforum/").get().json<SubforumModel[]>();

watchEffect(() => {
	if (data.value) {
		subforumList.value = data.value;
	}
});
</script>

<template>
	<main class="subforum">
		<div class="container">
			<h2>Subforums</h2>
			<router-link v-if="auth.user && auth.user.isAdmin" to="/subforum/create">
				<button>Create Subforum</button>
			</router-link>
		</div>
		<SubforumList :subforumList="subforumList" />
		<div v-if="isFetching">Loading...</div>
	</main>
</template>

<style scoped>
.subforum {
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: flex-start;
}

.container {
	width: 100%;
	display: flex;
	justify-content: space-between;
	margin: 0 0 16px 0;
}
</style>
