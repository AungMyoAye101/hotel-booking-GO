import { getRoomById } from "@/service/room-service"
import { useQuery } from "@tanstack/react-query"

export const useGetRoomById = (roomId: string) => {
    return useQuery({
        queryKey: ['room_by_id', roomId],
        queryFn: () => getRoomById(roomId),
        enabled: !!roomId && roomId.length === 24
    })
}