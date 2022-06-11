const Header = () => {
    // states
    const [state, setState] = useState({
        title:  "API Router"
    });



    return (
        <header>
            <div className='title-main'>
                <div className="container flex-container">
                    <h1 className="title">{state.title}</h1>
                </div>
            </div>
        </header>
    );
};