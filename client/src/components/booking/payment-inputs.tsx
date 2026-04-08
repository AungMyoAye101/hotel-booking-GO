import { Badge, Card, Input } from '@heroui/react';
import { CheckIcon } from 'lucide-react';
import Image from 'next/image';
import { useState } from 'react';

const cardRegex = /^\d{4} \d{4} \d{4} \d{4}$/;
const phoneRegex = /^\+\d{1,4}\d{6,12}$/;

export const CardInputs = () => {
    const [demo, setDemo] = useState({
        card: "",
        expired: '',
        cvc: ""
    })
    return (
        <div className='py-4 space-y-4 my-4'>
            <Input type='text'
                placeholder='xxxx xxxx xxxx xxxx'
                label="Card number"
                labelPlacement='outside'
                radius='sm'
                isRequired
                value={demo.card}
                onValueChange={(v) => setDemo((pre) => ({ ...pre, card: v }))}
                validate={(value) => {
                    if (!cardRegex.test(value)) return "Invalid card number"
                }}
            />
            <div className='flex justify-between gap-4'>

                <Input
                    type='text'
                    name='expired'
                    label='Expired date'
                    labelPlacement='outside'
                    placeholder='MM/YY'
                    radius='sm'
                    isRequired
                    value={demo.expired}
                    onValueChange={(v) => setDemo((pre) => ({ ...pre, expired: v }))}
                    validate={(value) => {
                        const regex = /^(0[1-9]|1[0-2])\/?([0-9]{2})$/;
                        if (!regex.test(value)) return "Use MM.YY format";
                        return null;
                    }}
                />
                <Input
                    type='text'
                    placeholder='CVC'
                    label="Security code"
                    labelPlacement='outside'
                    radius='sm'
                    isRequired
                    value={demo.cvc}
                    onValueChange={(v) => setDemo((pre) => ({ ...pre, cvc: v }))}
                    validate={(value) => {
                        if (value.length !== 3 || typeof Number(value) !== "number") {
                            return "Invalid CVC fromat."
                        }
                    }}
                />
            </div>
            <button
                type='button'
                onClick={() => setDemo({
                    card: "1220 3445 6578 8800",
                    expired: '02/29',
                    cvc: "576"
                })}
                className='text-sm text-secondary'>Set deomo values</button>

        </div>
    )
}

//mobile banking inputs sections

const mobileBanks = [
    {
        name: "KBZ pay",
        image: "/kbz.webp"
    },
    {
        name: "AYA pay",
        image: "/aya.webp"
    },
    {
        name: "WAVE pay",
        image: "/wave.webp"
    },

];

export const MobileBankingInputs = () => {

    const [mobile, setMobile] = useState("KBZ pay")
    const [demo, setDemo] = useState({
        name: "",
        phone: ''
    })
    return (
        <div className='space-y-4 py-4  my-4'>
            <Input
                type='text'
                placeholder='name'
                label="Account name"
                labelPlacement='outside'
                radius='sm'
                isRequired
                value={demo.name}
                onValueChange={(v) => setDemo((pre) => ({ ...pre, name: v }))}
                validate={(value) => {
                    if (value.length < 2) return "Name is too short"
                }} />

            <Input

                type='text'
                placeholder='+959xxxxxxxx'
                label="Phone"
                labelPlacement='outside'
                radius='sm'
                value={demo.phone}
                onValueChange={(v) => setDemo((pre) => ({ ...pre, phone: v }))}
                validate={(v) => {
                    if (!phoneRegex.test(v)) return "Invalid phone number"
                }}
            />

            <div className='space-y-2'>
                <p className='text-sm'>Choose a provider</p>
                <div className="flex gap-4">
                    {mobileBanks.map((bank) => (
                        <Badge
                            key={bank.name}
                            isOneChar
                            isInvisible={mobile !== bank.name}
                            color='success'
                            content={<CheckIcon className='text-white' />}
                            placement='top-right'

                        >
                            <Card

                                isPressable
                                onPress={() => setMobile(bank.name)}
                                className={`p-0 overflow-hidden border-2 ${mobile === bank.name
                                    ? 'border-success'
                                    : ' border-primary-100'}
                                                     `}
                                shadow='sm'

                            >



                                <Image
                                    src={bank.image}
                                    alt={bank.name + "icon"}
                                    width={45}
                                    height={45}
                                    className='object-cover aspect-square'
                                />

                            </Card>
                        </Badge>
                    ))}
                </div>
            </div>
            <button
                type='button'
                onClick={() => setDemo({
                    name: "Saw Yu Panthra",
                    phone: "+95944055066"
                })}
                className='text-sm text-secondary'>Set deomo values</button>
        </div>
    )
}


//Banking

export const BankingInputs = () => {
    const [demo, setDemo] = useState({
        name: "",
        card: ''
    })

    return (<div className='flex flex-col gap-4 py-4  my-4'>
        <Input
            type='text'
            placeholder='name'
            label="Account name"
            labelPlacement='outside'
            radius='sm'
            isRequired
            value={demo.name}
            onValueChange={(v) => setDemo(pre => ({ ...pre, name: v }))}
            validate={(value) => {
                if (value.length < 2) return "Name is too short"
            }} />
        <Input
            type='text'
            placeholder='xxxx xxxx xxxx xxxx'
            label="Card number"
            labelPlacement='outside'
            radius='sm'
            isRequired
            value={demo.card}
            onValueChange={(v) => setDemo(pre => ({ ...pre, card: v }))}
            validate={(value) => {
                if (!cardRegex.test(value)) return "Invalid card number"
            }}
        />
        <button
            type='button'
            className='text-sm text-secondary text-start'
            onClick={() => setDemo({
                name: "Saw Yu Panthra",
                card: "1220 3445 6578 8800",
            })}>
            Set demo values
        </button>

    </div>
    )
}