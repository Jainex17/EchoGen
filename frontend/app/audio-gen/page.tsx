"use client";

import { useState, useRef, useEffect } from "react";
import { useAudio } from "../components/AudioProvider";
import BottomPlayer from "../components/BottomPlayer";

const EXAMPLE_PROMPTS = [
  { icon: "🎧", text: "Summarize the history of AI in 3 minutes for a podcast" },
  { icon: "📚", text: "Explain React hooks to a beginner with analogies and examples" },
  { icon: "🧰", text: "Teach me Docker basics with practical steps and common pitfalls" },
  { icon: "🔗", text: "Compare REST vs GraphQL for building modern APIs" },
] as const;

export default function AudioGenPage() {
  const [prompt, setPrompt] = useState("");
  const [isLoading, setIsLoading] = useState(false);
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
    try {
      const response = await fetch('http://localhost:8080/api/tts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          prompt,
        }),
      });

      if (!response.ok) {
        console.error("Error generating audio:", response.statusText);
        return;
      }

      const audioBlob = await response.blob();
      
      const audioUrl = URL.createObjectURL(audioBlob);
      
      setGeneratedAudioUrls(prev => [...prev, audioUrl]);
      
      setTrack(audioUrl, `Generated Audio: ${prompt.slice(0, 50)}...`);
      
      console.log("Audio generated successfully!");
      
    } catch (error) {
      console.error("Error generating audio:", error);
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
    <main className="mx-auto max-w-3xl px-6 py-12 sm:py-16">
      <h1 className="text-2xl sm:text-5xl font-semibold tracking-tight text-center">Create learning audio</h1>
      <p className="mt-2 text-sm sm:text-lg text-black/70 dark:text-white/70 text-center">
        Describe a topic and pick a style. Your audio will be generated here.
      </p>

      <form onSubmit={handleSubmit} className="mt-6 rounded-lg space-y-4">
        <div className="space-y-2">
          <textarea
            id="prompt"
            ref={textareaRef}
            value={prompt}
            onChange={(e) => setPrompt(e.target.value)}
            rows={8}
            placeholder="e.g. Teach me the basics of Docker in a tutorial style"
            className="w-full resize-y rounded-md border border-black/10 dark:border-white/15 bg-transparent px-3 py-2 text-sm outline-none focus:ring-2 focus:ring-black/10 dark:focus:ring-white/15"
          />
        </div>

        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-center gap-3">
          
          <button
            type="submit"
            disabled={isLoading}
            className="inline-flex items-center justify-center gap-2 rounded-md bg-foreground py-3 text-sm font-medium text-background disabled:opacity-60 w-2/6"
          >
            {isLoading && (
              <span className="h-4 w-4 animate-spin rounded-full border-2 border-background border-r-transparent" aria-hidden />
            )}
            {isLoading ? "Generating…" : "Start Generating"}
          </button>
        </div>
      </form>

      <section className="mt-8">
        <h2 className="text-sm font-medium">Example prompts</h2>
        <div className="mt-3 grid grid-cols-1 sm:grid-cols-2 gap-3 sm:gap-4">
          {EXAMPLE_PROMPTS.map((ex) => (
            <button
              key={ex.text}
              type="button"
              onClick={() => {
                setPrompt(ex.text);
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
              className="group flex items-center text-center gap-2 rounded-lg border border-black/10 dark:border-white/15 p-4 text-sm hover:bg-black/5 dark:hover:bg-white/5 focus:outline-none focus:ring-2 focus:ring-black/10 dark:focus:ring-white/15 transition-colors"
            >
              <span className="inline-flex h-10 w-10 shrink-0 items-center justify-center rounded-md bg-black/5 dark:bg-white/10 text-lg">
                {ex.icon}
              </span>
              <span className="leading-relaxed">{ex.text}</span>
            </button>
          ))}
        </div>
      </section>

      <BottomPlayer />
    </main>
  );
}


