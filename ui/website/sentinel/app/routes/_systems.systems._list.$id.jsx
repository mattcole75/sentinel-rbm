import { redirect, useNavigate } from "@remix-run/react";
import SystemForm from "../components/system/SystemForm";

import Modal from "../components/util/Modal";
import { deleteSystem, updateSystem } from "../data/systems.server";
import { validateSystemInput } from "../data/validation.server";
// import { getRequirement } from "../data/requirements.server";

export default function UpdateSystemPage() {

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

export async function action({params, request}) {
    const systemId = params.id;
    
    if (request.method === "PATCH") {
        const formData = await request.formData();
        const systemData = Object.fromEntries(formData);

        try {
            validateSystemInput(systemData);
        } catch(err) {
            return err;
        }

        await updateSystem(systemId, systemData);
        return redirect("/systems");
    } else {
        await deleteSystem(systemId);
        return redirect("/systems");
    }

}

// export function loader({ params }) {
//     const requirementId = params.id;
//     return getRequirement(requirementId)
// }