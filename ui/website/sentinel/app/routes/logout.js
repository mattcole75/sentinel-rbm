import { json } from "stream/consumers";
import { destroyUserSession } from "../data/auth.server";

export function action({ request }) {
    if (request.method !== "POST") {
        throw json({ message: "invalid request method" }, { status: 400 });
    }

    return destroyUserSession(request)
}
