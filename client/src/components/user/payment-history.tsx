"use client";
import { useGetReceipts } from "@/hooks/use-payment";
import { Button, Card, CardBody, Chip } from "@heroui/react";
import { Skeleton } from "@heroui/skeleton";
import { useParams } from "next/navigation";
import { PaymentReceiptPDF } from "../ui/receipt-pdf";
import { PDFDownloadLink } from "@react-pdf/renderer";

const PaymentHistory = () => {
    const { id } = useParams<{ id: string }>()
    const { data: receipts, isLoading, isError, error } = useGetReceipts(id);

    if (isLoading) {
        return <div className="space-y-4">
            <Skeleton className="w-full h-40 rounded-md" />
            <Skeleton className="w-full h-40 rounded-md" />
            <Skeleton className="w-full h-40 rounded-md" />
            <Skeleton className="w-full h-40 rounded-md" />
        </div>
    }
    if (isError || receipts?.length === 0) {
        console.error(error?.message)
    }
    return (
        <div className="min-h-screen">
            <div className="max-w-5xl mx-auto space-y-6">

                <h1 className="text-2xl font-semibold mb-6">
                    Payment History
                </h1>
                {
                    receipts?.map(receipt => (
                        <Card className="py-4" key={receipt._id}>
                            <CardBody>
                                <div className="flex justify-between items-center ">
                                    <div className="flex gap-2 items-center">
                                        <p >Receipt No</p>
                                        <p className="font-mono text-sm font-medium text-violet-600">
                                            {receipt.receiptNo}
                                        </p>
                                    </div>

                                    <Chip color={receipt.status === "PAID" ? "success" : "secondary"} radius="full" className="text-white" size="sm">
                                        {receipt.status}
                                    </Chip>
                                </div>
                                <div className="grid grid-cols-2 md:grid-cols-4 gap-4 text-sm bg-slate-200 p-2 rounded-md my-4">

                                    <div className="space-y-2"  >
                                        <p >Booking ID</p>
                                        <p className="font-medium text-violet-600"   >
                                            {receipt.bookingId.slice(0, 12)}...
                                        </p>
                                    </div>

                                    <div className="space-y-2"  >
                                        <p >Payment Method</p>
                                        <p className="font-medium">
                                            {receipt.paymentMethod}
                                        </p>
                                    </div>

                                    <div className="space-y-2"  >
                                        <p >Amount</p>
                                        <p className="font-semibold text-amber-500 text-lg">
                                            {receipt.amount} $
                                        </p>
                                    </div>

                                    <div className="space-y-2"  >
                                        <p >Paid At</p>
                                        <p className="font-medium">
                                            {new Date(receipt.paidAt).toLocaleString()}
                                        </p>
                                    </div>


                                </div>
                                <div className="flex justify-end">
                                    <Button
                                        as={PDFDownloadLink}
                                        document={<PaymentReceiptPDF receipt={receipt} />}
                                        fileName={"receipt"}
                                        size="sm"
                                        variant="solid"
                                        radius="sm"
                                        color="primary"
                                        className="w-fit "
                                    >

                                        Download PDF
                                    </Button>
                                </div>


                            </CardBody>
                        </Card>
                    ))
                }

            </div>
        </div>
    )
}

export default PaymentHistory