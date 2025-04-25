import { requireUserSession } from "../data/auth.server";
import { getSystems } from "../data/systems.server";

export async function loader({request}) {
    const userId = await requireUserSession(request);
    return getSystems(userId);
}