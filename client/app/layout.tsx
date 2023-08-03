import "../styles/globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Nav from "@/components/Nav";
import ThemeRegistry from "@/components/Themeregistry";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Arbitrage web app",
  description: "Get various arbitrage betting opportunities",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ThemeRegistry options={{ key: "mui" }}>
          <Nav />
          {children}
        </ThemeRegistry>
      </body>
    </html>
  );
}
