import { getSystems } from "../data/systems.server";

export function loader() {
    return getSystems();
}