"use client";
import { ReactNode } from "react"
import BookingDetail from "./booking-detail"
import Indicator from "./indicator"
import { useParams } from "next/navigation";
import { useGetBookingById } from "@/hooks/use-booking";
import { BookingInfoType } from "@/types";
import { Skeleton } from "@heroui/skeleton";


type Prop = {
  showSideBar?: boolean,
  children: (booking: BookingInfoType) => ReactNode
}


const BookingShell = ({
  showSideBar = true,
  children }: Prop) => {
  const bookingId = useParams().id;

  const { data: booking, isLoading, isError, error } = useGetBookingById(bookingId as string)

  if (isLoading) {
    return <section className="flex gap-4 ">
      <div className="flex flex-col gap-4 ">
        <Skeleton className="w-sm md:max-w-md h-60 rounded-lg" />
        <Skeleton className="w-sm md:max-w-md h-60 rounded-lg" />
        <Skeleton className="w-sm md:max-w-md h-60 rounded-lg" />
        <Skeleton className="w-sm md:max-w-md h-60 rounded-lg" />
      </div>
      <Skeleton className="flex-1 h-100 rounded-lg" />
    </section>
  }
  if (isError || !booking) {
    console.warn(error)
    throw new Error("Failed to load booking")
  }
  return (
    <>
      <Indicator status={booking?.status} />
      <main className="flex flex-col md:flex-row gap-6">

        {

          showSideBar && booking.status !== "CONFIRMED" && <BookingDetail booking={booking} />
        }


        <div className="flex-1">
          {
            children(booking)
          }
        </div>
      </main>
    </>
  )
}

export default BookingShell