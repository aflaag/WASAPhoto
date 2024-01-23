import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'
import ProfileView from '../views/ProfileView.vue'
import SelfView from '../views/SelfView.vue'
import NotFoundView from '../views/NotFoundView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/user/:uname/stream', component: StreamView},
		{path: '/user/:uname', component: ProfileView},
		{path: '/user/self', component: SelfView},
		{path: '/404', component: NotFoundView},
	]
})

export default router