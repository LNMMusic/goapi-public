const Test = ({endpoint}) => {
    // States
    const [state, setState] = useState({
        method: "GET",      path: "",
        parameters: {},
        body: {},           authorization: ""
    });
    useEffect(() => {
        setState(endpoint)
    }, [endpoint]);

    return (
        <div className="testing">
            <h1>Its working</h1>
            <p>{state.path}</p>
        </div>
    )
};
