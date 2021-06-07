const NodePolyfillPlugin = require("node-polyfill-webpack-plugin");

module.exports = function(context, options) {
  return {
    name: 'custom-docusaurus-plugin',
    configureWebpack(config, isServer, utils) {
      return {
        plugins: [
          new NodePolyfillPlugin()
        ]
      }
    }
  }
}