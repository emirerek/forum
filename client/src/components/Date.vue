<template>
	<span class="date">
		<time :datetime="date" :title="date">
			{{ formattedDate }}
		</time>
	</span>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { formatDate } from "../util/formatDate";

const props = defineProps<{ date: string }>();
const formattedDate = ref("");
onMounted(() => {
	formattedDate.value = formatDate(props.date);
});
</script>

<style scoped>
.date {
	display: flex;
	justify-content: center;
	align-items: center;
	color: var(--text-primary);
	cursor: help;
}
</style>
