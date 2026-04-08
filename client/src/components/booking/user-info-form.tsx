'use client';

import { useUpdateBooking } from '@/hooks/use-booking';
import { useAuth } from '@/stores/auth-store';
import { UpdateBookingType } from '@/types';
import { userSchemaType, userSchmea } from '@/validations/user-schema';
import { addToast, Button, Card, CardBody, CardHeader, Input } from '@heroui/react'
import { zodResolver } from '@hookform/resolvers/zod';
import { useParams } from 'next/navigation';

import { useForm } from 'react-hook-form';


const UserInfoForm = () => {
    const user = useAuth(s => s.user)

    const bookingId = useParams<{ id: string }>().id;
    const { register, handleSubmit, formState: { errors } } = useForm<userSchemaType>({
        resolver: zodResolver(userSchmea),
        defaultValues: {
            name: user?.name,
            email: user?.email,
            phone: user?.phone,
            city: user?.city,
            country: user?.country
        }
    })
    const { mutate, isPending } = useUpdateBooking()
    const onSubmit = async (data: userSchemaType) => {
        if (!bookingId || !data) {
            return addToast({
                title: "Booking Id or data is required.",
                color: "warning"
            })
        }


        const bookingData = {
            status: "PENDING",
            ...data
        }
        console.log(bookingData)
        mutate({ bookingId, booking: bookingData as UpdateBookingType })

    }


    return (
        <Card className='p-4'>
            <CardHeader>
                <div>
                    <h1 className="text-3xl font-semibold">Secure your booking</h1>
                    <p>Complete the form below to comfirm your hotel reservation.</p>
                </div>

            </CardHeader>
            <CardBody>


                <form className='flex flex-col gap-4 '
                    onSubmit={handleSubmit(onSubmit)}
                >
                    <Input
                        {...register('name')}
                        isInvalid={!!errors.name}
                        errorMessage={errors.name?.message}
                        type='text'
                        placeholder='Your name'
                        label="Name"
                        labelPlacement='outside'
                        radius='sm'
                        className='md:col-span-2'

                    />
                    <Input
                        {...register('email')}
                        isInvalid={!!errors.email}
                        errorMessage={errors.email?.message}
                        type='email'
                        placeholder='example@gmail.com'
                        label="Email"
                        labelPlacement='outside'
                        radius='sm'

                    />
                    <Input
                        {...register('phone')}
                        isInvalid={!!errors.phone}
                        errorMessage={errors.phone?.message}
                        type='text'
                        placeholder='+959xxxxxxx'
                        label="Phone"
                        labelPlacement='outside'
                        radius='sm'
                    />
                    <Input
                        {...register('city')}
                        isInvalid={!!errors.city}
                        errorMessage={errors.city?.message}
                        type='text'
                        placeholder='city'
                        label="City"
                        labelPlacement='outside'
                        radius='sm'

                    />
                    <Input
                        {...register('country')}
                        isInvalid={!!errors.country}
                        errorMessage={errors.country?.message}
                        type='text'
                        placeholder='country'
                        label="Country"
                        labelPlacement='outside'
                        radius='sm'

                    />
                    <Button
                        isLoading={isPending}
                        type='submit'
                        color='primary'
                        radius='sm'

                    >
                        NEXT STEP
                    </Button>
                </form>
            </CardBody>
        </Card>




    )
}

export default UserInfoForm