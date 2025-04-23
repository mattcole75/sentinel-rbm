import { redirect, useNavigate } from "@remix-run/react";

import SystemForm from "../components/system/SystemForm";
import Modal from "../components/util/Modal";
import { addSystem } from "../data/systems.server";
import { validateSystemInput } from "../data/validation.server";

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

    const formData = await request.formData();
    const systemData = Object.fromEntries(formData)

    try{
        validateSystemInput(systemData);
    } catch(err) {
        return err;
    }

    await addSystem(systemData);

    return redirect("/systems");
}