import * as z from 'zod'

export const userSchmea = z.object({
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
})

export type userSchemaType = z.infer<typeof userSchmea>;