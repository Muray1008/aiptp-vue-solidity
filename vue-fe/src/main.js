import Vue from 'vue';
import VueRouter from 'vue-router';
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.css'
import 'vue-material/dist/theme/default.css'
import Utils from './Utils/Utils';
import App from './App.vue';
import Blockchain from './components/Blockchain';
import Block from './components/Block';


Vue.config.productionTip = false;

Vue.use(VueMaterial);

Vue.use(VueRouter);

const routes = [
    {path: '/blockchain', component: Blockchain},
    {path: '/block/:id', component: Block}
];

const router = new VueRouter({
    routes
});

const plugin = {
    install() {
        Vue.utils = Utils;
        Vue.prototype.$utils = Utils;
    }
};

Vue.use(plugin);

new Vue({
    router,
    render: h => h(App),
}).$mount('#app');
