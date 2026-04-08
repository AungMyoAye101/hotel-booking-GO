import * as z from "zod"

export const bookingSchema = z.object({
    user_id: z
        .string("user id is required."),
    room_id: z
        .string("roomId is required."),
    hotel_id: z
        .string("HotelId is required."),
    name: z
        .string()
        .min(3, "Name must be contain 3 characters.")
        .optional(),
    email: z
        .email("Inavlid email")
        .transform(v => v.toLocaleLowerCase())
        .optional(),
    city: z
        .string()
        .min(1, "City is required.")
        .optional(),
    country: z
        .string()
        .min(1, "Country is required.")
        .optional(),
    phone: z
        .string()
        .regex(/^\+\d{1,4}\d{6,12}$/, "Invalid phone number")
        .optional(),
    total_price: z
        .number("Total price is required")
        .positive(),
    quantity: z.
        number("Quantity is required,")
        .positive(),
    guest: z
        .number("Guest is required")
        .positive(),
    check_in: z
        .iso.datetime("CheckIn must be valid date string.")
        .optional(),
    check_out: z
        .iso.datetime("CheckOut must be valid date string.")
        .optional(),
}).refine((data) => {
    if (data.check_in && data.check_out) {
        if (data.check_out > data.check_in) {
            return {
                message: "Check out must be at least one day after check-in",
                path: ['checkOut']
            }
        }
    }

})

export type createBookingType = z.infer<typeof bookingSchema>