<template>
	<div class="container">
		<div class="subforum">
			<a class="title" :href="subforumUrl">
				{{ subforum.title }}
			</a>
			<p class="description">
				{{ subforum.description }}
			</p>
			<button class="edit" v-if="auth.user && auth.user.isAdmin" @click="goToEdit">Edit</button>
			<button class="delete" v-if="auth.user && auth.user.isAdmin" @click="onDelete">Delete</button>
		</div>
		<div class="thread">
			<SubforumThread v-if="subforum.threads && subforum.threads.length > 0"
				:threadPreview="subforum.threads[0]" />
		</div>
	</div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import type { SubforumModel } from "@/models/subforum";
import SubforumThread from "./SubforumThread.vue";

const props = defineProps<{ subforum: SubforumModel }>();
const subforumUrl = computed(() => `/subforum/${props.subforum.id}`);
const auth = useAuthStore();
const router = useRouter();
function goToEdit() {
	router.push(`/subforum/${props.subforum.id}/edit`);
}
function onDelete() {
	if (confirm('Are you sure you want to delete this subforum? This cannot be undone.')) {
		fetch(`/api/subforum/${props.subforum.id}`, {
			method: 'DELETE',
			credentials: 'include'
		}).then(res => {
			if (res.ok) {
				router.push('/');
			} else {
				alert('Failed to delete subforum.');
			}
		});
	}
}
</script>

<style scoped>
.container {
	width: 100%;
	display: flex;
	flex-direction: column;
	justify-content: flex-start;
	align-items: flex-start;
}

.subforum {
	width: 100%;
	padding: 8px;
	background-color: var(--component-hover);
}

.description {
	margin: 4px 0 0 0;
	color: var(--text-secondary);
}

.thread {
	width: 100%;
	background-color: var(--component);
}

.title {
	font-size: 1.25rem;
	font-weight: 600;
	color: var(--link);
	text-decoration: none;
}

.title:hover {
	text-decoration: underline;
	color: var(--link-hover);
}

.edit {
	margin: 8px 4px 0 0;
}
</style>
