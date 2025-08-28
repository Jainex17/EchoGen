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
  return (
    <div className="fixed inset-x-0 bottom-4 z-40 flex justify-center pointer-events-none">
      <div className="pointer-events-auto w-full max-w-2xl mx-4 rounded-lg border border-black/10 dark:border-white/15 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/80 shadow-sm">
        <div className="px-4 py-3 flex items-center gap-4">
          <button
            type="button"
            className="h-10 w-10 inline-flex items-center justify-center rounded-full border border-black/10 dark:border-white/15 disabled:opacity-50"
            onClick={isPlaying ? pause : play}
            disabled={!hasTrack}
            aria-label={isPlaying ? "Pause" : "Play"}
          >
            {isPlaying ? (
              <svg
                className="w-4 h-4"
                viewBox="0 0 24 24"
                fill="currentColor"
                aria-hidden="true"
              >
                <rect x="6" y="5" width="4" height="14" rx="1" />
                <rect x="14" y="5" width="4" height="14" rx="1" />
              </svg>
            ) : (
              <svg
                className="w-4 h-4"
                viewBox="0 0 24 24"
                fill="currentColor"
                aria-hidden="true"
              >
                <path d="M8 5v14l11-7-11-7z" />
              </svg>
            )}
          </button>

          <div className="min-w-0 flex-1">
            <p className="truncate text-sm font-medium">
              {title ?? (hasTrack ? "Playing audio" : "No track")}
            </p>
            <div className="mt-1 flex items-center gap-3">
              <span className="text-xs tabular-nums text-black/60 dark:text-white/60">{formatTime(currentTime)}</span>
              <input
                type="range"
                min={0}
                max={duration || 0}
                step={0.01}
                value={Number.isFinite(duration) ? Math.min(currentTime, duration) : 0}
                onChange={(e) => seek(parseFloat((e.target as HTMLInputElement).value))}
                className="flex-1 accent-foreground"
                disabled={!hasTrack}
                aria-label="Seek"
              />
              <span className="text-xs tabular-nums text-black/60 dark:text-white/60">{formatTime(duration)}</span>
            </div>
          </div>

         
          <a
            href={audioSrc ?? undefined}
            download
            className="inline-flex items-center justify-center rounded-md border border-black/10 dark:border-white/15 p-2 disabled:opacity-50"
            aria-disabled={!hasTrack}
            aria-label="Download"
          >
            <svg
              className="w-4 h-4"
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


