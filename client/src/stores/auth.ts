import { defineStore } from "pinia";
import { ref } from "vue";
import type { AccountModel } from "@/models/account";

type SessionUser = {
	accountId: number;
	username: string;
	isAdmin?: boolean;
};

export const useAuthStore = defineStore("auth", () => {
	const user = ref<SessionUser | null>(null);
	const authenticated = ref(false);
	const loading = ref(false);
	const error = ref<null | string>(null);

	async function fetchMe() {
		loading.value = true;
		error.value = null;
		try {
			const res = await fetch("/api/auth/me", { credentials: "include" });
			if (res.ok) {
				const data = await res.json();
				authenticated.value = !!data.authenticated;
				user.value = data.authenticated ? data : null;
			} else {
				authenticated.value = false;
				user.value = null;
			}
		} catch (e) {
			error.value = String(e);
			authenticated.value = false;
			user.value = null;
		} finally {
			loading.value = false;
		}
	}

	async function logout() {
		await fetch("/api/auth/logout", { method: "POST", credentials: "include" });
		user.value = null;
		authenticated.value = false;
	}

	return { user, authenticated, loading, error, fetchMe, logout };
});
