import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import SubforumView from "../views/SubforumView.vue";
import ThreadView from "../views/ThreadView.vue";
import RegisterView from "@/views/RegisterView.vue";
import LoginView from "@/views/LoginView.vue";
import CreateThreadView from "@/views/CreateThreadView.vue";
import EditThreadView from "@/views/EditThreadView.vue";
import AccountView from "@/views/AccountView.vue";
import CreateSubforumView from "../views/CreateSubforumView.vue";
import EditSubforumView from "../views/EditSubforumView.vue";

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: "/",
			name: "home",
			component: HomeView,
		},
		{
			path: "/subforum/:subforumId",
			name: "subforum",
			component: SubforumView,
			props: true
		},
		{
			path: "/thread/:threadId",
			name: "thread",
			component: ThreadView,
			props: true
		},
		{
			path: "/register",
			name: "register",
			component: RegisterView,
		},
		{
			path: "/login",
			name: "login",
			component: LoginView,
		},
		{
			path: "/create-thread",
			name: "create-thread",
			component: CreateThreadView,
		},
		{
			path: "/thread/create",
			name: "create-thread",
			component: CreateThreadView,
			props: route => ({ subforumId: route.query.subforumId })
		},
		{
			path: "/thread/:threadId/edit",
			name: "edit-thread",
			component: EditThreadView,
			props: true
		},
		{
			path: "/account/:accountId",
			name: "account",
			component: AccountView,
			props: true
		},
		{
			path: "/subforum/create",
			name: "create-subforum",
			component: CreateSubforumView,
		},
		{
			path: "/subforum/:subforumId/edit",
			name: "edit-subforum",
			component: EditSubforumView,
			props: true
		},
	],
})

export default router
