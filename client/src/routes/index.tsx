import { useRoutes } from "react-router-dom";
import { paths } from "./paths";
import React from "react";
import AuthGuard from "../auth/guard/auth-guard";


const HomeView = React.lazy(() => import("../sections/home"));
const LoginView = React.lazy(() => import("../sections/auth/login"));
const RegisterView = React.lazy(() => import("../sections/auth/register"));

const routes = [
    {
        path: paths.home,
        element: (
            <AuthGuard>
            <HomeView />
            </AuthGuard>
        )
    },
    {
        path: paths.login,
        element: (
            <>
            <LoginView />
            </>
        )
    },
    {
        path: paths.register,
        element:  (
            <>
            <RegisterView />
            </>
        )
    },
    {
        path: paths.logout,
        element: (
            <>
            <LoginView />
            </>
        )
    },
    { path: "*", element: <div>404</div> },
]

const Router = () => {
    return useRoutes(routes);
};

export default Router;