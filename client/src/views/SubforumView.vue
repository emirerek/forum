<script setup lang="ts">
import { ref, watchEffect } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { useFetch } from "@vueuse/core";
import ThreadPreviewList from "@/components/ThreadPreviewList.vue";
import type { ThreadModel } from "@/models/thread";
import type { SubforumModel } from "@/models/subforum";

const route = useRoute();
const auth = useAuthStore();
const threadPreviewList = ref<ThreadModel[] | null>(null);
const subforum = ref<SubforumModel | null>(null);

const { data, isFetching } = useFetch(() => `/api/thread/?subforumId=${route.params.subforumId}`).get().json<ThreadModel[]>();
const { data: subforumData } = useFetch(() => `/api/subforum/${route.params.subforumId}`).get().json<SubforumModel>();

watchEffect(() => {
	if (data.value) {
		threadPreviewList.value = data.value;
	}
});

watchEffect(() => {
	if (subforumData.value) {
		subforum.value = subforumData.value;
	}
});
</script>

<template>
	<main class="subforum">
		<div class="container">
			<h2>{{ subforum?.title || "Subforum" }}</h2>
			<router-link v-if="auth.authenticated" to="/thread/create">
				<button>Create Thread</button>
			</router-link>
		</div>
		<ThreadPreviewList v-if="threadPreviewList" :threadPreviewList="threadPreviewList" />
	</main>
</template>

<style scoped>
h2 {
	margin: 0 0 16px 0;
}

.container {
	width: 100%;
	display: flex;
	justify-content: space-between;
	margin: 0 0 16px 0;
}
</style>
