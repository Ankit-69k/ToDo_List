import { AuthContext } from "./auth-context";

type Props = {
  children: React.ReactNode;
};

export function AuthConsumer({ children }: Props) {
  return (
    // eslint-disable-next-line react/react-in-jsx-scope
    <AuthContext.Consumer>
      {(auth) => (auth.loading ? <>wait..</>: children)}
    </AuthContext.Consumer>
  );
}
