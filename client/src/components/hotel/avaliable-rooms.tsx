'use client'

import { useMemo, useState } from 'react'
import Image from 'next/image'
import {
    addToast,
    Button,
    Card,
    CardBody,
    DateRangePicker,
    Input,
    Select,
    SelectItem,
    Spinner,
    Tooltip
} from '@heroui/react'

import { BedDouble, Check, UsersRound } from 'lucide-react'
import { parseDate, getLocalTimeZone } from '@internationalized/date'
import { useGetAvaliableRoom } from '@/hooks/use-hotel'
import { useAuth } from '@/stores/auth-store'
import RoomCardLoading from '../loading/room-loading'
import Empty from '../empty'
import { useUpdateParams } from '@/hooks/use-params'
import { useCreateBooking } from '@/hooks/use-booking'



type Props = {
    hotel_id: string
}




const AvaliableRooms = ({ hotel_id }: Props) => {
    const { updateParams, searchParams } = useUpdateParams();
    const user_id = useAuth(s => s.user?.id);


    const todayISO = useMemo(
        () => new Date().toISOString(),
        []
    )

    const tomorrowISO = useMemo(() => {
        const d = new Date()
        d.setDate(d.getDate() + 1)
        return d.toISOString()
    }, [])

    const checkIn = searchParams.get('checkIn') ?? todayISO;
    const checkOut = searchParams.get('checkOut') ?? tomorrowISO;
    const maxPeople = Number(searchParams.get('guest'));
    const guest = maxPeople < 1 ? 1 : maxPeople

    const [value, setValue] = useState<any>({
        start: parseDate(checkIn.split("T")[0]),
        end: parseDate(checkOut.split("T")[0]),
    });


    const [count, setCount] = useState(guest);
    const [quantity, setQuantity] = useState(1)



    const { isAuthenticated } = useAuth()

    /* ----------------------------- fetch rooms ----------------------------- */

    const { data: rooms = [], isLoading, } = useGetAvaliableRoom(hotel_id,
        {
            checkIn,
            checkOut,
            guest,
        }
    )
    const onFliter = () => {
        const utcCheckIn = value.start.toDate(getLocalTimeZone()).toISOString();
        const utcCheckOut = value.end.toDate(getLocalTimeZone()).toISOString();
        const params = {
            checkIn: utcCheckIn,
            checkOut: utcCheckOut,
            guest: count.toString()
        }
        updateParams(params);

    }

    //create booking
    const { mutate, isPending } = useCreateBooking()
    const handleCreateBooking = (room_id: string, quantity: number, price: number) => {
        if (!user_id) {
            addToast({
                title: "Please login first.",
                color: "warning"
            })
            throw new Error("You are not aunticated yet.")
        }


        const bookingData = {
            user_id,
            hotel_id,
            room_id,
            quantity,
            guest,
            total_price: price * quantity,
            check_in: checkIn,
            check_out: checkOut,
        }

        mutate(bookingData)
    }

    return (
        <section className="py-4 space-y-4 mb-6 relative">

            <h1 className="head-1">Available Rooms</h1>

            {/* Filters */}
            <div

                className="flex flex-col sm:flex-row gap-2">
                <DateRangePicker
                    name='dateRange'
                    aria-label="Date range"
                    value={value}
                    onChange={setValue}
                />

                <Input
                    type="number"
                    placeholder="Guests"
                    startContent={<UsersRound size={16} />}
                    aria-label="Guests"
                    defaultValue={String(Number(guest) || 1)}
                    onChange={(e) => setCount(Number(e.target.value))}
                />

                <Button
                    isLoading={isLoading}
                    onPress={() => onFliter()}
                    color="primary"
                >Apply</Button>
            </div>
            {
                isPending && <section className='absolute inset-0  z-10 flex justify-center items-center bg-white/80'>
                    <Spinner size='lg' variant='spinner' />
                </section>
            }
            {/* Empty state */}
            {!isLoading && rooms.length === 0 && <Empty />}

            {/* Rooms */}
            <div className="flex flex-col gap-4">
                {isLoading ? (
                    <RoomCardLoading />
                ) : (
                    rooms.map(room => (
                        <Card key={room.id}>
                            <CardBody className="p-2">
                                <div className="flex flex-col sm:flex-row gap-4">
                                    {/* Image */}
                                    <div className="relative  w-full md:max-w-sm  aspect-video">
                                        <Image
                                            src={room.photo?.secure_url || '/room.webp'}
                                            alt={room.name}
                                            fill
                                            className=" rounded-lg"
                                        />
                                        <div className="absolute left-1 bottom-1 bg-yellow-100 px-2 py-1 flex items-center gap-1 text-sm">
                                            <Check className="text-green-600" size={14} />
                                            Free cancellation
                                        </div>
                                    </div>

                                    {/* Info */}
                                    <div className="flex-1 space-y-2 p-4">
                                        <div className='flex justify-between flex-wrap'>
                                            <h2 className="head-1 truncate line-clamp-1 text-wrap">{room.name}</h2>

                                            <div >


                                                <Select
                                                    aria-label='quantity select box'
                                                    variant='bordered'
                                                    color='secondary'
                                                    radius='sm'
                                                    fullWidth={false}

                                                    selectedKeys={[quantity.toString()]}
                                                    onSelectionChange={(keys) => {
                                                        const val = [...keys][0]
                                                        setQuantity(Number(val))
                                                    }}
                                                    className='w-8'
                                                >
                                                    {
                                                        Array.from({ length: room.total_rooms }).map((_, i) => {
                                                            const val = (i + 1).toString();
                                                            return (
                                                                < SelectItem
                                                                    key={val}
                                                                    textValue={val}
                                                                >
                                                                    {val}

                                                                </SelectItem>
                                                            )


                                                        })
                                                    }

                                                </Select>

                                            </div>
                                        </div>


                                        <div className="flex gap-4 text-sm">
                                            <span className="flex items-center gap-1 capitalize">
                                                <BedDouble size={16} /> {room.bedTypes} Bed
                                            </span>
                                            <span className="flex items-center gap-1">
                                                <UsersRound size={16} /> {room.maxPeople} guests
                                            </span>
                                        </div>

                                        <p className="text-danger">
                                            {room.available_rooms} rooms left</p>

                                        <div className="flex justify-between items-end">
                                            <div>
                                                <p>
                                                    From{' '}
                                                    <span className="text-xl font-semibold">
                                                        ${room.price}
                                                    </span>{' '}
                                                    / night
                                                </p>
                                                <p className="text-sm">Includes taxes and fees</p>
                                            </div>

                                            <Tooltip
                                                color={isAuthenticated ? 'default' : 'danger'}
                                                content={
                                                    isAuthenticated
                                                        ? 'Reserve this room'
                                                        : 'Please login to reserve'
                                                }
                                            >
                                                <Button
                                                    isLoading={isPending}
                                                    color="primary"
                                                    radius="sm"
                                                    onPress={() => handleCreateBooking(room.id, quantity, room.price)}

                                                >
                                                    Reserve
                                                </Button>
                                            </Tooltip>
                                        </div>
                                    </div>
                                </div>
                            </CardBody>
                        </Card>
                    ))
                )}
            </div>
        </section >
    )
}

export default AvaliableRooms
