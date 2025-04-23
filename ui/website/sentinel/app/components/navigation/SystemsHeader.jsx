import { NavLink } from "@remix-run/react";
import Logo from "../util/Logo";

export default function SystemsHeader() {
    return (
        <header id="main-header">
            <Logo />
            <nav id="main-nav">
                <ul>
                    <li>
                        <NavLink to="/systems" end>Manage Systems</NavLink>
                    </li>
                    <li>
                        <NavLink to="/systems/analysis">Analise Systems</NavLink>
                    </li>
                </ul>
            </nav>
            <nav id="cta-nav">
                <button className="cta">Logout</button>
            </nav>
        </header>
    );
}