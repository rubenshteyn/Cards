import jwt_decode from "jwt-decode";
import dbUtils from "../indexedDButils.js";
import router from "@/router";
import apiV1 from '@/axios'

export default {
    namespaced: true,
    state: {
        user: {
            email: "",
            loggedIn: false,
            authenticationToken: "",
            authorizationToken: "",
            status: 0,
            name: "",
            role: 0,
            access: "login",
            isLoaded: false,
            timeZone: "",
        }
    },
    mutations: {
        LOGIN(state, { email, token, status, access, name, role, timeZone}) {
            state.user.loggedIn = true
            state.user.email = email
            state.user.authenticationToken = token
            state.user.status = status
            state.user.access = access
            state.user.name = name
            state.user.role = role
            state.user.timeZone = timeZone
        },
        LOGOUT(state) {
            state.user.loggedIn = false;
            state.user.email = ""
            state.user.authenticationToken = ""
            state.user.authorizationToken = ""
            state.user.status = 0
            state.user.access = "login"
            state.user.role = 0
            state.user.name = ""
            state.user.timeZone = ""
        },
        USER_LOADED(state) {
            state.user.isLoaded = true
        },
        SET_AUTHORIZATION(state, token) {
            state.user.authorizationToken = token
        }
    },
    actions: {
        async login(context, form) {
            return apiV1.post('auth/login', form)
                .then(response => {
                    if (!response.data.success) {
                        throw new Error(response.data.error);
                    }
                    return response.data;
                })
                .then(data => {
                    let user = {
                        email: form.email,
                        token: data.authentication_token,
                        status: data.user.status,
                        name: data.user.name,
                        role: data.user.role,
                        access: "login",
                        timeZone: data.user.timeZone
                    }
                    dbUtils.addUser(user);
                    context.commit("LOGIN", user);
                    context.commit("SET_AUTHORIZATION", data.authorization_token);
                })
                .catch(error => {
                    context.dispatch("logout");
                    throw error;
                });
        },
        async logout(context) {
            return dbUtils
                .removeUser({
                    email: context.getters.currentUser.email
                })
                .then(() => {
                    return apiV1.get('auth/logout', {
                        headers: {
                            Authorization:
                                "Bearer " + context.getters.getTokenAuthentication
                        }})
                        .then(() => {
                        context.commit("LOGOUT");
                        router.push("/login");
                    });
                });
        },
        async signup(context, form) {
            return apiV1.post('auth/registration', form)
                .then(response => {
                    if (!response.data.success) {
                        throw new Error(response.data.error);
                    }
                    return response.data;
                })
                .catch(error => {
                    console.log(error)
                    throw error;
                });
        },
        async loadUser(context) {
            dbUtils.getUser().then(user => {
                if (user && user !== {}) context.commit("LOGIN", user);
                context.commit("USER_LOADED")
            });
        },
        async access(context, token) {
            let headers = {
                headers: {
                    "Authorization": "Bearer " + token
                }
            }
            return apiV1.get('admin/access', headers)
                .then(response => {
                    if (!response.data.success) {
                        throw new Error(response.data.error);
                    }
                    return response.data;
                })
                .then(data => {
                    let user = {
                        email: data.user.email,
                        token: data.authentication_token,
                        status: data.user.status,
                        name: data.user.name,
                        role: data.user.role,
                        access: "admin",
                        timeZone: data.user.timeZone
                    }
                    dbUtils.addUser(user, true);
                    context.commit("LOGIN", user);
                    context.commit("SET_AUTHORIZATION", data.authorization_token);
                })
                .catch(error => {
                    console.log(error)
                    context.dispatch("logout");
                    throw error;
                });
        },
        async updateAuthorizationIfNeeded(context) {
            let exp = 0;
            if (context.getters.getTokenAuthorization !== "") {
                const data = jwt_decode(context.getters.getTokenAuthorization);
                exp = data.exp * 1000;
            }
            if (context.getters.getTokenAuthorization === "" || Date.now() > exp) {
                return apiV1.get("auth/token", {
                    headers: {
                        Authorization: "Bearer " + context.getters.getTokenAuthentication
                    }
                })
                    .then(response => {
                        if (!response.data.success) {
                            throw new Error(response.data.error);
                        }
                        return response.data;
                    })
                    .then(data => {
                        context.commit("SET_AUTHORIZATION", data.authorization_token);
                    })
                    .catch(error => {
                        console.log(error)
                        context.dispatch("logout");
                        throw error;
                    });
            } else {
                return new Promise(success => {
                    success([]);
                });
            }
        },
    },
    getters: {
        currentUser(state) {
            return state.user;
        },
        isLoggedIn(state) {
            if (!state.user) return false;
            return state.user.loggedIn;
        },
        isLoaded(state) {
            if (!state.user) return false;
            return state.user.isLoaded;
        },
        getTokenAuthorization(state) {
            return state.user.authorizationToken;
        },
        getTokenHeader(state) {
            return "Bearer " + state.user.authorizationToken;
        },
        getTokenAuthentication(state) {
            return state.user.authenticationToken;
        }
    }
};
