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
            name: "",
            role: 0,
            status: 0,
            isLoaded: false,
            timeZone: "",
            USDTBalance: 0
        }
    },
    mutations: {
        LOGIN(state, { email, token, status, name, role, timeZone, USDTBalance}) {
            state.user.loggedIn = true
            state.user.email = email
            state.user.authenticationToken = token
            state.user.status = status
            state.user.name = name
            state.user.role = role
            state.user.timeZone = timeZone
            state.user.USDTBalance = USDTBalance
        },
        LOGOUT(state) {
            state.user.loggedIn = false;
            state.user.email = ""
            state.user.authenticationToken = ""
            state.user.authorizationToken = ""
            state.user.status = 0
            state.user.role = 0
            state.user.name = ""
            state.user.timeZone = ""
            state.user.USDTBalance = 0
        },
        USER_LOADED(state) {
            state.user.isLoaded = true
        },
        SET_AUTHORIZATION(state, token) {
            state.user.authorizationToken = token;
        }
    },
    actions: {
        async login(context, form) {
            return apiV1.post('admin/login', form)
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
                        timeZone: data.user.timeZone
                    }
                    dbUtils.addUser(user);
                    context.commit("LOGIN", user);
                    context.commit("SET_AUTHORIZATION", data.authorization_token);
                })
                .catch(error => {
                    console.log(error)
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
                    return apiV1.get('admin/logout', {
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
        async loadUser(context) {
            dbUtils.getUser().then(user => {
                if (user && user !== {}) context.commit("LOGIN", user);
                context.commit("USER_LOADED")
            });
        },
        async updateAuthorizationIfNeeded(context) {
            let exp = 0;
            if (context.getters.getTokenAuthorization !== "") {
                const data = jwt_decode(context.getters.getTokenAuthorization);
                exp = data.exp * 1000;
            }
            if (context.getters.getTokenAuthorization === "" || Date.now() > exp) {
                return apiV1.get("admin/token", {
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
        }
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
