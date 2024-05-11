import React, { useCallback, useEffect, useState } from "react";
import { useAuthContext } from "../../hooks/auth";
import { useRouter } from "../../hooks/router";

function AuthGuard({ children }: { children: React.ReactNode }) {
    const router = useRouter();
    const { user } = useAuthContext();

    if (!user) {
        router.push("/login");
        return <></>
    }

    return <>{children}</>;
}

export default AuthGuard;
