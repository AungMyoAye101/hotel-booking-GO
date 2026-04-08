import * as z from "zod";

export const registerSchema = z.object({
    name: z.string("Name is required.").min(1, "Name must be at least 1 character long."),
    email: z.email("Invalid email."),
    password:
        z.string("Password is required.")
            .min(8, "Password must be at least 8 charaters long.")
});
export const loginSchema = z.object({
    email: z.email("Invalid email."),
    password:
        z.string("Password is required.")
            .min(8, "Password must be at least 8 charaters long.")
});

export type registerType = z.infer<typeof registerSchema>
export type loginType = z.infer<typeof loginSchema>