import AuthForm from "../components/auth/AuthForm";
import { validateCredentials } from "../data/validation.server";

import styles from "../styles/auth.css?url";

export default function AuthPage() {
    return (
        <AuthForm />
    );
}

export async function action({request}) {
    const searchParams = new URL(request.url).searchParams;
    const authMode = searchParams.get("mode") || "login";

    const formData = request.formData();
    const credentials = Object.fromEntries(formData);

    // validate user input
    try {
        validateCredentials(credentials);
    } catch(err) {
        return err;
    }


    if (authMode === "login") {
        // login logic
    } else {
        // signup logic
    }
}

export function links() {
    return [
        { rel: "stylesheet", href: styles }
    ]
}
