<script setup lang="ts">
import { ref, watchEffect, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useFetch } from '@vueuse/core';
import Thread from '@/components/Thread.vue';
import ReplyList from '@/components/ReplyList.vue';
import ReplyForm from '@/components/ReplyForm.vue';
import { useAuthStore } from '@/stores/auth';
import type { ThreadModel } from '@/models/thread';

const route = useRoute();
const threadPreview = ref<ThreadModel | null>(null);
const auth = useAuthStore();

const { data, isFetching } = useFetch(() => `/api/thread/${route.params.threadId}`).get().json<ThreadModel>();

watchEffect(() => {
	if (data.value) {
		threadPreview.value = data.value;
	}
});
</script>

<template>
	<main class="index">
		<Thread v-if="threadPreview" :threadData="threadPreview" />
		<ReplyList v-if="threadPreview && threadPreview.replies" :replyList="threadPreview.replies"
			:threadTitle="threadPreview.title" />
		<ReplyForm v-if="threadPreview && auth.authenticated" :threadId="threadPreview.id" />
	</main>
</template>

<style scoped></style>
