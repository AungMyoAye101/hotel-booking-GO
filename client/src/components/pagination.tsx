"use client";
import { MetaType } from '@/types';
import { Button, Pagination } from '@heroui/react'
import { ArrowLeft, ArrowRight, } from 'lucide-react';
import { useRouter, useSearchParams } from 'next/navigation';


type Props = {
    meta: MetaType
}
const Meta = ({ meta }: Props) => {

    const router = useRouter()

    const searchParams = useSearchParams();


    const handlePageChange = (page: number) => {

        const params = new URLSearchParams(searchParams.toString());
        params.set('page', page.toString());
        router.replace(`/search?${params.toString()}`);
    }
    return (
        <div className='flex justify-end py-6'>
            <div className='flex gap-4 items-center '>


                <Button
                    isIconOnly
                    color={meta.hasPrev ? 'secondary' : 'default'}
                    radius='sm'
                    disabled={!meta.hasPrev}
                    onPress={() => handlePageChange(meta.page - 1)}>
                    <ArrowLeft />
                </Button>
                <Pagination
                    initialPage={meta.page}
                    total={meta.totalPages}
                    color='secondary'
                    radius='sm'
                    onChange={handlePageChange} />
                <Button
                    isIconOnly
                    color={meta.hasNext ? 'secondary' : 'default'}
                    radius='sm'
                    disabled={!meta.hasNext}
                    onPress={() => handlePageChange(meta.page + 1)}>
                    <ArrowRight />
                </Button>
            </div>
        </div>
    )
}

export default Meta