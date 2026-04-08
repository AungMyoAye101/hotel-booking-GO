
'use client';
import { useAuth } from '@/stores/auth-store';
import { APIResponse } from '@/types';
import { User } from '@/types/user-type';
import { loginSchema, type loginType } from '@/validations/auth-schema';
import { addToast, Button, Card, CardBody, CardHeader, Input } from '@heroui/react'
import { zodResolver } from '@hookform/resolvers/zod';
import axios from 'axios';
import { Eye, EyeOff, Lock, Mail } from 'lucide-react';
import Image from 'next/image'
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { useForm } from 'react-hook-form';

const Login = () => {
    const setUser = useAuth(s => s.setUser)
    const [isLoading, setIsLoading] = useState(false)
    const [isVisible, setIsVisible] = useState<boolean>(false)

    const router = useRouter()
    const toggleVisible = () => setIsVisible(pre => !pre)
    const { register, handleSubmit, formState: { errors }, reset } = useForm<loginType>({
        resolver: zodResolver(loginSchema)
    })

    const onSubmit = async (fields: loginType) => {
        setIsLoading(true)
        try {
            const { data } = await axios.post('/api/auth/login', fields, { withCredentials: true })

            setUser(data)
            addToast({
                title: data.message || "login success.",
                color: 'success'
            })
            reset()
            router.push('/')

        } catch (error: any) {
            console.warn(error)
            addToast({
                title: error?.response?.data?.message || "Failed to login.",
                color: 'danger'
            })
            throw new Error("login failed.")
        } finally {
            setIsLoading(false)
        }

    }
    return (
        <section className='h-screen flex justify-center md:justify-end items-center px-4 '>
            <Image
                src={'/hotel-bg.webp'}
                alt='hotel photo'
                width={1400}
                height={1240}
                className='object-cover absolute inset-0 -z-10 h-screen w-full '

            />
            <Card className=' px-4 py-6 min-w-xs w-full max-w-100 bg-background'>
                <CardHeader className='flex flex-col gap-1'>
                    <h1 className='text-2xl font-semibold text-center'>Login</h1>
                    <p>Please login your account.</p>
                </CardHeader>
                <CardBody>
                    <form
                        className="w-full max-w-xs flex flex-col gap-4"
                        onSubmit={handleSubmit(onSubmit)}
                    >
                        <Input
                            {
                            ...register('email')
                            }
                            isInvalid={!!errors.email}
                            errorMessage={errors.email?.message}
                            label="Email"
                            labelPlacement="outside"
                            placeholder="Enter your email"
                            type="email"
                            startContent={<Mail size={20} className='text-default-400' />}
                        />

                        <Input
                            {
                            ...register('password')
                            }

                            isInvalid={!!errors.password}
                            errorMessage={errors.password?.message}
                            label="Password"
                            labelPlacement="outside"
                            placeholder="Enter your password"
                            type={isVisible ? 'text' : "password"}
                            startContent={<Lock size={20} className='text-default-400' />}
                            endContent={

                                <Button

                                    isIconOnly size='sm'

                                    onPress={() => toggleVisible()}>
                                    {isVisible ? <Eye /> : <EyeOff />}</Button>

                            }
                        />

                        <Button isLoading={isLoading} color="primary" type="submit">
                            Login
                        </Button>


                    </form>
                    <p className='text-sm my-2'>Don't  have an account ? <Link href={'/signup'} className='text-violet-800'>Signup</Link></p>
                </CardBody>
            </Card>
        </section>
    )
}

export default Login

