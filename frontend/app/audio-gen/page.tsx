"use client";

import { useState, useRef, useEffect } from "react";
import { useAudio } from "../components/AudioProvider";
import BottomPlayer from "../components/BottomPlayer";

const STYLE_OPTIONS = [
  { id: "podcast", name: "Podcast", icon: "🎧", description: "Conversational and structured" },
  { id: "tutorial", name: "Tutorial", icon: "🧰", description: "Step-by-step walkthrough" },
  { id: "youtuber", name: "YouTuber", icon: "🎬", description: "Energetic and engaging" },
] as const;

const EXAMPLE_PROMPTS = [
  { icon: "🎧", text: "Summarize the history of AI in 3 minutes for a podcast", style: "podcast" },
  { icon: "📚", text: "Explain React hooks to a beginner with analogies and examples", style: "tutorial" },
  { icon: "🧰", text: "Teach me Docker basics with practical steps and common pitfalls", style: "tutorial" },
  { icon: "🔗", text: "Compare REST vs GraphQL for building modern APIs", style: "youtuber" },
] as const;

export default function AudioGenPage() {
  const [prompt, setPrompt] = useState("");
  const [selectedStyle, setSelectedStyle] = useState<string>("podcast");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [generatedAudioUrls, setGeneratedAudioUrls] = useState<string[]>([]);
  const { setTrack } = useAudio();
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);

  useEffect(() => {
    return () => {
      generatedAudioUrls.forEach(url => {
        URL.revokeObjectURL(url);
      });
    };
  }, [generatedAudioUrls]);

  async function generateAudio() {
    setError(null);
    try {
      const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL || 'http://localhost:8080';
      const response = await fetch(`${backendUrl}/api/tts`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({
          prompt,
          style: selectedStyle,
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Failed to generate audio: ${errorText}`);
      }

      const audioBlob = await response.blob();
      
      if (audioBlob.size === 0) {
        throw new Error('Generated audio file is empty');
      }
      
      const audioUrl = URL.createObjectURL(audioBlob);
      
      setGeneratedAudioUrls(prev => [...prev, audioUrl]);
      
      const selectedStyleName = STYLE_OPTIONS.find(s => s.id === selectedStyle)?.name || selectedStyle;
      setTrack(audioUrl, `${selectedStyleName}: ${prompt.slice(0, 40)}...`);
      
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Unknown error occurred';
      setError(errorMessage);
      console.error("Error generating audio:", errorMessage);
    }
  }

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!prompt.trim()) return;
    
    setIsLoading(true);
    try {
      await generateAudio();
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <main className="mx-auto max-w-4xl px-6 py-12 sm:py-16">
      <div className="text-center mb-12">
        <h1 className="text-3xl sm:text-5xl font-bold tracking-tight text-indigo-600 mb-4">
          Create Audio Content
        </h1>
        <p className="text-lg sm:text-xl text-slate-600">
          Describe any topic and choose your preferred learning style
        </p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-8">
        {/* Style Selection */}
        <div className="space-y-4">
          <h2 className="text-lg font-semibold text-slate-900">Choose your style</h2>
          <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
            {STYLE_OPTIONS.map((style) => (
              <label
                key={style.id}
                className={`cursor-pointer rounded-xl border-2 p-4 transition-all duration-200 ${
                  selectedStyle === style.id
                    ? 'border-indigo-500 bg-indigo-50'
                    : 'border-slate-200 hover:border-slate-300'
                }`}
              >
                <input
                  type="radio"
                  name="style"
                  value={style.id}
                  checked={selectedStyle === style.id}
                  onChange={(e) => setSelectedStyle(e.target.value)}
                  className="sr-only"
                />
                <div className="text-center">
                  <div className="text-2xl mb-2">{style.icon}</div>
                  <div className="font-medium text-slate-900">{style.name}</div>
                  <div className="text-sm text-slate-500 mt-1">{style.description}</div>
                </div>
              </label>
            ))}
          </div>
        </div>

        {/* Prompt Input */}
        <div className="space-y-3">
          <label htmlFor="prompt" className="text-lg font-semibold text-slate-900">
            What would you like to learn about?
          </label>
          <textarea
            id="prompt"
            ref={textareaRef}
            value={prompt}
            onChange={(e) => setPrompt(e.target.value)}
            rows={6}
            placeholder="e.g. Explain the fundamentals of machine learning, including key concepts and real-world applications"
            className="w-full resize-y rounded-xl border-2 border-slate-200 bg-white px-4 py-3 text-base placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-0 transition-colors"
          />
        </div>

        {/* Error Display */}
        {error && (
          <div className="rounded-xl bg-amber-50 border border-amber-200 p-4">
            <div className="flex items-center gap-3">
              <div className="text-amber-500">⚠️</div>
              <div>
                <h3 className="font-medium text-amber-900">Generation Error</h3>
                <p className="text-sm text-amber-700 mt-1">{error}</p>
              </div>
            </div>
          </div>
        )}

        {/* Generate Button */}
        <div className="flex justify-center">
          <button
            type="submit"
            disabled={isLoading || !prompt.trim()}
            className="inline-flex items-center justify-center gap-3 rounded-xl bg-indigo-600 px-8 py-4 text-base font-semibold text-white hover:bg-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 shadow-lg hover:shadow-xl min-w-[200px]"
          >
            {isLoading && (
              <div className="h-5 w-5 animate-spin rounded-full border-2 border-white border-r-transparent" />
            )}
            <span>{isLoading ? "Generating..." : "Generate Audio"}</span>
          </button>
        </div>
      </form>

      {/* Example Prompts */}
      <section className="mt-12">
        <h2 className="text-lg font-semibold text-slate-900 mb-6 text-center">
          Need inspiration? Try these examples
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {EXAMPLE_PROMPTS.map((ex) => (
            <button
              key={ex.text}
              type="button"
              onClick={() => {
                setPrompt(ex.text);
                setSelectedStyle(ex.style);
                setTimeout(() => {
                  const el = textareaRef.current;
                  if (el) {
                    el.focus();
                    try {
                      el.setSelectionRange(el.value.length, el.value.length);
                    } catch {}
                  }
                }, 0);
              }}
              className="group text-left rounded-xl border-2 border-slate-200 p-5 hover:border-indigo-300 hover:shadow-md transition-all duration-200 bg-white"
            >
              <div className="flex items-start gap-4">
                <div className="text-2xl">{ex.icon}</div>
                <div className="flex-1">
                  <div className="font-medium text-slate-900 mb-2 leading-relaxed">{ex.text}</div>
                  <div className="inline-flex items-center gap-2 text-sm text-indigo-600">
                    <span className="capitalize">{ex.style} style</span>
                    <span>→</span>
                  </div>
                </div>
              </div>
            </button>
          ))}
        </div>
      </section>

      <BottomPlayer />
    </main>
  );
}


