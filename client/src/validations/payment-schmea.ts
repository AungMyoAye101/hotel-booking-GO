
import * as z from "zod";
export const createPaymentSchema = z.object({
    booking_id: z.string("Booking id is required."),
    user_id: z.string("user id is required."),
    payment_method: z.enum(['CARD', 'MOBILE_BANKING', 'BANK'], "Invalid payment method."),
    amount: z.number().positive(),

})

export const paymentSchema = z.discriminatedUnion("method", [
    z.object({
        method: z.literal("CARD"),
        cardNumber: z.string().regex(/^\d{16}$/),
        expiry: z.string().regex(/^(0[1-9]|1[0-2])\/\d{2}$/),
        cvc: z.string().regex(/^\d{3,4}$/),
    }),

    z.object({
        method: z.literal("MOBILE_BANKING"),
        provider: z.string().min(1),
    }),

    z.object({
        method: z.literal("BANK"),
        name: z.string().min(1),
        cardNumber: z.string().regex(/^\d{16}$/),

    })
]);

export type PaymentInput = z.infer<typeof paymentSchema>;
export type CreatePaymentType = z.infer<typeof createPaymentSchema>
