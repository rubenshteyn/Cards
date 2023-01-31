import apiV1 from "@/axios";

export default {
    namespaced: true,
    state: {
        suppliers: [],
        sessions: [],
        cards: [],
        transactions: [],
        tasks: [],
        bins: [],
        userInfo: {},
        suppliersLoaded: false,
        transactionsLoaded: false,
        cardsLoaded: false,
        sessionsLoaded: false,
        userInfoLoaded: false,
        tasksLoaded: false,
        binsLoaded: false
    },
    mutations: {
        ADD_BINS(state, bins) {
            state.binsLoaded = true
            state.bins = bins;
        },
        ADD_TASKS(state, tasks) {
            state.tasksLoaded = true
            state.tasks = tasks;
        },
        ADD_SUPPLIERS(state, suppliers) {
            state.suppliersLoaded = true
            state.suppliers = suppliers;
        },
        ADD_SESSIONS(state, sessions) {
            state.sessionsLoaded = true
            state.sessions = sessions;
        },
        ADD_USER_INFO(state, userInfo) {
            state.userInfoLoaded = true
            state.userInfo = userInfo;
        },
        ADD_CARDS(state, cards) {
            state.cardsLoaded = true
            state.cards = cards;
        },
        ADD_TRANSACTIONS(state, transactions) {
            state.transactionsLoaded = true
            state.transactions = transactions;
        }
    },
    actions: {
        async deleteSessionMethod(context, payload) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.post("user/deleteSessions", payload,
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
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.get("user/getSessions",
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
        async getTasksMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.get("user/getTasks",
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
                            console.log(response.data)
                            return response.data.tasks;
                        })
                        .then(data => {
                            context.commit("ADD_TASKS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getUserInfoMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.get("user/getUserInfo",
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
                            return response.data.user;
                        })
                        .then(data => {
                            context.commit("ADD_USER_INFO", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getBinsMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.get("user/getBins",
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
                            return response.data.bins;
                        })
                        .then(data => {
                            context.commit("ADD_BINS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getCardsMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
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
        },
        async getAllTransactionsMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.get("user/getAllTransactions",
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
                            return response.data.transactions;
                        })
                        .then(data => {
                            context.commit("ADD_TRANSACTIONS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getSuppliersMethod(context) {
            context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.get("user/getSuppliers",
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
                            return response.data.suppliers;
                        })
                        .then(data => {
                            context.commit("ADD_SUPPLIERS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async changeUserInfo(context, form) {
           return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.post("user/changeUserInfo", form,
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
                            context.dispatch("getUserInfoMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async createTaskCardCreation(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.post("user/createTaskCardCreation", form,
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
                            context.dispatch("getTasksMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async createTaskCardReplenishment(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.post("user/createTaskCardReplenishment", form,
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
                            context.dispatch("getTasksMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async cancelTask(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.post("user/cancelTask", form,
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
                            context.dispatch("getTasksMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async repeatTask(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, { root: true })
                .then(() => {
                    return apiV1.post("user/repeatTask", form,
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
                            context.dispatch("getTasksMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
    },

    getters: {
        getSuppliersLoaded(state) {
            return state.suppliersLoaded
        },
        getSuppliers(state) {
            return state.suppliers
        },
        getUserInfo(state) {
            return state.userInfo
        },
        getSessions(state) {
            return state.sessions
        },
        getSessionsLoaded(state) {
            return state.sessionsLoaded
        },
        getCards(state) {
            return state.cards
        },
        getCardsLoaded(state) {
            return state.cardsLoaded
        },
        getUserInfoLoaded(state) {
            return state.userInfoLoaded
        },
        getAllTransactionsMethod(state) {
            return state.transactions
        },
        getAllTransactionsLoaded(state) {
            return state.transactionsLoaded
        },
        getTasks(state) {
            console.log(state.tasks)
            return state.tasks
        },
        getTasksLoaded(state) {
            return state.tasksLoaded
        },
        getBins(state) {
            return state.bins
        },
        getBinsLoaded(state) {
            return state.binsLoaded
        },
    }
};
