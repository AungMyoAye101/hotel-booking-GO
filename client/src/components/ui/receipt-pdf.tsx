
import {
    Document,
    Page,
    Text,
    View,
    StyleSheet,

} from "@react-pdf/renderer"
import { ReceiptType } from "@/types"


const styles = StyleSheet.create({
    page: {
        padding: 40,
        fontSize: 12,
        fontFamily: "Helvetica"
    },

    header: {
        marginBottom: 30
    },

    title: {
        fontSize: 22,
        fontWeight: 700,
        marginBottom: 10
    },

    section: {
        marginBottom: 20
    },

    row: {
        flexDirection: "row",
        justifyContent: "space-between",
        marginBottom: 8
    },

    label: {
        color: "#666"
    },

    value: {
        fontWeight: 600
    },

    divider: {
        borderBottomWidth: 1,
        borderBottomColor: "#e5e5e5",
        marginVertical: 20
    },

    totalSection: {
        marginTop: 20,
        paddingTop: 10,
        borderTopWidth: 1,
        borderTopColor: "#000"
    },

    totalAmount: {
        fontSize: 16,
        fontWeight: 700
    },

    footer: {
        marginTop: 40,
        fontSize: 10,
        color: "#888"
    }
})




export const PaymentReceiptPDF = ({
    receipt
}: {
    receipt: ReceiptType
}) => (
    <Document>
        <Page size="A4" style={styles.page}>

            {/* Header */}
            <View style={styles.header}>
                <Text style={styles.title}>PAYMENT RECEIPT</Text>
                <Text>Receipt No: {receipt.receiptNo}</Text>
            </View>

            {/* Payment Info */}
            <View style={styles.section}>
                <View style={styles.row}>
                    <Text style={styles.label}>Payment ID</Text>
                    <Text style={styles.value}>{receipt.paymentId}</Text>
                </View>

                <View style={styles.row}>
                    <Text style={styles.label}>Booking ID</Text>
                    <Text style={styles.value}>{receipt.bookingId}</Text>
                </View>

                <View style={styles.row}>
                    <Text style={styles.label}>Payment Method</Text>
                    <Text style={styles.value}>{receipt.paymentMethod}</Text>
                </View>

                <View style={styles.row}>
                    <Text style={styles.label}>Status</Text>
                    <Text style={styles.value}>{receipt.status}</Text>
                </View>

                <View style={styles.row}>
                    <Text style={styles.label}>Paid At</Text>
                    <Text style={styles.value}>
                        {new Date(receipt.paidAt).toLocaleString()}
                    </Text>
                </View>
            </View>

            <View style={styles.divider} />

            {/* Total */}
            <View style={styles.totalSection}>
                <View style={styles.row}>
                    <Text>Total Amount Paid</Text>
                    <Text style={styles.totalAmount}>
                        {receipt.amount.toFixed(2)}
                    </Text>
                </View>
            </View>

            {/* Footer */}
            <View style={styles.footer}>
                <Text>Thank you for your payment.</Text>
                <Text>This is a system-generated receipt.</Text>
            </View>

        </Page>
    </Document>
)
