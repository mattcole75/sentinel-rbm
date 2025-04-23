import SystemListItem from './SystemListItem';

function SystemList({ systems }) {

  return (
    <ol id="systems-list">
      {systems && systems.map((system) => (
        <li key={system.id}>
          <SystemListItem
            id={system.id}
            name={system.name}
            description={system.description}
          />
        </li>
      ))}
    </ol>
  );
}

export default SystemList;
