"use client";

import Link from "next/link";
import { Headphones, Mic, Video, BookOpen } from "lucide-react";
import bg from "@/app/assets/bg.png"

export default function Home() {
  return (
    <div 
      className="h-screen flex items-center justify-center relative"
      style={{
        backgroundImage: `url(${bg.src})`,
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        backgroundRepeat: 'no-repeat'
      }}
    >
      {/* Overlay for better text readability */}
      <div className="absolute inset-0 bg-black/70 bg-opacity-40"></div>
      <div className="max-w-4xl mx-auto px-6 text-center relative z-10">
        {/* Icon */}
        <div className="flex justify-center mb-4">
          <div className="w-16 h-16 bg-[#8ECAE6]/20 rounded-2xl flex items-center justify-center">
            <Headphones className="w-8 h-8 text-[#FFB703]" />
          </div>
        </div>

        {/* Headline */}
        <h1 className="text-5xl sm:text-6xl lg:text-7xl font-bold text-white mb-6 tracking-tight">
          Learn anything
          <span className="block text-[#FFB703]">through audio</span>
        </h1>
        
        {/* Description */}
        <p className="text-xl text-gray-200 mb-12 max-w-2xl mx-auto leading-relaxed">
          Transform any topic into engaging audio content. Choose from podcast discussions, tutorial walkthroughs, or casual explanations.
        </p>

        {/* CTA Button */}
        <Link
          href="/audio-gen"
          className="inline-flex items-center justify-center px-8 py-4 text-lg font-semibold text-[#1E2D2F] bg-[#FFB703] rounded-full hover:bg-[#e6a502] transition-all duration-200 shadow-lg hover:shadow-xl transform hover:scale-105"
        >
          Create Audio Content
        </Link>

        {/* Features */}
        <div className="mt-16 grid grid-cols-1 sm:grid-cols-3 gap-8">
          <div className="text-center">
            <div className="flex justify-center mb-3">
              <Mic className="w-8 h-8 text-[#FFB703]" />
            </div>
            <h3 className="font-semibold text-white mb-2">Podcast Mode</h3>
            <p className="text-sm text-gray-200">Professional discussions</p>
          </div>
          <div className="text-center">
            <div className="flex justify-center mb-3">
              <Video className="w-8 h-8 text-[#FFB703]" />
            </div>
            <h3 className="font-semibold text-white mb-2">YouTuber Style</h3>
            <p className="text-sm text-gray-200">Casual explanations</p>
          </div>
          <div className="text-center">
            <div className="flex justify-center mb-3">
              <BookOpen className="w-8 h-8 text-[#FFB703]" />
            </div>
            <h3 className="font-semibold text-white mb-2">Tutorial Mode</h3>
            <p className="text-sm text-gray-200">Step-by-step guides</p>
          </div>
        </div>
      </div>
    </div>
  );
}

