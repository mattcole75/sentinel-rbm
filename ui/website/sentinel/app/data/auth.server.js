import { createCookieSessionStorage, redirect } from "@remix-run/node";
import { prisma } from "./database.server";
import { hash, compare } from "bcryptjs";

// eslint-disable-next-line no-undef
const SESSION_SECRETE = process.env.SESSION_SECRETE;

const sessionStorage = createCookieSessionStorage({
    cookie: {
        // eslint-disable-next-line no-undef
        secure: process.env.NODE_ENV === "production",
        secrets: [SESSION_SECRETE],
        sameSite: "lax",
        maxAge: 30 * 24 * 60 * 60, // 30 days
        httpOnly: true,
    },
});

async function createUserSession(userId, redirectPath) {
    const session = await sessionStorage.getSession();
    session.set("userId", userId);
    return redirect(redirectPath, {
        headers: {
            "Set-Cookie": await sessionStorage.commitSession(session),
        },
    });
}

export async function getUserFromSession(request) {
    const session = await sessionStorage.getSession(
        request.headers.get("Cookie")
    );

    const userId = session.get("userId");

    if (!userId) {
        return null;
    }

    return userId;
}

export async function destroyUserSession(request) {
    const session = await sessionStorage.getSession(
        request.headers.get("Cookie")
    );

    return redirect("/", {
        headers: {
            "Set-Cookie": await sessionStorage.destroySession(session),
        },
    });
}

export async function requireUserSession(request) {
    const userId = await getUserFromSession(request);

    if (!userId) {
        throw redirect("/auth?mode=login");
    }

    return userId;
}

export async function signup({ email, password }) {
    const existingUser = await prisma.user.findFirst({ where: { email } });

    if (existingUser) {
        const err = new Error("email address already exists");
        err.status = 422;
        throw err;
    }

    const passwordHash = await hash(password, 12);

    const user = await prisma.user.create({
        data: { email: email, password: passwordHash },
    });

    return createUserSession(user.id, "/systems");
}

export async function login({ email, password }) {
    const existingUser = await prisma.user.findFirst({ where: { email } });
    if (!existingUser) {
        const err = new Error(
            "could not log you in, please check the provided credentials"
        );
        err.status = 401;
        throw err;
    }

    // check password
    const passwordMatch = await compare(password, existingUser.password);
    if (!passwordMatch) {
        const err = new Error(
            "could not log you in, please check the provided credentials"
        );
        err.status = 401;
        throw err;
    }

    return createUserSession(existingUser.id, "/systems");
}
