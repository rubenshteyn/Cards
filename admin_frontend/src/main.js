import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router'
import store from '@/store/store'
import VueSweetalert2 from 'vue-sweetalert2'
import Popper from "vue3-popper"
import vSelect from 'vue-select'
import { useAccordion } from "vue3-rich-accordion";
import "vue3-rich-accordion/accordion-library-styles.css";
import 'sweetalert2/dist/sweetalert2.min.css'
import "@/sass/_init.scss"
import "@/sass/plugins.scss"
import "@/sass/style.scss"
import { VueReCaptcha } from 'vue-recaptcha-v3'
import xlsx from "xlsx"



createApp(App)
    .use(router)
    .use(store)
    .use(useAccordion)
    .use(VueSweetalert2)
    .use(xlsx)
    .use(VueReCaptcha, { siteKey: process.env.VUE_APP_RECAPTCHA_SITEKEY })
    .use(Popper)
    .component('v-select', vSelect)
    .mount('#app')
 