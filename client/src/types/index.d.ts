import { hotelType, photoType } from "./hotel-types"

export interface APIResponse<T> {
    success: boolean,
    status: number,
    message: string,
    result: T
}

export type token = string

export type AuthResType = {
    token?: string,
    user: User
}

export type QueryType = {
    page?: number,
    limit?: number,
    search?: string,
    sort?: 'asc' | 'desc'
}

export type AvaliableRoomsType = {
    checkIn?: string,
    checkOut?: string,
    guest?: number,
    page?: number,
    limit?: number,
}
export type Bed_Types = "king" | "queen" | "full" | "twin" | "single";

export type RoomType = {
    id: string,
    name: string,
    maxPeople: number,
    price: number,
    hotelId: string,
    total_rooms: number,
    photo?: photoType,
    bed_types: Bed_Types,
    available_rooms?: number

}

export interface BookingRoomType extends RoomType {
    hotelId: Partial<hotelType>,
}
export interface TokenPayload {
    id: string,
    email: string,

}

export type MetaType = {
    page: number,
    limit: number,
    total: number;
    totalPages: number;
    hasPrev: boolean;
    hasNext: boolean
}

export type BookingStatus = "DRAFT" | "PENDING" | "CONFIRMED" | "STAYED" | "CANCELLED" | "EXPIRED";

export type BookingType = {
    id: string,
    userId: string,
    roomId: string,
    hotelId: string,
    totalPrice: number,
    quantity: number,
    status: BookingStatus,
    checkIn: Date,
    checkOut: Date,

}
export type UpdateBookingType = {

    name: string,
    email: string,
    city: string,
    country: string,
    phone: string,
    status: BookingStatus
}
export type PaymentMethodType = "MOBILE_BANKING" | "CARD" | "BANK"
export type PaymentCreateType = {
    id: string,
    bookingId: string,
    userId: string,
    paymentMethod: PaymentMethodType
    status: "PENDING" | "PAID" | "FAILED",
    amount: number,
    paidAt: Date,

}

export type createPaymentType = {
    bookingId: string,
    userId: string,
    paymentMethod: PaymentMethodType,
    amount: number,
    payNow: boolean
}

export type BookingInfoType = {
    id: string,
    name?: string,
    email?: string,
    city?: string,
    country?: string,
    phone?: string,
    check_in: Date,
    check_out: Date
    hotel: {
        address: string,
        city: string,
        name: string,
        star: number,
        rating: number,
        photo: photoType
    }
    quantity: number,
    room: Partial<RoomType>,
    status: BookingStatus
    total_price: number,
    guest: number,
    user: {
        id: string,
        name?: string,
    }

}
export type PaymentStatusType = "PAID" | "PENDING" | "FAILED";

export type PaymentType = {
    id: string,

    bookingId: {
        _id: string,
        name: string,
        hotelId: {
            id: string,
            name: string,
            city: string
        },
        checkIn: Date,
        checkOut: Date
    },
    userId: { id: string, name: string },
    status: PaymentStatusType,
    paymentMethod: PaymentMethodType,
    amount: number,
    paidAt: Date

}

export type ReceiptType = {
    id: string,
    receiptNo: string,
    user_id: string,
    payment_id: string,
    booking_id: string,
    status: PaymentStatusType,
    paymentMethod: PaymentMethodType,
    amount: number,
    paidAt: Date,
}

export type ReviewType = {
    id: string,
    user: {
        id: string,
        name: string,
    },

    review: string,
    rating: number,
    createdAt: Date
}
export type CreateReviewType = {
    userId: string,
    hotelId: string,
    review: string,
    rating: number,

}
