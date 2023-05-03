import axios from "axios"

const API_URL = import.meta.env.VITE_WEB_API_URL;

const instance = axios.create({
    baseURL: API_URL,
})

// instance.interceptors.request.use(function (config) {
//     // Setting Bearer token to each request header
//     config.headers.Authorization = `Bearer ${KeyCloakService.UserToken()}`;
//     return config;
// }, function (error) {
//     return Promise.reject(error);
// });


export default instance