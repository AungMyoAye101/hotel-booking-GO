'use client';
import { Button, DatePicker, Input } from "@heroui/react";
import { MapPin, UsersRound } from "lucide-react";
import { useRouter, useSearchParams } from "next/navigation";
import { useState } from "react";

const InputCon = () => {
    const searchParam = useSearchParams();
    const router = useRouter();


    const [destination, setDestination] = useState(searchParam.get('destination') || '')
    const [guest, setGuest] = useState(searchParam.get('guest') || '')
    const [checkIn, setCheckIn] = useState<any>(searchParam.get('checkIn') ? new Date(searchParam.get('checkIn')!) : null)
    const [checkOut, setCheckOut] = useState<any>(searchParam.get('checkOut') ? new Date(searchParam.get('checkOut')!) : null)

    const handlesearch = () => {
        const params = new URLSearchParams(searchParam);
        params.set('destination', destination);
        const checkInStr = checkIn ? (checkIn instanceof Date ? checkIn.toISOString() : String(checkIn)) : '';
        const checkOutStr = checkOut ? (checkOut instanceof Date ? checkOut.toISOString() : String(checkOut)) : '';
        params.set('checkIn', checkInStr);
        params.set('checkOut', checkOutStr);
        params.set('guest', guest);
        router.push(`/search?${params.toString()}`);
    }

    return (
        <div className="grid grid-cols-1  sm:grid-cols-2 md:grid-cols-5 gap-2  items-center  p-2 bg-white rounded-lg   overflow-hidden ">

            <Input
                name="destination"
                type="text"
                placeholder="destination"
                size="sm"
                radius="sm"
                aria-label="destination"
                startContent={<MapPin size={18} />}
                onChange={(e) => setDestination(e.target.value)}
            />
            <DatePicker
                name="checkIn"
                value={checkIn}
                onChange={(val: any) => setCheckIn(val)}
                size="sm"
                radius="sm"
                aria-label="check in"
            />
            <DatePicker
                name="checkOut"
                value={checkOut}
                onChange={(val: any) => setCheckOut(val)}
                size="sm"
                radius="sm"
                aria-label="check out"
            />
            <Input
                name="guest"
                type="number"
                placeholder="guest"
                size="sm"
                radius="sm"
                aria-label="guest"
                startContent={<UsersRound size={18} />}
                onChange={(e) => setGuest(e.target.value)}
            />
            <Button
                onPress={() => handlesearch()}
                color="primary"
                className="sm:col-span-2 md:col-span-1">Search</Button>
        </div>
    )
}

export default InputCon