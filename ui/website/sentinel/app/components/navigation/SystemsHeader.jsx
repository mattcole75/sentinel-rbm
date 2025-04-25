import { Form, NavLink } from "@remix-run/react";
import Logo from "../util/Logo";

export default function SystemsHeader() {
    return (
        <header id="main-header">
            <Logo />
            <nav id="main-nav">
                <ul>
                    <li>
                        <NavLink to="/systems" end>
                            Manage Systems
                        </NavLink>
                    </li>
                    <li>
                        <NavLink to="/systems/analysis">
                            Analise Systems
                        </NavLink>
                    </li>
                </ul>
            </nav>
            <nav id="cta-nav">
                <Form method="post" action="logout">
                    <button className="cta-alt">Logout</button>
                </Form>
            </nav>
        </header>
    );
}
