import api from "@/hooks/axios-api"
import { APIResponse, AvaliableRoomsType, RoomType } from "@/types"
import { hotelType } from "@/types/hotel-types"

export const getHotelDetails = async (hotelId: string) => {

    const { data } = await api.get<APIResponse<{ hotel: hotelType }>>(`/hotel/${hotelId}`)
    if (!data.success) {
        throw new Error("Faild to get hotel details")
    }

    return data.result.hotel
}

export const getAvaliableRooms = async (hotelId: string, params: AvaliableRoomsType) => {
    console.log(hotelId)
    const { data } = await api.get<APIResponse<RoomType[]>>(`/rooms/hotel/${hotelId}`, { params })

    if (!data.success) {
        throw new Error(data.message || "Failed to get room")
    }

    return data.result;
}

//=============get all hotel =========

export const getAllHotels = async () => {
    const { data } = await api.get<APIResponse<hotelType[]>>('/hotels');
    return data.result;
}