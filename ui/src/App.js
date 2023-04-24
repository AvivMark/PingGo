import React from 'react';
import GroupsManager from './components/groupsManager';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Sidebar from './components/Sidebar';

function App(){
    return(
    <React.Fragment>
        <Router>
            <Sidebar />
            <Switch>
                <Route exact path="/" component={GroupsManager}/>
            </Switch>
        </Router>
    </React.Fragment>
    );
}
export default App;