'use client'

import FiveStars from '@/components/star'
import { useGetBookingById } from '@/hooks/use-booking'
import {
    Button,
    Card,
    CardBody,
    CardHeader,
    Skeleton
} from '@heroui/react'
import {
    BedDouble,
    Calendar,
    CalendarCheck,
    CheckCircle2,
    CreditCard,
    Mail,
    Settings,
    UserRound
} from 'lucide-react'
import Image from 'next/image'
import Link from 'next/link'
import { useParams } from 'next/navigation'


const Complete = () => {

    const bookingId = useParams().id;

    const {
        data: booking,
        isLoading,
        error,
        isError
    } = useGetBookingById(bookingId as string)

    if (isLoading) {
        return <Loading />
    }
    if (isError || !booking) {
        throw new Error(error?.message)
    }
    return (
        <div className="px-4 py-8 space-y-6 max-w-5xl mx-auto">

            <div className='flex flex-col  justify-center items-center'>
                <Button color='success' radius='full' size='lg' className='mx-auto' isIconOnly><CheckCircle2 className='text-white ' /></Button>
                {/* Heading */}
                <h1 className="text-3xl font-semibold text-success">
                    Booking Confirmed!
                </h1>
                <p>Your reservation is complete</p>
            </div>

            <div className='bg-primary-50 rounded flex  gap-4 items-center p-2'>
                <Mail className='w-10 h-auto' />
                <span>


                    <span className='font-semibold'>Thanks you, {booking.name}!</span>{' '}
                    your booking is now confirmed.
                    A confirmation email has been sent to {' '}
                    <span className='font-semibold'>{booking.email}</span>
                    .
                </span>
            </div>

            {/* Booking details */}
            <Card className=' border border-slate-300' >
                <CardHeader className='text-lg font-semibold'>
                    Booking Details
                </CardHeader>
                <CardBody>
                    <div className='flex flex-col md:flex-row gap-4'>
                        <div>
                            <Image
                                src={booking.hotel?.photo?.secure_url || '/hotel-card.webp'}
                                alt={booking?.hotel?.name + "photo "}
                                width={300}
                                height={260}
                                className='object-cover aspect-video rounded-lg w-full md:max-w-sm'
                            />
                        </div>

                        <div className='flex-1 flex flex-col gap-2'>
                            <div>
                                <h1 className='text-2xl font-bold'>
                                    {booking.hotel?.name}
                                </h1>
                                <p className='text-slate-700'>
                                    {booking.hotel?.city} City
                                </p>
                                <FiveStars count={booking.hotel?.star} />
                            </div>
                            <div className='w-full h-0.5 bg-slate-200' />

                            {/* Checkin and out */}
                            <div className='flex gap-4 flex-wrap'>
                                <div className='flex items-center gap-1'>
                                    <Calendar />
                                    <span className='font-semibold text-sm'>
                                        Check in :  {new Date(booking.checkIn).toDateString()}
                                    </span>
                                </div>
                                <div className='flex items-center gap-1'>
                                    <Calendar />
                                    <span className='font-semibold text-sm'>
                                        Check Out : {new Date(booking.checkOut).toDateString()}
                                    </span>
                                </div>
                            </div>
                            <div className='w-full h-0.5 bg-slate-200' />

                            {/* guest*/}
                            <div className='flex items-center gap-1'>
                                <UserRound />
                                <span className='font-semibold text-sm'>
                                    Guest : {3} Adults
                                </span>
                            </div>
                            <div className='w-full h-0.5 bg-slate-200' />
                            {/* room */}
                            <div className='flex items-center gap-1'>
                                <BedDouble />
                                <span className='font-semibold text-sm'>
                                    Room : {booking.room?.name}
                                </span>
                            </div>

                            <div className='w-full h-0.5 bg-slate-200' />
                            {/* Payment */}

                            <div className='flex items-center gap-1'>
                                <CreditCard />
                                <span className='font-semibold '>
                                    Total : ${booking.totalPrice}
                                </span>
                            </div>
                        </div>


                    </div>
                </CardBody>



            </Card>
            {/* 
            Payment info */}
            <Card radius='sm' shadow='sm' className='bg-success-50'>
                <CardBody>
                    <div className='flex gap-4 items-start '>
                        <CheckCircle2 className='size-12 text-success' />
                        <div>
                            <h2 className='font-semibold text-lg text-success'>Payment successful!</h2>
                            <p className='text-sm'>
                                Your payment has been comfrimed . We look forward to your stay.
                            </p>
                        </div>
                    </div>
                </CardBody>
            </Card>

            {/* actions sections */}
            <div className='flex flex-col gap-4 items-center py-6'>
                <h3 className='text-2xl text-center font-bold'>What's Next?</h3>
                <div className='flex flex-wrap gap-6'>
                    <Card className='min-w-72 border border-slate-300'>
                        <CardBody>
                            <div className='flex flex-col items-center gap-2'>


                                <Settings className='size-16 text-center text-gray-500' />
                                <h4 className='font-semibold'>
                                    Manage Booking
                                </h4>
                                <p>
                                    Make changes or request extra
                                </p>
                                <Button
                                    as={Link}
                                    href={`/user/booking`}
                                    variant='solid'
                                    fullWidth
                                    color='primary'
                                    radius='sm' >
                                    Modify Booking
                                </Button>
                            </div>
                        </CardBody>
                    </Card>
                    <Card className='min-w-72 border border-slate-300'>
                        <CardBody>
                            <div className='flex flex-col items-center gap-2'>


                                <CalendarCheck className='size-16 text-center text-gray-500' />
                                <h4 className='font-semibold'>
                                    View Booking
                                </h4>
                                <p>
                                    Check Booking information
                                </p>
                                <Button
                                    as={Link}
                                    href={`/user/booking/${booking._id}`}
                                    variant='solid'
                                    fullWidth
                                    color='primary'
                                    radius='sm' >
                                    View Booking
                                </Button>
                            </div>
                        </CardBody>
                    </Card>
                    <Card className='min-w-72 border border-slate-300'>
                        <CardBody>
                            <div className='flex flex-col items-center gap-2'>


                                <CreditCard className='size-16 text-center text-gray-500' />
                                <h4 className='font-semibold'>
                                    Payment
                                </h4>
                                <p>
                                    Check Payment information
                                </p>
                                <Button
                                    as={Link}
                                    href={`/user/${booking.user._id}/payment`}
                                    variant='solid'
                                    fullWidth
                                    color='primary'
                                    radius='sm'>
                                    View Payment
                                </Button>
                            </div>
                        </CardBody>
                    </Card>

                </div>

                <Link href={'/'} className='text-sm text-violet-500'>Back to home page</Link>
                <h5>
                    <span className='font-semibold'>
                        Need Help?
                    </span>  Our support team is 24/7 .
                    <span className='font-semibold'>

                        Contact Us
                    </span>
                </h5>
            </div>

        </div>
    )
}


const Loading = () => {
    return (
        <section className="px-4 py-8 space-y-6 max-w-5xl mx-auto">
            <div className='flex flex-col  justify-center items-center gap-2'>
                <Skeleton className='w-20 h-20 rounded-full' />
                <Skeleton className='w-60 h-8 rounded-md' />
                <Skeleton className='w-50 h-8 rounded-md' />
            </div>
            <Skeleton className='w-full h-12 rounded-md' />
            <Skeleton className='w-full h-100 rounded-md' />
            <Skeleton className='w-full h-16 rounded-md' />
        </section>
    )
}

export default Complete


