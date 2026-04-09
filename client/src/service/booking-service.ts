import api from "@/hooks/axios-api";
import { APIResponse, BookingInfoType, BookingType, UpdateBookingType } from "@/types";
import { createBookingType } from "@/validations/booking-schema";


export const createBooking = async (booking: createBookingType) => {


    const { data } = await api.post<APIResponse<BookingType>>(
        '/bookings', booking
    )


    return data.result;
}

type updateBookingParam = {
    bookingId: string, booking: UpdateBookingType
}
export const updateBooking = async ({ bookingId, booking }: updateBookingParam) => {
    const { data } = await api.put<APIResponse<BookingType>>(
        `/bookings/${bookingId}`, booking
    )


    return data.result
}

export const getBookingById = async (bookingId: string) => {
    const { data } = await api.get<APIResponse<BookingInfoType>>(`/bookings/${bookingId}`)
    console.log(data)
    return data.result
}

export const getBookingByUseridService = async (userId: string) => {
    const { data } = await api.get<APIResponse<{ booking: BookingInfoType[] }>>(`/bookings/user/${userId}`)

    return data.result.booking;
}
export const cancelBookingService = async (bookingId: string) => {
    const { data } = await api.put<APIResponse<{ booking: BookingInfoType }>>(`/bookings/cancel/${bookingId}`, { bookingId })
    return data.result;
}