import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'
import ProfileView from '../views/ProfileView.vue'
import SelfView from '../views/SelfView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/session', component: LoginView},
		{path: '/stream', component: StreamView}, // TODO: DA CAMBIARE
		{path: '/profile', component: ProfileView}, // TODO: DA CAMBIARE
		{path: '/self', component: SelfView}, // TODO: DA CAMBIARE
	]
})

export default router