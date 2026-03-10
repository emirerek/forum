<template>
	<div class="account">
		<p>
			<a :href="accountURL">
				{{ accountPreview.username }}
			</a>
		</p>
		<p>
			Joined at:
		</p>
		<Date :date="accountPreview.createdAt" />
	</div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import type { AccountModel } from "@/models/account";
import Date from "./Date.vue";

const props = defineProps<{ accountPreview: AccountModel }>();
const accountURL = computed(() => `/account/${props.accountPreview.id}`);
</script>

<style scoped>
.account {
	width: 100%;
	display: flex;
	flex-direction: row;
	justify-content: flex-start;
	align-items: center;
	padding: 8px;
	background-color: var(--component);
}
</style>
