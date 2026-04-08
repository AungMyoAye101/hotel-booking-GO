import { AirVent, Bath, Dumbbell, ParkingCircle, Utensils, Wifi } from "lucide-react";

export const AMENITIES = [
    { label: "Wi-Fi", value: "wifi", icon: <Wifi /> },
    { label: "Swimming Pool", value: "pool", icon: <Bath /> },
    { label: "Parking", value: "parking", icon: <ParkingCircle /> },
    { label: "Air Conditioning", value: "ac", icon: <AirVent /> },
    { label: "Gym", value: "gym", icon: <Dumbbell /> },
    { label: "Breakfast", value: "breakfast", icon: <Utensils /> },
];