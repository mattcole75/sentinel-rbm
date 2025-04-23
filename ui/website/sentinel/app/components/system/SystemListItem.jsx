import { Link, useFetcher } from "@remix-run/react";

function SystemListItem({ id, name, description }) {
    const fetcher = useFetcher();
    
    function deleteSystemItemHandler() {
        const proceed = confirm('Are you sure you want to delete this item?');
        if(!proceed)
            return;
        
        fetcher.submit(null, { method: "delete", action: `/systems/${id}` });
    }

    if (fetcher.state !== "idle") {
        return <article className="system-item locked">
            <p>Deleting...</p>
        </article>
    }

return (
    <article className="system-item">
        <div>
            <h2 className="system-title">{ name }</h2>
            <p className="requirement-statement">{ description }</p>
        </div>
        <menu className="system-actions">
            <button onClick={ deleteSystemItemHandler }>Delete</button>
            <Link to={ "/systems/"+ id }>Edit</Link>
        </menu>
    </article>
  );
}

export default SystemListItem;
