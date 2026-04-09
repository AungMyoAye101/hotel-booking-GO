'use client'

import { HeroUIProvider, ToastProvider } from '@heroui/react'


import {
    QueryClient,
    QueryClientProvider,
} from '@tanstack/react-query'

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            refetchOnWindowFocus: false,

            staleTime: 5 * 60 * 1000, // 5 minutes
        }
    }
})

export function Providers({ children }: { children: React.ReactNode }) {
    return (
        <QueryClientProvider client={queryClient}>
            <HeroUIProvider>
                <ToastProvider />
                {children}
            </HeroUIProvider>
        </QueryClientProvider>

    )
}