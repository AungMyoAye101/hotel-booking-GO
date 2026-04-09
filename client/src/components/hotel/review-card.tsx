import { Avatar, Card, CardBody, Skeleton } from '@heroui/react'
import FiveStars from '../star'
import { useGetReviewByHotelId } from '@/hooks/use-review'
type Props = {
    hotelId: string
}
const ReviewCard = ({ hotelId }: Props) => {
    const { data: reviews, isLoading, isError, error } = useGetReviewByHotelId(hotelId)
    if (isLoading) {
        return <section className='space-y-4'>

            <h1 className='text-2xl font-semibold'>Guest Reviews</h1>
            <div className='flex flex-wrap items-center justify-center gap-4'>


                {
                    Array(4).fill(null).map((_, i) => (
                        <Card key={i} shadow='sm' className=' min-w-72  max-w-xs border border-slate-300 '>

                            <CardBody >
                                <div className='flex items-start gap-2'>
                                    <Skeleton className='w-14 h-14 rounded-full ' />
                                    <div className='space-y-2'>

                                        <Skeleton className='w-52 h-6 rounded-md' />
                                        <Skeleton className='w-52 h-6 rounded-md' />
                                        <Skeleton className='w-52 h-16 rounded-md' />
                                    </div>
                                </div>
                            </CardBody>
                        </Card>
                    ))
                }
            </div>
        </section>
    }

    if (isError || reviews?.length === 0) {
        console.error(error?.message)
        throw new Error(error?.message);
    }

    return (
        <section className='space-y-4 py-4'>


            <h1 className='text-2xl font-semibold'>Guest Reviews</h1>
            <div className='grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-4'>

                {
                    reviews?.map(value => (
                        <Card key={value.id} shadow='sm' className=' w-fll border border-slate-300 '>

                            <CardBody >
                                <div className='flex items-start gap-2'>
                                    <Avatar
                                        src='/user.jpg'
                                        alt='user photo'
                                        size='md'
                                        isBordered
                                        className='w-8 h-8'
                                    />
                                    <div className='space-y-2 flex-1'>
                                        <h1 className='text-lg font-semibold'>{value.user.name}</h1>
                                        <FiveStars count={value.rating} />
                                        <p className='text-sm'>Reviewed : {new Date(value.createdAt).toDateString()}</p>
                                        <p className='text-sm line-clamp-2'>
                                            {value.review}
                                        </p>
                                    </div>
                                </div>
                            </CardBody>
                        </Card>
                    ))
                }


            </div>

        </section>
    )
}

export default ReviewCard