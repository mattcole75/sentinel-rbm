import {
    isRouteErrorResponse,
    Link,
    Links,
    Meta,
    Outlet,
    Scripts,
    ScrollRestoration,
    useMatches,
    useRouteError,
} from "@remix-run/react";

import Error from "./components/util/Error";

import styles from "./styles/shared.css?url";


function Layout({ children }) {
    const matches = useMatches();
    const disableJS = matches.some(match => match.handle?.disableJS);

  return (
    <html lang="en">
        <head>
            <meta charSet="utf-8" />
            <meta name="viewport" content="width=device-width, initial-scale=1" />
            <Meta />
            <Links />
        </head>
        <body>
            { children }
            <ScrollRestoration />
            { !disableJS && <Scripts /> }
        </body>
    </html>
  );
}

export default function App() {
  return (
    <Layout>
        <Outlet />;
    </Layout>
  );
}

export function links() {
    return [
        {
            rel: "preconnect",
            href: "https://fonts.googleapis.com" 
        },
        {
            rel: "preconnect",
            href: "https://fonts.gstatic.com",
            crossOrigin: "anonymous",
        },
        {
            rel: "stylesheet",
            href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
        },
        { 
            rel: "stylesheet",
            href: styles 
        }
    ]
}

export function ErrorBoundary() {
    console.log("Entered Error Boundary");

    const error = useRouteError();

    if (isRouteErrorResponse(error)) {
        return (
            <Layout>
                <main>
                    <Error title={error.statusText}>
                        <p>{error.data?.message || "Something went wrong! Please try again later."}</p>
                        <p>Back to <Link to="/">safety</Link>.</p>
                    </Error>
                </main>
            </Layout>
        );
    } else if (error instanceof Error) {
        return (
            <Layout>
                <main>
                    <Error title={error.statusText}>
                        <p>{error.data?.message || "Something went wrong! Please try again later."}</p>
                        {/* <p>The stack trace is:</p>
                        <pre>{error.stack}</pre> */}
                        <p>Back to <Link to="/">safety</Link>.</p>
                    </Error>
                </main>
            </Layout>
        );
    } else {
        return (
            <Layout>
                <main>
                    <Error title={error.statusText}>
                        <p>{error.data?.message || "Something went wrong! Please try again later."}</p>
                        {/* <p>The stack trace is:</p> */}
                        {/* <pre>{error.stack}</pre> */}
                        <p>Back to <Link to="/">safety</Link>.</p>
                    </Error>
                </main>
            </Layout>
        );
    }
}
