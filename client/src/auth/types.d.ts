

type CanRemove = {
    logout: () => Promise<void>;
  };
  
  export type FirebaseContextType = CanRemove & {
    user: firebase.User;
    loading: boolean;
    authenticated: boolean;
    loginWithGoogle: () => Promise<void>;
  };