/** In this file, we create a React component which incorporates components provided by material-ui */

var React = require('react'),
  mui = require('material-ui'),
  RaisedButton = mui.RaisedButton;

var Home = React.createClass({

  render: function() {

    return (
      <div className="example-page">

        <h1>T T</h1>
        <h2>s  nimenenienie </h2>

        <RaisedButton label="Super Secret Password" primary={true} onTouchTap={this._handleTouchTap} />

      </div>
    );
  },

  _handleTouchTap: function() {
    alert('1-2-3-4-5');
  }
  
});

module.exports = Home;
