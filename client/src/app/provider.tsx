'use client'

import { HeroUIProvider, ToastProvider } from '@heroui/react'
import dynamic from 'next/dynamic';

const ThemeProvider = dynamic(() => import("next-themes").then(mod => mod.ThemeProvider), { ssr: false })

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
        <ThemeProvider attribute="class" defaultTheme='light'>
            <QueryClientProvider client={queryClient}>
                <HeroUIProvider>
                    <ToastProvider />
                    {children}
                </HeroUIProvider>
            </QueryClientProvider>
        </ThemeProvider>
    )
}