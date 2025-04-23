import { Link, Outlet, useLoaderData } from "@remix-run/react";
import SystemList from "../components/system/SystemList";

import { FaPlus, FaDownload } from "react-icons/fa";
import { getSystems } from "../data/systems.server";

export default function SystemListLayout() {
    const systems = useLoaderData();

    return (
        <>
            <Outlet />
            <main>
                <section id="systems-actions">
                    <Link to="add">
                        <FaPlus />
                        <span>Add System</span>
                    </Link>
                    <a href="/systems/raw" target="_blank" rel="noopener noreferrer">
                        <FaDownload />
                        <span>Download Raw Data</span>
                    </a>
                </section>
                <SystemList systems={ systems } />
            </main>
        </>
    );
}

export  function loader() {
    return getSystems();
}
