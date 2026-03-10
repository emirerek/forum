<template>
	<div class="title">
		<h2>{{ threadData.title }}</h2>
		<div class="title-subforum" v-if="subforum">
			<span>posted in </span>
			<a :href="subforumUrl" :title="subforum?.title">
				{{ subforum?.title }}
			</a>
		</div>
		<button v-if="canEdit" class="edit" @click="goToEdit">Edit</button>
		<button v-if="canEdit" class="delete" @click="onDelete">Delete</button>
	</div>
	<article class="thread">
		<div class="info">
			<a :href="accountUrl">
				{{ threadData.account.username }}
			</a>
			<span class="at">posted at</span>
			<Date :date="threadData.account.createdAt" />
		</div>
		<div class="content">{{ threadData.content }}</div>
	</article>
</template>

<script setup lang="ts">
import { ref, computed, watchEffect } from "vue";
import { useFetch } from '@vueuse/core';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';
import Date from "./Date.vue";
import type { ThreadModel } from "@/models/thread";
import type { SubforumModel } from "@/models/subforum";

const props = defineProps<{ threadData: ThreadModel }>();
const subforum = ref<SubforumModel | null>(null);
const { data, isFetching } = useFetch(() => `/api/subforum/${props.threadData.subforumId}`).get().json<SubforumModel>();
watchEffect(() => {
	if (data.value) {
		subforum.value = data.value;
	}
});
const accountUrl = computed(() => `/account/${props.threadData.account.id}`);
const subforumUrl = computed(() => `/subforum/${props.threadData.subforumId}`);
const auth = useAuthStore();
const router = useRouter();
const canEdit = computed(() => {
	if (!auth.authenticated || !auth.user) return false;
	return auth.user.accountId === props.threadData.account.id || auth.user.isAdmin;
});
function goToEdit() {
	router.push(`/thread/${props.threadData.id}/edit`);
}
function onDelete() {
	if (confirm("Are you sure you want to delete this thread? This cannot be undone.")) {
		fetch(`/api/thread/${props.threadData.id}`, {
			method: "DELETE",
			credentials: "include"
		}).then(res => {
			if (res.ok) {
				// Redirect to subforum after delete
				router.push(`/subforum/${props.threadData.subforumId}`);
			} else {
				alert("Failed to delete thread.");
			}
		});
	}
}
</script>

<style scoped>
.title {
	display: flex;
	justify-content: flex-start;
	align-items: center;
	margin-bottom: 16px;
	gap: 8px;
}

.title-subforum {

	margin: 0 0 0 6px;
}

.thread {
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: flex-start;
	align-items: flex-start;
	margin-bottom: 8px;
	background-color: var(--bg-secondary);
}

.container {
	display: flex;
	flex-direction: row;
	justify-content: flex-start;
	align-items: flex-start;
}

.content {
	padding: 16px;
}

.info {
	width: 100%;
	display: flex;
	flex-direction: row;
	justify-content: flex-start;
	align-items: center;
	padding: 8px;
	background-color: var(--component);
}

.at {
	margin: 0 4px;
}

.edit,
.delete {
	margin-left: 8px;
}
</style>