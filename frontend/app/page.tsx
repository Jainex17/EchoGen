"use client";

import Link from "next/link";
import Image from "next/image";
import { useState } from "react";
import { useAuth } from "./AuthContext";

export default function Home() {
  const [prompt, setPrompt] = useState("");
  const [style, setStyle] = useState("podcast");
  const [isLoading, setIsLoading] = useState(false);
  const { user, login, logout, isLoading: isAuthLoading } = useAuth();

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setIsLoading(true);
    setTimeout(() => setIsLoading(false), 1200);
  }

  return (
    <div className="font-sans">
      <main className="mx-auto max-w-6xl px-6 py-16 sm:py-24">
        <section className="grid gap-10 sm:grid-cols-1 sm:gap-16 items-center">
          <div>
            <h1 className="text-3xl sm:text-5xl font-semibold tracking-tight">
              Learn anything through clear, focused audio
            </h1>
            <p className="mt-4 text-base sm:text-lg text-black/70 dark:text-white/70">
              Turn your prompt into a concise podcast, a YouTuber-style explanation, or a step-by-step tutorial.
            </p>
            <div className="mt-8 flex items-center gap-3">
              <Link
                href="/audio-gen"
                className="inline-flex items-center justify-center rounded-md bg-foreground px-5 py-2.5 text-sm font-medium text-background hover:opacity-90 transition"
              >
                Create audio
              </Link>
              {
                isAuthLoading ? (
                  <p className="text-sm text-black/70 dark:text-white/70">
                    Loading...
                  </p>
                ) : user ? (
                  <>
                  <p className="text-sm text-black/70 dark:text-white/70">
                    Logged in as {user.name}
                  </p>
                  <button
                    onClick={logout}
                    className="inline-flex items-center justify-center rounded-md bg-foreground px-5 py-2.5 text-sm font-medium text-background hover:opacity-90 transition"
                  >
                    Logout
                  </button>
                  </>
                ) : (
                  <button
                    onClick={login}
                    className="inline-flex items-center justify-center rounded-md bg-foreground px-5 py-2.5 text-sm font-medium text-background hover:opacity-90 transition"
                  >
                    Login
                  </button>
                )}
            </div>
            </div>
        </section>

        <section id="features" className="mt-20 sm:mt-12">
          <h2 className="text-xl sm:text-2xl font-semibold">Built for focused learning</h2>
          <div className="mt-6 grid gap-4 sm:gap-6 sm:grid-cols-3">
            <FeatureCard
              icon="/file.svg"
              title="Podcast mode"
              desc="Digest topics as short, structured episodes."
            />
            <FeatureCard
              icon="/window.svg"
              title="YouTuber style"
              desc="Casual, high-energy explanations that keep you engaged."
            />
            <FeatureCard
              icon="/globe.svg"
              title="Tutorial walkthrough"
              desc="Clear, step-by-step guides for hands-on learning."
            />
          </div>
        </section>
      </main>
    </div>
  );
}

function FeatureCard({
  icon,
  title,
  desc,
}: {
  icon: string;
  title: string;
  desc: string;
}) {
  return (
    <div className="rounded-lg border border-black/10 dark:border-white/15 p-5">
      <div className="flex items-center gap-3">
        <span className="inline-flex h-9 w-9 items-center justify-center rounded-md border border-black/10 dark:border-white/15">
          <Image src={icon} alt="" width={16} height={16} />
        </span>
        <p className="font-medium">{title}</p>
      </div>
      <p className="mt-3 text-sm text-black/70 dark:text-white/70">{desc}</p>
    </div>
  );
}
