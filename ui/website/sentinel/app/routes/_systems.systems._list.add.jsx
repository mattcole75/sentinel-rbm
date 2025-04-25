import { redirect, useNavigate } from "@remix-run/react";

import SystemForm from "../components/system/SystemForm";
import Modal from "../components/util/Modal";
import { addSystem } from "../data/systems.server";
import { validateSystemInput } from "../data/validation.server";
import { requireUserSession } from "../data/auth.server";

export default function AddSystemPage() {

    const navigate = useNavigate();

    function closeHandler() {
        navigate("..");
    }

    return (
        <Modal onClose={closeHandler}>
            <SystemForm />
        </Modal>
    );
}

export async function action({request}) {

    // don't need to check for cookie only getting the userID
    const userId = await requireUserSession(request);

    const formData = await request.formData();
    const systemData = Object.fromEntries(formData);

    try{
        validateSystemInput(systemData);
    } catch(err) {
        return err;
    }

    await addSystem(systemData, userId);

    return redirect("/systems");
}