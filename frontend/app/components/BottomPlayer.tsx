"use client";

import { useAudio } from "./AudioProvider";

function formatTime(sec: number) {
  if (!Number.isFinite(sec)) return "0:00";
  const m = Math.floor(sec / 60);
  const s = Math.floor(sec % 60);
  return `${m}:${s.toString().padStart(2, "0")}`;
}

export default function BottomPlayer() {
  const { audioSrc, title, isPlaying, duration, currentTime, play, pause, seek } = useAudio();
  const hasTrack = Boolean(audioSrc);
  
  if (!hasTrack) return null;
  
  return (
    <div className="fixed inset-x-0 bottom-6 z-40 flex justify-center pointer-events-none">
      <div className="pointer-events-auto w-full max-w-3xl mx-6 rounded-2xl border border-slate-200 bg-white/95 backdrop-blur supports-[backdrop-filter]:bg-white/90 shadow-2xl shadow-slate-900/10">
        <div className="px-6 py-4 flex items-center gap-6">
          <button
            type="button"
            className="h-12 w-12 inline-flex items-center justify-center rounded-full bg-[#FFB703] hover:bg-[#e6a502] text-[#1E2D2F] shadow-lg hover:shadow-xl transition-all duration-200 transform hover:scale-105 disabled:opacity-50 disabled:hover:scale-100"
            onClick={isPlaying ? pause : play}
            disabled={!hasTrack}
            aria-label={isPlaying ? "Pause" : "Play"}
          >
            {isPlaying ? (
              <svg
                className="w-5 h-5"
                viewBox="0 0 24 24"
                fill="currentColor"
                aria-hidden="true"
              >
                <rect x="6" y="4" width="4" height="16" rx="2" />
                <rect x="14" y="4" width="4" height="16" rx="2" />
              </svg>
            ) : (
              <svg
                className="w-5 h-5 ml-1"
                viewBox="0 0 24 24"
                fill="currentColor"
                aria-hidden="true"
              >
                <path d="M8 5v14l11-7-11-7z" />
              </svg>
            )}
          </button>

          <div className="min-w-0 flex-1">
            <p className="truncate text-base font-semibold text-[#1E2D2F] mb-2">
              {title ?? "Playing audio"}
            </p>
            <div className="flex items-center gap-4">
              <span className="text-sm font-mono text-[#4F5D56] min-w-[40px]">{formatTime(currentTime)}</span>
              <div className="flex-1 relative">
                <input
                  type="range"
                  min={0}
                  max={duration || 0}
                  step={0.01}
                  value={Number.isFinite(duration) ? Math.min(currentTime, duration) : 0}
                  onChange={(e) => seek(parseFloat((e.target as HTMLInputElement).value))}
                  className="w-full h-2 bg-slate-200 rounded-lg appearance-none cursor-pointer slider"
                  disabled={!hasTrack}
                  aria-label="Seek"
                  style={{
                    background: `linear-gradient(to right, #219EBC 0%, #219EBC ${duration ? (currentTime / duration) * 100 : 0}%, rgb(226 232 240) ${duration ? (currentTime / duration) * 100 : 0}%, rgb(226 232 240) 100%)`
                  }}
                />
              </div>
              <span className="text-sm font-mono text-[#4F5D56] min-w-[40px]">{formatTime(duration)}</span>
            </div>
          </div>

          <a
            href={audioSrc ?? undefined}
            download
            className="inline-flex items-center justify-center rounded-xl border-2 border-slate-300 p-3 hover:bg-[#8ECAE6]/10 hover:border-[#8ECAE6] transition-all duration-200 disabled:opacity-50"
            aria-disabled={!hasTrack}
            aria-label="Download audio file"
            title="Download"
          >
            <svg
              className="w-5 h-5 text-[#4F5D56]"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              aria-hidden="true"
            >
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
              <polyline points="7,10 12,15 17,10" />
              <line x1="12" y1="15" x2="12" y2="3" />
            </svg>
          </a>
        </div>
      </div>
    </div>
  );
}


