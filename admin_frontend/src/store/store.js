import { createStore } from 'vuex'
import usersModule from "./userStore";
import authModule from "./authStore";
import adminFunctionsModule from "./adminFunctionsStore"

const store = createStore({
    state: {  },
    mutations: {  },
    actions: {  },
    modules: { users: usersModule, auth: authModule, admin: adminFunctionsModule }
})

export default store