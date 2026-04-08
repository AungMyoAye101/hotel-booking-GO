'use client';

import { useCreatePayment } from '@/hooks/use-payment';
import { BookingInfoType, PaymentMethodType } from '@/types';
import { addToast, Button, Card, CardBody, CardHeader, cn, Radio, RadioGroup } from '@heroui/react'
import Image from 'next/image';
import { FormEvent, useState } from 'react';
import { createPaymentSchema, CreatePaymentType } from '@/validations/payment-schmea';
import { BankingInputs, CardInputs, MobileBankingInputs } from './payment-inputs';


const paymentMethodOptions = [
    {
        method: "CARD",
        image: "/card.png",
        text: "Card / debait card",
        slug: "pay with card"
    },
    {
        method: "MOBILE_BANKING",
        image: "/mobile-banking.svg",
        text: "Mobile Banking",
        slug: "pay via e-wallet"
    },
    {
        method: "BANK",
        image: "/bank.svg",
        text: "Bank Transfer",
        slug: "pay via bank"
    },
]


type Prop = {
    booking: BookingInfoType
}
const PaymentDetailsForm = ({ booking }: Prop) => {

    const [payment, setPayment] = useState<PaymentMethodType>('CARD')
    const [payNow, setPayNow] = useState("paynow")
    const { mutate, isPending } = useCreatePayment()





    const onSubmit = (e: FormEvent) => {
        e.preventDefault()
        if (!booking) return;

        const payStatus = payNow === "paynow";
        const paymentData = {
            bookingId: booking._id,
            userId: booking.user._id,
            paymentMethod: payment,
            amount: booking.totalPrice,
            payNow: payStatus
        }

        const { success, error, data } = createPaymentSchema.safeParse(paymentData);

        if (!success) {
            addToast({
                title: error.issues.map(e => e.message),
                color: "danger"
            })
        }

        mutate(data as CreatePaymentType)
    }
    return (
        <section>


            <Card className='p-4'>
                <CardHeader>
                    <div>
                        <h1 className='head-1 mb-2'>Choose Payment Method </h1>
                        <p className='text-sm'>Select a payment option to proceed</p>
                    </div>
                </CardHeader>
                <CardBody>

                    <form onSubmit={onSubmit}>


                        {/* -----------Payment option------------ */}
                        <div className='flex flex-col sm:flex-row items-center justify-between gap-4  '>
                            {
                                paymentMethodOptions.map(value => (
                                    <Card
                                        key={value.method}
                                        isPressable
                                        onPress={() => setPayment(value.method as PaymentMethodType)}
                                        className={`
                                    ${payment === value.method ? 'bg-success-50 border-2 border-success-400' : 'bg-deafult'}
                                     min-w-44 w-full max-w-xs h-36 flex items-center justify-center`} >
                                        <CardBody >
                                            <div className={`flex flex-col items-center gap-2 `} >


                                                <Image
                                                    src={value.image}
                                                    alt={value.method + "icon"}
                                                    width={value.method === "CARD" ? 80 : 60}
                                                    height={value.method === "CARD" ? 80 : 60}


                                                />
                                                <div className='text-center'>
                                                    <h1 className=' font-medium capitalize'>
                                                        {value.text}
                                                    </h1>
                                                    <p className='text-sm capitalize'>
                                                        {value.slug}

                                                    </p>
                                                </div>
                                            </div>
                                        </CardBody>
                                    </Card>

                                ))
                            }





                        </div>
                        {
                            payment === "CARD" &&
                            <CardInputs />

                        }

                        {
                            payment === "MOBILE_BANKING" &&
                            <MobileBankingInputs />
                        }

                        {
                            payment === "BANK" &&
                            <BankingInputs />
                        }
                        <div className='my-4 space-y-2'>
                            <RadioGroup orientation="horizontal"
                                value={payNow}
                                onValueChange={setPayNow}
                            >
                                <Radio

                                    value={"paynow"}
                                    size='sm'

                                    classNames={{
                                        base: cn(
                                            'mr-2 flex item-center gap-2  py-1.5 px-3  rounded-lg bg-slate-200 border-transparent',
                                            "data-[selected=true]:border-primary data-[selected=true]:bg-primary-100",
                                        )
                                    }}


                                >Pay Now</Radio>
                                <Radio value={"payAtProperty"}
                                    size='sm'
                                    classNames={{
                                        base: cn(
                                            'flex item-center gap-2  py-1.5 px-3   rounded-lg bg-slate-200 border-transparent',
                                            "data-[selected=true]:border-primary data-[selected=true]:bg-primary-100",
                                        )
                                    }}

                                >
                                    Pay at property

                                </Radio>
                            </RadioGroup>

                            <label htmlFor='secure' className=' flex items-center  gap-1 mt-4'>
                                <input type='checkbox' id='secure' />
                                <span className='text-sm '>
                                    Secure and encrypted
                                </span>

                            </label>
                            <label htmlFor='terms' className=' flex items-center gap-1'>
                                <input type='checkbox' id='terms' />
                                <span className='text-sm '>
                                    I agree to the cancellation policy and terms
                                </span>

                            </label>
                        </div>
                        <Button
                            type='submit'
                            isLoading={isPending}
                            variant='solid'
                            color='primary'
                            radius='sm'
                            fullWidth

                        >
                            Proceed to Payment
                        </Button>
                    </form>
                </CardBody>
            </Card>

        </section>
    )
}

export default PaymentDetailsForm





