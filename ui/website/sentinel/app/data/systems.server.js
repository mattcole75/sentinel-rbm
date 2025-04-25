import { prisma } from "./database.server";

export async function addSystem(data, userId) {
    try {
        return await prisma.system.create({
            data: {
                name: data.name,
                description: data.description,
                user: { connect: { id: userId } },
            },
        });
    } catch (err) {
        throw new Error("failed to add new system");
    }
}

export async function getSystems(userId) {
    if (!userId) {
        throw new Error("failed to get systems");
    }

    try {
        const systems = await prisma.system.findMany({
            where: { userId: userId },
            orderBy: { name: "desc" },
        });
        return systems;
    } catch (err) {
        throw new Error("failed to get systems");
    }
}

export async function getSystem(id) {
    try {
        const system = await prisma.system.findFirst({
            where: { id },
        });
        return system;
    } catch (err) {
        throw new Error("failed to get system");
    }
}

export async function updateSystem(id, data) {
    try {
        return await prisma.system.update({
            where: { id },
            data: {
                name: data.name,
                description: data.description,
            },
        });
    } catch (err) {
        throw new Error("failed to update system");
    }
}

export async function deleteSystem(id) {
    try {
        return await prisma.system.delete({ where: { id } });
    } catch (err) {
        throw new Error("failed to delete systems");
    }
    // prisma.system
    //     .delete({ where: { id } })
    //     .then((res) => {
    //         console.log(`${id} delete successfully!`, res);
    //     })
    //     .catch((err) => {
    //         throw new Error("failed to delete systems");
    //     });
}
