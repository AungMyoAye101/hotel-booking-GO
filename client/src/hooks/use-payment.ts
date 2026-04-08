import { confirmPaymentService, createPayment, getPaymentById, getReceiptService } from "@/service/payment-service"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"

import { useRouter } from "next/navigation"

export const useCreatePayment = () => {
    const qc = useQueryClient();
    const router = useRouter();
    return useMutation({
        mutationKey: ['create_payment'],
        mutationFn: createPayment,
        onSuccess: (data) => {
            qc.invalidateQueries({
                queryKey: ['booking_by_id', data.bookingId]
            })
            router.push(`/booking/${data.bookingId}/complete`, { scroll: true })
        },
        onError: (error) => {
            console.log(error)
        }
    })
}
export const useConfirmPayment = () => {

    const queryClient = useQueryClient()

    return useMutation({
        mutationKey: ['confirm_payment'],
        mutationFn: confirmPaymentService,
        onSuccess: (data) => {

            queryClient.invalidateQueries({
                queryKey: ['payment_id', data.id]
            })

        },
        onError: (error) => {
            console.log(error)
        }
    })
}

export const useGetPaymentById = (id: string) => {
    return useQuery({
        queryKey: ['payment_id', id],
        queryFn: () => getPaymentById(id),
        enabled: !!id
    })
}

export const useGetReceipts = (userId: string) => {
    return useQuery({
        queryKey: ["receipts"],
        queryFn: () => getReceiptService(userId),
        enabled: !!userId
    })
}