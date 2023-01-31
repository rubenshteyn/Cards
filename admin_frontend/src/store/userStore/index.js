import apiV1 from "@/axios";

export default {
    namespaced: true,
    state: {
        sessions: [],
        cards: [],
        sessionsLoaded: false
    },
    mutations: {
        ADD_SESSIONS(state, sessions) {
            state.sessionsLoaded = true
            state.sessions = sessions;
        },
        ADD_CARDS(state, cards) {
            state.cards = cards;
        }
    },
    actions: {
        async deleteSessionMethod(context, payload) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/deleteSessions", payload,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                throw new Error(response.data.error);
                            }
                            return response.data;
                        })
                        .then(() => {
                            context.dispatch("getSessionsMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async getSessionsMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("admin/getSessions",
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"],
                                "Content-Type": "application/json"
                            },
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                throw new Error(response.data.error);
                            }
                            return response.data.sessions;
                        })
                        .then(data => {
                            context.commit("ADD_SESSIONS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getCardsMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("user/getCards",
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"],
                                "Content-Type": "application/json"
                            },
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                throw new Error(response.data.error);
                            }
                            return response.data.cards;
                        })
                        .then(data => {
                            context.commit("ADD_CARDS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        }
    },

    getters: {
        getSessions(state) {
            return state.sessions
        },
        getSessionsLoaded(state) {
            return state.sessionsLoaded
        },
        getCards(state) {
            return state.cards
        },
    }
};
