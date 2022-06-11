// APP COMPONENT    ->      [Global for Pages + Components] [as HTML Body]
class App extends React.Component {
    render() {
        return (
            <>
                <Header />
                <div className="divider">
                    <hr className="uk-divider-icon"></hr>
                </div>
                <Router />
            </>
        )
    }
}


// RENDER
ReactDOM.render(
    <App />,
    document.getElementById('root')
)