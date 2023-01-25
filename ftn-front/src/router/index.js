import Vue from 'vue'
import VueRouter from 'vue-router'

import HomePage from '../views/HomePage'

import RegisterStudent from '../components/RegisterStudent'
import RegisterProfesor from '../components/RegisterProfesor'

Vue.use(VueRouter)

const routes = [
	{
		path: "/",
		name: HomePage,
		component: HomePage,
		children: [
			{
				path: "RegisterStudent",
				name: "RegisterStudent",
				component: RegisterStudent
			},
			{
				path: "RegisterProfesor",
				name: "RegisterProfesor",
				component: RegisterProfesor
			},
		]
	},
	{
		path: '*',
		redirect: "/"
	}
]

const router = new VueRouter({
	mode: 'history',
	base: process.env.BASE_URL,
	routes
})

export default router