const RouterNav = ({groups, handler}) => {
    return (
        <div className="router-nav">
            <ul uk-tab="true" className="uk-tab uk-text-bold">
                {groups.map((group, index) => {return (
                    <li key={index}>
                        <a href="#" value={group} onClick={(e) => handler(e)}>{group}</a>
                    </li>
                )})}
            </ul>
        </div>
    );
};
