import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/session', component: LoginView},
		{path: '/stream', component: StreamView}, // TODO: DA CAMBIARE
	]
})

export default router