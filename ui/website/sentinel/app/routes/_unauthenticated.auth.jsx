import AuthForm from "../components/auth/AuthForm";
import { login, signup } from "../data/auth.server";
import { validateCredentials } from "../data/validation.server";

import styles from "../styles/auth.css?url";

export default function AuthPage() {
    return <AuthForm />;
}

export async function action({ request }) {
    const searchParams = new URL(request.url).searchParams;
    const authMode = searchParams.get("mode") || "login";

    const formData = await request.formData();
    const credentials = Object.fromEntries(formData);

    // validate user input
    try {
        validateCredentials(credentials);
    } catch (err) {
        return err;
    }

    try {
        if (authMode === "login") {
            return await login(credentials);
        } else {
            return await signup(credentials);
        }
    } catch (err) {
        if (err.status === 422) {
            return { credentials: err.message };
        }
    }
    console.log("fell through")
}

export function links() {
    return [{ rel: "stylesheet", href: styles }];
}
