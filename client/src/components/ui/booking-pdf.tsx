"use client"
import { BookingInfoType } from '@/types';
import { Document, Page, Text, View, StyleSheet, } from '@react-pdf/renderer';


type Prop = {
    booking: BookingInfoType
}




const styles = StyleSheet.create({
    page: {
        padding: 32,
        fontSize: 11,
        fontFamily: "Helvetica",
        color: "#111827",
        position: 'relative'
    },

    header: {
        flexDirection: "row",
        justifyContent: "space-between",
        marginBottom: 24,
    },

    brand: {
        fontSize: 20,
        fontWeight: "bold",
        color: "#2563eb",
    },

    section: {
        marginBottom: 16,
    },

    sectionTitle: {
        fontSize: 13,
        fontWeight: "bold",
        marginBottom: 8,
        borderBottom: "1 solid #e5e7eb",
        paddingBottom: 4,
    },

    row: {
        flexDirection: "row",
        justifyContent: "space-between",
        marginBottom: 4,
    },

    label: {
        color: "#6b7280",
    },

    value: {
        fontWeight: "bold",
    },

    hotelBox: {
        border: "1 solid #e5e7eb",
        borderRadius: 6,
        padding: 12,
    },

    totalBox: {
        marginTop: 16,
        padding: 12,
        borderRadius: 6,
        backgroundColor: "#f9fafb",
    },

    totalText: {
        fontSize: 14,
        fontWeight: "bold",
    },

    footer: {
        marginTop: 40,
        textAlign: "center",
        color: "#6b7280",
        fontSize: 10,
    },
});

type Props = {
    booking: BookingInfoType;
};


const BookingPDF = ({ booking }: Prop) => {
    return (
        <Document>
            <Page size="A4" style={styles.page}>

                {/* Header */}
                <View style={styles.header}>
                    <Text style={styles.brand}>Booking</Text>
                    <View>
                        <Text>Booking ID</Text>
                        <Text style={{ fontWeight: "bold", marginTop: 4 }}>{booking._id}</Text>
                    </View>
                </View>


                {/* Guest Info */}
                <View style={styles.section}>
                    <Text style={styles.sectionTitle}>Guest Information</Text>

                    <View style={styles.row}>
                        <Text style={styles.label}>Name</Text>
                        <Text style={styles.value}>{booking.name || booking.user.name}</Text>
                    </View>

                    <View style={styles.row}>
                        <Text style={styles.label}>Email</Text>
                        <Text>{booking.email || "-"}</Text>
                    </View>

                    <View style={styles.row}>
                        <Text style={styles.label}>Phone</Text>
                        <Text>{booking.phone || ''}</Text>
                    </View>
                </View>

                {/* Hotel Info */}
                <View style={styles.section}>
                    <Text style={styles.sectionTitle}>Hotel Details</Text>

                    <View style={styles.hotelBox}>
                        <Text style={{ fontSize: 12, fontWeight: "bold", marginBottom: 4 }}>
                            {booking.hotel.name}
                        </Text>

                        <Text>
                            {booking.hotel.adddress}, {booking.hotel.city}
                        </Text>


                    </View>
                </View>

                {/* Booking Info */}
                <View style={styles.section}>
                    <Text style={styles.sectionTitle}>Booking Details</Text>

                    <View style={styles.row}>
                        <Text style={styles.label}>Check-in</Text>
                        <Text>{new Date(booking.checkIn).toDateString()}</Text>
                    </View>

                    <View style={styles.row}>
                        <Text style={styles.label}>Check-out</Text>
                        <Text>{new Date(booking.checkOut).toDateString()}</Text>
                    </View>

                    <View style={styles.row}>
                        <Text style={styles.label}>Room Quantity</Text>
                        <Text>{booking.quantity}</Text>
                    </View>

                    <View style={styles.row}>
                        <Text style={styles.label}>Booking Status</Text>
                        <Text>{booking.status}</Text>
                    </View>
                </View>

                {/* Price */}
                <View style={styles.totalBox}>
                    <View style={styles.row}>
                        <Text>Total Price</Text>
                        <Text style={styles.totalText}>
                            ${booking.totalPrice.toFixed(2)}
                        </Text>
                    </View>
                </View>

                {/* Footer */}
                <Text style={styles.footer}>
                    Thank you for booking with Booking.
                    This document is generated electronically.
                </Text>
            </Page>
        </Document>
    )
}

export default BookingPDF 