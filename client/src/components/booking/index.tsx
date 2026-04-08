"use client";

import Complete from "./booking-complete";
import BookingShell from "./booking-shell"
import PaymentDetailsForm from "./payment-details";
import UserInfoForm from "./user-info-form"

export const BookingCreateProvider = () => {
    return (
        <BookingShell>
            {booking => <UserInfoForm />}
        </BookingShell>
    )
}
export const BookingPaymentProvider = () => {
    return (
        <BookingShell>
            {booking => <PaymentDetailsForm booking={booking} />}
        </BookingShell>
    )
}
export const BookingCompleteProvider = () => {
    return (
        <BookingShell showSideBar={false}>
            {booking => <Complete />}
        </BookingShell>
    )
}