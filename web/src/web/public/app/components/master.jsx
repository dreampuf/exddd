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
    var githubButton = (
      <IconButton
        className="github-icon-button"
        iconClassName="muidocs-icon-custom-github"
        href="https://github.com/callemall/material-ui"
        linkButton={true} />
    );

    return (
      <AppCanvas predefinedLayout={1}>

        <AppBar
          className="mui-dark-theme"
          onMenuIconButtonTouchTap={this._onMenuIconButtonTouchTap}
          title={title}
          zDepth={0}>
          {githubButton}
        </AppBar>


        <RouteHandler />

        <div className="footer full-width-section mui-dark-theme">
          <p>
            专灭善男信女 Yeah哈哈哈，直接刷新
          </p>
          {githubButton}
        </div>

      </AppCanvas>
    );
  },

  _onMenuIconButtonTouchTap: function() {
    this.refs.leftNav.toggle();
  }

});

module.exports = Master;

