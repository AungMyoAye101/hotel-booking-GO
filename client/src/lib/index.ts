
export const getBaseUrl = () => {
    const BASE_URL = process.env.BASE_URL;
    if (!BASE_URL) {
        throw new Error("Base url is required.")
    }
    return BASE_URL;

}

export const PROTECTED_ROUTE = [
    '/booking',
    '/user',
    '/payment'
]