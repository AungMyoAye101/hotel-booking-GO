import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import api from "./axios-api"
import { APIResponse } from "@/types"
import { User } from "@/types/user-type"
import { currentUser, updateUser } from "@/service/user-service"
import { useAuth } from "@/stores/auth-store"


export const useGetUserById = (id: string) => {
    return useQuery({
        queryKey: ['user_id', id],
        queryFn: async () => {
            try {
                const { data } = await api.get<APIResponse<User>>(`/users/${id}`)
                return data.result;
            } catch (error) {
                console.warn(error)
            }
        },
        enabled: !!id
    })
}
export const useGetMe = () => {
    return useQuery({
        queryKey: ['me',],
        queryFn: currentUser,
        retry: false,


    })
}

export const useUpdateUser = () => {
    const setUser = useAuth.getState().setUser;
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['update_user'],
        mutationFn: updateUser,
        onSuccess: (updatedUser) => {
            setUser(updatedUser!);
            queryClient.invalidateQueries({
                queryKey: ['user_id'],
                exact: false,
            })
        }


    })
}