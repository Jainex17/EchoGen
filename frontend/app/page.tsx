"use client";

import Link from "next/link";
import Image from "next/image";
import { useState } from "react";

export default function Home() {
  const [prompt, setPrompt] = useState("");
  const [style, setStyle] = useState("podcast");
  const [isLoading, setIsLoading] = useState(false);

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
            <div className="text-center">
              <h1 className="text-4xl sm:text-6xl font-bold tracking-tight text-indigo-600">
                EchoGen
              </h1>
              <p className="mt-2 text-lg sm:text-xl text-slate-600 font-medium">
                Learn anything through audio
              </p>
              <p className="mt-4 text-base sm:text-lg text-slate-700 max-w-2xl mx-auto">
                Transform any topic into engaging audio content. Choose from podcast-style discussions, tutorial walkthroughs, or casual explanations tailored to your learning style.
              </p>
            </div>
            <div className="mt-8 flex justify-center">
              <Link
                href="/audio-gen"
                className="inline-flex items-center justify-center rounded-lg bg-indigo-600 px-8 py-3 text-base font-semibold text-white hover:bg-indigo-700 transition-all duration-200 shadow-lg hover:shadow-xl transform hover:scale-105"
              >
                🎧 Create Audio Content
              </Link>
            </div>
            </div>
        </section>

        <section id="features" className="mt-20 sm:mt-16">
          <div className="text-center mb-12">
            <h2 className="text-2xl sm:text-3xl font-bold text-slate-900">Built for focused learning</h2>
            <p className="mt-3 text-lg text-slate-600">Choose the perfect format for your learning style</p>
          </div>
          <div className="mt-8 grid gap-6 sm:gap-8 md:grid-cols-3">
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
    <div className="group rounded-xl border border-slate-200 p-6 hover:border-indigo-300 transition-all duration-200 hover:shadow-lg hover:shadow-indigo-100 bg-white/80 backdrop-blur">
      <div className="text-center">
        <div className="inline-flex h-12 w-12 items-center justify-center rounded-xl bg-indigo-100 group-hover:bg-indigo-200 transition-all duration-200">
          <Image src={icon} alt="" width={20} height={20} className="opacity-80" />
        </div>
        <h3 className="mt-4 font-semibold text-lg text-slate-900">{title}</h3>
        <p className="mt-2 text-sm text-slate-600 leading-relaxed">{desc}</p>
      </div>
    </div>
  );
}
