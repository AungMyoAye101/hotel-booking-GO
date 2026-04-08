'use client'

import { useCreateReview } from '@/hooks/use-review'
import { useAuth } from '@/stores/auth-store'
import { CreateReviewType } from '@/types'
import { addToast, Button, Card, CardBody, Textarea, Tooltip } from '@heroui/react'
import { FormEvent, useState } from 'react'
import { Star } from 'lucide-react'


type Props = {
    hotelId: string
}
const HotelReview = ({ hotelId }: Props) => {
    const userId = useAuth(s => s.user?.id);

    const [reviewData, setReviewData] = useState({
        rating: 1,
        description: ''
    })

    const { mutate, isPending } = useCreateReview();

    const onSubmit = (e: FormEvent) => {
        e.preventDefault();
        if (!userId) {
            return addToast({
                title: "Please login first.",
                color: "warning"
            })

        }
        if (!hotelId) {
            console.error("Hotel id is required.")
            return addToast({
                title: "Something went worng",
                color: "danger"
            })
        }
        const data = {
            hotelId,
            userId,
            rating: Number(reviewData.rating),
            review: reviewData.description
        } as CreateReviewType;

        mutate(data, {
            onSuccess: () => {
                addToast({
                    title: "Review post successful.",
                    color: "success"
                })
            },
            onError: (error) => {
                addToast({
                    title: error?.message || "Failed to post review.",
                    color: "danger"
                })
            }

        })
    }



    return (
        <Card className='min-w-75 max-w-sm '>
            <CardBody>
                <form
                    onSubmit={onSubmit}
                    className=" space-y-3  ">
                    <h1 className='font-semibold text-lg capitalize'>Please leave your review</h1>

                    <div className='flex items-center gap-0.5'>
                        {
                            Array.from({ length: 5 }).map((_, i) => (

                                <Star
                                    key={i}
                                    fill={i + 1 <= reviewData.rating ? "oklch(79.5% 0.184 86.047)" : "white"}
                                    size={20}
                                    className="text-yellow-500"
                                    onClick={() => setReviewData(pre => ({ ...pre, rating: i + 1 }))}
                                />

                            ))
                        }
                    </div>
                    <Textarea
                        isRequired
                        validate={(value) => {
                            if (value.trim().length < 3) return "Description is too short."
                        }}
                        placeholder='Your review'
                        minRows={4}
                        value={reviewData.description}
                        onValueChange={(value) => setReviewData(pre => ({ ...pre, description: value }))}
                    />
                    <Tooltip
                        color={userId ? 'default' : 'danger'}
                        content={
                            userId
                                ? 'Review'
                                : 'Please login to reserve'
                        }
                    >


                        <Button type='submit'
                            isLoading={isPending}
                            fullWidth variant='solid'
                            color={userId ? 'primary' : 'default'}
                        >
                            Submit Review
                        </Button>
                    </Tooltip>
                </form>
            </CardBody>
        </Card>

    )
}



export default HotelReview