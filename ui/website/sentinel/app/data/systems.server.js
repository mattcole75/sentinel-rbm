import { prisma } from "./database.server";

export async function addSystem(data) {
    prisma.System.create();
}
