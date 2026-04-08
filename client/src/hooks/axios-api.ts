"use client";

import { useAuth } from "@/stores/auth-store";
import axios from "axios";





// export const apiClient = axios.create({
//     baseURL: "/api",
//     withCredentials: true
// });
// export const api = axios.create({
//     baseURL: "/api",
//     withCredentials: true
// });

// apiClient.interceptors.request.use(config => {
//     const token = useAuth.getState().token;
//     console.log(token, "api")

//     if (token) {
//         config.headers.Authorization = `Bearer ${token}`
//     }
//     return config;
// },
//     error => {
//         return Promise.reject(error)
//     })

// let isRefeshing = false;

// type failedQueue = {
//     resolve: (token: string) => void,
//     reject: (error: AxiosError) => void
// }
// let failedQueue: failedQueue[] = [];

// const processQueue = (error: AxiosError | null, token?: string) => {
//     failedQueue.forEach(({ resolve, reject }) => {
//         if (error) {
//             reject(error)
//         } else {
//             resolve(token ?? "")
//         }
//     })
//     failedQueue = [];
// }

// apiClient.interceptors.response.use(
//     (response) => response,
//     async (error) => {
//         const clearAuth = useAuth.getState().reset;
//         const original_request = error.config;
//         if (error.response.status === 401 && !original_request._retry) {
//             if (
//                 original_request.url.includes('/login') ||
//                 original_request.url.includes('/signup') ||
//                 original_request.url.includes('/refresh')

//             ) {
//                 return Promise.reject(error);
//             }
//             original_request._retry = true;

//             //refreshing 
//             if (isRefeshing) {
//                 return new Promise((resolve, reject) => {
//                     failedQueue.push({ resolve, reject })
//                 }).then((token) => {
//                     original_request.headers.Authorization = `Bearer ${token}`
//                 })
//             }
//             isRefeshing = true;
//             try {

//                 const { data } = await axios.post<APIResponse<{ token: string }>>(`${BASE_URL}/auth/refresh`, "", {
//                     withCredentials: true
//                 })
//                 const token = data.result.token;

//                 processQueue(null, token)
//                 useAuth.getState().setToken(token);

//                 original_request.headers.Authorization = `Bearer ${token}`;
//                 return apiClient(original_request)

//             } catch (refreshError) {

//                 processQueue(refreshError as AxiosError);
//                 failedQueue = [];
//                 clearAuth();
//                 isRefeshing = false;
//                 return Promise.reject(refreshError)
//             } finally {
//                 isRefeshing = false;
//             }
//         }

//         return Promise.reject(error)
//     }
// )


const api = axios.create({
    baseURL: "/api/server",
    withCredentials: true
})

api.interceptors.response.use(
    (response) => response,
    async (error) => {
        if (error.status === 401) {
            useAuth.getState().logout();
        }
    }
)


export default api;

