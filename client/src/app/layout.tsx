import type { Metadata } from "next";
import "./globals.css";
import { Providers } from "./provider";
import Header from "@/components/header";
import Footer from "@/components/footer";
import AuthInitializer from "@/components/auth/auth-provider";




export const metadata: Metadata = {
  icons: '/brand-logo.svg',
  title: "Booking website",
  description: "Book hotel ,reserve room and find perfect palace",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {


  return (
    <html lang="en">
      <body>
        <Providers>
          <Header />
          <div className="max-w-7xl mx-auto">
            <AuthInitializer />
            {children}
          </div>
          <Footer />
        </Providers>

      </body>
    </html >
  )
}
