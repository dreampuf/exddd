var React = require('react');
var Router = require('react-router');
var RouteHandler = Router.RouteHandler;
var mui = require('material-ui');
var AppBar = mui.AppBar;
var AppCanvas = mui.AppCanvas;
var Menu = mui.Menu;
var IconButton = mui.IconButton;

var Master = React.createClass({

  mixins: [Router.State],

  render: function() {

    var title =
      this.isActive('get-started') ? 'Get Started' :
      this.isActive('css-framework') ? 'Css Framework' :
      this.isActive('components') ? 'Components' : '';

    return (
      <AppCanvas predefinedLayout={1}>

        <AppBar
          className="mui-dark-theme"
          onMenuIconButtonTouchTap={this._onMenuIconButtonTouchTap}
          title={title}
          zDepth={0}>

          <IconButton
            className="weibo-icon-button"
            iconClassName="icon-weibo"
            href="/connect/weibo"
            linkButton={true}
            tooltip="Weibo Sign In"/>
        </AppBar>

        <RouteHandler />

        <div className="footer full-width-section mui-dark-theme">
          <p>
            placehost<br />
            nini<br />
            ninini<br />
          </p>
        </div>

      </AppCanvas>
    );
  },

  _onMenuIconButtonTouchTap: function() {
    //this.refs.leftNav.toggle();
  }

});

module.exports = Master;

