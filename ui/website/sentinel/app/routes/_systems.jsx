import { Outlet } from "@remix-run/react";

import SystemsHeader from "../components/navigation/SystemsHeader";
import styles from "../styles/systems.css?url";
import { requireUserSession } from "../data/auth.server";

export default function SystemsLayout() {
    return (
        <>
            <SystemsHeader />
            <Outlet />
        </>
    );
}

export async function loader({request}) {
    return await requireUserSession(request);
}

export function links() {
    return [
        { rel: "stylesheet", href: styles }
    ]
}