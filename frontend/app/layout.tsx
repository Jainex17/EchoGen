import { AuthProvider } from "./AuthContext";
import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { AudioProvider } from "./components/AudioProvider";
import BottomPlayer from "./components/BottomPlayer";
import Navigation from "./components/Navigation";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "EchoGen - Learn through Audio",
  description: "Turn any topic into engaging audio content. Learn through podcasts, tutorials, and explanations tailored to your style.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        <div className="min-h-dvh">
          <AuthProvider>
            <AudioProvider>
              <Navigation />
              {children}
            </AudioProvider>
          </AuthProvider>
        </div>
      </body>
    </html>
  );
}
