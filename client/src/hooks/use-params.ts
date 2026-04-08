import { usePathname, useRouter, useSearchParams } from "next/navigation"
import { useCallback } from "react";

export const useUpdateParams = () => {
    const router = useRouter();
    const searchParams = useSearchParams();
    const pathname = usePathname();

    const updateParams = useCallback(
        (params: Record<string, string | null>) => {
            const newParams = new URLSearchParams(searchParams.toString());
            Object.entries(params).forEach(([key, value]) => {
                if (value === null) {
                    newParams.delete(key)
                } else {
                    newParams.set(key, value)
                }
                router.replace(`${pathname}?${newParams.toString()}`, { scroll: false })
            })
        },
        [router, pathname, searchParams])
    return { searchParams, updateParams }
}