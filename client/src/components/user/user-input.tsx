'use cleint';
import { useUpdateUser } from '@/hooks/use-user'
import { User } from '@/types/user-type'
import { userSchemaType, userSchmea } from '@/validations/user-schema'
import { addToast, Button, Input } from '@heroui/react'
import { zodResolver } from '@hookform/resolvers/zod'

import { useForm } from 'react-hook-form'

type props = {
    user: User,
    readonly?: boolean
}

const UserInfoInput = ({ user, readonly }: props) => {


    const { register, handleSubmit, formState: { errors } } = useForm<userSchemaType>({
        resolver: zodResolver(userSchmea),
        defaultValues: user
    })
    const { mutate, isPending } = useUpdateUser()
    const onSubmit = async (data: userSchemaType) => {
        mutate({ id: user.id, user: data as User }, {
            onSuccess: () => {
                addToast({
                    title: "Updated Successfully",
                    color: "success",
                })
            }
        })
    }
    return (
        <form
            onSubmit={handleSubmit(onSubmit)}
            className='grid grid-cols-1  gap-4 bg-white/60'>
            <Input
                {...register('name')}
                isInvalid={!!errors.name}
                errorMessage={errors.name?.message}
                type='text'
                placeholder='Your name'
                label="Name"
                labelPlacement='outside'
                radius='sm'
                readOnly={readonly}
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

                readOnly={readonly}
            />
            <Input
                {...register('phone')}
                isInvalid={!!errors.phone}
                errorMessage={errors.phone?.message}
                type='text'
                placeholder='phone'
                label="Phone"
                labelPlacement='outside'
                radius='sm'

                readOnly={readonly}
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

                readOnly={readonly}
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

                readOnly={readonly}
            />
            <Button
                isLoading={isPending}
                type='submit'
                disabled={readonly}
                color={readonly ? "default" : 'primary'}
                radius='sm'
                className='place-items-end'>
                Submit
            </Button>
        </form>
    )
}

export default UserInfoInput

