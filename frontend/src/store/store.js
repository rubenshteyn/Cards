import { createStore } from 'vuex'
import usersModule from "./userStore";
import authModule from "./authStore";

// Create a new store instance.
const store = createStore({
    state: {  },
    mutations: {  },
    actions: {  },
    modules: { users: usersModule, auth: authModule }
})

export default store