import api from "@/hooks/axios-api";
import { APIResponse, PaymentType, ReceiptType } from "@/types";
import { CreatePaymentType } from "@/validations/payment-schmea";


export const createPayment = async (payment: CreatePaymentType) => {

    const { data } = await api.post<APIResponse<{ payment: CreatePaymentType }>>(
        '/payment/create', payment
    )
    if (!data.success) {
        throw new Error("Failed to create payment.")
    }

    return data.result.payment;
}

type confirmPaymentType = {
    userId: string;
    paymentId: string;
    bookingId: string;
}

export const confirmPaymentService = async (payment: confirmPaymentType) => {
    const { data } = await api.put<APIResponse<any>>(
        '/payment/update', payment
    )

    if (!data.success) {
        throw new Error("Failed to confirm payment.")
    }

    return data.result;
}

export const getPaymentById = async (id: string) => {
    const { data } = await api.get<APIResponse<{ payment: PaymentType }>>(
        `/payment/${id}`
    )


    return data.result.payment;
}


export const getReceiptService = async (userId: string) => {
    const { data } = await api.get<APIResponse<{ receipts: ReceiptType[] }>>(`/receipt/${userId}`)

    return data.result.receipts;

}