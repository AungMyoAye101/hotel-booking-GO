import api from "@/hooks/axios-api";
import { APIResponse, CreateReviewType, ReviewType } from "@/types";

export const createReviewService = async (review: CreateReviewType) => {
    const { data } = await api.post<APIResponse<{ review: ReviewType }>>('/review/create', review)

    return data.result;
}
export const updateReviewService = async (review: ReviewType) => {
    const { data } = await api.put<APIResponse<{ review: ReviewType }>>('/review/create', review)

    return data.result;
}
export const getReviewsByHotelIdService = async (hotelId: string) => {
    const { data } = await api.get<APIResponse<ReviewType[]>>('/reviews/hotel/' + hotelId, {
        params: { limit: 4 }
    })

    return data.result;
}