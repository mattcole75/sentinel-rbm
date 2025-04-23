import { Form, Link, useActionData, useMatches, useNavigation, useParams } from "@remix-run/react";

function SystemForm() {
    // const today = new Date().toISOString().slice(0, 10); // yields something like 2023-09-10
    const validationErrors = useActionData();

    // instead of using a loader to query the database we can pull the object from the list in the systems list as follows
    const params = useParams(); // pull the route parameter, in the case it will be id
    const matches = useMatches(); // matches contains all current active routes
    const systems = matches.find(match => match.id === "routes/_systems.systems._list").data; // accesses the array in this route, instead of querying the database for the record
    const systemData = systems && systems.find(system => system.id === parseInt(params.id)) // find the object in the list based on its id
    
    const navigation = useNavigation();

    const defaultValues = systemData 
        ? {
            name: systemData.name,
            description: systemData.description
        }
        : {
            name: "",
            description: ""
        }


    const isSubmitting = navigation.state !== "idle";

    return (
        <Form method={systemData ? "patch" : "delete"} className="form" id="system-form">
            <p>
                <label htmlFor="name">System Name</label>
                <input type="text" id="name" name="name" required minLength={1} maxLength={64} defaultValue={defaultValues.name} />
            </p>

            <p>
                <label htmlFor="description">System Description</label>
                <textarea type="" id="description" name="description" rows="8" required minLength={1} maxLength={512} defaultValue={defaultValues.description} />
            </p>

            {/* <p>
                <label htmlFor="reference">Requirement Reference</label>
                <input type="text" id="reference" name="reference" required minLength={1} maxLength={64} defaultValue={defaultValues.reference} />
            </p>

            <p>
                <label htmlFor="referenceSource">Requirement Reference Source</label>
                <input type="text" id="referenceSource" name="referenceSource" required defaultValue={defaultValues.referenceSource} />
            </p> */}

            {/* <div className="form-row">
                <p>
                    <label htmlFor="amount">Amount</label>
                    <input
                        type="number"
                        id="amount"
                        name="amount"
                        min="0"
                        step="0.01"
                        required
                    />
                </p>
                <p>
                    <label htmlFor="date">Date</label>
                    <input type="date" id="date" name="date" max={today} required />
                </p>
            </div> */}
            {validationErrors && 
                <ul>
                    {Object.values(validationErrors).map((err) => (
                        <li key={err}>{err}</li>
                    ))}
                </ul>}

            <div className="form-actions">
                <button disabled={isSubmitting}>{isSubmitting ? "Saving..." : "Save Requirement"}</button>
                <Link to="..">Cancel</Link>
            </div>
        </Form>
    );
}

export default SystemForm;
