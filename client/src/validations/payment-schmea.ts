
import * as z from "zod";
export const createPaymentSchema = z.object({
    bookingId: z.string("Booking id is required."),
    userId: z.string("user id is required."),
    paymentMethod: z.enum(['CARD', 'MOBILE_BANKING', 'BANK'], "Invalid payment method."),
    amount: z.number().positive(),
    payNow: z.boolean(),
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
