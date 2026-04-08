'use client'
import { ArrowLeft, MapPin } from "lucide-react"
import Image from "next/image"
import FiveStars from "../star"
import { Button, Card, CardBody, Skeleton } from "@heroui/react"
import { useCancelBooking, useGetBookingById } from "@/hooks/use-booking"
import { useParams } from "next/navigation"
import { PDFDownloadLink } from "@react-pdf/renderer"
import BookingPDF from "../ui/booking-pdf"
import Link from "next/link"

const BookingInfo = () => {
    const { bookingId } = useParams()
    const { data: booking, isLoading, isError, error } = useGetBookingById(bookingId as string)
    if (isLoading) {
        return <section className="flex justify-center mt-14 w-full min-h-screen">
            <div className="w-full max-w-sm space-y-4">

                <Skeleton className="w-full h-96 max-w-md rounded-lg " />
                <div className="flex gap-4">
                    <Skeleton className="h-12 w-full rounded" />
                    <Skeleton className="h-12 w-full rounded" />
                </div>
            </div>
        </section>
    }

    if (isError || !booking) {
        console.warn(error?.message);
        throw new Error("Failed to get booking data.")
    }


    return (
        <>
            <Button
                as={Link}
                href="/user/booking"
                isIconOnly
                radius="full"
                className="absolute left-8"

            ><ArrowLeft /></Button>
            <section className="space-y-4 w-full max-w-md">

                <Card className=" h-fit">
                    <CardBody>
                        <div className='  overflow-hidden'>
                            <div
                                className="relative w-full aspect-video rounded-md">
                                <Image
                                    src={booking.hotel?.photo?.secure_url || "/hotel-card.webp"}
                                    alt={booking.hotel?.name + "photo "}
                                    fill

                                    className="aspect-video rounded-md"
                                />
                                <div
                                    className="absolute left-0 bottom-0 p-2 text-white font-semibold text-shadow bg-black/10">
                                    <h1 className="head-1">
                                        {booking.hotel?.name}
                                    </h1>
                                    <p className="flex items-center gap-1">
                                        <MapPin size={18} />  {booking.hotel?.adddress}
                                    </p>
                                    <FiveStars count={booking.hotel?.star} />
                                </div>
                            </div>
                            <div className="p-4 space-y-2">
                                <div className="grid grid-cols-2 gap-4  place-content-between">
                                    <span>Name </span>
                                    <span className="font-semibold text-end">
                                        {booking.user.name}
                                    </span>
                                    <span >Check In </span>
                                    <span className="font-semibold  text-end">
                                        {new Date(booking.checkIn).toLocaleDateString()}
                                    </span>

                                    <span>Check Out </span>
                                    <span className="font-semibold text-end">
                                        {new Date(booking.checkOut).toLocaleDateString()}
                                    </span>
                                    <span>Guests </span>
                                    <span className="font-semibold text-end">
                                        {booking.guest} Adult
                                    </span>
                                    <span>Status </span>
                                    <span className="font-semibold text-sm text-end">
                                        {booking.status}
                                    </span>

                                    <span>Room </span>
                                    <span className="font-semibold text-sm text-end">
                                        {booking.room?.name}
                                    </span>
                                    <span>Quantity </span>
                                    <span className="font-semibold text-sm text-end">
                                        {booking.quantity} Room
                                    </span>

                                </div>
                                <div
                                    className="flex justify-between gap-4 text-lg  py-2 border-t border-slate-700">
                                    <span >Total price</span>
                                    <span className="font-semibold text-end text-2xl text-amber-600">
                                        $ {booking.totalPrice}
                                    </span>
                                </div>
                            </div>
                        </div>
                    </CardBody>
                </Card>
                <div className="flex gap-4">
                    <CancelBooking bookingId={booking._id} />
                    <Button
                        as={PDFDownloadLink}
                        document={<BookingPDF booking={booking} />}
                        fileName={booking.name + "booking.pdf"}
                        variant="solid"
                        radius="sm"
                        color="primary"
                        fullWidth
                    >

                        Download PDF
                    </Button>

                </div>
            </section>
        </>
    )
}

export default BookingInfo


const CancelBooking = ({ bookingId }: { bookingId: string }) => {
    const { mutate, isPending } = useCancelBooking()

    const handleCancel = () => {
        if (!bookingId) {
            throw new Error("Booking Id is required")
        }
        mutate(bookingId)
    }
    return <Button
        isLoading={isPending}
        variant='bordered'
        color='danger'
        radius="sm"
        fullWidth
        onPress={() => handleCancel()}
    >
        Cancel Booking
    </Button>
}