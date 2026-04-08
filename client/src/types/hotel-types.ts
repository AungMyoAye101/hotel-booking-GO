export type hotelTypes = "hotel" | "motel" | "guest-house"

export type hotelType = {
    id: string,
    name: string,
    description: string,
    rating: number,
    star: number,
    type: hotelTypes,
    address: string,
    price: number,
    amenities: string[],
    city: string,
    country: string,
    createdAt: Date,
    photo?: photoType
}
export type photoType = {
    id: string,
    secure_url: string,
    public_id: string,

}