"use client";

import React, { createContext, useCallback, useContext, useEffect, useMemo, useState } from "react";

type User = {
  name: string;
  email: string;
  picture: string;
};

type AuthContextValue = {
  user: User | null;
  login: () => void;
  logout: () => void;
  isLoading: boolean;
};

const AuthCtx = createContext<AuthContextValue | undefined>(undefined);

export function useAuth() {
  const ctx = useContext(AuthCtx);
  if (!ctx) throw new Error("useAuth must be used within AuthProvider");
  return ctx;
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8080";

  useEffect(() => {
    fetch(`${backendUrl}/auth/me`, {
      credentials: "include",
    })
      .then((res) => {
        if (!res.ok) throw new Error("Unauthorized");
        return res.json();
      })
      .then(setUser)
      .catch(() => console.log("Failed to fetch user data"))
      .finally(() => setIsLoading(false));
  }, [backendUrl]);

  const login = useCallback(() => {
    window.location.href = `${backendUrl}/auth/google/login`;
  }, [backendUrl]);

  const logout = useCallback(() => {
    fetch(`${backendUrl}/auth/logout`, {
      method: "POST",
      credentials: "include",
    })
      .then((res) => {
        if (!res.ok) throw new Error("Logout failed");
        setUser(null);
      })
      .catch(console.error);
  }, [backendUrl]);

  const value = useMemo<AuthContextValue>(() => ({
    user,
    login,
    logout,
    isLoading,
  }), [user, login, logout, isLoading]);

  return (
    <AuthCtx.Provider value={value}>
      {children}
    </AuthCtx.Provider>
  );
}
