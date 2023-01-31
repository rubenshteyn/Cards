import { createApp } from 'vue'
import App from '@/App.vue'
import i18n from '@/i18n'
import router from '@/router'
import store from '@/store/store'
import VueCookies from 'vue-cookies'
import vSelect from 'vue-select'
import VueSimpleContextMenu from 'vue-simple-context-menu'
import VueSweetalert2 from 'vue-sweetalert2'
import Popper from "vue3-popper"
import 'sweetalert2/dist/sweetalert2.min.css'
import 'vue-simple-context-menu/dist/vue-simple-context-menu.css'
import "@/sass/_init.scss"
import "@/sass/plugins.scss"
import "@/sass/style.scss"
import Select2 from 'vue3-select2-component';
import xlsx from "xlsx"
import {VueReCaptcha} from 'vue-recaptcha-v3'

createApp(App)
    .use(i18n)
    .use(router)
    .use(store)
    .use(VueCookies)
    .use(Popper)
    .use(VueSweetalert2)
    .use(xlsx)
    .use(VueReCaptcha, { siteKey: process.env.VUE_APP_RECAPTCHA_SITEKEY })
    .component('v-select', vSelect)
    .component('Select2', Select2)
    .component('vue-simple-context-menu', VueSimpleContextMenu)
    .mount('#app')
