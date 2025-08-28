"use client";

import React, { createContext, useCallback, useContext, useEffect, useMemo, useRef, useState } from "react";

type AudioContextValue = {
  audioSrc: string | null;
  title: string | null;
  isPlaying: boolean;
  duration: number;
  currentTime: number;
  setTrack: (src: string, title?: string) => void;
  play: () => void;
  pause: () => void;
  seek: (time: number) => void;
};

const AudioCtx = createContext<AudioContextValue | undefined>(undefined);

export function useAudio() {
  const ctx = useContext(AudioCtx);
  if (!ctx) throw new Error("useAudio must be used within AudioProvider");
  return ctx;
}

export function AudioProvider({ children }: { children: React.ReactNode }) {
  const audioRef = useRef<HTMLAudioElement | null>(null);
  const [audioSrc, setAudioSrc] = useState<string | null>(null);
  const [title, setTitle] = useState<string | null>(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [duration, setDuration] = useState(0);
  const [currentTime, setCurrentTime] = useState(0);

  const ensureAudio = useCallback(() => {
    if (!audioRef.current) {
      audioRef.current = new Audio();
      audioRef.current.preload = "metadata";
    }
    return audioRef.current;
  }, []);

  const setTrack = useCallback((src: string, t?: string) => {
    const audio = ensureAudio();
    setAudioSrc(src);
    setTitle(t ?? null);
    audio.src = src;
    audio.load();
    void audio.play().catch(() => {});
    setIsPlaying(true);
  }, [ensureAudio]);

  const play = useCallback(() => {
    const audio = ensureAudio();
    void audio.play().catch(() => {});
    setIsPlaying(true);
  }, [ensureAudio]);

  const pause = useCallback(() => {
    const audio = ensureAudio();
    audio.pause();
    setIsPlaying(false);
  }, [ensureAudio]);

  const seek = useCallback((time: number) => {
    const audio = ensureAudio();
    audio.currentTime = Math.max(0, Math.min(time, audio.duration || time));
  }, [ensureAudio]);

  useEffect(() => {
    const audio = ensureAudio();
    const onLoaded = () => setDuration(Number.isFinite(audio.duration) ? audio.duration : 0);
    const onTime = () => setCurrentTime(audio.currentTime);
    const onEnded = () => setIsPlaying(false);
    const onPlay = () => setIsPlaying(true);
    const onPause = () => setIsPlaying(false);
    const onError = () => setIsPlaying(false);
    audio.addEventListener("loadedmetadata", onLoaded);
    audio.addEventListener("timeupdate", onTime);
    audio.addEventListener("ended", onEnded);
    audio.addEventListener("play", onPlay);
    audio.addEventListener("pause", onPause);
    audio.addEventListener("error", onError);
    return () => {
      audio.removeEventListener("loadedmetadata", onLoaded);
      audio.removeEventListener("timeupdate", onTime);
      audio.removeEventListener("ended", onEnded);
      audio.removeEventListener("play", onPlay);
      audio.removeEventListener("pause", onPause);
      audio.removeEventListener("error", onError);
    };
  }, [ensureAudio]);

  const value = useMemo<AudioContextValue>(() => ({
    audioSrc,
    title,
    isPlaying,
    duration,
    currentTime,
    setTrack,
    play,
    pause,
    seek,
  }), [audioSrc, title, isPlaying, duration, currentTime, setTrack, play, pause, seek]);

  return (
    <AudioCtx.Provider value={value}>
      {children}
    </AudioCtx.Provider>
  );
}


