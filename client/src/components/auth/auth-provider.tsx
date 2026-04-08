// components/AuthInitializer.tsx
"use client";

import { useGetMe } from '@/hooks/use-user';
import { useAuth } from '@/stores/auth-store';
import { useEffect } from 'react';



export default function AuthInitializer() {
    const { setUser, logout } = useAuth(s => s)
    const { data: user, isSuccess, isError } = useGetMe()


    useEffect(() => {
        if (isSuccess) {
            if (user) {
                setUser(user)
            }
        }
        if (isError) {
            logout()
        }

    }, [isSuccess, isError])

    return null;
}