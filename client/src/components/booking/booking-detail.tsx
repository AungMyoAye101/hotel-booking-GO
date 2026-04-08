import { BookingInfoType } from '@/types'
import { Card, CardBody, CardHeader } from '@heroui/react'
import { BedDouble, Check, CircleDollarSign, Map, MapPin } from 'lucide-react'
import Image from 'next/image'



type Props = {
    booking: BookingInfoType

}
const BookingDetail = ({ booking }: Props) => {
    console.log(booking, "ninini")

    return (
        <section className='w-full md:max-w-sm space-y-4'>

            {/* hotel details */}

            <Card >
                <CardBody>

                    <div className="relative rounded-lg aspect-video w-full">
                        <Image
                            src={booking?.hotel?.photo?.secure_url || "/hotel-card.webp"}
                            alt="hotel photo"
                            fill

                            className='rounded-lg'
                        />
                    </div>
                    <div className='space-y-1 py-2'>
                        <h1 className='head-1'>
                            {booking?.hotel?.name}


                        </h1>

                        <p className='flex gap-1 text-sm items-center '>

                            <MapPin size={16} />
                            {booking?.hotel?.adddress}
                            Address

                        </p>
                        <p className='flex flex-wrap items-center  gap-1'>
                            <Map size={16} />
                            <span className='font-semibold'>
                                {booking?.hotel?.city}
                            </span>
                            city

                        </p>
                    </div>

                </CardBody>
            </Card>

            {/* room detail  */}

            <Card className='border border-warning-400'>
                <CardHeader className='pb-0'>
                    <h1 className='font-semibold'>
                        Room Details
                    </h1>

                </CardHeader>
                <CardBody className='py-2 space-y-1'>
                    <h2 className='head-1'>{booking?.room?.name}</h2>
                    <p className='flex gap-1 items-center '>
                        <BedDouble size={16} />
                        Bed Type : {booking?.room?.bedTypes}
                    </p>
                    <div>
                        <span className='text-warning-500 text-lg font-semibold'>
                            $ {booking?.room?.price}
                        </span> / night
                    </div>
                </CardBody>
            </Card>

            {/* Booking info */}
            <Card className='border border-success-400'>
                <CardHeader className='pb-0'>
                    <h1 className='font-semibold'>
                        Your Booking Information
                    </h1>

                </CardHeader>
                <CardBody className='pt-2 pb-4 grid grid-cols-2 gap-2'>
                    <div className='bg-slate-200 p-2 rounded-lg flex-col gap-2'>
                        <p>Check In </p>
                        <p className='font-semibold'>{new Date(booking.checkIn).toDateString()}</p>
                    </div>
                    <div className='bg-slate-200 p-2 rounded-lg flex-col gap-2'>
                        <p>Check Out </p>
                        <p className='font-semibold'>{new Date(booking.checkOut).toDateString()}</p>
                    </div>
                    <div className='bg-slate-200 p-2 rounded-lg flex-col gap-2'>
                        <p>Quantity </p>
                        <p className='font-semibold'>{booking.quantity}</p>
                    </div>
                    <div className='bg-slate-200 p-2 rounded-lg flex-col gap-2 border border-warning-400'>
                        <p>Price</p>
                        <p className='font-semibold text-xl text-amber-500'>{booking.totalPrice} $</p>
                    </div>


                </CardBody>
            </Card>

            {/* Promotion */}
            <Card>
                <CardBody >
                    <div className="head-1 flex gap-2 items-center">
                        <CircleDollarSign size={26} />
                        Free Cancellation
                    </div>
                    <p className='text-sm py-2'>
                        Book now and pay later!
                        You can cancel your reservation free of charge up to 24 hours before check-in.
                        No hidden fees, no upfront payment required.
                    </p>
                    <p
                        className="flex gap-1 p-2 bg-green-100 rounded-sm text-sm">
                        <Check className="text-green-500" />
                        You're booking is protect by SSL encryption.
                    </p>

                </CardBody>
            </Card>

        </section >
    )
}

export default BookingDetail