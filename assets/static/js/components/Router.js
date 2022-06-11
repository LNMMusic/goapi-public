const Router = () => {
    // States
    const [state, setState] = useState({
        groups: Object.keys(data),  groupSelected:  "root",
        endpoints: data["root"],    endpointSelected: null,
    });

    // Handlers
    function handleGroupSelected(e) {
        let group = e.target.getAttribute("value");

        setState(Object.assign({},
            state,
            {
                ["groupSelected"]:  group,
                ["endpoints"]:      data[group]
            }
        ))
    }
    function handleEndpointSelected(e) {
        let endpoint = e.target.getAttribute("value");

        setState(Object.assign({},
            state,
            {["endpointSelected"]:data[state.groupSelected][endpoint]}
        ))
    }


    return (
        <div className="section-router">
            <div className="router">
                <RouterNav
                    groups={state.groups}
                    handler={handleGroupSelected}
                />
                <RouterContent
                    endpoints={state.endpoints}
                    handler={handleEndpointSelected}
                />
            </div>
            <hr className="uk-divider-vertical"></hr>
            <div className="fetch">
                {/* <Test endpoint={state.endpointSelected} /> */}
                <Fetch endpoint={state.endpointSelected}/>
            </div>
        </div>
    );
};