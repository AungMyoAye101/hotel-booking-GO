'use client';

import { getAllHotels, getAvaliableRooms, getHotelDetails } from "@/service/hotel-service"
import { AvaliableRoomsType } from "@/types";
import { useQuery } from "@tanstack/react-query"

export const useHotelDetail = (hotelId: string) => {
    return useQuery({
        queryKey: ['hotel_details', hotelId],
        queryFn: () => getHotelDetails(hotelId),
        enabled: !!hotelId
    })
}
export const useGetAvaliableRoom = (hotel_id: string, param: AvaliableRoomsType) => {
    return useQuery({
        queryKey: ['rooms', hotel_id, param],
        queryFn: () => getAvaliableRooms(hotel_id, param),
        enabled: !!hotel_id
    })
}

//===========get all hotels 

export const useGetAllHotels = () => {
    return useQuery({
        queryKey: ['hotels'],
        queryFn: () => getAllHotels()
    })
}