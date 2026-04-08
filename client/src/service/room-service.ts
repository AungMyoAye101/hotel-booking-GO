import api from "@/hooks/axios-api"
import { APIResponse, BookingRoomType } from "@/types"

export const getRoomById = async (roomId: string) => {
    const { data } = await api.get<APIResponse<{ room: BookingRoomType }>>(
        `/room/${roomId}`
    )
    if (!data.success) {
        throw new Error(data.message || "Failed to room")
    }
    return data.result.room;
}