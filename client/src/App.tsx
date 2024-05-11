import './App.css'
import { AuthConsumer } from './auth/context/auth-consumer';
import AuthProvider from './auth/context/auth-provider';
import Router from './routes';


function App() {
  return (
    <AuthProvider>
      <AuthConsumer>
        <Router />
      </AuthConsumer>
    </AuthProvider>
  ) ;
}

export default App
