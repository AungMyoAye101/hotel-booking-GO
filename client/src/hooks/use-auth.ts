import { useMutation } from "@tanstack/react-query"
import { useAuth } from "@/stores/auth-store"
import { APIResponse } from "@/types"
import { User } from "@/types/user-type"
import { addToast } from "@heroui/react"
import axios from "axios"

export const useLogout = () => {
    const reset = useAuth.getState().logout
    return useMutation({
        mutationKey: ['logout'],
        mutationFn: async () => {
            const { data } = await axios.post<APIResponse<{ user: User }>>('/api/user/logout');
            console.log(data)
            if (!data.success) {
                addToast({
                    title: data.message || "Failed to logout",
                    color: "danger"
                })
                return data;
            }
            addToast({
                title: data.message || "User logout",
                color: 'warning'
            })
            reset();
            return data
        }
    })
}