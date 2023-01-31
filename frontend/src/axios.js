import axios from "axios";

const apiV1 = axios.create({
    baseURL: process.env.VUE_APP_API_URL,
    timeout: 100000,
    headers: {
        Accept: "application/json;charset=UTF-8",
        "Content-Type": "application/json;charset=UTF-8",
    }
})

export default apiV1