import apiV1 from "@/axios";

export default {
    namespaced: true,
    state: {
        users: [],
        cards: [],
        cardsLoaded: false,
        transactions: [],
        suppliers: [],
        finLogs: [],
        actionLog: [],
        tasks: [],
    },
    mutations: {
        ADD_USERS(state, users) {
            state.users = users;
        },
        ADD_CARDS(state, cards) {
            state.cardsLoaded = true;
            state.cards = cards;
        },
        ADD_TRANSACTIONS(state, transactions) {
            state.transactions = transactions;
        },
        ADD_SUPPLIER(state, suppliers) {
            state.suppliers = suppliers;
        },
        ADD_LOGS(state, finLogs) {
            state.finLogs = finLogs;
        },
        ADD_TASKS(state, tasks) {
            state.tasks = tasks;
        }
    },
    actions: {
        async getAccess(context, user) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/getAccess", {email: user.email},
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
                            return response;
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getUsersMethod(context) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("admin/getUsers",
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
                            return response.data.users;
                        })
                        .then(data => {
                            context.commit("ADD_USERS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getCardsMethod(context) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("admin/getCards",
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
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("admin/getAllTransactions",
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
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("admin/getSuppliers",
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
                            context.commit("ADD_SUPPLIER", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async getTasksMethod(context) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("admin/getTasks",
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
        async getAllFinLogsMethod(context) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.get("admin/getFinLogs",
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
                            return response.data.finLogs;
                        })
                        .then(data => {
                            context.commit("ADD_LOGS", data);
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                });
        },
        async addUser(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/addUser", form,
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
                            return response.data.users;
                        })
                        .then(() => {
                            context.dispatch("getUsersMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async createCards(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/createCards", form,
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
                            return response.data.cards;
                        })
                        .then(() => {
                            context.dispatch("getCardsMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async deleteUser(context, userID) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/deleteUser", {userID: userID},
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
                            return response.data.users;
                        })
                        .then(() => {
                            context.dispatch("getUsersMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async updateUser(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/updateUser", form,
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
                            return response.data.users;
                        })
                        .then(() => {
                            context.dispatch("getUsersMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async updateUsersCommission(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/updateUsersCommission", form,
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
                            return response.data.users;
                        })
                        .then(() => {
                            context.dispatch("getUsersMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async updateSupplier(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/updateSupplier", form,
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
                            return response.data.suppliers;
                        })
                        .then(() => {
                            context.dispatch("getSuppliersMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async topUpBalance(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/topUpBalance", form,
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
                        })
                        .then(() => {
                            context.dispatch("getUsersMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async topUpCardBalance(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/topUpCardBalance", form,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
                        })
                        .then(() => {
                            context.dispatch("getCardsMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async changeCardStatus(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/changeCardStatus", form,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
                        })
                        .then(() => {
                            context.dispatch("getCardsMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async changeCardLimit(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/changeCardLimit", form,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
                        })
                        .then(() => {
                            context.dispatch("getCardsMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async changeCardInfo(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/changeCardInfo", form,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
                        })
                        .then(() => {
                            context.dispatch("getCardsMethod")
                        })
                        .catch(error => {
                            console.log(error)
                            throw error;
                        });
                })
        },
        async completeTask(context, id) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/completeTask", id,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
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
        async repeatTask(context, id) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/repeatTask", id,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
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
        async createCardByTask(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/createCardByTask", form,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
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
        async cardReplenishmentByTask(context, form) {
            return context
                .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
                .then(() => {
                    return apiV1.post("admin/cardReplenishmentByTask", form,
                        {
                            headers: {
                                Authorization: context.rootGetters["auth/getTokenHeader"]
                            }
                        }
                    )
                        .then(response => {
                            if (!response.data.success) {
                                console.log(response)
                                throw new Error(response.data.error);
                            }
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
        // async getFinLog(context, form) {
        //     return context
        //         .dispatch("auth/updateAuthorizationIfNeeded", {}, {root: true})
        //         .then(() => {
        //             return apiV1.post("admin/getActionLog", form,
        //                 {
        //                     headers: {
        //                         Authorization: context.rootGetters["auth/getTokenHeader"]
        //                     }
        //                 }
        //             )
        //                 .then(response => {
        //                     if (!response.data.success) {
        //                         throw new Error(response.data.error);
        //                     }
        //                     console.log(response.data)
        //                     return response.data.finLogs;
        //                 })
        //                 .then(() => {
        //                     context.dispatch("getAllFinLogsMethod")
        //                 })
        //                 .catch(error => {
        //                     console.log(error)
        //                     throw error;
        //                 });
        //         })
        // },
    },
    getters: {
        getUsers(state) {
            return state.users
        },
        getCards(state) {
            return state.cards
        },
        getCardsLoaded(state) {
            return state.cardsLoaded
        },
        getAllTransactions(state) {
            return state.transactions
        },
        getSuppliers(state) {
            return state.suppliers
        },
        getAllFinLogs(state) {
            return state.finLogs
        },
        getTasks(state) {
            return state.tasks
        }
    }
};
