import { createReviewService, getReviewsByHotelIdService } from "@/service/review-service"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"

export const useCreateReview = () => {
    const qc = useQueryClient()
    return useMutation({
        mutationKey: ['create_review'],
        mutationFn: createReviewService,
        onSuccess: (data) => {
            qc.invalidateQueries({
                queryKey: ['review_by_hotelId'],
                exact: false
            })
        }
    })
}

export const useGetReviewByHotelId = (hotelId: string) => {
    return useQuery({
        queryKey: ['review_by_hotelId', hotelId],
        queryFn: () => getReviewsByHotelIdService(hotelId),
        enabled: !!hotelId
    })
}