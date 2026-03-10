<script setup lang="ts">
import { ref, watchEffect } from "vue";
import { useRoute } from "vue-router";
import { useFetch } from "@vueuse/core";
import ThreadPreviewList from "@/components/ThreadPreviewList.vue";
import ReplyList from "@/components/ReplyList.vue";
import type { ThreadModel } from "@/models/thread";
import type { ReplyModel } from "@/models/reply";
import type { AccountModel } from "@/models/account";
import { useAuthStore } from "@/stores/auth";

const route = useRoute();
const auth = useAuthStore();
const account = ref<AccountModel | null>(null);
const threads = ref<ThreadModel[] | null>(null);
const replies = ref<ReplyModel[] | null>(null);

const { data, isFetching } = useFetch(() => `/api/account/${route.params.accountId}`).get().json<AccountModel>();

watchEffect(() => {
	if (data.value) {
		account.value = data.value;
		threads.value = data.value.threads;
		replies.value = data.value.replies;
	}
});

function onDeleteAccount() {
	if (confirm('Are you sure you want to delete this user? This cannot be undone.')) {
		fetch(`/api/account/${account.value?.id}`, {
			method: 'DELETE',
			credentials: 'include'
		}).then(res => {
			if (res.ok) {
				window.location.href = '/';
			} else {
				alert('Failed to delete user.');
			}
		});
	}
}
</script>

<template>
	<main class="account">
		<h2 v-if="account">{{ account.username }}'s Profile</h2>
		<button v-if="auth.user && auth.user.isAdmin && account" class="delete" @click="onDeleteAccount">Delete
			User</button>
		<div class="lists">
			<div class="container">
				<h3>Threads</h3>
				<ThreadPreviewList v-if="threads && threads.length > 0" :threadPreviewList="threads" />
				<p v-else>No threads found.</p>
			</div>
			<div class="container">
				<h3>Replies</h3>
				<ReplyList v-if="replies && replies.length > 0" :replyList="replies" />
				<p v-else>No replies found.</p>
			</div>
		</div>
	</main>
</template>

<style scoped>
h2 {
	margin: 0 0 16px 0;
}

h3 {
	margin: 0 0 8px 0;
}

.lists {
	display: flex;
	justify-content: space-between;
}

.container {
	width: 100%;
	padding: 8px;
}

.delete {
	background-color: #ff4d4f;
	color: white;
	border: none;
	padding: 8px 16px;
	cursor: pointer;
}

.delete:hover {
	background-color: #ff7875;
}
</style>
