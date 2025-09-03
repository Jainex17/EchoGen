"use client";

import Link from "next/link";

export default function Home() {
  return (
    <div className="h-screen flex items-center justify-center bg-gradient-to-br from-slate-50 to-indigo-50">
      <div className="max-w-4xl mx-auto px-6 text-center">
        {/* Icon */}
        <div className="flex justify-center mb-8">
          <div className="w-16 h-16 bg-indigo-100 rounded-2xl flex items-center justify-center">
            <span className="text-2xl">🎧</span>
          </div>
        </div>

        {/* Headline */}
        <h1 className="text-5xl sm:text-6xl lg:text-7xl font-bold text-slate-900 mb-6 tracking-tight">
          Learn anything
          <span className="block text-indigo-600">through audio</span>
        </h1>
        
        {/* Description */}
        <p className="text-xl text-slate-600 mb-12 max-w-2xl mx-auto leading-relaxed">
          Transform any topic into engaging audio content. Choose from podcast discussions, tutorial walkthroughs, or casual explanations.
        </p>

        {/* CTA Button */}
        <Link
          href="/audio-gen"
          className="inline-flex items-center justify-center px-8 py-4 text-lg font-semibold text-white bg-indigo-600 rounded-full hover:bg-indigo-700 transition-all duration-200 shadow-lg hover:shadow-xl transform hover:scale-105"
        >
          Create Audio Content
        </Link>

        {/* Features */}
        <div className="mt-16 grid grid-cols-1 sm:grid-cols-3 gap-8">
          <div className="text-center">
            <div className="text-3xl mb-3">🎙️</div>
            <h3 className="font-semibold text-slate-900 mb-2">Podcast Mode</h3>
            <p className="text-sm text-slate-600">Professional discussions</p>
          </div>
          <div className="text-center">
            <div className="text-3xl mb-3">🎬</div>
            <h3 className="font-semibold text-slate-900 mb-2">YouTuber Style</h3>
            <p className="text-sm text-slate-600">Casual explanations</p>
          </div>
          <div className="text-center">
            <div className="text-3xl mb-3">📚</div>
            <h3 className="font-semibold text-slate-900 mb-2">Tutorial Mode</h3>
            <p className="text-sm text-slate-600">Step-by-step guides</p>
          </div>
        </div>
      </div>
    </div>
  );
}

