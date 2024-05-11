import React, { useCallback, useEffect, useMemo, useState } from 'react';
import { GoogleAuthProvider, getAuth, onAuthStateChanged, signInWithPopup, signOut, User } from 'firebase/auth';
import { firebaseApp } from './auth';
import { AuthContext } from './auth-context';

const AUTH = getAuth(firebaseApp);

type AuthProviderProps = {
  children: React.ReactNode;
};

export function AuthProvider({ children }: AuthProviderProps): JSX.Element {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  const loginWithGoogle = useCallback(async () => {
    const provider = new GoogleAuthProvider();
    await signInWithPopup(AUTH, provider);
  }, []);

  const logout = useCallback(async () => {
    await signOut(AUTH);
  }, []);

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(AUTH, (currentUser) => {
      setUser(currentUser); // Update the local state with the current user
      setLoading(false); // Update the loading status once we've checked authentication
      if (currentUser) {
        console.log("User is authenticated:", currentUser);

      } else {
        console.log("User is not authenticated.");
       
      }
    });

    return () => {
      unsubscribe(); // Clean up the listener to prevent memory leaks
    };
  }, []); // Empty dependency array ensures this effect runs only once on component mount

  const contextValue = useMemo(
    () => ({
      user,
      loading,
      authenticated:!!user,
      loginWithGoogle,
      logout,
    }),
    [user, loading, loginWithGoogle, logout]
  );

  return (
    <AuthContext.Provider value={contextValue}>
      {loading ? <p>Wait...</p> : children}
    </AuthContext.Provider>
  );
}

export default AuthProvider;
