var React = require('react');
var Router = require('react-router');
var Route = Router.Route;
var Redirect = Router.Redirect;
var DefaultRoute = Router.DefaultRoute;

// Here we define all our material-ui ReactComponents.
var Master = require('./components/Master.jsx');
var Main = require('./components/main.jsx');
var Home = require('./components/pages/home.jsx');


/** Routes: https://github.com/rackt/react-router/blob/master/docs/api/components/Route.md
 *
 * Routes are used to declare your view hierarchy.
 *
 * Say you go to http://material-ui.com/#/components/paper
 * The react router will search for a route named 'paper' and will recursively render its
 * handler and its parent handler like so: Paper > Components > Master
 */

var AppRoutes = (
    <Route name="root" path="/" handler={Master}>
        <Route name="home" handler={Home} />
        <Route name="view" handler={Home} />

        <DefaultRoute handler={Home}/>
    </Route>
);

module.exports = AppRoutes;