import type { Metadata } from "next"
import { Inter } from "next/font/google"
import { Toaster } from 'react-hot-toast'
import "./globals.css"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "License Key System",
  description: "A license key system capable of managing large amounts of users and handling product licensing and availability",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`dark ${inter.className} antialiased`}
      >
          <Toaster
              position="top-center"
          />
        {children}
      </body>
    </html>
  );
}
