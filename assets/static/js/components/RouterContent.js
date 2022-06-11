const RouterContent = ({endpoints, handler}) => {
    return (
        <div className="router-content">
            <ul className="uk-list">
                {endpoints.map((endpoint, index) => {return (
                    <li key={index} value={index} className="endpoints-item" onClick={(e) => handler(e)}>
                        {endpoint.path}
                    </li>
                )})}
            </ul>
        </div>
    );
};
 