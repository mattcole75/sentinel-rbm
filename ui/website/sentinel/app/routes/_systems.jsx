import { Outlet } from "@remix-run/react";

import SystemsHeader from "../components/navigation/SystemsHeader";
import styles from "../styles/systems.css?url";

export default function SystemsLayout() {
    return (
        <>
            <SystemsHeader />
            <Outlet />
        </>
    );
}

export function links() {
    return [
        { rel: "stylesheet", href: styles }
    ]
}