import api from "@/hooks/axios-api";
import { APIResponse } from "@/types";
import { User } from "@/types/user-type";

export const currentUser = async () => {

    const { data } = await api.get<APIResponse<User>>(`/auth/me`)

    return data.result;

}

type updateUserType = {
    id: string,
    user: User
}
export const updateUser = async ({ id, user }: updateUserType) => {
    try {
        const { data } = await api.put<APIResponse<User>>(
            `/users/${id}`, user
        )
        if (!data.success) {
            throw new Error("Faied to update user")
        }

        return data.result;
    } catch (error) {
        console.warn(error)
    }
}