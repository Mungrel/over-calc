import React from 'react'
import { Switch, Route } from 'react-router-dom'
import SignUp from './pages/SignUp'
import SignIn from './pages/SignIn'
import NotFound from './pages/NotFound'
import Calc from './pages/Calc';

const Routes: React.FC = () => {
    return (
        <Switch>
            <Route exact path="/sign_up" component={SignUp} />
            <Route exact path="/sign_in" component={SignIn} />
            <Route exact path="/calc" component={Calc} />
            <Route component={NotFound} />
        </Switch>
    )
}

export default Routes
