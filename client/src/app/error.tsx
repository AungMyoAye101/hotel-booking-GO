"use client"

import { Button } from "@heroui/button"
import { useEffect } from "react"

const Error = (
    {
        error,
        reset
    }: {
        error: Error & { digest?: string },
        reset: () => void
    }
) => {

    useEffect(() => {
        console.warn(error)

    }, [error])

    let title = "Something went worng!";
    let message = error.message;

    return (
        <section className='h-screen px-4 py-12 flex justify-center items-center '>
            <div className="flex flex-col items-center gap-4">
                <h1 className='text-6xl font-bold text-center'>
                    SORRY
                </h1>
                <h2 className='text-2xl font-bold text-center'>{title}</h2>

                <Button
                    variant="solid"
                    radius="sm"
                    size="sm"
                    color="primary"
                    onPress={() => reset()}
                >
                    Try again
                </Button>
            </div>
        </section>
    )
}

export default Error